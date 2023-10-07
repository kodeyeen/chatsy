<script setup>
import ChatLayout from '@/layouts/chat.vue'
import ChatCard from '@/components/Chat/Card.vue'
import ChatContent from '@/components/Chat/Content.vue'

import SpinnerIcon from '@/icons/Spinner.vue'
</script>

<template>
    <ChatLayout>
        <template #leftbar>
            <div class="flex flex-col h-full">
                <!-- <ChatSearchBar class="shrink-0" v-model:filter="title"/> -->

                <div class="relative grow bg-primary-brand-white">
                    <div ref="scrollableContainer" class="absolute top-0 left-0 w-full h-full overflow-x-hidden overflow-y-auto scrollbar p-[8px]">
                        <div class="h-full">
                            <div>
                                <ChatCard
                                    v-for="i in 10"
                                    :key="i"
                                    xchat="chat"
                                    xclass="{
                                        'active': activeChat && activeChat.id === chat.id,
                                    }"
                                    xclick="onChatEnter(chat, toggleContent)"
                                />
                            </div>

                            <div xv-if="chats && chats.has_more" class="flex justify-center items-center my-[16px]">
                                <span ref="spinner" class="text-primary-london-120 animate-spin">
                                    <SpinnerIcon/>
                                </span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </template>

        <template #content>
            <ChatContent
                xv-if="store.currentState === store.State.CHAT_MESSAGES"
                xchat="activeChat"
                xforwardMessages="onMessagesForward"
            />
            <!-- <ChatGroupDetail v-else-if="store.currentState === store.State.GROUP_DETAIL" :chat="activeChat"/> -->
            <!-- <ChatGroupEditForm v-else-if="store.currentState === store.State.GROUP_EDIT_FORM" :chat="activeChat"/> -->
            <!-- <ChatGroupRequests v-else-if="store.currentState === store.State.GROUP_REQUESTS"/> -->
        </template>
    </ChatLayout>
</template>
