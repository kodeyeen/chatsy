<template>
    <div ref="root"></div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { Picker } from 'emoji-mart'

const props = withDefaults(
    defineProps<{
        class: string
        dynamicWidth?: boolean
        locale?: string
        theme?: string
        previewPosition?: string
        searchPosition?: string
    }>(),
    {
        dynamicWidth: false,
        locale: 'en',
        theme: 'auto',
        previewPosition: 'bottom',
        searchPosition: 'sticky',
    },
)

const emit = defineEmits(['emojiSelect', 'clickOutside'])

const root = ref<HTMLDivElement | null>(null)

onMounted(() => {
    const picker = new Picker({
        data: async () => {
            const response = await fetch('https://cdn.jsdelivr.net/npm/@emoji-mart/data')

            return response.json()
        },

        dynamicWidth: props.dynamicWidth,
        locale: props.locale,
        previewPosition: props.previewPosition,
        searchPosition: props.searchPosition,
        theme: props.theme,

        onEmojiSelect: (...args: any[]) => {
            emit('emojiSelect', ...args)
        },

        onClickOutside: (...args: any[]) => {
            emit('clickOutside', ...args)
        },
    }) as unknown as HTMLElement

    if (props.class) {
        picker.classList.add(...props.class.split(' '))
    }

    setTimeout(() => {
        if (root.value) {
            root.value.append(picker)
        }
    })
})
</script>
