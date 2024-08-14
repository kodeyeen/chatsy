<script lang="ts" setup>
import { useAuthStore } from '@/stores/auth'

import Avatar from '@/components/Avatar.vue'
import Counter from '@/components/Counter.vue'
import CheckIcon from '@/components/icons/Check.vue'
import MuteIcon from '@/components/icons/Mute.vue'
import ReadIcon from '@/components/icons/Read.vue'

const props = defineProps<{
    chat: any
    active?: boolean
}>()

const auth = useAuthStore()
</script>

<template>
    <div
        class="chat-card flex gap-x-[9px] py-[8px] pl-[9px] pr-[12px] rounded-[10px] hover:bg-primary-brand-wildSand cursor-pointer transition-colors"
        :class="{
            'items-center': chat.lastMessage,
            '!bg-primary-brand-accent': active,
        }"
    >
        <div class="shrink-0">
            <Avatar class="w-[56px] h-[56px] rounded-full" />
        </div>

        <div class="grow">
            <div class="flex justify-between items-center gap-x-[8px]">
                <span class="grow flex items-center gap-x-[4px]">
                    <span
                        class="text-l-long-16 font-bold text-primary-brand-onPrimary line-clamp-1"
                        :class="{ '!text-primary-brand-white': active }"
                    >
                        {{ chat.title }}
                    </span>

                    <MuteIcon
                        v-if="!chat.areNotificationsEnabled"
                        class="text-primary-seattle-100"
                        :class="{ '!text-primary-brand-white': active }"
                        :width="24"
                        :height="24"
                    />
                </span>

                <span v-if="chat.lastMessage" class="shrink-0 flex items-center gap-x-[3px]">
                    <ReadIcon
                        v-if="
                            chat.lastMessage.senderId === auth.currentUser?.id &&
                            chat.lastMessage.isViewed
                        "
                        class="text-utilitarian-geneva-100"
                        :class="{ '!text-primary-brand-white': active }"
                        :width="24"
                        :height="24"
                    />
                    <CheckIcon
                        v-else-if="chat.lastMessage.senderId === auth.currentUser?.id"
                        class="text-utilitarian-geneva-100"
                        :class="{ '!text-primary-brand-white': active }"
                        :width="24"
                        :height="24"
                    />

                    <time
                        v-if="chat.lastMessage"
                        class="text-s-12 text-primary-seattle-100"
                        :class="{ '!text-primary-brand-white': active }"
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
                            chat.lastMessage.senderId !== auth.currentUser?.id
                        "
                        class="text-primary-brand-onPrimary whitespace-nowrap"
                        :class="{ '!text-primary-brand-white': active }"
                    >
                        {{ chat.lastMessage.senderName }}:
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
                        class="text-primary-seattle-100 line-clamp-1"
                        :class="{ '!text-primary-brand-white': active }"
                    >
                        {{ chat.lastMessage.text }}
                    </span>
                </div>

                <div v-if="chat.unread_message_count > 0" class="shrink-0">
                    <Counter
                        v-if="active"
                        :value="chat.unread_message_count"
                        size="default"
                        variant="neutral"
                    />
                    <Counter
                        v-else
                        :value="chat.unread_message_count"
                        size="default"
                        variant="accent"
                    />
                </div>
            </div>
        </div>
    </div>
</template>
