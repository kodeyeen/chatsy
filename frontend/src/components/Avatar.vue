<script setup>
const props = defineProps({
    src: {
        type: String,
        default: '/assets/images/default_avatar.svg',
    },
})
</script>

<template>
    <span class="block border border-primary-seattle-10 overflow-hidden">
        <img :src="src" class="w-full h-full object-cover" alt="Avatar" loading="lazy">
    </span>
</template>
