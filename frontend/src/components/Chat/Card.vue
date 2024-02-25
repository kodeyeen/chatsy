<script lang="ts" setup>
import { useAuthStore } from '@/stores/auth'

import Avatar from '@/components/Avatar.vue'
import Counter from '@/components/Counter.vue'

const props = defineProps<{
    chat: any
}>()

const auth = useAuthStore()
</script>

<template>
    <div
        class="chat-card flex gap-x-[9px] py-[8px] pl-[9px] pr-[12px] rounded-[10px] hover:bg-primary-brand-wildSand [&.active]:bg-primary-brand-accent cursor-pointer transition-colors"
        :class="{
            'items-center': chat.lastMessage,
        }"
    >
        <div class="shrink-0">
            <Avatar class="w-[56px] h-[56px] rounded-full" />
        </div>

        <div class="grow">
            <div class="flex justify-between items-center gap-x-[8px]">
                <span class="grow flex items-center gap-x-[4px]">
                    <span
                        class="text-l-long-16 font-bold text-primary-brand-onPrimary [.chat-card.active_&]:text-primary-brand-white line-clamp-1"
                    >
                        {{ chat.title }}
                    </span>

                    <svg
                        v-if="!chat.areNotificationsEnabled"
                        class="text-primary-seattle-100 [.chat-card.active_&]:text-primary-brand-white"
                        width="24"
                        height="24"
                    >
                        <use xlink:href="#muteIcon"></use>
                    </svg>
                </span>

                <span v-if="chat.lastMessage" class="shrink-0 flex items-center gap-x-[3px]">
                    <svg
                        v-if="chat.lastMessage.authorId === auth.currentUser?.id"
                        class="text-utilitarian-geneva-100 [.chat-card.active_&]:text-primary-brand-white"
                        width="24"
                        height="24"
                    >
                        <use v-if="chat.lastMessage.isViewed" xlink:href="#readIcon"></use>
                        <use v-else xlink:href="#checkIcon"></use>
                    </svg>

                    <time
                        v-if="chat.lastMessage"
                        class="text-s-12 text-primary-seattle-100 [.chat-card.active_&]:text-primary-brand-white"
                        :datetime="chat.lastMessage.sentAt"
                    >
                        {{
                            new Date(chat.lastMessage.sentAt).toLocaleTimeString('ru', {
                                hour: '2-digit',
                                minute: '2-digit',
                            })
                        }}
                    </time>
                </span>
            </div>

            <div v-if="chat.lastMessage" class="flex items-center gap-x-[8px]">
                <div class="grow flex items-center gap-x-[4px] mt-[2px] text-l-short-16">
                    <span
                        v-if="
                            chat.type === 'group' &&
                            chat.lastMessage.authorId !== auth.currentUser?.id
                        "
                        class="text-primary-brand-onPrimary [.chat-card.active_&]:text-primary-brand-white whitespace-nowrap"
                    >
                        {{ chat.lastMessage.authorFullName }}:
                    </span>

                    <span
                        v-if="false"
                        class="shrink-0 w-[24px] h-[24px] border border-primary-seattle-10 rounded-[4px] overflow-hidden"
                    >
                        <img
                            class="w-full h-full object-cover"
                            src="https://telegra.ph/file/75085f39c9315c6eeefb6.jpg"
                            alt="Attachment"
                        />
                    </span>

                    <span
                        class="text-primary-seattle-100 [.chat-card.active_&]:text-primary-brand-white line-clamp-1"
                    >
                        {{ chat.lastMessage.text }}
                    </span>
                </div>

                <div v-if="chat.unread_message_count > 0" class="shrink-0">
                    <Counter
                        class="accent [.chat-card.active_&]:hidden"
                        :value="chat.unread_message_count"
                    />
                    <Counter
                        class="neutral hidden [.chat-card.active_&]:block"
                        :value="chat.unread_message_count"
                    />
                </div>
            </div>
        </div>
    </div>
</template>
