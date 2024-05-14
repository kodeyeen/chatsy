<script lang="ts" setup>
import { ref, watch, onMounted } from 'vue'
import { useRouter } from 'vue-router'

import { useChatsStore } from '@/stores/chats'
import ChatCard from '@/components/Chat/Card.vue'
import ChatSearchBar from '@/components/Chat/SearchBar.vue'
import ChatContent from '@/components/Chat/Content.vue'

const isContentActive = ref(false)

const toggleContent = (force: boolean | null = null) => {
    isContentActive.value = force ?? !isContentActive.value
}

// import { useIntersectionObserver } from '@vueuse/core'

// const { currentUser } = useAuth()
// const route = useRoute()
const router = useRouter()
const chatsStore = useChatsStore()

const chats = ref<any>(null)
const activeChat = ref<any>(null)
const title = ref('')

// const filteredChats = computed(() => {
//   if (!chats.value) {
//     return []
//   }

//   return chats.value.results.filter((chat) => {
//     return chat.title.toLowerCase().includes(title.value.toLowerCase())
//   })
// })

// const sortedChats = computed(() => {
//   return filteredChats.value.sort((chat1, chat2) => {
//     const sentAt1 = chat1.last_message ? new Date(chat1.last_message.sent_at) : 0
//     const sentAt2 = chat2.last_message ? new Date(chat2.last_message.sent_at) : 0

//     return sentAt2 - sentAt1
//   })
// })

const eventHandlers: any = {
    connected: (event: any) => {
        chats.value = event.payload.chats
    },

    //   fetch_or_create_chat(event) {
    //     if (!event.is_created) {
    //       const index = chats.value.results.findIndex(chat => chat.id === event.chat.id)
    //       if (index !== -1) {
    //         chats.value.results.splice(index, 1)
    //       }
    //     }
    //     chats.value.results.unshift(event.chat)
    //     activeChat.value = event.chat
    //     store.currentState = store.State.CHAT_MESSAGES
    //   },
    //   send_message_form(event) {
    //     // удаляем чат из списка
    //     const chatIndex = chats.value.results.findIndex(chat => chat.id === event.chat.id)
    //     if (chatIndex !== -1) {
    //       chats.value.results.splice(chatIndex, 1)
    //     }
    //     // и добавляем обновленную версию
    //     chats.value.results.push(event.chat)
    //     if (event.messages[0].senderId !== currentUser.value.id) {
    //       currentUser.value.unseen_message_count += event.messages.length
    //     }
    //   },
    //   see_message(event) {
    //     // берем чат, которому принадлежит просмотренное сообщение
    //     const chat = chats.value.results.find(chat => chat.id === event.chat_id)
    //     if (!chat) {
    //       return
    //     }
    //     // если последнее сообщение этого чата === просмотренному сообщению
    //     if (chat.last_message.id === event.message_id) {
    //       // помечаем последнее сообщение в объекте чата как просмотренное, чтобы обновилась иконка в сайдбаре
    //       chat.last_message.is_seen = true
    //     }
    //     if (currentUser.value.id === event.seen_by_id) {
    //       chat.unread_message_count -= 1
    //       currentUser.value.unseen_message_count -= 1
    //     }
    //   },
    //   toggle_notifications(event) {
    //     const chat = chats.value.results.find(chat => chat.id === event.chat_id)
    //     if (chat) {
    //       chat.areNotificationsEnabled = event.areNotificationsEnabled
    //     }
    //     if (activeChat.value && activeChat.value.id === event.chat_id) {
    //       activeChat.value.areNotificationsEnabled = event.areNotificationsEnabled
    //     }
    //   },
    //   create_group(event) {
    //     chats.value.results.unshift(event.chat)
    //     if (event.owner_id === currentUser.value.id) {
    //       activeChat.value = event.chat
    //       store.currentState = store.State.GROUP_DETAIL
    //     }
    //   },
    //   update_group(event) {
    //     if (activeChat.value && activeChat.value.id === event.chat.id) {
    //       Object.assign(activeChat.value, event.chat)
    //     }
    //     const chat = chats.value.results.find(chat => chat.id === event.chat.id)
    //     if (chat) {
    //       Object.assign(chat, event.chat)
    //     }
    //     store.currentState = store.State.GROUP_DETAIL
    //   },
    //   delete_messages(event) {
    //     const chat = chats.value.results.find((chat) => {
    //       if (!chat.last_message) {
    //         return false
    //       }
    //       return event.message_ids.includes(chat.last_message.id)
    //     })
    //     if (chat) {
    //       store.refreshChat(chat.id)
    //     }
    //   },
    //   refresh_chat(event) {
    //     const chatIndex = chats.value.results.findIndex(chat => chat.id === event.chat.id)
    //     if (chatIndex !== -1) {
    //       chats.value.results.splice(chatIndex, 1)
    //     }
    //     chats.value.results.push(event.chat)
    //   },
    //   delete_chat(event) {
    //     const chatIndex = chats.value.results.findIndex(chat => chat.id === event.chat_id)
    //     if (chatIndex !== -1) {
    //       chats.value.results.splice(chatIndex, 1)
    //     }
    //     if (activeChat.value && activeChat.value.id === event.chat_id) {
    //       store.currentState = null
    //       activeChat.value = null
    //     }
    //   },
    //   join_group(event) {
    //     if (event.new_membership.user.id === currentUser.value.id) {
    //       activeChat.value.is_joined = true
    //     }
    //     activeChat.value.member_count += 1
    //   },
    //   leave_group(event) {
    //     if (event.user_id === currentUser.value.id) {
    //       const chatIndex = chats.value.results.findIndex(chat => chat.id === event.chat_id)
    //       if (chatIndex !== -1) {
    //         chats.value.results.splice(chatIndex, 1)
    //         store.currentState = null
    //         activeChat.value = null
    //       }
    //       return
    //     }
    //     if (activeChat.value && activeChat.value.id === event.chat_id) {
    //       activeChat.value.member_count -= 1
    //     }
    //   },
    //   add_members(event) {
    //     if (activeChat.value && activeChat.value.id === event.chat_id) {
    //       activeChat.value.member_count += event.memberships.length
    //     }
    //   },
    //   added_to_group(event) {
    //     chats.value.results.unshift(event.chat)
    //   },
}

