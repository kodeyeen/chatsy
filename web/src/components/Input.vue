<script lang="ts" setup>
const props = defineProps<{
    value?: string
    type: 'text' | 'search'
    size: 'l' | 'm' | 's'
    variant: 'outlined' | 'filled'
    placeholder?: string
    disabled?: boolean
    readonly?: boolean
}>()

const style: any = {
    base: 'block w-full border focus:border-primary-seattle-140 hover:border-primary-seattle-140 rounded-[8px] text-primary-brand-primary bg-primary-brand-white focus-visible:outline-none focus:placeholder:text-transparent transition-colors outline-none disabled:border-primary-seattle-30 disabled:bg-primary-brand-wildSand disabled:cursor-not-allowed disabled:border-primary-seattle-30 disabled:bg-primary-brand-wildSand disabled:cursor-not-allowed hover:border-primary-seattle-140 focus:border-primary-seattle-140',
    size: {
        l: 'h-[44px] px-[11px] py-[9px] text-l-long-16',
        m: 'h-[36px] px-[11px] py-[7px] text-m-14',
        s: 'px-[11px] py-[5px] text-m-14',
    },
    variant: {
        outlined: 'border-primary-seattle-30 bg-primary-brand-white',
        filled: 'border-primary-brand-wildSand bg-primary-brand-wildSand',
    },
    spacing: {
        l: 'pl-[44px]',
        m: 'pl-[40px]',
        s: '',
    },
}
</script>

<template>
    <div class="relative">
        <input
            :class="[
                style.base,
                style.size[size],
                style.variant[variant],
                $slots.start ? style.spacing[size] : null,
            ]"
            :value="value"
            :type="type"
            :placeholder="placeholder"
            :disabled="disabled"
            :readonly="readonly"
        />

        <slot name="start"></slot>

        <slot name="end"></slot>
    </div>
</template>
