<script lang="ts" setup>
import DoneCheckIcon from '@/components/icons/DoneCheck.vue'
// import MinusIcon from '@/components/icons/Minus.vue'

const props = defineProps<{
    boxClass?: string
    label?: string
    size: string
    checked?: boolean
    disabled?: boolean
}>()

const style: any = {
    box: {
        base: 'flex justify-center items-center border-2 border-primary-seattle-30 rounded-[4px] transition-colors bg-primary-brand-white text-primary-brand-primary hover:border-primary-brand-primary peer-disabled:bg-primary-seattle-10 peer-disabled:border-primary-seattle-10 peer-disabled:hover:border-primary-seattle-10 peer-disabled:peer-checked:text-primary-brand-primary peer-disabled:cursor-default peer-checked:bg-primary-brand-primary peer-checked:border-primary-brand-primary peer-checked:text-primary-brand-white peer-checked:hover:border-primary-london-40 peer-indeterminate:border-primary-brand-primary peer-indeterminate:bg-primary-brand-primary peer-indeterminate:text-primary-brand-white peer-indeterminate:hover:border-primary-london-40 peer-invalid:bg-utilitarian-moscow-10 peer-invalid:border-utilitarian-moscow-100 peer-invalid:peer-checked:text-primary-brand-primary peer-checked:[&_svg]:scale-100 peer-indeterminate:[&_svg]:hidden',
        size: {
            l: 'w-[24px] h-[24px] text-[20px]',
            m: 'w-[20px] h-[20px] text-[16px]',
            s: 'w-[16px] h-[16px] text-[12px]',
        },
    },
    icon: {
        base: 'absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 scale-0 transition-transform',
        size: {
            l: 'w-[20px] h-[20px]',
            m: 'w-[16px] h-[16px]',
            s: 'w-[12px] h-[12px]',
        },
    },
    label: {
        base: 'text-primary-brand-primary peer-disabled:text-primary-london-120',
        size: {
            l: 'ml-[12px] text-[16px] leading-[24px]',
            m: 'ml-[8px] text-[14px] leading-[20px]',
            s: 'ml-[8px] text-[12px] leading-[16px]',
        },
    },
}
</script>

<template>
    <label class="relative flex items-center cursor-pointer">
        <input
            class="peer absolute opacity-0"
            type="checkbox"
            :checked="checked"
            :disabled="disabled"
        />

        <span class="relative" :class="[style.box.base, style.box.size[size], boxClass]">
            <DoneCheckIcon :class="`${style.icon.base} ${style.icon.size[size]} `" />
            <!-- <MinusIcon :class="`${style.icon.base} ${style.icon.size[size]}`" /> -->
        </span>

        <span v-if="label" :class="`${style.label.base} ${style.label.size[size]}`">
            {{ label }}
        </span>
    </label>
</template>
