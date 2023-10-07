<script setup>
import { Tippy } from 'vue-tippy'

import MessageParent from '@/components/Message/Parent.vue'
import MentionMenu from '@/components/MentionMenu.vue'
import SendIcon from '@/icons/Send.vue'
import CornerIcon from '@/icons/Corner.vue'
</script>

<template>
    <form
        ref="root"
        class="relative flex items-end max-w-[720px]"
        action="#"
        method="POST"
        xsubmit.prevent="submit()"
    >
        <Tippy
            xv-if="chat.type === 'group'"
            ref="mentionTooltip"
            placement="top-start"
            trigger="manual"
            interactive
            :arrow="false"
            :offset="[0, 4]"
        >
            <template #content="{ hide, state }">
                <MentionMenu
                    xchat="chat"
                    xfilter="usernameFilter"
                    xactive="state.isShown"
                    xselect="onUserMention"
                    xnoresults="hide()"
                />
            </template>
        </Tippy>


        <div class="relative grow flex items-end">
            <div class="grow px-[12px] py-[8px] bg-primary-brand-white rounded-[12px] rounded-br-none">
                <div xv-if="messageToReply" class="flex items-end mb-[12px]">
                    <MessageParent class="grow" xmessage="messageToReply"/>

                    <button class="shrink-0 button m ghost only-icon p-[5px]" type="button" xclick="$emit('cancelReplying')">
                        <i class="icon dd-Delete-Disabled-small text-[24px]"></i>
                    </button>
                </div>
                <!-- <div v-else-if="store.messagesToForward.length !== 0">
                {{ $tc('forwardedMessage', store.messagesToForward.length) }}
                </div> -->

                <div class="flex items-end gap-x-[12px]">
                    <!-- <div class="shrink-0">
                        <Tippy
                            theme="none"
                            placement="top-start"
                            :arrow="false"
                            :duration="[0, 0]"
                            :offset="[-12, 12]"
                            :zIndex="35"
                            trigger="mouseenter click"
                            interactive
                        >
                            <button id="emojiPickerTrigger" class="button l ghost p-[5px]" type="button" @click="onEmojiPickerToggle">
                                <svg width="24" height="24">
                                    <use xlink:href="#emojiIcon"></use>
                                </svg>
                            </button>

                            <template #content>
                                <EmojiPicker
                                    class="hidden md:flex"
                                    locale="ru"
                                    previewPosition="none"
                                    searchPosition="none"
                                    theme="light"
                                    xemojiSelect="onEmojiSelect"
                                />
                            </template>
                        </Tippy>
                    </div> -->

                    <div class="grow">
                        <textarea
                            ref="textarea"
                            class="block w-full max-h-[284px] py-[8px] text-l-short-16 text-primary-brand-onPrimary placeholder:text-primary-seattle-60 outline-none resize-none"
                            rows="1"
                            placeholder="Напишите сообщение"
                            xv-model="data.text"
                            xv-autoresize
                            xv-focus
                            xkeyup="onKeyUp"
                            xkeydown="onKeyDown"
                        ></textarea>
                    </div>

                    <div v-if="false" class="shrink-0">
                        <label class="button l ghost relative p-[5px]">
                            <input
                                class="absolute opacity-0 pointer-events-none"
                                type="file"
                                multiple
                                xchange="onFileSelect"
                            >

                            <svg width="24" height="24">
                                <use xlink:href="#attachIcon"></use>
                            </svg>
                        </label>
                    </div>
                </div>
            </div>

            <CornerIcon class="shrink-0 text-primary-brand-white"/>
        </div>

        <div class="relative shrink-0">
            <div class="absolute bottom-[calc(100%+12px)] pointer-events-none">
                <slot name="actions"></slot>
            </div>

            <button v-if="true" xv-if="data.text" class="button xl accent round only-icon relative" type="submit">
                <SendIcon/>
            </button>
            <button v-else class="button xl accent round only-icon relative" type="button">
                <svg class="shrink-0" width="24" height="24">
                    <use xlink:href="#cameraIcon"></use>
                </svg>
            </button>
        </div>

        <!-- <Teleport to="#popupsContainer">
            <Popup id="filesPopup" contentClass="p-[16px] pt-[8px]">
                <form action="#" method="POST">
                <div class="flex items-center gap-x-[16px]">
                    <div class="shrink-0">
                    <button class="text-primary-seattle-60" type="button" data-popup-closer>
                        <i class="icon dd-Delete-Disabled-big text-[24px]"></i>
                    </button>
                    </div>

                    <div class="grow">
                    <span class="text-primary-brand-onPrimary text-l-long-16 font-bold">
                        Отправить 2 фото
                    </span>
                    </div>

                    <div class="shrink-0">
                    <button class="button m primary px-[28px]" type="button">
                        <span>
                        Отправить
                        </span>
                    </button>
                    </div>
                </div>

                <div class="mt-[16px]">
                    <ChatMessageAttachment
                    v-for="file in data.attachments"
                    :file="file"
                    />
                </div>
                </form>
            </Popup>
        </Teleport>

        <Teleport to="#emojiPanel">
            <EmojiPicker
                class="w-full h-full"
                locale="ru"
                previewPosition="none"
                searchPosition="none"
                theme="light"
                dynamicWidth
                @emojiSelect="onEmojiSelect"
                @clickOutside="onPickerClickOutside"
            />
        </Teleport> -->
    </form>
</template>