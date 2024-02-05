<script lang="ts" setup>
const props = withDefaults(
    defineProps<{
        src?: string
    }>(),
    {
        src: '/avatar',
    },
)
</script>

<template>
    <span class="block border border-primary-seattle-10 overflow-hidden">
        <img class="w-full h-full object-cover" :src="src" alt="Avatar" loading="lazy" />
    </span>
</template>
