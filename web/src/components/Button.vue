<script lang="ts" setup>
const props = defineProps<{
    text?: string
    type?: string
    size: string
    variant: string
    disabled?: boolean
    round?: boolean
}>()

const style: any = {
    base: 'box-border flex justify-center items-center border border-transparent font-medium transition-[transform,background-color,border-color,box-shadow] outline-none cursor-pointer shrink-0 disabled:cursor-not-allowed [&.loading]:pointer-events-none [&:not(.loading):not(:disabled)]:active:scale-[.96] focus-visible:border focus-visible:border-primary-atlantic-120',
    round: 'rounded-full',
    size: {
        xl: 'px-[23px] py-[13px] rounded-[10px] text-l-long-16',
        l: 'px-[19px] py-[9px] rounded-[8px] text-l-long-16',
        m: 'px-[19px] py-[7px] rounded-[8px] text-m-14',
        s: 'px-[15px] py-[5px] rounded-[8px] text-m-14',
        xs: 'px-[11px] py-[3px] rounded-[4px] text-s-12',
    },
    variant: {
        accent: 'border-primary-brand-accent bg-primary-brand-accent text-primary-brand-white [&:not(.loading):not(:disabled)]:hover:bg-primary-brand-onAccent [&:not(.loading):not(:disabled)]:active:bg-primary-brand-onAccent disabled:opacity-30 disabled:bg-primary-brand-accent [&.loading]:bg-primary-london-100 [&.loading]:border-primary-london-100 [&.loading]:text-primary-london-120',
        primary:
            'border-primary-brand-primary bg-primary-brand-primary text-primary-brand-white [&:not(.loading):not(:disabled)]:hover:bg-primary-brand-onPrimary [&:not(.loading):not(:disabled)]:active:bg-primary-brand-onPrimary disabled:bg-primary-london-110 disabled:text-primary-london-120 disabled:border-transparent [&.loading]:bg-primary-london-110 [&.loading]:border-primary-london-110 [&.loading]:text-primary-london-120',
        outlined:
            'border-primary-seattle-30 bg-transparent text-primary-brand-primary [&:not(.loading):not(:disabled)]:hover:border-primary-seattle-60 [&:not(.loading):not(:disabled)]:active:border-primary-seattle-60 disabled:bg-transparent disabled:text-primary-london-120 disabled:border-primary-seattle-10 [&.loading]:border-primary-seattle-10 [&.loading]:text-primary-london-120 [&.loading]:bg-transparent',
        critical:
            '!border-primary-accentOpacity-14 bg-primary-accentOpacity-7 text-utilitarian-moscow-100 [&:not(.loading):not(:disabled)]:hover:bg-primary-accentOpacity-14 [&:not(.loading):not(:disabled)]:active:bg-primary-accentOpacity-14 disabled:bg-primary-london-100 disabled:text-primary-london-120 disabled:border-transparent [&.loading]:border-primary-london-100 [&.loading]:text-primary-london-120 [&.loading]:bg-primary-london-100',
        clear: 'border-transparent bg-primary-london-100 text-primary-brand-primary [&:not(.loading):not(:disabled)]:hover:bg-primary-london-110 [&:not(.loading):not(:disabled)]:active:bg-primary-london-110 disabled:bg-primary-london-100 disabled:text-primary-london-120 disabled:border-transparent [&.loading]:border-primary-london-100 [&.loading]:text-primary-london-120 [&.loading]:bg-primary-london-100',
        ghost: 'border-transparent bg-transparent text-primary-seattle-100 [&:not(.loading):not(:disabled)]:hover:border-transparent [&:not(.loading):not(:disabled)]:hover:bg-primary-london-100 [&:not(.loading):not(:disabled)]:hover:text-primary-brand-primary [&:not(.loading):not(:disabled)]:active:bg-primary-london-100 [&:not(.loading):not(:disabled)]:active:text-primary-brand-primary disabled:bg-transparent disabled:text-primary-london-120 disabled:border-transparent [&.loading]:border-transparent [&.loading]:text-primary-london-120 [&.loading]:bg-transparent focus-visible:text-primary-brand-primary',
        white: 'border-transparent bg-primary-brand-white text-primary-brand-primary [&:not(.loading):not(:disabled)]:hover:bg-primary-seattle-1 [&:not(.loading):not(:disabled)]:active:bg-primary-seattle-1 disabled:bg-primary-seattle-1 disabled:text-primary-london-120 [&.loading]:text-primary-london-120 focus-visible:bg-primary-brand-white shadow-[0_0_0_1px_rgba(32,35,51,0.1),0_2px_4px_-1px_rgba(32,35,51,0.15)]',
    },
}
</script>

<template>
    <button
        v-if="type === 'button' || type === 'submit' || type === 'reset'"
        :class="[
            style.base,
            style.size[size],
            style.variant[variant],
            props.round ? style.round : null,
        ]"
        :type="type"
        :disabled="disabled"
    >
        <slot></slot>
    </button>
</template>
