<script lang="ts" setup>
import { computed } from 'vue'

import Avatar from '@/components/Avatar.vue'
import GhostButton from '@/components/GhostButton.vue'
import Checkbox from '@/components/Checkbox.vue'
import CheckIcon from '@/components/icons/Check.vue'
import CornerIcon from '@/components/icons/Corner.vue'
import MessageParent from '@/components/MessageParent.vue'
import ReadIcon from '@/components/icons/Read.vue'

const props = withDefaults(
    defineProps<{
        message: any
        modelValue: any
        selectable?: boolean
        own?: boolean
        firstInGroup?: boolean
        lastInGroup?: boolean
        showUserInfo?: boolean
    }>(),
    {
        modelValue: [],
    },
)

const emit = defineEmits(['update:modelValue', 'parentClick'])

const model = computed({
    get() {
        return props.modelValue
    },

    set(newValue) {
        emit('update:modelValue', newValue)
    },
})

const isSelected = computed(() => {
    return model.value.some((message: any) => message.id === props.message.id)
})

const formattedText = computed(() => {
    const pattern = /\B@[A-Za-z0-9_-]+/g
    const cls =
        'inline bg-transparent border-0 transition-colors cursor-pointer text-primary-atlantic-140 hover:text-primary-brand-onPrimary'

    return props.message.text.replaceAll(pattern, (mention: any) => {
        const username = mention.slice(1)

        return `<a class="${cls}" href="/users/${username}/" target="_blank">${mention}</a>`
    })
})

const onSelect = () => {
    if (!props.selectable) {
        return
    }

    if (isSelected.value) {
        model.value = model.value.filter((message: any) => message.id !== props.message.id)
    } else {
        model.value = [...model.value, props.message]
    }
}

const onParentClick = () => {
    emit('parentClick', {
        message: props.message.parentId,
    })
}
</script>

<template>
    <div
        ref="root"
        class="block px-[6px] 2md:px-[22px] py-[2px] transition-colors [&.selected]:bg-[rgba(45,164,48,0.20)]"
        :class="{
            selected: isSelected,
            'cursor-pointer': selectable,
        }"
        @click="onSelect"
    >
        <div
            class="relative flex [.own_&]:justify-end max-w-[720px] mx-auto 2md:pr-[52px]"
            :class="{
                'pl-[53px]': own && showUserInfo,
            }"
        >
            <Checkbox
                class="!absolute top-1/2 right-[calc(100%+20px)] -translate-y-1/2 opacity-0"
                :class="{
                    'opacity-100': selectable,
                }"
                boxClass="border-primary-seattle-100 hover:border-primary-brand-primary peer-checked:border-primary-brand-primary transition-opacity duration-300"
                size="l"
                :checked="isSelected"
            />

            <div
                class="relative flex items-end max-w-max ml-0"
                :class="{
                    'pl-[44px]': !own && showUserInfo,
                }"
            >
                <RouterLink
                    v-if="!own && lastInGroup && showUserInfo"
                    class="absolute bottom-0 left-0"
                    to="/chats"
                >
                    <Avatar class="w-[44px] h-[44px] rounded-full" />
                </RouterLink>

                <div class="flex [.own_&]:flex-row-reverse items-end max-w-[426px]">
                    <CornerIcon
                        class="shrink-0 text-primary-brand-white [.own_&]:text-[rgba(241,254,225,1)] [.own_&]:-scale-x-100"
                        :class="{
                            invisible: !lastInGroup,
                        }"
                        :width="9"
                        :height="16"
                    />

                    <div
                        class="relative grow flex flex-col items-start max-w-[417px] px-[12px] py-[6px] rounded-[12px] [.own_&]:rounded-bl-[12px] bg-primary-brand-white [.own_&]:bg-[rgba(241,254,225,1)]"
                        :class="{
                            'rounded-bl-none [.own_&]:rounded-br-none': lastInGroup,
                        }"
                    >
                        <div
                            v-if="(firstInGroup && showUserInfo) || message.originalId"
                            class="flex items-center gap-x-[4px]"
                            :class="{
                                '[.own_&]:hidden': !message.originalId,
                            }"
                        >
                            <span v-if="message.originalId" class="flex items-center gap-x-[4px]">
                                <span class="text-m-14 font-medium text-primary-atlantic-140">
                                    Переслано от:
                                </span>

                                <GhostButton
                                    class="!font-bold"
                                    type="link"
                                    size="m-medium"
                                    variant="link"
                                    to="/chats"
                                >
                                    <span>
                                        {{ message.authorName }}
                                    </span>
                                </GhostButton>
                            </span>
                            <span v-else class="flex items-center gap-x-[4px]">
                                <GhostButton
                                    class="!font-bold"
                                    type="link"
                                    size="m-medium"
                                    variant="link"
                                    to="/chats"
                                >
                                    {{ message.senderName }}
                                </GhostButton>
                            </span>
                        </div>

                        <div>
                            <MessageParent
                                v-if="message.parentId"
                                :senderName="message.parentSenderName"
                                :text="message.parentText"
                                @click="onParentClick"
                            />

                            <div class="text-[16px] leading-[21px] text-primary-brand-onPrimary">
                                <pre
                                    class="inline font-[inherit] whitespace-pre-wrap"
                                    v-html="formattedText"
                                ></pre>
                                <span class="inline-block w-[44px] [.own_&]:w-[61px]"></span>
                            </div>

                            <!-- <div class="flex flex-col items-start mt-[8px]">
                                <ChatMessageAttachment/>
                            </div> -->
                        </div>

                        <span class="absolute bottom-[6px] right-[12px] flex items-end gap-x-[3px]">
                            <time
                                class="text-s-12 text-primary-seattle-100"
                                :datetime="message.sentAt"
                            >
                                {{
                                    new Date(message.sentAt).toLocaleTimeString('ru', {
                                        hour: '2-digit',
                                        minute: '2-digit',
                                    })
                                }}
                            </time>

                            <ReadIcon
                                v-if="message.isViewed"
                                class="hidden [.own_&]:block text-utilitarian-geneva-100"
                                :width="17"
                                :height="17"
                            />
                            <CheckIcon
                                v-else
                                class="hidden [.own_&]:block text-utilitarian-geneva-100"
                                :width="17"
                                :height="17"
                            />
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
