<script lang="ts" setup>
import { computed } from 'vue'

import Avatar from '@/components/Avatar.vue'
import CheckIcon from '@/components/icons/Check.vue'
import CornerIcon from '@/components/icons/Corner.vue'
import MessageParent from '@/components/Message/Parent.vue'
import ReadIcon from '@/components/icons/Read.vue'

const props = withDefaults(
    defineProps<{
        message: any
        modelValue: any
        selectable: boolean
        own: boolean
        firstInGroup: boolean
        lastInGroup: boolean
        showUserInfo: boolean
    }>(),
    {
        modelValue: [],
        selectable: false,
        own: false,
        firstInGroup: false,
        lastInGroup: false,
        showUserInfo: false,
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

// const author = props.message.forward_original
//     ? props.message.forward_original.author
//     : props.message.author

const isSelected = computed(() => {
    return model.value.some((message: any) => message.id === props.message.id)
})

const formattedText = computed(() => {
    const pattern = /\B@[A-Za-z0-9_-]+/g

    return props.message.text.replaceAll(pattern, (mention: any) => {
        const username = mention.slice(1)

        return `<a class="button-ghost link inline" href="/users/${username}/" target="_blank">${mention}</a>`
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
            <input
                class="l absolute top-1/2 right-[calc(100%+20px)] -translate-y-1/2 opacity-0 border-primary-seattle-100 hover:border-primary-brand-primary checked:border-primary-brand-primary transition-opacity duration-300"
                :class="{
                    'opacity-100': selectable,
                }"
                type="checkbox"
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
                    to="$userLink(message.author)"
                >
                    <Avatar class="w-[44px] h-[44px] rounded-full" />
                </RouterLink>

                <div class="flex [.own_&]:flex-row-reverse items-end max-w-[426px]">
                    <CornerIcon
                        class="shrink-0 text-primary-brand-white [.own_&]:text-[rgba(241,254,225,1)] [.own_&]:-scale-x-100"
                        :class="{
                            invisible: !lastInGroup,
                        }"
                        width="9"
                        height="16"
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

                                <RouterLink class="button-ghost m-medium link !font-bold" to="$userLink(author)">
                                    <span>
                                        {{ message.authorFullName }}
                                    </span>
                                </RouterLink>
                            </span>
                            <span v-else class="flex items-center gap-x-[4px]">
                                <RouterLink class="button-ghost m-medium link !font-bold" to="$userLink(author)">
                                    <span>
                                        {{ message.senderFullName }}
                                    </span>
                                </RouterLink>
                            </span>
                        </div>

                        <div>
                            <MessageParent
                                v-if="message.parentId"
                                :message="message.parentId"
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
                                width="17"
                                height="17"
                            />
                            <CheckIcon
                                v-else
                                class="hidden [.own_&]:block text-utilitarian-geneva-100"
                                width="17"
                                height="17"
                            />
                        </span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
