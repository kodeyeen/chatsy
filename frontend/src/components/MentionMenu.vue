<script lang="ts" setup>
import { ref, watch, onMounted } from 'vue'
import { ofetch } from 'ofetch'
import { useEventListener } from '@vueuse/core'

import Avatar from '@/components/Avatar.vue'

const props = withDefaults(
    defineProps<{
        chat: any
        filter: string | null
        active: boolean
    }>(),
    {
        active: false,
    },
)

const emit = defineEmits(['select', 'noresults'])

const users = ref<any[]>([])
const activeUserIndex = ref(0)
const root = ref<HTMLElement | null>(null)

onMounted(() => {
    useEventListener(document, 'keydown', (event) => {
        if (!props.active) {
            return
        }

        if (event.key === 'Enter' || event.key === 'Tab') {
            event.preventDefault()

            select(users.value[activeUserIndex.value])
        }

        if (event.key === 'ArrowUp') {
            activeUserIndex.value =
                (activeUserIndex.value - 1 + users.value.length) % users.value.length
        } else if (event.key == 'ArrowDown') {
            activeUserIndex.value = (activeUserIndex.value + 1) % users.value.length
        }

        if (event.key === 'ArrowUp' || event.key === 'ArrowDown') {
            event.preventDefault()

            if (!root.value) {
                return
            }

            const buttons = root.value.querySelectorAll('button')
            const currentButton = buttons[activeUserIndex.value]

            currentButton.scrollIntoView({ behavior: 'smooth', block: 'center' })
        }
    })
})

watch(
    () => props.filter,
    async (newFilter) => {
        if (newFilter === null) {
            return
        }

        const response = await ofetch('/users/', {
            baseURL: 'http://localhost:8080',
            params: {
                chat_id: props.chat.id,
                username: newFilter,
                limit: 10,
                offset: 0,
            },
        })

        activeUserIndex.value = 0
        users.value = response.results

        if (response.results.length === 0) {
            emit('noresults')
        }
    },
)

const select = (user: any) => {
    emit('select', {
        user,
    })
}

defineExpose({
    users,
})
</script>

<template>
    <div
        ref="root"
        class="min-w-[288px] max-h-[240px] overflow-x-hidden overflow-y-auto scrollbar p-[4px] rounded-[6px] bg-primary-brand-white shadow-elevation-3"
    >
        <ul>
            <li v-for="(user, index) in users" :key="user.id">
                <button
                    class="option-button gap-x-[8px] px-[10px] py-[8px]"
                    :class="{
                        active: index === activeUserIndex,
                    }"
                    type="button"
                    @click="select(user)"
                >
                    <span class="shrink-0">
                        <Avatar
                            class="w-[36px] h-[36px] rounded-full"
                            src="user.photo ? $media(user.photo.url) : undefined"
                        />
                    </span>

                    <span class="grow flex flex-col items-start">
                        <span class="text-m-14 font-medium text-primary-brand-primary">
                            <!-- {{ user.display_name }} -->
                            Full Name
                        </span>

                        <span class="text-s-12 text-primary-seattle-120">
                            <!-- @{{ user.username }} -->
                            @username
                        </span>
                    </span>
                </button>
            </li>
        </ul>
    </div>
</template>
