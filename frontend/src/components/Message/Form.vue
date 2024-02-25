<script lang="ts" setup>
import { reactive, ref, computed, watch } from 'vue'
import { Tippy } from 'vue-tippy'

import { adjustTextarea } from '@/services/form'
import { useChatsStore } from '@/stores/chats'

import EmojiPicker from '@/components/EmojiPicker.vue'
import MessageParent from '@/components/Message/Parent.vue'
import MentionMenu from '@/components/MentionMenu.vue'
import AttachIcon from '@/components/icons/Attach.vue'
import CameraIcon from '@/components/icons/Camera.vue'
import CornerIcon from '@/components/icons/Corner.vue'
import DeleteDisabledSmallIcon from '@/components/icons/DeleteDisabledSmall.vue'
import EmojiIcon from '@/components/icons/Emoji.vue'
import SendIcon from '@/components/icons/Send.vue'

const props = defineProps<{
    chat: any
    parent: any
}>()

const emit = defineEmits(['cancelReplying', 'toggleEmojiPicker', 'closeEmojiPicker'])

const store = useChatsStore()

const getInitialData = () => ({
    parentId: null,
    text: '',
    attachments: [],
})

const data = reactive(getInitialData())
const usernameFilter = ref(null)
const root = ref<HTMLFormElement | null>(null)
const textarea = ref<HTMLTextAreaElement | null>(null)
const mentionTooltip = ref<any | null>(null)

const showMentionTooltip = () => {
    mentionTooltip.value?.setProps({
        getReferenceClientRect: () => root.value?.getBoundingClientRect(),
    })

    mentionTooltip.value.show()
}

watch(
    () => props.parent,
    (newParent) => {
        data.parentId = newParent ? newParent.id : null
    },
)

watch(
    () => props.chat,
    () => {
        focus()
    },
)

const isComplete = computed(() => {
    return data.text.trim() !== ''
})

const focus = () => {
    textarea.value?.focus({ preventScroll: true })
}

const onEmojiSelect = (emoji: any) => {
    const field = textarea.value

    if (!field) {
        return
    }

    field.setRangeText(emoji.native, field.selectionStart, field.selectionEnd, 'end')
    // триггерим событие input, чтобы обновилось data.text (v-model)
    field.dispatchEvent(new Event('input', { bubbles: true }))
    field.focus()
}

const onEmojiPickerToggle = () => {
    emit('toggleEmojiPicker')
}

const onPickerClickOutside = (event: any) => {
    const target = event.target

    if (!target.closest('#emojiPickerTrigger')) {
        emit('closeEmojiPicker')
    }
}

const onKeyUp = async (event: any) => {
    if (props.chat.type === 'personal') {
        return
    }

    const field = event.currentTarget
    const text = field.value

    const closestSpaceIndex = text.lastIndexOf(' ', field.selectionStart - 1)
    const currentWord = text.slice(closestSpaceIndex + 1, field.selectionStart)

    usernameFilter.value = currentWord.startsWith('@') ? currentWord.slice(1) : null

    if (usernameFilter.value !== null) {
        showMentionTooltip()
    } else {
        mentionTooltip.value.hide()
    }
}

const onKeyDown = (event: any) => {
    if (event.keyCode === 13 && !event.shiftKey) {
        if (mentionTooltip.value && mentionTooltip.value.state.isShown) {
            return
        }

        event.preventDefault()

        submit()
    }
}

const onFileSelect = (event: any) => {
    // const input = event.currentTarget
    // data.attachments.push(...input.files)
    // $showPopup('filesPopup')
}

const onUserMention = (event: any) => {
    usernameFilter.value = event.user.username

    const field = textarea.value

    if (!field) {
        return
    }

    const atSignIndex = data.text.lastIndexOf('@', field.selectionStart)

    field.setRangeText(`@${event.user.username} `, atSignIndex, field.selectionStart, 'end')
    field.focus()

    mentionTooltip.value.hide()
}

const submit = () => {
    const newMessage = {
        chatId: props.chat.id,
        parentId: data.parentId,
        originalId: null,
        authorName: null,
        text: data.text.trim(),
    }

    const forwardedMessages = store.messagesToForward.map((message) => ({
        chatId: props.chat.id,
        parentId: message.parentId,
        originalId: message.id,
        authorName: message.senderName,
        text: message.text,
    }))

    if (isComplete.value) {
        store.client.sendEvent({
            type: 'send_messages',
            payload: {
                chatId: props.chat.id,
                messages: [newMessage, ...forwardedMessages],
            },
        })
    } else if (store.messagesToForward.length !== 0) {
        store.sendMessageForm(props.chat.id, null, store.messagesToForward)
    }

    reset()
}

