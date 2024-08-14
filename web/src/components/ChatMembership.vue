<script lang="ts" setup>
import Avatar from '@/components/Avatar.vue'
import DotsMenuIcon from '@/components/icons/DotsMenu.vue'

const props = withDefaults(
    defineProps<{
        membership: any
        readonly: boolean
    }>(),
    {
        readonly: false,
    },
)
</script>

<template>
    <div class="flex items-center gap-x-[9px] p-[9px]">
        <div class="shrink-0">
            <RouterLink to="/chats">
                <Avatar class="w-[44px] h-[44px] rounded-full" />
            </RouterLink>
        </div>

        <div class="grow">
            <RouterLink class="text-primary-brand-onPrimary text-l-long-16 font-medium" to="/chats">
                {{ membership.user.display_name }}
            </RouterLink>
        </div>

        <div class="shrink-0">
            <span class="text-primary-seattle-100 text-[14px] leading-[18px]">
                {{ membership.role }}
            </span>
        </div>

        <div v-if="!readonly && membership.role !== 'owner'" class="shrink-0">
            <button class="text-primary-seattle-100" type="button">
                <DotsMenuIcon class="rotate-90" :width="24" :height="24" />
            </button>
        </div>
    </div>
</template>
