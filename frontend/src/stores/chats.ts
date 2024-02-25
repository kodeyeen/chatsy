// @ts-nocheck
import { ref, computed } from 'vue'
import { useWebSocket } from '@vueuse/core'
import { defineStore } from 'pinia'
import { useChatClient } from '@/composables/chatClient'

enum State {
    DEFAULT = 'default',
    CHAT_MESSAGES = 'chat_messages',
    GROUP_DETAIL = 'group_detail',
    GROUP_EDIT_FORM = 'group_edit_form',
}

export const useChatsStore = defineStore('chats', () => {
    // const client = ref(null)
    // сообщения, которые собираемся переслать
    const messagesToForward = ref<any[]>([])
    // просматриваемый чат
    // состояние, определяет, что показывается справа в области контента
    const currentState = ref<State>(null)

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

    const fetchChat = (chatId) => {
        return sendJson({
            type: 'fetch_chat',
            chat_id: chatId,
        })
    }

    const refreshChat = (chatId) => {
        return sendJson({
            type: 'refresh_chat',
            chat_id: chatId,
        })
    }

    const fetchMessages = (chatId, params = {}) => {
        return sendJson({
            type: 'fetch_messages',
            chat_id: chatId,
            is_searching: params.is_searching ?? false,
            include: params.include ?? null,
            highlight: params.highlight ?? true,
            text: params.text ?? '',
            limit: params.limit ?? 30,
            offset: params.offset ?? 0,
        })
    }

    const sendMessageForm = (
        chatId: number,
        newMessageData: any,
        messagesToForward: any[] = [],
    ) => {
        const messageIds = messagesToForward.map((message) => message.id)

        return sendJson({
            type: 'send_message_form',
            chat_id: chatId,
            new_message_data: newMessageData,
            message_ids_to_forward: messageIds,
        })
    }

    const seeMessage = (messageId) => {
        return sendJson({
            type: 'see_message',
            message_id: messageId,
        })
    }

    const createGroup = (data) => {
        return sendJson({
            type: 'create_group',
            ...data,
        })
    }

    const updateGroup = (chatId, newData) => {
        return sendJson({
            type: 'update_group',
            chat_id: chatId,
            new_data: newData,
        })
    }

    const toggleNotifications = (chatId, enable) => {
        return sendJson({
            type: 'toggle_notifications',
            chat_id: chatId,
            are_notifications_enabled: enable,
        })
    }

    const deleteMessages = (chatId, messageIds, deleteForEveryone) => {
        return sendJson({
            type: 'delete_messages',
            chat_id: chatId,
            message_ids: messageIds,
            delete_for_everyone: deleteForEveryone,
        })
    }

    const openChat = (chatId) => {
        return sendJson({
            type: 'open_chat',
            chat_id: chatId,
        })
    }

    const fetchOrCreateChat = (userId) => {
        return sendJson({
            type: 'fetch_or_create_chat',
            user_id: userId,
        })
    }

    const forwardMessages = (targetChatId, messages, newMessageData = null) => {
        const messageIds = messages.map((message) => message.id)

        return sendJson({
            type: 'forward_messages',
            target_chat_id: targetChatId,
            message_ids: messageIds,
            new_message_data: newMessageData,
        })
    }

    const deleteChat = (chatId) => {
        return sendJson({
            type: 'delete_chat',
            chat_id: chatId,
        })
    }

    const leaveGroup = (chatId) => {
        return sendJson({
            type: 'leave_group',
            chat_id: chatId,
        })
    }

    const joinGroup = (chatId) => {
        return sendJson({
            type: 'join_group',
            chat_id: chatId,
        })
    }

    const addMembers = (chatId, userIds) => {
        return sendJson({
            type: 'add_members',
            chat_id: chatId,
            user_ids: userIds,
        })
    }

    const fetchMemberships = (chatId, params = {}) => {
        return sendJson({
            type: 'fetch_memberships',
            chat_id: chatId,
            limit: params.limit ?? 3,
            offset: params.offset ?? 0,
            is_staff: params.is_staff ?? false,
        })
    }

    return {
        client,
        init,
        currentState,
        messagesToForward,

        isFirstInGroup,
        isLastInGroup,
        State,

        fetchChat,
        fetchMessages,
        sendMessageForm,
        seeMessage,
        createGroup,
        updateGroup,
        toggleNotifications,
        deleteMessages,
        openChat,
        fetchOrCreateChat,
        forwardMessages,
        deleteChat,
        leaveGroup,
        joinGroup,
        addMembers,
        fetchMemberships,
        refreshChat,
    }
})
