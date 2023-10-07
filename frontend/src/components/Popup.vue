<template>
    <div
      ref="root"
      :id="id"
      class="popup fade-in z-[2147483111]"
      :class="{ 'active': store.isActive(id) }"
      @click="onClick"
    >
      <slot name="dialog">
        <div
          class="popup-dialog slide-up-zoom-in"
          :class="[dialogClass, { 'active': store.isActive(id) }]"
        >
          <button class="popup-closer absolute right-[16px] md:right-0 bottom-[calc(100%+8px)]" type="button" data-popup-closer>
            <i class="icon dd-Delete-Disabled-small text-[16px]"></i>
  
            <span>
              Закрыть
            </span>
          </button>
  
          <div class="popup-content scrollbar" :class="[contentClass]">
            <slot></slot>
          </div>
        </div>
      </slot>
    </div>
  </template>
  
<script setup>
import { ref, watch, onMounted, onUnmounted } from 'vue'

import { usePopupsStore } from '@/stores/popups'

const props = defineProps({
    id: {
        type: String,
        required: true,
    },

    dialogClass: {
        type: String,
    },

    contentClass: {
        type: String,
    },
})
  
const emit = defineEmits(['show', 'shown', 'hide', 'hidden'])
  
// const { $lockPageScroll, $unlockPageScroll } = useNuxtApp()
const store = usePopupsStore()
  
watch(() => store.isActive(props.id), (isActive) => {
    if (isActive) {
        emit('show')
  
        // const style = getComputedStyle(root.value)
  
        // if (style['display'] !== 'none') {
        //     $lockPageScroll()
        // }

        setTimeout(() => emit('shown'), 500)
    } else {
        emit('hide')
  
        // setTimeout($unlockPageScroll, 200)
        setTimeout(() => emit('hidden'), 500)
    }
})

const root = ref(null)

const onClick = (event) => {
    // клик по оверлею
    const root = event.currentTarget
    const target = event.target

    if (target === root) {
        store.hide(props.id)
    }
}
  
const onAnyClick = (event) => {
    const target = event.target
    const closer = target.closest('[data-popup-closer]')

    if (closer) {
        const closestPopup = closer.closest('.popup')

        if (closestPopup === root.value) {
            store.hide(props.id)
        }
    }
}
  
onMounted(() => {
    if (props.id) {
        document.addEventListener('click', onAnyClick)
    }
})

onUnmounted(() => {
    if (props.id) {
        document.removeEventListener('click', onAnyClick)
    }
})
</script>
  
<style>
.popup {
    @apply
    fixed
    top-0
    left-0
    z-[2147483111]
    w-full
    h-full
    pt-[64px]
    md:pb-[32px]
    flex
    md:overflow-y-auto

    bg-[rgba(0,0,0,0.5)]
}
  
.popup.fade-in {
    transition-delay: 200ms;
}

.popup.fade-in.active {
    transition-delay: 0ms;
}
  
.popup-dialog {
    @apply
    relative

    flex
    max-w-[480px]
    w-full
    max-h-full
    md:max-h-max
    md:h-auto
    m-auto
    mb-0
    md:mb-auto

    rounded-t-[12px]
    md:rounded-[16px]
    shadow-elevation-0
    md:shadow-elevation-5
    bg-primary-brand-white
    ;
}
  
.popup-dialog.slide-up-zoom-in {
    transform: translateY(100%);
    transition: transform 200ms;
    transition-delay: 0ms;
}

@media (min-width: 712px) {
    .popup-dialog.slide-up-zoom-in {
        transform: scale(0.9);
        visibility: hidden;
        opacity: 0;
        transition: transform 200ms,
                    visibility 200ms,
                    opacity 200ms;
        transition-delay: 0ms;
    }
}

.popup-dialog.slide-up-zoom-in.active {
    transform: translateY(0);
    transition-delay: 200ms;
}

@media (min-width: 712px) {
    .popup-dialog.slide-up-zoom-in.active {
        transform: scale(1);
        visibility: visible;
        opacity: 1;
        transition-delay: 200ms;
    }
}

.popup-content {
    @apply
    w-full
    overflow-y-auto
    px-[16px]
    pt-[20px]
    pb-[32px]
    md:p-[32px]
    ;
}
  
.popup-closer {
    @apply
    flex
    items-center
    gap-[4px]
    pr-[11px]
    pl-[9px]
    py-[4px]
    rounded-[26px]
    bg-[rgba(43,45,51,0.3)]
    text-[11px]
    leading-[14px]
    font-medium
    text-primary-seattle-1
    ;
}

.fade-in {
    @apply
    invisible
    opacity-0
    transition-[visibility,opacity]
    duration-[200ms]
    ;
}

.fade-in.active {
    @apply
    visible
    opacity-100
    ;
}
</style>
