import { useWebSocket } from '@vueuse/core'
import type { MaybeRefOrGetter } from 'vue'
import type { UseWebSocketOptions } from '@vueuse/core'

interface ChatEvent {
    type: string
    payload: any
}

interface Chat {
    id: number
    type: string
    title: string
    description: string
    inviteHash: string
    joinByLinkCount: number
    isPrivate: boolean
    joinByRequest: boolean
}

interface Page<T> {
    items: T[]
    count: number
    limit: number
    offset: number
}

export const useChatClient = (
    url: MaybeRefOrGetter<string | URL | undefined>,
    options?: UseWebSocketOptions | undefined,
) => {
    const { data, close, send, open, ws } = useWebSocket(url, options)

    const sendEvent = (event: ChatEvent): boolean => {
        return send(JSON.stringify(event))
    }

    // const sendMessageForm = (
    //     chatId: number,
    //     newMessageData: any,
    //     messagesToForward: any[] = [],
    // ) => {
    //     const messageIds = messagesToForward.map((message) => message.id)

    //     return sendJson({
    //         type: 'send_message_form',
    //         chat_id: chatId,
    //         new_message_data: newMessageData,
    //         message_ids_to_forward: messageIds,
    //     })
    // }

    // const createMessages = (newMessage: any, messagesToForward: any[] = []): boolean => {
    //     return sendEvent({
    //         type: 'create_messages',
    //         payload: {
    //             newMessage,
    //         },
    //     })
    // }

    return {
        data,

        close,
        sendEvent,
        open,
    }
}
