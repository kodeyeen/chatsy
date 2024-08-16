<script lang="ts" setup>
import { ref, computed, watch, nextTick, onMounted } from 'vue'
import { useIntersectionObserver } from '@vueuse/core'
import { Tippy } from 'vue-tippy'

import type { Chat } from '@/models/chat'

import { useAuthStore } from '@/stores/auth'
import { useChatsStore } from '@/stores/chats'
import { usePopupsStore } from '@/stores/popups'

import Avatar from '@/components/Avatar.vue'
import CopyText from '@/components/CopyText.vue'
import Button from '@/components/Button.vue'
import GhostButton from '@/components/GhostButton.vue'
import OptionButton from '@/components/OptionButton.vue'
import OptionsMenu from '@/components/OptionsMenu.vue'
import OptionItem from '@/components/OptionItem.vue'
import ChatTopBar from '@/components/ChatTopBar.vue'
import ArrowDownIcon from '@/components/icons/ArrowDown.vue'
import CopyIcon from '@/components/icons/Copy.vue'
import CheckmarkBigCircleIcon from '@/components/icons/CheckmarkBigCircle.vue'
import DeleteDisabledBigIcon from '@/components/icons/DeleteDisabledBig.vue'
import TrashDeleteBinIcon from '@/components/icons/TrashDeleteBin.vue'
import ReplyIcon from '@/components/icons/Reply.vue'
import ShareArrowIcon from '@/components/icons/ShareArrow.vue'
import LockIcon from '@/components/icons/Lock.vue'
import MuteIcon from '@/components/icons/Mute.vue'
import DotsMenuIcon from '@/components/icons/DotsMenu.vue'
import SpinnerIcon from '@/components/icons/Spinner.vue'
import SearchIcon from '@/components/icons/Search.vue'
import SingleUserIcon from '@/components/icons/SingleUser.vue'
import LoginEnterArrowIcon from '@/components/icons/LoginEnterArrow.vue'
import PenEditWriteIcon from '@/components/icons/PenEditWrite.vue'
import InfoInformationCircleIcon from '@/components/icons/InfoInformationCircle.vue'
import VolumeDisabledIcon from '@/components/icons/VolumeDisabled.vue'
import MessageBubble from '@/components/MessageBubble.vue'
import MessageSendForm from '@/components/MessageSendForm.vue'
import MessageDeleteForm from '@/components/MessageDeleteForm.vue'
import Popup from '@/components/Popup.vue'
import PopupConfirm from '@/components/PopupConfirm.vue'
import PopupShare from '@/components/PopupShare.vue'

const props = defineProps<{
    chat: Chat,
}>()

const emit = defineEmits(['forwardMessages'])

const auth = useAuthStore()
const store = useChatsStore()
const popupsStore = usePopupsStore()

const now = new Date()

