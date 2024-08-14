<script lang="ts" setup>
import { computed } from 'vue'

import { useAuthStore } from '@/stores/auth'
import Button from '@/components/Button.vue'
import CopyText from '@/components/CopyText.vue'
import Popup from '@/components/Popup.vue'
import SocialMediaLinks from '@/components/SocialMediaLinks.vue'
import AttachmentLinkIcon from '@/components/icons/AttachmentLink.vue'
import ShareArrowIcon from '@/components/icons/ShareArrow.vue'

const auth = useAuthStore()
const props = defineProps<{
    id: string
    title?: string
    text?: string
    url?: string
}>()

const entityLink = computed(() => {
    let path = props.url

    if (path) {
        return new URL(path, window.location.origin).toString()
    } else {
        return ''
    }
})

const onShareButtonClick = () => {
    if (typeof navigator.share !== 'function') {
        return
    }

    navigator.share({
        title: props.title,
        text: props.text,
        url: props.url,
    })
}
</script>

<template>
    <Popup :id="id" dialogClass="md:max-w-[400px]" contentClass="md:pb-[24px]">
        <div>
            <div class="text-h5 md:text-h4 text-center text-primary-brand-primary">Share</div>
        </div>

        <div class="mt-[20px] md:mt-[24px]">
            <div class="mt-[20px] md:mt-[24px]">
                <div
                    class="flex flex-col gap-x-[12px] gap-y-[16px]"
                    :class="{
                        'md:flex-row': $slots.actions,
                    }"
                >
                    <SocialMediaLinks
                        :links="{
                            odnoklassniki: `https://connect.ok.ru/offer?url=${entityLink}&utm_source=share2`,
                            vkontakte: `https://vk.com/share.php?url=${entityLink}&utm_source=share2`,
                            telegram: `https://t.me/share/url?url=${entityLink}&utm_source=share2`,
                            whatsapp: `https://api.whatsapp.com/send?text=${entityLink}&utm_source=share2`,
                            viber: `viber://forward?text=${entityLink}&utm_source=share2`,
                        }"
                    />

                    <div class="grow flex gap-x-[12px]">
                        <CopyText v-slot="{ copy }" :source="entityLink">
                            <Button
                                class="grow gap-x-[8px] px-[16px]"
                                type="button"
                                size="l"
                                variant="clear"
                                round
                                @click="copy()"
                            >
                                <AttachmentLinkIcon :width="24" :height="24" />

                                <span> Copy link </span>
                            </Button>
                        </CopyText>

                        <Button
                            class="md:hidden w-[44px] h-[44px] !p-0"
                            type="button"
                            size="l"
                            variant="clear"
                            round
                            @click="onShareButtonClick"
                        >
                            <ShareArrowIcon :width="24" :height="24" />
                        </Button>
                    </div>
                </div>
            </div>
        </div>

        <div v-if="$slots.actions" class="mt-[24px]">
            <slot name="actions"></slot>
        </div>

        <div class="md:hidden mt-[16px]">
            <Button class="w-full" type="button" size="l" variant="clear" round data-popup-closer>
                <span> Close </span>
            </Button>
        </div>
    </Popup>
</template>
