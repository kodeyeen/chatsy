<script lang="ts" setup>
import type { RouteLocationRaw } from 'vue-router'

const props = defineProps<{
    text: string
    type: 'button' | 'submit' | 'reset' | 'link'
    to?: RouteLocationRaw
    active?: boolean
    disabled?: boolean
}>()

const style: any = {
    base: 'flex items-center w-full px-[10px] py-[8px] rounded-[4px] bg-primary-brand-white hover:bg-primary-seattle-10 transition-colors text-m-14 font-medium text-left text-primary-brand-primary',
}
</script>

<template>
    <button
        v-if="type === 'button' || type === 'submit' || type === 'reset'"
        :class="[style.base, active ? 'bg-primary-seattle-10' : null]"
        :type="type"
        :disabled="disabled"
    >
        <slot name="start"></slot>

        <slot name="content">
            <span v-if="text">{{ text }}</span>
        </slot>

        <slot name="end"></slot>
    </button>
    <RouterLink
        v-else-if="type === 'link' && to"
        :class="[style.base, active ? 'bg-primary-seattle-10' : null]"
        :to="to"
    >
        <slot name="start"></slot>

        <slot>
            <span v-if="text">{{ text }}</span>
        </slot>

        <slot name="end"></slot>
    </RouterLink>
</template>