const messages = ref<any | null>([
    {
        id: 1,
        authorName: 'test',
        senderName: 'test123',
        senderId: 1,
        text: 'msg text',
        sentAt: now.toString(),
        isViewed: false,
        // parentId:
        // parentSenderName:
        // parentText:
    },
    {
        id: 2,
        authorName: 'kek',
        senderName: 'kek123',
        senderId: 2,
        text: 'kek text',
        sentAt: new Date(2024, 7, 12).toString(),
        isViewed: false,
    },
    {
        id: 3,
        authorName: 'September',
        senderName: 'kek123',
        senderId: 3,
        text: 'September',
        sentAt: new Date(2024, 7, 12).toString(),
        isViewed: false,
    },
    {
        id: 4,
        authorName: 'September',
        senderName: 'kek123',
        senderId: 3,
        text: 'September',
        sentAt: new Date(2024, 7, 12).toString(),
        isViewed: false,
    },
    {
        id: 5,
        authorName: 'September',
        senderName: 'kek123',
        senderId: 3,
        text: 'September',
        sentAt: new Date(2024, 7, 12).toString(),
        isViewed: false,
    },
    {
        id: 6,
        authorName: 'September',
        senderName: 'kek123',
        senderId: 3,
        text: 'September',
        sentAt: new Date(2024, 7, 12).toString(),
        isViewed: false,
    },
    {
        id: 7,
        authorName: 'September',
        senderName: 'kek123',
        senderId: 3,
        text: 'September',
        sentAt: new Date(2024, 7, 12).toString(),
        isViewed: false,
    },
    {
        id: 8,
        authorName: 'September',
        senderName: 'kek123',
        senderId: 3,
        text: 'loremfsda lf jdsaof jsdauig hdafigu jdsoif jdsaui gfmadspof jasdi fundsaijfkdslajf lkasdj f',
        sentAt: new Date(2024, 7, 12).toString(),
        isViewed: false,
    },
    {
        id: 9,
        authorName: 'September',
        senderName: 'kek123',
        senderId: 3,
        text: 'September',
        sentAt: new Date(2024, 7, 12).toString(),
        isViewed: false,
    },
    {
        id: 10,
        authorName: 'September',
        senderName: 'kek123',
        senderId: 3,
        text: 'loremfsda lf jdsaof jsdauig hdafigu jdsoif jdsaui gfmadspof jasdi fundsaijfkdslajf lkasdj f',
        sentAt: new Date(2024, 7, 12).toString(),
        isViewed: false,
    },
    {
        id: 11,
        authorName: 'September',
        senderName: 'kek123',
        senderId: 3,
        text: 'loremfsda lf jdsaof jsdauig hdafigu jdsoif jdsaui gfmadspof jasdi fundsaijfkdslajf lkasdj f',
        sentAt: new Date(2024, 7, 12).toString(),
        isViewed: false,
    },
])
const limit = ref(7)
const offsetStart = ref(0)
const offsetEnd = ref(0)
const count = ref<number | null>(null)
// при загрузке сообщений, а затем скроллинге до результата поиска
// (см. ветку if (event.include) в fetch_messages)
// можем случайно триггернуть загрузку новых сообщений сверху или снизу
// этот флаг позволит предотвращать подгрузку, пока не дойдем до результата поиска
const canFetch = ref(true)
const isFetchingMessages = ref(false)

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
const MessageBubbles = ref([])
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

const onSelect = () => {
    selectedMessages.value.push(contextMenuMessage.value)
    isSelecting.value = selectedMessages.value.length > 0
}

