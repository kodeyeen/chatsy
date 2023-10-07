<script setup>
import { ref } from 'vue'

const props = defineProps({
    text: {
        type: String,
        required: true,
    },
})

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

<template>
    <button type="button" @click="onClick">
        <slot :isCooldown="isCooldown"></slot>
    </button>
</template>
