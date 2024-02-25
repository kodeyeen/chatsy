<script lang="ts" setup>
import { ref, computed, watch, nextTick, onMounted } from 'vue'
import { useIntersectionObserver } from '@vueuse/core'
import { Tippy } from 'vue-tippy'

import { useAuthStore } from '@/stores/auth'
import { useChatsStore } from '@/stores/chats'

import Avatar from '@/components/Avatar.vue'
import ChatTopBar from '@/components/Chat/TopBar.vue'
import CopyButton from '@/components/CopyButton.vue'
import ArrowDownIcon from '@/components/icons/ArrowDown.vue'
import CopyIcon from '@/components/icons/Copy.vue'
import CheckmarkBigCircleIcon from '@/components/icons/CheckmarkBigCircle.vue'
import TrashDeleteBinIcon from '@/components/icons/TrashDeleteBin.vue'
import ReplyIcon from '@/components/icons/Reply.vue'
import ShareArrowIcon from '@/components/icons/ShareArrow.vue'
import MessageCard from '@/components/Message/Card.vue'
import MessageForm from '@/components/Message/Form.vue'
import Popup from '@/components/Popup.vue'

const props = defineProps({
    chat: {
        type: Object,
        required: true,
    },
})

const emit = defineEmits(['forwardMessages'])

const auth = useAuthStore()
const store = useChatsStore()

const messages = ref<any | null>(null)
const limit = ref(30)
const offsetStart = ref(0)
const offsetEnd = ref(0)
const count = ref<number | null>(null)
// при загрузке сообщений, а затем скроллинге до результата поиска
// (см. ветку if (event.include) в fetch_messages)
// можем случайно триггернуть загрузку новых сообщений сверху или снизу
// этот флаг позволит предотвращать подгрузку, пока не дойдем до результата поиска
const canFetch = ref(true)
const isFetchingMessages = ref(true)

const groupedMessages = computed(() => {
    if (!messages.value) {
        return {}
    }

    const groups: any = {}

    for (const message of messages.value) {
        const sentAt = new Date(message.sentAt)
        const formattedSentAt = formatDate(sentAt)

        if (!groups[formattedSentAt]) {
            groups[formattedSentAt] = []
        }

        groups[formattedSentAt].push(message)
    }

    return groups
})

const formatDate = (date: Date) => {
    const now = new Date()

    if (
        date.getDate() === now.getDate() &&
        date.getMonth() === now.getMonth() &&
        date.getFullYear() === now.getFullYear()
    ) {
        return 'Сегодня'
    } else if (date.getFullYear() === now.getFullYear()) {
        return date.toLocaleDateString('ru', {
            day: 'numeric',
            month: 'long',
        })
    } else {
        return date
            .toLocaleDateString('ru', {
                day: 'numeric',
                month: 'long',
                year: 'numeric',
            })
            .slice(0, -3)
    }
}

const selectedMessages = ref<any[]>([])
const messageCards = ref([])
const scrollableContainer = ref<HTMLElement | null>(null)
const messageForm = ref<HTMLFormElement | null>(null)
const topSpinner = ref(null)
const bottomSpinner = ref(null)
const scrollTop = ref(0)
const contextMenu = ref<any | null>(null)
const contextMenuMessage = ref<any | null>(null)
const messageToReply = ref(null)
const messageToDelete = ref(null)
const messageToForward = ref(null)
const isSelecting = ref(false)
const isPanelActive = ref(false)
const isSearching = ref(false)

const scrollToMessage = (messageId: number, highlight: boolean = true) => {
    const messageCard: any = messageCards.value.find((card: any) => {
        return card.$.props.message.id === messageId
    })

    if (!messageCard) {
        throw new Error('Message not found')
    }

    messageCard.$el.scrollIntoView({
        behavior: 'smooth',
        block: 'center',
    })

    if (!highlight) {
        return
    }

    messageCard.$el.classList.add('selected')

    setTimeout(() => {
        messageCard.$el.classList.add('duration-[2000ms]')
        messageCard.$el.classList.remove('selected')

        setTimeout(() => messageCard.$el.classList.remove('duration-[2000ms]'), 2000)
    }, 1000)
}

const scrollToMessageOrFetch = (messageId: number, highlight: boolean = true) => {
    try {
        scrollToMessage(messageId, highlight)
    } catch {
        store.fetchMessages(props.chat.id, {
            include: messageId,
            highlight,
        })

        canFetch.value = false
    }
}