const onContextMenu = (event: any, message: any) => {
    event.preventDefault()

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
        <div v-if="chat" class="relative grow flex flex-col h-full">
            <div
                id="messageSearchBar"
                class="2md:hidden"
                :class="{ hidden: !isSearching }"
            ></div>
            <div class="shrink-0" :class="{ 'hidden 2md:block': isSearching }">
                <ChatTopBar
                    class="flex items-center gap-x-[16px] px-[12px] 2md:px-[20px] pt-[6px] pb-[5px]"
                >
                    <div class="shrink-0">
                        <RouterLink v-if="chat.type === 'dialog'" class="block" to="/chats">
                            <Avatar class="w-[44px] h-[44px] rounded-full" />
                        </RouterLink>
                        <button
                            v-else-if="chat.type === 'group'"
                            class="block"
                            type="button"
                            @click="store.currentState = store.State.CHAT_DETAIL"
                        >
                            <Avatar class="w-[44px] h-[44px] rounded-full" />
                        </button>
                    </div>

                    <div class="grow flex flex-col items-start">
                        <div class="flex items-center gap-x-[4px]">
                            <RouterLink
                                v-if="chat.type === 'dialog'"
                                class="whitespace-nowrap"
                                to="/chats"
                            >
                                <span>
                                    {{ chat.title }}
                                </span>
                            </RouterLink>
                            <button
                                v-else
                                class="whitespace-nowrap"
                                type="button"
                                @click="store.currentState = store.State.CHAT_DETAIL"
                            >
                                <span>
                                    {{ chat.title }}
                                </span>
                            </button>

                            <MuteIcon
                                v-if="!chat.areNotificationsEnabled"
                                class="text-primary-seattle-100"
                                :width="24"
                                :height="24"
                            />
                        </div>

                        <button
                            v-if="chat.type === 'group'"
                            class="text-primary-seattle-100 text-[14px] leading-[18px]"
                            type="button"
                            @click="store.currentState = store.State.CHAT_DETAIL"
                        >
                            <!-- {{ chat.memberCount }} participants -->
                            12 members
                        </button>
                    </div>

                    <div class="shrink-0 flex items-center gap-x-[8px]">
                        <Button
                            class="hidden 2md:flex w-[44px] h-[44px] !p-0"
                            type="button"
                            size="l"
                            variant="ghost"
                            @click="isSearching = true"
                        >
                            <SearchIcon :width="24" :height="24" />
                        </Button>

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
                            <Button
                                class="2md:hidden w-[36px] h-[36px] !p-0"
                                type="button"
                                size="m"
                                variant="ghost"
                                :disabled="isSelecting"
                                @click="popupsStore.show('topBarMenuPopup')"
                            >
                                <DotsMenuIcon
                                    class="rotate-90 w-[20px] h-[20px] 2md:w-[24px] 2md:h-[24px]"
                                />
                            </Button>
                            <Button
                                class="hidden 2md:flex w-[44px] h-[44px] !p-0"
                                type="button"
                                size="l"
                                variant="ghost"
                                :disabled="isSelecting"
                                @click="popupsStore.show('topBarMenuPopup')"
                            >
                                <DotsMenuIcon
                                    class="rotate-90 w-[20px] h-[20px] 2md:w-[24px] 2md:h-[24px]"
                                />
                            </Button>

                            <template #content="{ hide }">
                                <OptionsMenu
                                    v-if="chat.type === 'dialog'"
                                    class="hidden md:block min-w-[260px] max-h-max"
                                >
                                    <OptionItem @click="hide()">
                                        <OptionButton
                                            class="gap-x-[6px]"
                                            text="User profile"
                                            type="link"
                                            to="/chats"
                                        >
                                            <template #start>
                                                <SingleUserIcon :width="20" :height="20" />
                                            </template>
                                        </OptionButton>
                                    </OptionItem>

                                    <OptionItem @click="hide()">
                                        <OptionButton
                                            class="gap-x-[6px]"
                                            text="Select messages"
                                            type="button"
                                            @click="isSelecting = true"
                                        >
                                            <template #start>
                                                <CheckmarkBigCircleIcon
                                                    :width="20"
                                                    :height="20"
                                                />
                                            </template>
                                        </OptionButton>
                                    </OptionItem>

                                    <OptionItem @click="hide()">
                                        <OptionButton
                                            class="gap-x-[6px]"
                                            :text="
                                                chat.areNotificationsEnabled
                                                    ? 'Mute notifications'
                                                    : 'Unmute notifications'
                                            "
                                            type="button"
                                            click="onNotificationsToggle"
                                        >
                                            <template #start>
                                                <VolumeDisabledIcon
                                                    v-if="!chat.areNotificationsEnabled"
                                                    :width="20"
                                                    :height="20"
                                                />
                                            </template>
                                        </OptionButton>
                                    </OptionItem>

                                    <OptionItem v-if="false" @click="hide()">
                                        <OptionButton
                                            class="gap-x-[6px]"
                                            text="Block user"
                                            type="button"
                                        >
                                            <template #start>
                                                <LockIcon :width="20" :height="20" />
                                            </template>
                                        </OptionButton>
                                    </OptionItem>

                                    <OptionItem @click="hide()">
                                        <OptionButton
                                            class="gap-x-[6px]"
                                            text="Share profile"
                                            type="button"
                                            @click="popupsStore.show('userSharePopup')"
                                        >
                                            <template #start>
                                                <ShareArrowIcon :width="20" :height="20" />
                                            </template>
                                        </OptionButton>
                                    </OptionItem>

                                    <OptionItem @click="hide()">
                                        <OptionButton
                                            class="gap-x-[6px] text-utilitarian-moscow-100"
                                            text="Delete chat"
                                            type="button"
                                            @click="
                                                popupsStore.show('confirmChatDeletionPopup')
                                            "
                                        >
                                            <template #start>
                                                <TrashDeleteBinIcon :width="20" :height="20" />
                                            </template>
                                        </OptionButton>
                                    </OptionItem>
                                </OptionsMenu>
                                <OptionsMenu
                                    v-else-if="chat.type === 'group'"
                                    class="hidden md:block min-w-[260px] max-h-max"
                                >
                                    <OptionItem
                                        xv-if="chat.owner_id === auth.currentUser?.id"
                                        @click="hide()"
                                    >
                                        <OptionButton
                                            class="gap-x-[6px]"
                                            text="Edit"
                                            type="button"
                                            @click="
                                                store.currentState = store.State.CHAT_EDIT_FORM
                                            "
                                        >
                                            <template #start>
                                                <PenEditWriteIcon :width="20" :height="20" />
                                            </template>
                                        </OptionButton>
                                    </OptionItem>

                                    <OptionItem @click="hide()">
                                        <OptionButton
                                            class="gap-x-[6px]"
                                            text="Details"
                                            type="button"
                                            @click="
                                                store.currentState = store.State.CHAT_DETAIL
                                            "
                                        >
                                            <template #start>
                                                <InfoInformationCircleIcon
                                                    :width="20"
                                                    :height="20"
                                                />
                                            </template>
                                        </OptionButton>
                                    </OptionItem>

                                    <OptionItem @click="hide()">
                                        <OptionButton
                                            class="gap-x-[6px]"
                                            :text="
                                                chat.areNotificationsEnabled
                                                    ? 'Mute notifications'
                                                    : 'Unmute notifications'
                                            "
                                            type="button"
                                            click="onNotificationsToggle"
                                        >
                                            <template #start>
                                                <VolumeDisabledIcon
                                                    v-if="!chat.areNotificationsEnabled"
                                                    :width="20"
                                                    :height="20"
                                                />
                                            </template>
                                        </OptionButton>
                                    </OptionItem>

                                    <OptionItem @click="hide()">
                                        <OptionButton
                                            xv-if="auth.currentUser?.id === chat.owner_id"
                                            class="gap-x-[6px] text-utilitarian-moscow-100"
                                            text="Delete and leave"
                                            type="button"
                                            click="onChatDelete"
                                        >
                                            <template #start>
                                                <TrashDeleteBinIcon :width="20" :height="20" />
                                            </template>
                                        </OptionButton>
                                        <!-- <OptionButton
                                            v-else
                                            class="gap-x-[6px] text-utilitarian-moscow-100"
                                            text="Leave group"
                                            type="button"
                                            click="onGroupLeave"
                                        >
                                            <template #start>
                                                <LoginEnterArrowIcon :width="20" :height="20" />
                                            </template>
                                        </OptionButton> -->
                                    </OptionItem>
                                </OptionsMenu>
                            </template>
                        </Tippy>
                    </div>
                </ChatTopBar>

                <ChatTopBar
                    xv-if="chat.type === 'group' && chat.join_request_count > 0"
                    class="px-[20px] pt-[6px] pb-[5px]"
                >
                    <GhostButton
                        class="w-full gap-x-[12px] justify-center"
                        type="button"
                        size="l-medium"
                        variant="link"
                        @click="store.currentState = store.State.CHAT_JOIN_REQUESTS"
                    >
                        <Avatar class="w-[28px] h-[28px] rounded-full" />

                        <span> 12 join requests </span>
                    </GhostButton>
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
                        <SpinnerIcon
                            class="text-primary-brand-accent animate-spin"
                            :width="32"
                            :height="32"
                        />
                    </div>
                    <div v-else>
                        <div
                            v-if="count && offsetEnd + limit < count"
                            class="flex justify-center items-center py-[16px]"
                        >
                            <span ref="topSpinner">
                                <SpinnerIcon
                                    class="text-primary-brand-accent animate-spin"
                                    :width="32"
                                    :height="32"
                                />
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
                                    <MessageBubble
                                        v-for="(message, index) in messages"
                                        :key="message.id"
                                        :class="{
                                            own: message.senderId === auth.currentUser?.id,
                                            selected: contextMenuMessage?.id === message.id,
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
                                        parentClick="onParentClick"
                                    />
                                </div>
                            </div>
                        </div>

                        <div
                            v-if="offsetStart > 0"
                            class="flex justify-center items-center py-[16px]"
                        >
                            <span ref="bottomSpinner">
                                <SpinnerIcon
                                    class="text-primary-brand-accent animate-spin"
                                    :width="32"
                                    :height="32"
                                />
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
                    <Button
                        class="w-full"
                        type="button"
                        size="xl"
                        variant="primary"
                        click="onGroupJoin"
                    >
                        <span> Join the group </span>
                    </Button>
                </div>
                <div
                    v-else-if="isSelecting"
                    class="flex justify-between items-center gap-x-[24px] max-w-[570px] mx-auto px-[12px] py-[14px] rounded-[12px] bg-primary-brand-white"
                >
                    <GhostButton
                        type="button"
                        size="l-medium"
                        variant="ghost-70"
                        @click="
                            selectedMessages = [];
                            isSelecting = false
                        "
                    >
                        <DeleteDisabledBigIcon :width="24" :height="24" />
                    </GhostButton>

                    <span
                        class="grow text-l-long-16 font-medium text-primary-brand-onPrimary overflow-hidden text-ellipsis whitespace-nowrap"
                    >
                        <!-- {{ $tc('message', selectedMessages.length) }} -->
                        15 messages
                    </span>

                    <GhostButton
                        v-if="selectedMessages.length > 0"
                        type="button"
                        size="l-medium"
                        variant="ghost-70"
                        @click="popupsStore.show('forwardPopup')"
                    >
                        <ShareArrowIcon :width="24" :height="24" />

                        <span> Forward </span>
                    </GhostButton>

                    <GhostButton
                        v-if="selectedMessages.length > 0"
                        type="button"
                        size="l-medium"
                        variant="primary"
                        @click="popupsStore.show('deletionConfirmationPopup')"
                    >
                        <TrashDeleteBinIcon :width="24" :height="24" />

                        <span> Delete </span>
                    </GhostButton>
                </div>
                <MessageSendForm
                    v-else
                    ref="messageForm"
                    class="mx-auto"
                    :chat="chat"
                    :parent="messageToReply"
                    @cancelReplying="messageToReply = null"
                    @toggleEmojiPicker="isPanelActive = !isPanelActive"
                    @closeEmojiPicker="isPanelActive = false"
                >
                    <template #actions>
                        <Button
                            class="translate-y-[calc(100%+12px)] pointer-events-auto w-[52px] px-[14px] rounded-full"
                            :class="{
                                'translate-y-0':
                                    Math.abs(scrollTop) > 500 ||
                                    (offsetEnd > 0 && Math.abs(scrollTop) > 500),
                            }"
                            type="button"
                            size="xl"
                            variant="white"
                            click="onGoDown"
                        >
                            <ArrowDownIcon :width="24" :height="24" />
                        </Button>
                    </template>
                </MessageSendForm>
            </div>

            <div
                id="emojiPanel"
                class="md:hidden h-0 overflow-hidden transition-[height]"
                :class="{
                    'h-[275px]': isPanelActive,
                }"
            ></div>
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
                    <OptionsMenu class="min-w-[180px] max-h-max">
                        <OptionItem @click="hide()">
                            <OptionButton
                                class="gap-x-[6px]"
                                text="Reply"
                                type="button"
                                click="onReply"
                            >
                                <template #start>
                                    <ReplyIcon :width="20" :height="20" />
                                </template>
                            </OptionButton>
                        </OptionItem>

                        <OptionItem @click="hide()">
                            <CopyText
                                v-slot="{ copy }"
                                :source="contextMenuMessage ? contextMenuMessage.text : ''"
                            >
                                <OptionButton
                                    class="gap-x-[6px]"
                                    text="Copy Text"
                                    type="button"
                                    @click="copy()"
                                >
                                    <template #start>
                                        <CopyIcon :width="20" :height="20" />
                                    </template>
                                </OptionButton>
                            </CopyText>
                        </OptionItem>

                        <OptionItem @click="hide()">
                            <OptionButton
                                class="gap-x-[6px]"
                                text="Forward"
                                type="button"
                                @pointerdown="messageToForward = contextMenuMessage"
                                @click="popupsStore.show('forwardPopup')"
                            >
                                <template #start>
                                    <ShareArrowIcon :width="20" :height="20" />
                                </template>
                            </OptionButton>
                        </OptionItem>

                        <OptionItem @click="hide()">
                            <OptionButton
                                class="gap-x-[6px]"
                                text="Select"
                                type="button"
                                @click="onSelect"
                            >
                                <template #start>
                                    <CheckmarkBigCircleIcon :width="20" :height="20" />
                                </template>
                            </OptionButton>
                        </OptionItem>

                        <OptionItem @click="hide()">
                            <OptionButton
                                class="gap-x-[6px] !text-utilitarian-moscow-100"
                                text="Delete"
                                type="button"
                                @pointerdown="messageToDelete = contextMenuMessage"
                                @click="popupsStore.show('deletionConfirmationPopup')"
                            >
                                <template #start>
                                    <TrashDeleteBinIcon :width="20" :height="20" />
                                </template>
                            </OptionButton>
                        </OptionItem>
                    </OptionsMenu>
                </template>
            </Tippy>

            <Popup v-if="chat" id="topBarMenuPopup" class="md:hidden" contentClass="pt-[10px] px-[6px]">
                <OptionsMenu v-if="chat.type === 'dialog'" class="max-h-max p-0 shadow-none">
                    <OptionItem>
                        <OptionButton
                            class="gap-x-[6px]"
                            text="Search"
                            type="button"
                            @click="
                                popupsStore.hide('topBarMenuPopup');
                                isSearching = true
                            "
                        >
                            <template #start>
                                <SearchIcon :width="20" :height="20" />
                            </template>
                        </OptionButton>
                    </OptionItem>

                    <OptionItem>
                        <OptionButton
                            class="gap-x-[6px]"
                            text="User profile"
                            type="link"
                            to="/chats"
                        >
                            <template #start>
                                <SingleUserIcon :width="20" :height="20" />
                            </template>
                        </OptionButton>
                    </OptionItem>

                    <OptionItem>
                        <OptionButton
                            class="gap-x-[6px]"
                            text="Select messages"
                            type="button"
                            @click="
                                popupsStore.hide('topBarMenuPopup');
                                isSelecting = true
                            "
                        >
                            <template #start>
                                <CheckmarkBigCircleIcon :width="20" :height="20" />
                            </template>
                        </OptionButton>
                    </OptionItem>

                    <OptionItem>
                        <OptionButton
                            class="gap-x-[6px]"
                            :text="
                                chat.areNotificationsEnabled
                                    ? 'Mute notifications'
                                    : 'Unmute notifications'
                            "
                            type="button"
                            click="onNotificationsToggle"
                            @click="popupsStore.hide('topBarMenuPopup')"
                        >
                            <template #start>
                                <VolumeDisabledIcon
                                    v-if="!chat.areNotificationsEnabled"
                                    :width="20"
                                    :height="20"
                                />
                            </template>
                        </OptionButton>
                    </OptionItem>

                    <OptionItem v-if="false">
                        <OptionButton
                            class="gap-x-[6px]"
                            text="Block user"
                            type="button"
                            @click="popupsStore.hide('topBarMenuPopup')"
                        >
                            <template #start>
                                <LockIcon :width="20" :height="20" />
                            </template>
                        </OptionButton>
                    </OptionItem>

                    <OptionItem>
                        <OptionButton
                            class="gap-x-[6px]"
                            text="Share profile"
                            type="button"
                            @click="
                                popupsStore.hide('topBarMenuPopup');
                                popupsStore.show('userSharePopup')
                            "
                        >
                            <template #start>
                                <ShareArrowIcon :width="20" :height="20" />
                            </template>
                        </OptionButton>
                    </OptionItem>

                    <OptionItem>
                        <OptionButton
                            class="gap-x-[6px] text-utilitarian-moscow-100"
                            text="Delete chat"
                            type="button"
                            @click="
                                popupsStore.hide('topBarMenuPopup');
                                popupsStore.show('confirmChatDeletionPopup')
                            "
                        >
                            <template #start>
                                <TrashDeleteBinIcon :width="20" :height="20" />
                            </template>
                        </OptionButton>
                    </OptionItem>
                </OptionsMenu>
                <OptionsMenu v-else-if="chat.type === 'group'" class="max-h-max p-0 shadow-none">
                    <OptionItem>
                        <OptionButton
                            class="gap-x-[6px]"
                            text="Search"
                            type="button"
                            @click="
                                popupsStore.hide('topBarMenuPopup');
                                isSearching = true
                            "
                        >
                            <template #start>
                                <SearchIcon :width="20" :height="20" />
                            </template>
                        </OptionButton>
                    </OptionItem>

                    <OptionItem xv-if="chat.owner_id === auth.currentUser?.id">
                        <OptionButton
                            class="gap-x-[6px]"
                            text="Edit"
                            type="button"
                            @click="
                                popupsStore.hide('topBarMenuPopup');
                                store.currentState = store.State.CHAT_EDIT_FORM
                            "
                        >
                            <template #start>
                                <PenEditWriteIcon :width="20" :height="20" />
                            </template>
                        </OptionButton>
                    </OptionItem>

                    <OptionItem>
                        <OptionButton
                            class="gap-x-[6px]"
                            text="Details"
                            type="button"
                            @click="
                                popupsStore.hide('topBarMenuPopup');
                                store.currentState = store.State.CHAT_DETAIL
                            "
                        >
                            <template #start>
                                <InfoInformationCircleIcon :width="20" :height="20" />
                            </template>
                        </OptionButton>
                    </OptionItem>

                    <OptionItem>
                        <OptionButton
                            class="gap-x-[6px]"
                            :text="
                                chat.areNotificationsEnabled
                                    ? 'Mute Notifications'
                                    : 'Unmute Notifications'
                            "
                            type="button"
                            click="onNotificationsToggle"
                            @click="popupsStore.hide('topBarMenuPopup')"
                        >
                            <template #start>
                                <VolumeDisabledIcon
                                    v-if="!chat.areNotificationsEnabled"
                                    :width="20"
                                    :height="20"
                                />
                            </template>
                        </OptionButton>
                    </OptionItem>

                    <OptionItem>
                        <OptionButton
                            xv-if="auth.currentUser?.id === chat.owner_id"
                            class="gap-x-[6px] text-utilitarian-moscow-100"
                            text="Delete and leave"
                            type="button"
                            click="onChatDelete"
                            @click="popupsStore.hide('topBarMenuPopup')"
                        >
                            <template #start>
                                <TrashDeleteBinIcon :width="20" :height="20" />
                            </template>
                        </OptionButton>
                        <!-- <OptionButton
                            v-else
                            class="gap-x-[6px] text-utilitarian-moscow-100"
                            text="Leave group"
                            type="button"
                            @click="popupsStore.hide('topBarMenuPopup')"
                            click="onGroupLeave"
                        >
                            <template #start>
                                <LoginEnterArrowIcon :width="20" :height="20" />
                            </template>
                        </OptionButton> -->
                    </OptionItem>
                </OptionsMenu>

                <div class="mt-[16px] px-[10px]">
                    <Button
                        class="round w-full"
                        type="button"
                        size="m"
                        variant="clear"
                        @click="popupsStore.hide('topBarMenuPopup')"
                    >
                        <span> Close </span>
                    </Button>
                </div>
            </Popup>

            <Popup
                id="deletionConfirmationPopup"
                dialogClass="md:max-w-[330px]"
                contentClass="md:px-[16px] md:py-[20px]"
                @hidden="messageToDelete = null"
            >
                <MessageDeleteForm
                    :messages="messageToDelete ? [messageToDelete] : selectedMessages"
                    submit="onMessagesDelete"
                    @cancel="popupsStore.hide('deletionConfirmationPopup')"
                />
            </Popup>

            <Popup id="forwardPopup">
                <!-- <ChatMessageForwardForm @submit="onMessagesForward" /> -->
            </Popup>

            <PopupConfirm
                id="confirmChatDeletionPopup"
                confirmText="Delete"
                cancelText="Cancel"
                title="Delete Chat"
                description="Are you sure you want to delete this chat?"
                confirm="onChatDelete"
            />

            <PopupShare
                v-if="chat && chat.type === 'dialog'"
                id="userSharePopup"
                url="/"
                title="Full Name"
                text="Text"
            />
        </Teleport>
    </div>
</template>
