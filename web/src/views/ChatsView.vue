<script lang="ts" setup>
import { ref, watch, onMounted } from 'vue'
import { useRouter } from 'vue-router'

import { useChatsStore } from '@/stores/chats'
import ChatLayout from '@/components/layouts/Chat.vue'
import ChatCard from '@/components/ChatCard.vue'
import ChatSearchBar from '@/components/ChatSearchBar.vue'
import ChatContent from '@/components/ChatContent.vue'
import ChatDetail from '@/components/ChatDetail.vue'
import ChatEditForm from '@/components/ChatEditForm.vue'
import ChatJoinRequests from '@/components/ChatJoinRequests.vue'
import SpinnerIcon from '@/components/icons/Spinner.vue'

// tmp
const chatsStore = useChatsStore()
const title = ref<string>('')
const chats = ref<any[]>([
    {
        id: 1,
        title: 'First Chat',
        description: 'descr',
        type: 'group',
        areNotificationsEnabled: false,
        participantCount: 3,
        isJoined: true,
        owner_id: 1,
        join_request_count: 12,
        lastMessage: {
            text: "i'm the last",
            sentAt: new Date().toString(),
        },
    },
    {
        id: 2,
        title: 'Second Chat',
        description: 'Kek',
        type: 'group',
        areNotificationsEnabled: false,
        participantCount: 5,
        isJoined: true,
        owner_id: 1,
        join_request_count: 0,
        lastMessage: {
            text: "i'm the second",
            sentAt: new Date().toString(),
        },
    },
])
const activeChat = ref<any>(chats.value[0])
</script>

<template>
    <ChatLayout>
        <template #leftbar="{ toggleContent }">
            <div class="flex flex-col h-full">
                <ChatSearchBar class="shrink-0" v-model:filter="title" />

                <div class="relative grow bg-primary-brand-white">
                    <div
                        ref="scrollableContainer"
                        class="absolute top-0 left-0 w-full h-full overflow-x-hidden overflow-y-auto scrollbar p-[8px]"
                    >
                        <div class="h-full">
                            <div>
                                <div v-if="chats.length > 0">
                                    <ChatCard
                                        v-for="chat in chats"
                                        :key="chat.id"
                                        :chat="chat"
                                        :active="activeChat && activeChat.id === chat.id"
                                        click="onChatOpen(chat)"
                                    />
                                </div>

                                <div v-if="true" class="flex justify-center items-center my-[16px]">
                                    <SpinnerIcon
                                        ref="spinner"
                                        class="text-primary-brand-accent animate-spin"
                                        :width="32"
                                        :height="32"
                                    />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </template>

        <template #content>
            <ChatContent
                v-if="chatsStore.currentState === chatsStore.State.CHAT_MESSAGES"
                :chat="activeChat"
                forwardMessages="onMessagesForward"
            />
            <ChatDetail
                v-else-if="chatsStore.currentState === chatsStore.State.CHAT_DETAIL"
                :chat="activeChat"
            />
            <ChatEditForm
                v-else-if="chatsStore.currentState === chatsStore.State.CHAT_EDIT_FORM"
                :chat="activeChat"
            />
            <ChatJoinRequests
                v-else-if="chatsStore.currentState === chatsStore.State.CHAT_JOIN_REQUESTS"
            />
        </template>
    </ChatLayout>
</template>