const eventHandlers: any = {
    'chat_opened': (event: any) => {
        // if (event.is_searching) {
        //   return
        // }

        if (!messages.value || event.payload.include) {
            messages.value = event.payload.messages.items
        } else {
            // после подгрузки сообщений скролл прыгает
            // данный кусок кода исправляет это

            // сохраняем старую позицию скролла
            const oldScrollTop = scrollableContainer.value?.scrollTop || 0

            if (event.payload.offset === offsetEnd.value + limit.value) {
                // подгрузка сообщений сверху (когда скроллим наверх)
                messages.value.push(...event.payload.items)

                // после вставки сообщений, когда они отрисуются
                // устанавливаем старую позицию скролла
                nextTick(() => {
                    if (scrollableContainer.value) {
                        scrollableContainer.value.scrollTop = oldScrollTop
                    }
                })

                offsetEnd.value = event.payload.offset
            } else if (event.payload.offset === Math.max(0, offsetStart.value - limit.value)) {
                // сохраняем старую высоту контейнера с сообщениями
                const oldHeight = scrollableContainer.value?.scrollHeight || 0

                // подгрузка сообщений снизу (когда скроллим вниз)
                // (может быть при поиске сообщений)
                messages.value.unshift(...event.payload.items)

                // после вставки сообщений, когда они отрисуются
                // высчитываем позицию скролла с учетом нового контента
                nextTick(() => {
                    const newHeight = scrollableContainer.value?.scrollHeight || 0
                    const heightDiff = Math.abs(oldHeight - newHeight)

                    if (scrollableContainer.value) {
                        scrollableContainer.value.scrollTop = oldScrollTop - heightDiff
                    }
                })

                offsetStart.value = event.payload.offset
            }
        }

        count.value = event.payload.count

        if (event.payload.include) {
            offsetStart.value = offsetEnd.value = event.payload.offset

            try {
                nextTick(() => {
                    scrollToMessage(event.payload.include, event.payload.highlight)

                    setTimeout(() => (canFetch.value = true), 1000)
                })
            } catch {}
        }

        isFetchingMessages.value = false
    },

    'new_messages': (event: any) => {
        if (props.chat.id === event.payload.chatId) {
            messages.value.unshift(...event.payload.messages)
        }

        // store.messagesToForward = []
    },

    //   see_message(event: any) {
    //     if (props.chat.id !== event.chat_id) {
    //       return
    //     }

    //     // если просмотренное сообщение принадлежит этому чату

    //     // находим это сообщение
    //     const message = messages.value.find((message: any) => {
    //       return message.id === event.message_id
    //     })

    //     // если оно уже подгружено
    //     if (message) {
    //       // помечаем как просмотренное
    //       message.is_seen = true
    //     }
    //   },

    //   delete_messages(event: any) {
    //     if (props.chat.id !== event.chat_id) {
    //       return
    //     }

    //     for (const messageId of event.message_ids) {
    //       const index = messages.value.findIndex((message: any) => {
    //         return message.id === messageId
    //       })

    //       if (index !== -1) {
    //         messages.value.splice(index, 1)
    //       }
    //     }
    //   },
}

watch(
    () => props.chat,
    (newChat) => {
        isFetchingMessages.value = true
        messages.value = null

        store.client.sendEvent({
            type: 'open_chat',
            payload: {
                chatId: newChat.id,
            },
        })
    },
    {
        immediate: true,
    },
)

onMounted(() => {
    watch(
        () => store.client.data,
        (data) => {
            const message = JSON.parse(data)

            if (eventHandlers[message.type]) {
                eventHandlers[message.type](message)
            }
        },
    )
})

const onMessageView = (entries: IntersectionObserverEntry[], observer: IntersectionObserver) => {
    // console.log('ENTRIES', entries)

    for (const entry of entries) {
        if (!entry.isIntersecting) {
            continue
        }

        const messageElement = entry.target as HTMLElement
        observer.unobserve(messageElement)

        const messageId = Number(messageElement.dataset.messageId)
        const message = messages.value?.find((message: any) => message.id === messageId)

        if (!message.is_new || message.senderId === auth.currentUser?.id) {
            continue
        }

        store.seeMessage(messageId)
        message.is_new = false
    }
}

useIntersectionObserver(messageCards, onMessageView, {
    root: scrollableContainer,
})

