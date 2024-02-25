<template>
    <div class="counter">
        <span class="counter__value" :class="{ overflowed: isOverflowed }">
            {{ formattedValue }}
        </span>
    </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
const props = defineProps<{
    value: number
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
</script>

<style>
.counter {
    @apply inline-flex
      justify-center
      items-center
      max-w-[30px]
      px-[6px]
      py-[2px]
      rounded-2xl
      pointer-events-none
  
      min-w-[20px]
      h-[20px]
      text-[12px]
      leading-[16px];
}

.counter.small {
    @apply min-w-[16px]
      h-[16px]
      py-[1px]
      px-[4.5px]
      text-[11px]
      leading-[14px];
}

.counter.critical {
    @apply bg-utilitarian-moscow-100
      text-primary-brand-white;
}

.counter.warning {
    @apply bg-utilitarian-osaka-100
      text-primary-brand-white;
}

.counter.success {
    @apply bg-utilitarian-geneva-100
      text-primary-brand-white;
}

.counter.accent {
    @apply bg-primary-brand-accent
      text-primary-brand-white;
}

.counter.neutral {
    @apply bg-primary-brand-primary
      text-primary-brand-white;
}

.counter.inactive {
    @apply bg-primary-seattle-30
      text-primary-brand-primary;
}

.counter__value.overflowed {
    margin-top: -50%;
}
</style>