const reset = () => {
    Object.assign(data, getInitialData())

    setTimeout(() => {
        if (textarea.value) {
            adjustTextarea(textarea.value)
        }
    })

    emit('cancelReplying')
}

defineExpose({
    focus,
})
</script>

<template>
    <form
        ref="root"
        class="relative flex items-end max-w-[720px]"
        action="#"
        method="POST"
        @submit.prevent="submit()"
    >
        <Tippy
            v-if="chat.type === 'group'"
            ref="mentionTooltip"
            theme="none"
            placement="top-start"
            trigger="manual"
            interactive
            :arrow="false"
            :offset="[0, 4]"
        >
            <template #content="{ hide, state }">
                <MentionMenu
                    :chat="chat"
                    :filter="usernameFilter"
                    :active="state.isShown"
                    @select="onUserMention"
                    @noresults="hide()"
                />
            </template>
        </Tippy>

        <div class="relative grow flex items-end shadow-elevation-1">
            <div
                class="grow px-[12px] py-[8px] bg-primary-brand-white rounded-[12px] rounded-br-none"
            >
                <div v-if="parent" class="flex items-end mb-[12px]">
                    <MessageParent class="grow" :senderName="parent.senderName" :text="parent.text" />

                    <button
                        class="shrink-0 button m ghost only-icon !p-[5px]"
                        type="button"
                        @click="$emit('cancelReplying')"
                    >
                        <DeleteDisabledSmallIcon width="24" height="24"/>
                    </button>
                </div>
                <div v-else-if="store.messagesToForward.length !== 0">
                    <!-- {{ $tc('forwardedMessage', store.messagesToForward.length) }} -->
                    19 messages
                </div>
                <div class="flex items-end gap-x-[12px]">
                    <div class="shrink-0">
                        <Tippy
                            theme="none"
                            placement="top-start"
                            :arrow="false"
                            :duration="[0, 0]"
                            :offset="[-12, 12]"
                            :zIndex="35"
                            trigger="click"
                            interactive
                            inlinePositioning
                        >
                            <button
                                id="emojiPickerTrigger"
                                class="button l ghost !p-[5px]"
                                type="button"
                                @click="onEmojiPickerToggle"
                            >
                                <EmojiIcon width="24" height="24" />
                            </button>

                            <template #content>
                                <EmojiPicker
                                    class="hidden md:flex"
                                    locale="ru"
                                    previewPosition="none"
                                    searchPosition="none"
                                    theme="light"
                                    @emojiSelect="onEmojiSelect"
                                />
                            </template>
                        </Tippy>
                    </div>

                    <div class="grow">
                        <textarea
                            ref="textarea"
                            class="block w-full max-h-[284px] py-[8px] text-l-short-16 text-primary-brand-onPrimary placeholder:text-primary-seattle-60 outline-none resize-none"
                            rows="1"
                            placeholder="Напишите сообщение"
                            v-model="data.text"
                            v-autoresize
                            v-focus
                            @keyup="onKeyUp"
                            @keydown="onKeyDown"
                        ></textarea>
                    </div>

                    <div v-if="false" class="shrink-0">
                        <label class="button l ghost relative p-[5px]">
                            <input
                                class="absolute opacity-0 pointer-events-none"
                                type="file"
                                multiple
                                @change="onFileSelect"
                            />

                            <AttachIcon width="24" height="24" />
                        </label>
                    </div>
                </div>
            </div>
            <CornerIcon
                class="shrink-0 text-primary-brand-white -scale-x-100"
                width="9"
                height="16"
            />
        </div>

        <div class="relative shrink-0">
            <div class="absolute bottom-[calc(100%+12px)] pointer-events-none">
                <slot name="actions"></slot>
            </div>

            <button
                v-if="true"
                xv-if="data.text"
                class="button xl accent round only-icon relative"
                type="submit"
            >
                <SendIcon class="shrink-0" width="24" height="24" />
            </button>
            <button v-else class="button xl accent round only-icon relative" type="button">
                <CameraIcon class="shrink-0" width="24" height="24" />
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
