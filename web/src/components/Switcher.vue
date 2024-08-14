<script lang="ts" setup>
import { computed } from 'vue'

const props = defineProps<{
    name?: string
    value?: string
    text?: string
    size: string
    disabled?: boolean
    required?: boolean
    active?: boolean
}>()

const model = computed({
    get() {},

    set(newValue: any) {},
})

const style: any = {
    box: {
        base: 'shrink-0 relative p-[3px] rounded-2xl bg-primary-seattle-30 peer-disabled:bg-primary-seattle-10 peer-disabled:peer-checked:bg-primary-seattle-30 peer-checked:bg-primary-brand-primary transition-colors peer-checked:[&_span]:left-full peer-checked:[&_span]:-translate-y-1/2 peer-checked:[&_span]:translate-x-[calc(-100%-3px)] peer-disabled:[&_span]:bg-[rgba(32,35,51,0.28)] peer-disabled:peer-checked:[&_span]:bg-primary-brand-white',
        size: {
            l: 'w-10 h-6',
            m: 'w-8 h-5',
        },
    },
    icon: {
        base: 'absolute top-1/2 left-[3px] -translate-y-1/2 rounded-full bg-primary-brand-white transition-[left,transform]',
        size: {
            l: 'w-[18px] h-[18px]',
            m: 'w-[14px] h-[14px]',
        },
    },
    label: {
        base: 'ml-2 empty:ml-0 peer-disabled:text-primary-london-120',
        size: {
            l: 'text-[16px]',
            m: 'text-[14px]',
        },
    },
}
</script>

<template>
    <label class="relative flex cursor-pointer">
        <input
            class="peer absolute opacity-0 pointer-events-none"
            type="checkbox"
            v-model="model"
            :name="name"
            :value="value"
            :checked="active"
            :disabled="disabled"
            :required="required"
        />

        <span :class="[style.box.base, style.box.size[size]]">
            <span :class="[style.icon.base, style.icon.size[size]]"></span>
        </span>

        <span v-if="text" :class="[style.label.base, style.label.size[size]]">{{ text }}</span>
    </label>
</template>
