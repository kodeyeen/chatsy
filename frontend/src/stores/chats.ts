// @ts-nocheck
import { ref } from 'vue'
import { useWebSocket } from '@vueuse/core'
import { defineStore } from 'pinia'

export const useChatsStore = defineStore('chats', () => {
    const State = {
        CHAT_MESSAGES: 'chat_messages',
        GROUP_DETAIL: 'group_detail',
        GROUP_EDIT_FORM: 'group_edit_form',
    }

    const client = ref(null)
    // сообщения, которые собираемся переслать
    const messagesToForward = ref([])
    // просматриваемый чат
    // состояние, определяет, что показывается справа в области контента
    const currentState = ref(null)
    const back = ref('/')

    const isFirstInGroup = (messages, message, index) => {
        return index === messages.length - 1 || messages[index + 1].author.id !== message.author.id
    }

    const isLastInGroup = (messages, message, index) => {
        return index === 0 || messages[index - 1].author.id !== message.author.id
    }

    const init = (ticket) => {
        // const url = new URL(runtimeConfig.chatsSocketBaseUrl)
        const url = new URL('ws://localhost:8080/ws')
        url.searchParams.set('ticketID', ticket.id)

        client.value = useWebSocket(url, {
            // autoReconnect: true,

            onMessage: (ws, event) => {
                const message = JSON.parse(event.data)

                console.log('WS MESSAGE', message)
            },
        })
        console.log('INITED', client.value)
    }

    const sendJson = (content) => {
        return client.value.send(JSON.stringify(content))
    }

    const fetchChats = (params = {}) => {
        return sendJson({
            type: 'fetch_chats',
            payload: {
                // is_forward_form: params.is_forward_form ?? false,
                limit: params.limit ?? 10,
                offset: params.offset ?? 0,
            },
        })
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

    const sendMessageForm = (chatId, newMessageData, messagesToForward = []) => {
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
        back,
        messagesToForward,

        isFirstInGroup,
        isLastInGroup,
        State,

        fetchChats,
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
