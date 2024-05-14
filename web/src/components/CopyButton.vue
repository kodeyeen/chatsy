<template>
    <button type="button" @click="onClick">
        <slot :isCooldown="isCooldown"></slot>
    </button>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const props = defineProps<{
    text: string
}>()

const isCooldown = ref(false)

const onClick = () => {
    if (isCooldown.value) {
        return
    }

    navigator.clipboard.writeText(props.text)

    isCooldown.value = true

    setTimeout(() => {
        isCooldown.value = false
    }, 1000)
}
</script>
