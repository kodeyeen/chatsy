<script lang="ts" setup>
const props = defineProps<{
    senderName: string
    text: string
}>()
</script>

<template>
    <div class="flex gap-x-[8px] pt-[6px] pb-[4px] cursor-pointer">
        <div class="shrink-0 w-[2px] rounded-[10px] bg-primary-brand-accent"></div>

        <span v-if="false" class="shrink-0 w-[34px] h-[34px] rounded-[6px] overflow-hidden">
            <img
                class="w-full object-cover aspect-square"
                src="https://telegra.ph/file/d522a1a85381bb5aafe6e.jpg"
                alt="Media"
            />
        </span>

        <div class="grow">
            <div class="text-m-14 font-medium text-primary-brand-primary">
                {{ senderName }}
            </div>

            <div class="mt-[-2px] text-s-12 line-clamp-1 text-primary-brand-onPrimary">
                {{ text }}
            </div>
        </div>
    </div>
</template>
