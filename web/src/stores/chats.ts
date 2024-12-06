// @ts-nocheck
import { ref, computed } from 'vue'
import { useWebSocket } from '@vueuse/core'
import { defineStore } from 'pinia'
import { useChatClient } from '@/composables/chatClient'

enum State {
    DEFAULT = 'default',
    CHAT_MESSAGES = 'chat_messages',
    CHAT_DETAIL = 'chat_detail',
    CHAT_EDIT_FORM = 'chat_edit_form',
    CHAT_JOIN_REQUESTS = 'chat_join_requests',
}

export const useChatsStore = defineStore('chats', () => {
    // const client = ref(null)
    // сообщения, которые собираемся переслать
    const messagesToForward = ref<any[]>([])
    // просматриваемый чат
    // состояние, определяет, что показывается справа в области контента
    const currentState = ref<State | null>(null)

    const ticket = ref(null)
    const wsURL = computed(() => {
        const url = new URL('ws://localhost:8080/ws')

        if (ticket.value) {
            url.searchParams.set('ticket', ticket.value)
        }

        return url
    })

    const client = useChatClient(wsURL, {
        // autoReconnect: true,
        immediate: false,
        onConnected: (ws) => {
            console.log('CONNECTED')
        },
        //     onMessage: (ws, event) => {
        //         const message = JSON.parse(event.data)

        //         console.log('WS MESSAGE', message)
        //     },
    })

    const init = (t) => {
        ticket.value = t
        client.open()
    }

    const isFirstInGroup = (messages, message, index) => {
        return index === messages.length - 1 || messages[index + 1].senderId !== message.senderId
    }

    const isLastInGroup = (messages, message, index) => {
        return index === 0 || messages[index - 1].senderId !== message.senderId
    }

    const sendJson = (content) => {
        return client.send(JSON.stringify(content))
    }

    return {
        client,
        init,
        currentState,
        messagesToForward,

        isFirstInGroup,
        isLastInGroup,
        State,
    }
})
