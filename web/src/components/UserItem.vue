<script lang="ts" setup>
import Avatar from '@/components/Avatar.vue'

const props = defineProps<{
    user: any
}>()
</script>

<template>
    <div class="flex items-center gap-x-[12px]">
        <div class="shrink-0">
            <RouterLink to="/chats">
                <Avatar class="w-[44px] h-[44px] rounded-full" />
            </RouterLink>
        </div>

        <div class="grow">
            <div class="flex gap-x-[6px]">
                <RouterLink
                    class="line-clamp-1 text-m-14 md:text-l-long-16 font-medium text-primary-brand-primary"
                    to="/chats"
                >
                    {{ user.display_name }}
                </RouterLink>
            </div>
        </div>

        <div v-if="$slots.end" class="shrink-0">
            <slot name="end"></slot>
        </div>
    </div>
</template>
