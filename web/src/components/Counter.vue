<script lang="ts" setup>
import { computed } from 'vue'
const props = defineProps<{
    value: number
    size: string
    variant: string
}>()

const isOverflowed = computed(() => {
    return props.value >= 100_000
})

const formattedValue = computed(() => {
    if (props.value <= 999) {
        return String(props.value)
    }

    if (isOverflowed.value) {
        return '...'
    }

    const thousands = Math.floor(props.value / 1000)

    return `${thousands}K`
})

const style: any = {
    base: 'inline-flex justify-center items-center max-w-[30px] rounded-2xl pointer-events-none',
    size: {
        default: 'min-w-[20px] h-[20px] px-[6px] py-[2px] text-[12px] leading-[16px]',
        small: 'min-w-[16px] h-[16px] py-[1px] px-[4.5px] text-[11px] leading-[14px]',
    },
    variant: {
        critical: 'bg-utilitarian-moscow-100 text-primary-brand-white',
        warning: 'bg-utilitarian-osaka-100 text-primary-brand-white',
        success: 'bg-utilitarian-geneva-100 text-primary-brand-white',
        accent: 'bg-primary-brand-accent text-primary-brand-white',
        neutral: 'bg-primary-brand-primary text-primary-brand-white',
        inactive: 'bg-primary-seattle-30 text-primary-brand-primary',
    },
}
</script>

<template>
    <div :class="[style.base, style.size[size], style.variant[variant]]">
        <span :class="{ 'mt-[-50%]': isOverflowed }">
            {{ formattedValue }}
        </span>
    </div>
</template>
