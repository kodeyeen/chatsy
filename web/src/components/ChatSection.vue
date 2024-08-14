<script lang="ts" setup>
const props = defineProps<{
    title: string
    subtitle?: string
}>()
</script>

<template>
    <div class="px-[20px] py-[16px] rounded-[20px] bg-primary-brand-white shadow-elevation-1">
        <div>
            <div class="text-h5 text-primary-brand-primary">
                {{ title }}
            </div>

            <div v-if="subtitle" class="mt-[8px] text-m-14 text-primary-seattle-100">
                {{ subtitle }}
            </div>
        </div>

        <div class="mt-[24px]">
            <slot></slot>
        </div>
    </div>
</template>