useIntersectionObserver(
    topSpinner,
    ([entry]) => {
        if (!entry.isIntersecting || !canFetch.value) {
            return
        }

        store.fetchMessages(props.chat.id, {
            limit: limit.value,
            offset: offsetEnd.value + limit.value,
        })
    },
    {
        root: scrollableContainer,
    },
)

useIntersectionObserver(
    bottomSpinner,
    ([entry]) => {
        if (!entry.isIntersecting || !canFetch.value) {
            return
        }

        store.fetchMessages(props.chat.id, {
            limit: limit.value,
            offset: Math.max(0, offsetStart.value - limit.value),
        })
    },
    {
        root: scrollableContainer,
    },
)

const onReply = () => {
    messageToReply.value = contextMenuMessage.value
    messageForm.value?.focus()
}

const onSelect = () => {
    selectedMessages.value.push(contextMenuMessage.value)
    isSelecting.value = selectedMessages.value.length > 0
}

const onNotificationsToggle = () => {
    store.toggleNotifications(props.chat.id, !props.chat.areNotificationsEnabled)
}

const onMessagesDelete = (deleteForEveryone: boolean) => {
    const messages = messageToDelete.value ? [messageToDelete.value] : selectedMessages.value
    const messageIds = messages.map((message) => message.id)
    store.deleteMessages(props.chat.id, messageIds, deleteForEveryone)

    selectedMessages.value = []
    isSelecting.value = false
}

const onMessagesForward = (targetChat: any) => {
    store.messagesToForward = messageToForward.value
        ? [messageToForward.value]
        : selectedMessages.value

    selectedMessages.value = []
    isSelecting.value = false

    emit('forwardMessages', {
        targetChat,
    })
}

const onChatDelete = () => {
    store.deleteChat(props.chat.id)
}

const onGroupLeave = () => {
    store.leaveGroup(props.chat.id)
}

const onGroupJoin = () => {
    store.joinGroup(props.chat.id)
}

const onParentClick = (event: any) => {
    scrollToMessageOrFetch(event.message.id)
}

const onSearchResultSelect = (event: any) => {
    scrollToMessageOrFetch(event.message.id)
}

const onGoDown = () => {
    scrollToMessageOrFetch(props.chat.last_message.id, false)
}

const onContextMenu = (event: any, message: any) => {
    event.preventDefault()
    console.log('CONTEXT MENU')
    contextMenuMessage.value = message

    contextMenu.value.setProps({
        getReferenceClientRect: () => ({
            width: 0,
            height: 0,
            top: event.clientY,
            bottom: event.clientY,
            left: event.clientX,
            right: event.clientX,
        }),
    })

    contextMenu.value.show()
}
</script>

