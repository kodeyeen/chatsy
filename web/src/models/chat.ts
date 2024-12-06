import type { Message } from './message'

export interface Chat {
    id: number
    type: string
    title: string
    description?: string
    inviteHash?: string
    joinByLinkCount: number
    isPrivate: boolean
    joinByRequest: boolean
    isJoined: boolean
    memberCount: number
    areNotificationsEnabled: number
    lastMessage?: Message
}
