<script lang="ts" setup>
import { watch } from 'vue'

import { usePopupsStore } from '@/stores/popups'
import DeleteDisabledSmallIcon from '@/components/icons/DeleteDisabledSmall.vue'

const props = defineProps<{
    id: string
    dialogClass?: string
    contentClass?: string
}>()

const emit = defineEmits(['show', 'shown', 'hide', 'hidden'])

const popupsStore = usePopupsStore()

watch(
    () => popupsStore.isActive(props.id),
    (isActive: boolean) => {
        if (isActive) {
            emit('show')

            setTimeout(() => emit('shown'), 500)
        } else {
            emit('hide')

            setTimeout(() => emit('hidden'), 500)
        }
    },
)
</script>

<template>
    <div
        :id="id"
        class="popup fade-in fixed top-0 left-0 z-[2147483111] w-full h-full pt-[64px] md:pb-[32px] flex md:overflow-y-auto bg-[rgba(0,0,0,0.5)]"
        :class="{
            active: popupsStore.isActive(id),
        }"
        @click.self="popupsStore.hide(props.id)"
    >
        <slot name="dialog">
            <div
                class="popup-dialog slide-up-zoom-in relative flex max-w-[480px] w-full max-h-full md:max-h-max md:h-auto m-auto mb-0 md:mb-auto rounded-t-[12px] md:rounded-[16px] shadow-elevation-0 md:shadow-elevation-5 bg-primary-brand-white"
                :class="[dialogClass, { active: popupsStore.isActive(id) }]"
            >
                <button
                    class="flex items-center gap-[4px] pr-[11px] pl-[9px] py-[4px] rounded-[26px] bg-[rgba(43,45,51,0.3)] text-[11px] leading-[14px] font-medium text-primary-seattle-1 absolute right-[16px] md:right-0 bottom-[calc(100%+8px)]"
                    type="button"
                    @click="popupsStore.hide(props.id)"
                >
                    <DeleteDisabledSmallIcon :width="16" :height="16" />

                    <span> Close </span>
                </button>

                <div
                    class="scrollbar w-full overflow-y-auto px-[16px] pt-[20px] pb-[32px] md:p-[32px]"
                    :class="contentClass"
                >
                    <slot></slot>
                </div>
            </div>
        </slot>
    </div>
</template>

<style lang="scss">
.popup.fade-in {
    transition-delay: 200ms;

    &.active {
        transition-delay: 0ms;
    }
}

.popup-dialog.slide-up-zoom-in {
    transform: translateY(100%);
    transition: transform 200ms;
    transition-delay: 0ms;

    @media (min-width: 712px) {
        transform: scale(0.9);
        visibility: hidden;
        opacity: 0;
        transition:
            transform 200ms,
            visibility 200ms,
            opacity 200ms;
        transition-delay: 0ms;
    }

    &.active {
        transform: translateY(0);
        transition-delay: 200ms;

        @media (min-width: 712px) {
            transform: scale(1);
            visibility: visible;
            opacity: 1;
            transition-delay: 200ms;
        }
    }
}

.fade-in {
    @apply invisible
  opacity-0
  transition-[visibility,opacity]
  duration-[200ms];

    &.active {
        @apply visible
    opacity-100;
    }
}
</style>