<template>
    <div class="relative flex h-full">
        <div class="relative bg-[#EEEAE3] grow">
            <div
                class="absolute top-0 left-0 w-full h-full bg-[url('@/assets/images/chat/bg.png')] bg-[length:670px] bg-[0_0] opacity-[.25] pointer-events-none"
            ></div>

            <div class="relative flex flex-col h-full">
                <div
                    id="messageSearchBar"
                    class="2md:hidden"
                    :class="{ hidden: !isSearching }"
                ></div>
                <div class="shrink-0" :class="{ 'hidden 2md:block': isSearching }">
                    <ChatTopBar
                        class="flex items-center gap-x-[16px] px-[12px] 2md:px-[20px] py-[6px]"
                    >
                        <div class="shrink-0">
                            <RouterLink v-if="chat.type === 'personal'" to="/chats">
                                <Avatar
                                    class="w-[44px] h-[44px] rounded-full"
                                    src="chat.photo ? $media(chat.photo.url) : undefined"
                                />
                            </RouterLink>
                            <button
                                v-else-if="chat.type === 'group'"
                                type="button"
                                @click="store.currentState = store.State.GROUP_DETAIL"
                            >
                                <Avatar
                                    class="w-[44px] h-[44px] rounded-full"
                                    src="chat.photo ? $media(chat.photo.url) : undefined"
                                />
                            </button>
                        </div>

                        <div class="grow flex flex-col items-start">
                            <div class="flex items-center gap-x-[4px]">
                                <RouterLink
                                    v-if="chat.type === 'personal'"
                                    class="whitespace-nowrap"
                                    to="$userLink(chat.other_user)"
                                >
                                    <span>
                                        {{ chat.title }}
                                    </span>
                                </RouterLink>
                                <button
                                    v-else
                                    class="whitespace-nowrap"
                                    type="button"
                                    @click="store.currentState = store.State.GROUP_DETAIL"
                                >
                                    <span>
                                        {{ chat.title }}
                                    </span>
                                </button>

                                <svg
                                    v-if="!chat.areNotificationsEnabled"
                                    class="text-primary-seattle-100"
                                    width="24"
                                    height="24"
                                >
                                    <use xlink:href="#muteIcon"></use>
                                </svg>
                            </div>

                            <button
                                v-if="chat.type === 'group'"
                                class="text-primary-seattle-100 text-[14px] leading-[18px]"
                                type="button"
                                @click="store.currentState = store.State.GROUP_DETAIL"
                            >
                                {{ chat.participantCount }} participants
                            </button>
                        </div>

                        <div class="shrink-0 flex items-center gap-x-[8px]">
                            <button
                                class="button l ghost only-icon hidden 2md:flex"
                                type="button"
                                @click="isSearching = true"
                            >
                                <i class="icon dd-Search text-[24px]"></i>
                            </button>

                            <Tippy
                                v-if="chat.isJoined"
                                theme="none"
                                placement="bottom-end"
                                trigger="click"
                                :arrow="false"
                                :duration="[0, 0]"
                                :offset="[0, 4]"
                                :zIndex="35"
                                interactive
                            >
                                <button
                                    class="button m 2md:l ghost only-icon"
                                    type="button"
                                    :disabled="isSelecting"
                                    click="$showPopup('topBarMenuPopup')"
                                >
                                    <i
                                        class="icon dd-Dots-menu rotate-90 text-[20px] 2md:text-[24px]"
                                    ></i>
                                </button>

                                <template #content="{ hide }">
                                    <div
                                        v-if="chat.type === 'personal'"
                                        class="options-menu hidden md:block min-w-[260px] max-h-max"
                                    >
                                        <ul>
                                            <li @click="hide">
                                                <RouterLink
                                                    class="option-button gap-x-[6px]"
                                                    to="/chats"
                                                >
                                                    <i class="icon dd-Single-user text-[20px]"></i>

                                                    <span> Профиль пользователя </span>
                                                </RouterLink>
                                            </li>
                                            <li @click="hide">
                                                <button
                                                    class="option-button gap-x-[6px]"
                                                    type="button"
                                                    @click="isSelecting = true"
                                                >
                                                    <i
                                                        class="icon dd-Checkmark-big-circle text-[20px]"
                                                    ></i>

                                                    <span> Выбрать сообщения </span>
                                                </button>
                                            </li>
                                            <li @click="hide">
                                                <button
                                                    class="option-button gap-x-[6px]"
                                                    type="button"
                                                    @click="onNotificationsToggle"
                                                >
                                                    <i
                                                        v-if="!chat.areNotificationsEnabled"
                                                        class="icon dd-Volume-Disabled text-[20px]"
                                                    ></i>

                                                    <span v-if="chat.areNotificationsEnabled">
                                                        Выключить уведомления
                                                    </span>
                                                    <span v-else> Включить уведомления </span>
                                                </button>
                                            </li>
                                            <li v-if="false" @click="hide">
                                                <button
                                                    class="option-button gap-x-[6px]"
                                                    type="button"
                                                >
                                                    <i class="icon dd-Lock text-[20px]"></i>

                                                    <span> Заблокировать пользователя </span>
                                                </button>
                                            </li>
                                            <li @click="hide">
                                                <button
                                                    class="option-button gap-x-[6px]"
                                                    type="button"
                                                    click="$showPopup('userSharePopup')"
                                                >
                                                    <i class="icon dd-Share-arrow text-[20px]"></i>

                                                    <span> Поделиться профилем </span>
                                                </button>
                                            </li>
                                            <li @click="hide">
                                                <button
                                                    class="option-button gap-x-[6px] text-utilitarian-moscow-100"
                                                    type="button"
                                                    click="$showPopup('confirmChatDeletionPopup')"
                                                >
                                                    <i
                                                        class="icon dd-Trash-Delete-Bin text-[20px]"
                                                    ></i>

                                                    <span> Удалить чат </span>
                                                </button>
                                            </li>
                                        </ul>
                                    </div>
                                    <div
                                        v-else-if="chat.type === 'group'"
                                        class="options-menu hidden md:block min-w-[260px] max-h-max"
                                    >
                                        <ul>
                                            <li
                                                v-if="chat.owner_id === auth.currentUser?.id"
                                                @click="hide"
                                            >
                                                <button
                                                    class="option-button gap-x-[6px]"
                                                    type="button"
                                                    @click="
                                                        store.currentState =
                                                            store.State.GROUP_EDIT_FORM
                                                    "
                                                >
                                                    <i
                                                        class="icon dd-Pen-Edit-Write text-[20px]"
                                                    ></i>

                                                    <span> Редактировать </span>
                                                </button>
                                            </li>
                                            <li @click="hide">
                                                <button
                                                    class="option-button gap-x-[6px]"
                                                    type="button"
                                                    @click="
                                                        store.currentState =
                                                            store.State.GROUP_DETAIL
                                                    "
                                                >
                                                    <i
                                                        class="icon dd-Info-information-circle text-[20px]"
                                                    ></i>

                                                    <span> Информация </span>
                                                </button>
                                            </li>
                                            <li @click="hide">
                                                <button
                                                    class="option-button gap-x-[6px]"
                                                    type="button"
                                                    @click="onNotificationsToggle"
                                                >
                                                    <i
                                                        v-if="!chat.areNotificationsEnabled"
                                                        class="icon dd-Volume-Disabled text-[20px]"
                                                    ></i>

                                                    <span v-if="chat.areNotificationsEnabled">
                                                        Выключить уведомления
                                                    </span>
                                                    <span v-else> Включить уведомления </span>
                                                </button>
                                            </li>
                                            <li @click="hide">
                                                <button
                                                    v-if="auth.currentUser?.id === chat.owner_id"
                                                    class="option-button gap-x-[6px] text-utilitarian-moscow-100"
                                                    type="button"
                                                    @click="onChatDelete"
                                                >
                                                    <i
                                                        class="icon dd-Trash-Delete-Bin text-[20px]"
                                                    ></i>

                                                    <span> Удалить и покинуть </span>
                                                </button>
                                                <button
                                                    v-else
                                                    class="option-button gap-x-[6px] text-utilitarian-moscow-100"
                                                    type="button"
                                                    @click="onGroupLeave"
                                                >
                                                    <i
                                                        class="icon dd-login-enter-arrow text-[20px]"
                                                    ></i>

                                                    <span> Покинуть группу </span>
                                                </button>
                                            </li>
                                        </ul>
                                    </div>
                                </template>
                            </Tippy>
                        </div>
                    </ChatTopBar>

                    <ChatTopBar
                        v-if="chat.type === 'group' && chat.join_request_count > 0"
                        class="px-[20px] py-[6px]"
                    >
                        <button
                            class="button-ghost l-medium link w-full gap-x-[12px] justify-center"
                            type="button"
                        >
                            <Avatar class="w-[28px] h-[28px] rounded-full" />

                            <span>
                                <!-- {{ $tc('membershipRequest', chat.join_request_count) }} на вступление -->
                                12 join requests
                            </span>
                        </button>
                    </ChatTopBar>
                </div>

                <div class="relative grow">
                    <div
                        ref="scrollableContainer"
                        class="absolute top-0 left-0 w-full h-full flex flex-col-reverse overflow-x-hidden overflow-y-auto scrollbar pb-[10px]"
                        @scroll="scrollTop = ($event.currentTarget as HTMLDivElement).scrollTop"
                    >
                        <div
                            v-if="isFetchingMessages"
                            class="flex justify-center items-center h-full"
                        >
                            <i
                                class="icon dd-Spiner-l text-[32px] text-primary-brand-accent animate-spin"
                            ></i>
                        </div>
                        <div v-else>
                            <div
                                v-if="count && offsetEnd + limit < count"
                                class="flex justify-center items-center py-[16px]"
                            >
                                <span ref="topSpinner">
                                    <i
                                        class="icon dd-Spiner-l text-[32px] text-primary-brand-accent animate-spin"
                                    ></i>
                                </span>
                            </div>

                            <div class="flex flex-col-reverse">
                                <div v-for="(messages, date) in groupedMessages" :key="date">
                                    <div
                                        class="sticky top-[10px] z-[11] flex justify-center my-[10px] pointer-events-none"
                                    >
                                        <span
                                            class="pointer-events-auto px-[8px] py-[2px] rounded-[6px] text-m-14 text-primary-seattle-120 bg-primary-brand-white shadow-[0px_1px_1px_0px_rgba(0,0,0,0.16)]"
                                        >
                                            {{ date }}
                                        </span>
                                    </div>

                                    <div class="flex flex-col-reverse">
                                        <MessageCard
                                            v-for="(message, index) in messages"
                                            :key="message.id"
                                            ref="messageCards"
                                            :class="{
                                                own: message.senderId === auth.currentUser?.id,
                                                selected:
                                                    contextMenuMessage &&
                                                    contextMenuMessage.id === message.id,
                                            }"
                                            :message="message"
                                            :selectable="isSelecting"
                                            :own="message.senderId === auth.currentUser?.id"
                                            :showUserInfo="chat.type === 'group'"
                                            :firstInGroup="
                                                store.isFirstInGroup(messages, message, index)
                                            "
                                            :lastInGroup="
                                                store.isLastInGroup(messages, message, index)
                                            "
                                            v-model="selectedMessages"
                                            :data-message-id="message.id"
                                            @contextmenu="onContextMenu($event, message)"
                                            @parentClick="onParentClick"
                                        />
                                    </div>
                                </div>
                            </div>

                            <div
                                v-if="offsetStart > 0"
                                class="flex justify-center items-center py-[16px]"
                            >
                                <span ref="bottomSpinner">
                                    <i
                                        class="icon dd-Spiner-l text-[32px] text-primary-brand-accent animate-spin"
                                    ></i>
                                </span>
                            </div>
                        </div>
                    </div>
                </div>

                <div
                    id="messageSearchNav"
                    class="2md:hidden"
                    :class="{ hidden: !isSearching }"
                ></div>
                <div
                    class="shrink-0 px-[6px] 2md:px-[22px] pb-[6px] 2md:pb-[22px]"
                    :class="{ 'hidden 2md:block': isSearching }"
                >
                    <div v-if="!chat.isJoined" class="max-w-[720px] mx-auto">
                        <button class="button xl primary w-full" type="button" @click="onGroupJoin">
                            <span> Вступить в группу </span>
                        </button>
                    </div>
                    <div
                        v-else-if="isSelecting"
                        class="flex justify-between items-center gap-x-[24px] max-w-[570px] mx-auto px-[12px] py-[14px] rounded-[12px] bg-primary-brand-white"
                    >
                        <button
                            class="button-ghost l-medium ghost-70"
                            type="button"
                            @click="
                                selectedMessages = [];
                                isSelecting = false
                            "
                        >
                            <i class="icon dd-Delete-Disabled-big text-[24px]"></i>
                        </button>

                        <span
                            class="grow text-l-long-16 font-medium text-primary-brand-onPrimary overflow-hidden text-ellipsis whitespace-nowrap"
                        >
                            <!-- {{ $tc('message', selectedMessages.length) }} -->
                            15 сообщений
                        </span>

                        <button
                            v-if="selectedMessages.length > 0"
                            class="button-ghost l-medium ghost-70 gap-x-[8px]"
                            type="button"
                            click="$showPopup('forwardPopup')"
                        >
                            <i class="icon dd-Share-arrow text-[24px]"></i>

                            <span> Переслать </span>
                        </button>

                        <button
                            v-if="selectedMessages.length > 0"
                            class="button-ghost l-medium primary gap-x-[8px]"
                            type="button"
                            click="$showPopup('deletionConfirmationPopup')"
                        >
                            <i class="icon dd-Trash-Delete-Bin text-[24px]"></i>

                            <span> Удалить </span>
                        </button>
                    </div>
                    <MessageForm
                        v-else
                        ref="messageForm"
                        class="mx-auto"
                        :chat="chat"
                        :messageToReply="messageToReply"
                        @cancelReplying="messageToReply = null"
                        @toggleEmojiPicker="isPanelActive = !isPanelActive"
                        @closeEmojiPicker="isPanelActive = false"
                    >
                        <template #actions>
                            <button
                                class="button xl white round only-icon translate-y-[calc(100%+12px)] pointer-events-auto"
                                :class="{
                                    'translate-y-0':
                                        Math.abs(scrollTop) > 500 ||
                                        (offsetEnd > 0 && Math.abs(scrollTop) > 500),
                                }"
                                type="button"
                                @click="onGoDown"
                            >
                                <ArrowDownIcon width="24" height="24" />
                            </button>
                        </template>
                    </MessageForm>
                </div>

                <div
                    id="emojiPanel"
                    class="md:hidden h-0 overflow-hidden transition-[height]"
                    :class="{
                        'h-[275px]': isPanelActive,
                    }"
                ></div>
            </div>
        </div>

        <div
            class="shrink-0 hidden 2md:block absolute xl:static top-0 right-0 z-[12] w-full h-full max-w-0 transition-[max-width] duration-[350ms] border-t border-l border-t-primary-london-110 border-l-primary-london-110"
            :class="{
                '!max-w-[350px]': isSearching,
            }"
        >
            <!-- <ChatMessageSearchResults
                :chat="chat"
                @messageSelect="onSearchResultSelect"
                @close="isSearching = false"
            /> -->
            Search Results
        </div>

        <Teleport to="#popupsContainer">
            <Tippy
                ref="contextMenu"
                theme="none"
                placement="right-start"
                trigger="manual"
                interactive
                :arrow="false"
                :offset="[0, 0]"
                @hide="contextMenuMessage = null"
            >
                <template #content="{ hide }">
                    <div class="options-menu min-w-[180px] max-h-max">
                        <ul>
                            <li @click="hide">
                                <button
                                    class="option-button gap-x-[6px]"
                                    type="button"
                                    @click="onReply"
                                >
                                    <ReplyIcon width="20" height="20" />

                                    <span> Ответить </span>
                                </button>
                            </li>
                            <li @click="hide">
                                <CopyButton
                                    class="option-button gap-x-[6px]"
                                    :text="contextMenuMessage ? contextMenuMessage.text : ''"
                                >
                                    <CopyIcon width="20" height="20" />

                                    <span> Копировать текст </span>
                                </CopyButton>
                            </li>
                            <li @click="hide">
                                <button
                                    class="option-button gap-x-[6px]"
                                    type="button"
                                    @pointerdown="messageToForward = contextMenuMessage"
                                    click="$showPopup('forwardPopup')"
                                >
                                    <ShareArrowIcon width="20" height="20" />

                                    <span> Переслать </span>
                                </button>
                            </li>
                            <li @click="hide">
                                <button
                                    class="option-button gap-x-[6px]"
                                    type="button"
                                    @click="onSelect"
                                >
                                    <CheckmarkBigCircleIcon width="20" height="20" />

                                    <span> Выбрать </span>
                                </button>
                            </li>
                            <li @click="hide">
                                <button
                                    class="option-button gap-x-[6px] text-utilitarian-moscow-100"
                                    type="button"
                                    @pointerdown="messageToDelete = contextMenuMessage"
                                    click="$showPopup('deletionConfirmationPopup')"
                                >
                                    <TrashDeleteBinIcon width="20" height="20" />

                                    <span> Удалить </span>
                                </button>
                            </li>
                        </ul>
                    </div>
                </template>
            </Tippy>

            <Popup id="topBarMenuPopup" class="md:hidden" contentClass="pt-[10px] px-[6px]">
                <div v-if="chat.type === 'dialog'" class="options-menu max-h-max p-0 shadow-none">
                    <ul>
                        <li>
                            <button
                                class="option-button gap-x-[6px]"
                                type="button"
                                data-popup-closer
                                @click="isSearching = true"
                            >
                                <i class="icon dd-Search text-[20px]"></i>

                                <span> Поиск </span>
                            </button>
                        </li>
                        <li>
                            <RouterLink class="option-button gap-x-[6px]" to="$userLink(chat)">
                                <i class="icon dd-Single-user text-[20px]"></i>

                                <span> Профиль пользователя </span>
                            </RouterLink>
                        </li>
                        <li>
                            <button
                                class="option-button gap-x-[6px]"
                                type="button"
                                data-popup-closer
                                @click="isSelecting = true"
                            >
                                <i class="icon dd-Checkmark-big-circle text-[20px]"></i>

                                <span> Выбрать сообщения </span>
                            </button>
                        </li>
                        <li>
                            <button
                                class="option-button gap-x-[6px]"
                                type="button"
                                data-popup-closer
                                @click="onNotificationsToggle"
                            >
                                <i
                                    v-if="!chat.areNotificationsEnabled"
                                    class="icon dd-Volume-Disabled text-[20px]"
                                ></i>

                                <span v-if="chat.areNotificationsEnabled">
                                    Выключить уведомления
                                </span>
                                <span v-else> Включить уведомления </span>
                            </button>
                        </li>
                        <li v-if="false">
                            <button
                                class="option-button gap-x-[6px]"
                                type="button"
                                data-popup-closer
                            >
                                <i class="icon dd-Lock text-[20px]"></i>

                                <span> Заблокировать пользователя </span>
                            </button>
                        </li>
                        <li>
                            <button
                                class="option-button gap-x-[6px]"
                                type="button"
                                data-popup-closer
                                click="$showPopup('userSharePopup')"
                            >
                                <i class="icon dd-Share-arrow text-[20px]"></i>

                                <span> Поделиться профилем </span>
                            </button>
                        </li>
                        <li>
                            <button
                                class="option-button gap-x-[6px] text-utilitarian-moscow-100"
                                type="button"
                                data-popup-closer
                                click="$showPopup('confirmChatDeletionPopup')"
                            >
                                <i class="icon dd-Trash-Delete-Bin text-[20px]"></i>

                                <span> Удалить чат </span>
                            </button>
                        </li>
                    </ul>
                </div>
                <div
                    v-else-if="chat.type === 'group'"
                    class="options-menu max-h-max p-0 shadow-none"
                >
                    <ul>
                        <li>
                            <button
                                class="option-button gap-x-[6px]"
                                type="button"
                                data-popup-closer
                                @click="isSearching = true"
                            >
                                <i class="icon dd-Search text-[20px]"></i>

                                <span> Поиск </span>
                            </button>
                        </li>
                        <li v-if="chat.owner_id === auth.currentUser?.id">
                            <button
                                class="option-button gap-x-[6px]"
                                type="button"
                                data-popup-closer
                                @click="store.currentState = store.State.GROUP_EDIT_FORM"
                            >
                                <i class="icon dd-Pen-Edit-Write text-[20px]"></i>

                                <span> Редактировать </span>
                            </button>
                        </li>
                        <li>
                            <button
                                class="option-button gap-x-[6px]"
                                type="button"
                                data-popup-closer
                                @click="store.currentState = store.State.GROUP_DETAIL"
                            >
                                <i class="icon dd-Info-information-circle text-[20px]"></i>

                                <span> Информация </span>
                            </button>
                        </li>
                        <li>
                            <button
                                class="option-button gap-x-[6px]"
                                type="button"
                                data-popup-closer
                                @click="onNotificationsToggle"
                            >
                                <i
                                    v-if="!chat.areNotificationsEnabled"
                                    class="icon dd-Volume-Disabled text-[20px]"
                                ></i>

                                <span v-if="chat.areNotificationsEnabled">
                                    Выключить уведомления
                                </span>
                                <span v-else> Включить уведомления </span>
                            </button>
                        </li>
                        <li>
                            <button
                                v-if="auth.currentUser?.id === chat.owner_id"
                                class="option-button gap-x-[6px] text-utilitarian-moscow-100"
                                type="button"
                                data-popup-closer
                                @click="onChatDelete"
                            >
                                <i class="icon dd-Trash-Delete-Bin text-[20px]"></i>

                                <span> Удалить и покинуть </span>
                            </button>
                            <button
                                v-else
                                class="option-button gap-x-[6px] text-utilitarian-moscow-100"
                                type="button"
                                data-popup-closer
                                @click="onGroupLeave"
                            >
                                <i class="icon dd-login-enter-arrow text-[20px]"></i>

                                <span> Покинуть группу </span>
                            </button>
                        </li>
                    </ul>
                </div>

                <div class="mt-[16px] px-[10px]">
                    <button class="button m clear round w-full" type="button" data-popup-closer>
                        <span> Закрыть </span>
                    </button>
                </div>
            </Popup>

            <Popup
                id="deletionConfirmationPopup"
                dialogClass="md:max-w-[330px]"
                contentClass="md:px-[16px] md:py-[20px]"
                @hidden="messageToDelete = null"
            >
                <ChatMessageDeletionForm
                    :messages="messageToDelete ? [messageToDelete] : selectedMessages"
                    @submit="onMessagesDelete"
                />
            </Popup>

            <Popup id="forwardPopup">
                <ChatMessageForwardForm @submit="onMessagesForward" />
            </Popup>

            <PopupConfirm
                id="confirmChatDeletionPopup"
                confirmText="Удалить"
                cancelText="Отмена"
                title="Удалить чат"
                description="Вы уверены, что хотите удалить чат?"
                @confirm="onChatDelete"
            />

            <PopupShare
                v-if="chat.type === 'personal'"
                id="userSharePopup"
                url="$userLink(chat.other_user)"
                :title="chat.other_user.display_name"
                :text="chat.other_user.display_name"
            />
        </Teleport>
    </div>
</template>
