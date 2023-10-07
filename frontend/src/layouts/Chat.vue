<script setup>
import { ref } from 'vue'

const isContentActive = ref(false)

const toggleContent = (force = null) => {
    isContentActive.value = force ?? !isContentActive.value
}
</script>

<template>
    <div class="grow flex flex-col h-full">
        <div style="height: 60px; border: 1px solid #000">HEADER</div>

        <div class="grow flex overflow-hidden">
            <div
                class="shrink-0 w-full 2md:w-[430px] transition-transform duration-[500ms]"
                :class="{ '-translate-x-full 2md:translate-x-[-430px] 2lg:translate-x-0': isContentActive }"
            >
                <slot name="leftbar" :toggleContent="toggleContent"></slot>
            </div>

            <div
                class="grow shrink-0 w-full 2lg:w-auto transition-transform duration-[500ms]"
                :class="{ '-translate-x-full 2md:translate-x-[-430px] 2lg:translate-x-0': isContentActive }"
                @click="isContentActive = true"
            >
                <slot name="content" :toggleContent="toggleContent"></slot>
            </div>
        </div>
    </div>
</template>
