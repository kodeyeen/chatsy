<script lang="ts" setup>
import { ref } from 'vue'
import { useIntersectionObserver } from '@vueuse/core'
import { ofetch } from 'ofetch'

import { useAuthStore } from '@/stores/auth'
import { useChatsStore } from '@/stores/chats'
import Button from '@/components/Button.vue'
import Input from '@/components/Input.vue'
import UserOption from '@/components/UserOption.vue'
import SearchIcon from '@/components/icons/Search.vue'
import SpinnerIcon from '@/components/icons/Spinner.vue'

const props = defineProps<{
    chat: any
    memberships: any
}>()

const auth = useAuthStore()
const store = useChatsStore()

const filter = ref('')
const scrollableContainer = ref(null)
const spinner = ref(null)

const getInitialData = () => ({
    members: [],
})

const data = ref(getInitialData())

const isOptionDisabled = (user: any) => {
    if (user.id === auth.currentUser?.id) {
        return true
    }

    const isMember = props.memberships.results.some((membership: any) => {
        return membership.user.id === user.id
    })

    return isMember
}

let usersOffset = 0
const usersLimit = 10
//   const { data: users, pending } = useLazyFetch('/users/users/', {
//     // baseURL: runtimeConfig.apiBaseUrl,
//     params: {
//       offset: usersOffset,
//       limit: usersLimit,
//       display_name: filter,
//     },
//   })
const users = ref<any>({
    results: [
        {
            id: 1,
            display_name: 'Test User',
        },
    ],
    count: 0,
    next: 1,
})
const pending = false

useIntersectionObserver(
    spinner,
    async ([entry], observer) => {
        if (!entry.isIntersecting) {
            return
        }

        usersOffset += usersLimit

        const moreUsers = await ofetch('/users/users/', {
            //   baseURL: runtimeConfig.apiBaseUrl,
            params: {
                offset: usersOffset,
                limit: usersLimit,
                display_name: filter.value,
            },
        })

        users.value.results.push(...moreUsers.results)
        users.value.count = moreUsers.count
        users.value.next = moreUsers.next

        if (!moreUsers.next) {
            observer.unobserve(entry.target)
        }
    },
    {
        root: scrollableContainer,
    },
)

const reset = () => {
    filter.value = ''
    Object.assign(data, getInitialData())
}

const submit = async () => {
    const memberIds = data.value.members.map((member: any) => member.id)

    // store.addMembers(props.chat.id, memberIds)

    reset()
}
</script>

<template>
    <form class="flex flex-col h-full" action="#" method="POST" @submit.prevent="submit()">
        <div class="shrink-0">
            <div class="text-h5 md:text-h4 text-primary-brand-primary">Select Users</div>
        </div>

        <div class="grow flex flex-col mt-[16px]">
            <div class="shrink-0">
                <Input
                    type="search"
                    size="l"
                    variant="outlined"
                    placeholder="Whom would you like to invite?"
                    v-model="filter"
                >
                    <template #start>
                        <SearchIcon
                            class="absolute top-1/2 left-[12px] -translate-y-1/2 text-primary-seattle-60"
                            :width="24"
                            :height="24"
                        />
                    </template>
                </Input>
            </div>

            <div ref="scrollableContainer" class="relative grow">
                <div
                    class="absolute top-0 bottom-[-32px] left-0 w-full pt-[16px] pb-[100px] overflow-x-hidden overflow-y-auto scrollbar"
                >
                    <div v-if="pending" class="flex justify-center items-center">
                        <SpinnerIcon
                            class="text-primary-brand-accent animate-spin"
                            :width="32"
                            :height="32"
                        />
                    </div>
                    <div v-else>
                        <div>
                            <UserOption
                                v-for="user in users.results"
                                :key="user.id"
                                class="mt-[16px] first:mt-0"
                                :user="user"
                                v-model="data.members"
                                :disabled="isOptionDisabled(user)"
                            />
                        </div>
                        <div v-if="false" class="flex justify-center items-center mt-[16px]">
                            <SpinnerIcon
                                ref="spinner"
                                class="text-primary-brand-accent animate-spin"
                                :width="32"
                                :height="32"
                            />
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div
            class="absolute right-[16px] md:right-[32px] left-[16px] md:left-[32px] bottom-[32px] flex justify-center items-center pointer-events-none"
        >
            <Button
                class="min-w-[288px] pointer-events-auto"
                size="l"
                variant="primary"
                type="submit"
                round
                data-popup-closer
            >
                <span> Add </span>
            </Button>
        </div>
    </form>
</template>
