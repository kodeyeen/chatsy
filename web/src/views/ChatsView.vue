<script lang="ts" setup>
import { ref, watch, onMounted } from 'vue'

import type { Chat } from '@/models/chat'

import { useChatsStore } from '@/stores/chats'

import ChatLayout from '@/layouts/Chat.vue'
import ChatsBar from '@/components/ChatsBar.vue'
import ChatContent from '@/components/ChatContent.vue'
import ChatDetail from '@/components/ChatDetail.vue'
import ChatEditForm from '@/components/ChatEditForm.vue'
import ChatJoinRequests from '@/components/ChatJoinRequests.vue'

// tmp
const chatsStore = useChatsStore()

const activeChat = ref<Chat | null>(null)
</script>

<template>
    <ChatLayout>
        <template #leftbar="{ toggleContent }">
            <ChatsBar
                v-model:activeChat="activeChat"
                @update:activeChat="toggleContent(true)"
            />
        </template>

        <template v-if="activeChat" #content>
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
