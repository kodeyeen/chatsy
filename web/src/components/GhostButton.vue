<script lang="ts" setup>
import { computed } from 'vue'
import type { RouteLocationRaw } from 'vue-router'

const props = defineProps<{
    type: 'button' | 'submit' | 'reset' | 'file' | 'link'
    size: 'l-medium' | 'm-medium' | 's-medium' | 'l-regular' | 'm-regular' | 's-regular'
    variant: 'accent' | 'primary' | 'ghost-10' | 'ghost-30' | 'ghost-70' | 'link'
    to?: RouteLocationRaw
    caps?: boolean
    disabled?: boolean
}>()

const style: any = {
    base: 'relative flex justify-center items-center gap-x-[8px] bg-transparent border-0 disabled:cursor-not-allowed transition-colors cursor-pointer',
    caps: 'uppercase',
    underlined: 'border-b',
    size: {
        'l-medium': 'text-base font-medium leading-6',
        'm-medium': 'text-sm font-medium leading-5',
        's-medium': 'text-xs font-medium leading-4',
        'l-regular': 'text-[16px] leading-6',
        'm-regular': 'text-[14px] leading-5',
        's-regular': 'text-[12px] leading-4',
    },
    variant: {
        accent: 'text-primary-brand-accent hover:text-primary-brand-onPrimary disabled:text-primary-london-120',
        primary:
            'text-primary-brand-primary hover:text-primary-brand-onAccent disabled:text-primary-london-120',
        'ghost-10':
            'text-primary-seattle-60 hover:text-primary-brand-onPrimary disabled:text-primary-london-110',
        'ghost-30':
            'text-primary-seattle-100 hover:text-primary-brand-onPrimary disabled:text-primary-london-110',
        'ghost-70':
            'text-primary-seattle-120 hover:text-primary-brand-onPrimary disabled:text-primary-london-110',
        link: 'text-primary-atlantic-140 hover:text-primary-brand-onPrimary disabled:text-primary-london-110',
    },
}

const cls = computed(() => [
    style.base,
    style.size[props.size],
    style.variant[props.variant],
    props.caps ? style.caps : null,
])
</script>

<template>
    <button
        v-if="type === 'button' || type === 'submit' || type === 'reset'"
        :class="cls"
        :type="type"
    >
        <slot></slot>
    </button>
    <RouterLink v-else-if="type === 'link' && to" :class="cls" :to="to">
        <slot></slot>
    </RouterLink>
    <label v-else-if="type === 'file'" :class="cls">
        <input class="absolute opacity-0 pointer-events-none" :type="type" />

        <slot></slot>
    </label>
</template>
