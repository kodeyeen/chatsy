<script lang="ts" setup>
import { computed, reactive, watch } from 'vue'

import { useAuthStore } from '@/stores/auth'
import Button from '@/components/Button.vue'
import Checkbox from '@/components/Checkbox.vue'

const props = defineProps<{
    messages: any[]
}>()

const emit = defineEmits(['submit', 'cancel'])

const auth = useAuthStore()

const hasOthersMessages = computed(() => {
    return props.messages.some((message: any) => message.author?.id !== auth.currentUser?.id)
})

const data = reactive({
    deleteForEveryone: true,
})

watch(hasOthersMessages, (newValue) => {
    data.deleteForEveryone = !newValue
})

const submit = () => {
    emit('submit', data.deleteForEveryone)
}

const onCancel = () => {
    emit('cancel')
}
</script>

<template>
    <form action="#" method="POST" @submit.prevent="submit()">
        <div class="flex flex-col items-center text-center">
            <svg
                width="52"
                height="52"
                viewBox="0 0 52 52"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
            >
                <g clip-path="url(#clip0_11_19710)">
                    <path
                        d="M0 24.96C0 13.1937 0 7.31062 3.65531 3.65531C7.31062 0 13.1937 0 24.96 0H27.04C38.8063 0 44.6893 0 48.3447 3.65531C52 7.31062 52 13.1937 52 24.96V27.04C52 38.8063 52 44.6893 48.3447 48.3447C44.6893 52 38.8063 52 27.04 52H24.96C13.1937 52 7.31062 52 3.65531 48.3447C0 44.6893 0 38.8063 0 27.04V24.96Z"
                        fill="url(#paint0_radial_11_19710)"
                    />
                    <path
                        d="M25.8604 26.1899C25.8604 28.951 23.6098 31.16 20.8746 31.16C18.1393 31.16 15.8888 28.951 15.8888 26.1899C15.8888 23.4287 18.1047 21.2197 20.8746 21.2197C23.6445 21.2197 25.8604 23.4632 25.8604 26.1899Z"
                        fill="white"
                    />
                    <path
                        d="M23.9009 22.2398C22.5106 21.9157 20.3912 21.8482 19.1436 23.8774C17.3745 26.7549 19.9457 30.6031 23.6399 30.325C22.8465 30.8531 21.8943 31.16 20.8747 31.16C18.1394 31.16 15.8889 28.951 15.8889 26.1899C15.8889 23.4287 18.1048 21.2197 20.8747 21.2197C22.0164 21.2197 23.0639 21.6008 23.9009 22.2398Z"
                        fill="#E9EAEA"
                    />
                    <path
                        d="M25.7913 40.4446H19.109C17.5856 40.4446 16.3391 39.202 16.3391 37.6834C16.3391 36.1647 17.5856 34.9222 19.109 34.9222H25.7567C30.7079 34.9222 34.7242 30.9185 34.7242 25.9829C34.7242 21.0818 30.7079 17.078 25.7913 17.078H19.109C17.5856 17.078 16.3391 15.8355 16.3391 14.3168C16.3391 12.7982 17.5856 11.5557 19.109 11.5557H25.7567C33.7548 11.5557 40.264 18.0444 40.264 26.0174C40.264 33.9558 33.7894 40.4446 25.7913 40.4446Z"
                        fill="white"
                    />
                    <path
                        d="M26.581 17.1128C31.8499 17.1463 36.1438 21.4134 36.1438 26.7077C36.1438 32.8168 31.1581 37.787 25.0297 37.787H16.7547C16.6133 37.787 16.4747 37.7741 16.3399 37.7494C16.3394 37.7275 16.3391 37.7055 16.3391 37.6835C16.3391 36.1648 17.5856 34.9223 19.109 34.9223H25.7567C30.7079 34.9223 34.7242 30.9186 34.7242 25.9829C34.7242 21.347 31.1308 17.514 26.581 17.1128Z"
                        fill="#E9EAEA"
                    />
                </g>
                <defs>
                    <radialGradient
                        id="paint0_radial_11_19710"
                        cx="0"
                        cy="0"
                        r="1"
                        gradientUnits="userSpaceOnUse"
                        gradientTransform="translate(-0.810073 7.29063) rotate(40.2798) scale(62.6484 62.6484)"
                    >
                        <stop stop-color="#FF6712" />
                        <stop offset="1" stop-color="#FF5132" />
                    </radialGradient>
                    <clipPath id="clip0_11_19710">
                        <rect width="52" height="52" fill="white" />
                    </clipPath>
                </defs>
            </svg>

            <div class="mt-[20px] text-primary-brand-primary text-l-long-16 font-bold">
                {{
                    messages.length === 1
                        ? 'Do you want to delete selected message?'
                        : 'Do you want to delete selected messages?'
                }}
            </div>
        </div>

        <div class="flex flex-col gap-y-[8px] mt-[20px]">
            <Button type="submit" size="m" variant="accent" round data-popup-closer>
                <span> Delete </span>
            </Button>

            <Button
                size="m"
                variant="clear"
                type="button"
                round
                data-popup-closer
                @click="onCancel"
            >
                <span> Cancel </span>
            </Button>
        </div>

        <div class="flex justify-center mt-[20px]">
            <Checkbox
                label="Delete for everyone"
                size="m"
                v-model="data.deleteForEveryone"
                :disabled="hasOthersMessages"
            />
        </div>
    </form>
</template>