onMounted(() => {
    watch(
        () => chatsStore.client.data,
        (data) => {
            const message = JSON.parse(data)

            if (eventHandlers[message.type]) {
                eventHandlers[message.type](message)
            }
        },
    )

    // chatsStore.client.getChats(10, 0)

    //   if (route.hash) {
    //     const chatId = Number(route.hash.slice(1))
    //     store.openChat(chatId)
    //   } else if (route.query.user_id) {
    //     const userId = Number(route.query.user_id)
    //     store.fetchOrCreateChat(userId)
    //   }
})

// const scrollableContainer = ref(null)
// const spinner = ref(null)

// useIntersectionObserver(spinner, ([entry]) => {
//   if (!entry.isIntersecting) {
//     return
//   }

//   chats.value.offset += chats.value.limit

//   store.fetchChats({
//     limit: chats.value.limit,
//     offset: chats.value.offset,
//   })
// }, {
//   root: scrollableContainer,
// })

// const onMessagesForward = (event) => {
//   activeChat.value = event.targetChat
// }

const onChatOpen = (chat: any) => {
    toggleContent(true)

    if (activeChat.value && activeChat.value.id === chat.id) {
        return
    }

    router.push({
        hash: `#${chat.id}`,
        query: {},
    })

    activeChat.value = chat
    chatsStore.currentState = chatsStore.State.CHAT_MESSAGES
}
</script>

<template>
    <div class="flex flex-col h-full">
        <header>HEADER</header>

        <main class="grow flex overflow-hidden">
            <div
                class="shrink-0 w-full 2md:w-[430px] transition-transform duration-[500ms]"
                :class="{
                    '-translate-x-full 2md:translate-x-[-430px] 2lg:translate-x-0': isContentActive,
                }"
            >
                <div class="flex flex-col h-full">
                    <ChatSearchBar class="shrink-0" v-model:filter="title" />

                    <div class="relative grow bg-primary-brand-white">
                        <div
                            ref="scrollableContainer"
                            class="absolute top-0 left-0 w-full h-full overflow-x-hidden overflow-y-auto scrollbar p-[8px]"
                        >
                            <div class="h-full">
                                <div>
                                    <div v-if="chats">
                                        <ChatCard
                                            v-for="chat in chats.items"
                                            :key="chat.id"
                                            :chat="chat"
                                            :class="{
                                                active: activeChat && activeChat.id === chat.id,
                                            }"
                                            @click="onChatOpen(chat)"
                                        />
                                    </div>

                                    <!-- <div v-if="chats && chats.has_more" class="flex justify-center items-center my-[16px]">
                                        <i ref="spinner" class="icon dd-Spiner-l text-[32px] text-primary-brand-accent animate-spin"></i>
                                    </div> -->
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div
                class="grow shrink-0 w-full 2lg:w-auto transition-transform duration-[500ms]"
                :class="{
                    '-translate-x-full 2md:translate-x-[-430px] 2lg:translate-x-0': isContentActive,
                }"
                @click="isContentActive = true"
            >
                <ChatContent
                    v-if="chatsStore.currentState === chatsStore.State.CHAT_MESSAGES"
                    :chat="activeChat"
                    forwardMessages="onMessagesForward"
                />
                <!-- 
                <ChatGroupDetail v-else-if="store.currentState === store.State.GROUP_DETAIL" :chat="activeChat"/>
                <ChatGroupEditForm v-else-if="store.currentState === store.State.GROUP_EDIT_FORM" :chat="activeChat"/>
                <ChatGroupRequests v-else-if="store.currentState === store.State.GROUP_REQUESTS"/> -->
            </div>
        </main>
    </div>
</template>
