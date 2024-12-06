export interface Message {
    id: number
    chatId: number
    senderId: number
    senderName: string
    authorName: string
    originalId: number
    parentId: number
    parentSenderName: string
    parentText: string
    text: string
    sentAt: string
    isViewed: boolean
}
