<script lang="ts" setup>
import { useClipboard } from '@vueuse/core'

const props = defineProps<{
    source: string
}>()

const { copy, copied } = useClipboard({ source: props.source })
</script>

<template>
    <slot :copy="copy" :copied="copied"></slot>
</template>
