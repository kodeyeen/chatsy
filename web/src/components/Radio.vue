<script lang="ts" setup>
const props = defineProps<{
    text?: string
    name?: string
    value?: string
    size: string
    disabled?: boolean
}>()

const style: any = {
    box: {
        base: 'flex justify-center items-center border-2 border-primary-seattle-30 rounded-full bg-primary-brand-white transition-colors hover:border-primary-brand-primary peer-checked:border-primary-brand-primary peer-checked:bg-primary-brand-primary peer-checked:hover:border-primary-london-40 peer-disabled:border-primary-seattle-10 peer-disabled:bg-primary-seattle-10 peer-disabled:hover:border-primary-seattle-10 peer-disabled:cursor-default peer-checked:[&_span]:bg-primary-brand-white peer-checked:[&_span]:scale-100 peer-checked:peer-disabled:[&_span]:bg-primary-london-120',
        size: {
            l: 'w-[24px] h-[24px]',
            m: 'w-[20px] h-[20px]',
            s: 'w-[16px] h-[16px]',
        },
    },
    icon: {
        base: 'scale-0 shrink-0 pointer-events-none rounded-full transition-transform',
        size: {
            l: 'w-2.5 h-2.5',
            m: 'w-2 h-2',
            s: 'w-1.5 h-1.5',
        },
    },
    label: {
        base: 'peer-disabled:text-primary-london-120',
        size: {
            l: 'ml-[12px] text-l-long-16',
            m: 'ml-[8px] text-m-14',
            s: 'ml-[8px] text-s-12',
        },
    },
}
</script>

<template>
    <label class="relative flex items-center cursor-pointer">
        <input
            class="peer absolute opacity-0 pointer-events-none"
            type="radio"
            :name="name"
            :value="value"
        />

        <span :class="[style.box.base, style.box.size[size]]">
            <span :class="[style.icon.base, style.icon.size[size]]"></span>
        </span>

        <span v-if="text" :class="[style.label.base, style.label.size[size]]">{{ text }}</span>
    </label>
</template>
