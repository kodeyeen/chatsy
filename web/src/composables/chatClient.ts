import { useWebSocket } from '@vueuse/core'
import type { MaybeRefOrGetter } from 'vue'
import type { UseWebSocketOptions } from '@vueuse/core'

export interface ChatEvent {
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

    return {
        data,

        close,
        sendEvent,
        open,
    }
}
