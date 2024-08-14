<script lang="ts" setup>
import Avatar from '@/components/Avatar.vue'
import Button from '@/components/Button.vue'
</script>

<template>
    <div class="flex gap-x-[9px] p-[9px] pr-[12px]">
        <div class="shrink-0">
            <RouterLink to="/users/ivansizz">
                <Avatar class="w-[44px] h-[44px] rounded-full" />
            </RouterLink>
        </div>

        <div class="grow">
            <div class="flex justify-between items-center">
                <RouterLink
                    class="text-primary-brand-onPrimary text-l-long-16 font-medium"
                    to="/users/ivansizz/"
                >
                    <span> Anastasia Sizikh </span>
                </RouterLink>

                <time
                    class="text-primary-seattle-100 text-[14px] leading-[18px]"
                    datetime="2018-12-03"
                >
                    21.02.2022
                </time>
            </div>

            <div class="flex gap-x-[8px] mt-[6px]">
                <Button class="px-[11px]" type="button" size="s" variant="accent" round>
                    <span> Approve </span>
                </Button>
                <Button class="px-[11px]" type="button" size="s" variant="ghost" round>
                    <span> Refuse </span>
                </Button>
            </div>
        </div>
    </div>
</template>
