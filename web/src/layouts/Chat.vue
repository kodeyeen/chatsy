<script lang="ts" setup>
import { ref } from 'vue'

import TheHeader from '@/components/TheHeader.vue'

const isContentActive = ref(false)

const toggleContent = (force = null) => {
    isContentActive.value = force ?? !isContentActive.value
}
</script>

<template>
    <div class="flex flex-col h-full">
        <TheHeader />

        <main class="grow flex overflow-hidden">
            <div
                class="shrink-0 w-full 2md:w-[430px] transition-transform duration-[500ms]"
                :class="{
                    '-translate-x-full 2md:translate-x-[-430px] 2lg:translate-x-0': isContentActive,
                }"
            >
                <slot name="leftbar" :toggleContent="toggleContent"></slot>
            </div>

            <div
                class="relative grow shrink-0 w-full 2lg:w-auto bg-[#EEEAE3] transition-transform duration-[500ms]"
                :class="{
                    '-translate-x-full 2md:translate-x-[-430px] 2lg:translate-x-0': isContentActive,
                }"
                @click="isContentActive = true"
            >
                <div
                    class="absolute top-0 left-0 w-full h-full bg-[url('@/assets/images/bg.png')] bg-[length:670px] bg-[0_0] opacity-[.25] pointer-events-none"
                ></div>

                <slot name="content" :toggleContent="toggleContent"></slot>
            </div>
        </main>
    </div>
</template>
