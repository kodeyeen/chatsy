<script lang="ts" setup>
import { ref, reactive, computed } from 'vue'
import { useIntersectionObserver } from '@vueuse/core'

import { useAuthStore } from '@/stores/auth'
import { useChatsStore } from '@/stores/chats'
import Avatar from '@/components/Avatar.vue'
import Button from '@/components/Button.vue'
import GhostButton from '@/components/GhostButton.vue'
import Input from '@/components/Input.vue'
import Divider from '@/components/Divider.vue'
import UserOption from '@/components/UserOption.vue'
import SearchIcon from '@/components/icons/Search.vue'
import SpinnerIcon from '@/components/icons/Spinner.vue'

const props = defineProps<{
    extraData?: any
}>()

const store = useChatsStore()
const auth = useAuthStore()

const step = ref(0)
const filter = ref('')
const scrollableContainer = ref(null)
const spinner = ref(null)

const getInitialData = () => ({
    photo: null,
    title: '',
    members: [
        {
            id: 1,
            display_name: 'User Name',
        },
    ],
})

const data = reactive(getInitialData())

const isOptionDisabled = (user: any) => {
    if (user.id === auth.currentUser?.id) {
        return true
    }

    // if (!store.activeChat) {
    //   return false
    // }

    // const isMember = store.activeChat.memberships.some((membership) => {
    //   return membership.user.id === user.id
    // })

    return false
}

const isFirstStepComplete = computed(() => {
    return data.members.length !== 0
})

const isSecondStepComplete = computed(() => {
    return data.title !== ''
})

const photoUrl = computed(() => {
    if (data.photo) {
        return URL.createObjectURL(data.photo)
    }

    return undefined
})

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
const users = ref({
    results: [
        {
            id: 1,
            display_name: 'User Name',
        },
    ],
    count: 0,
    next: 1,
})

useIntersectionObserver(
    spinner,
    async ([entry], observer) => {
        if (!entry.isIntersecting) {
            return
        }

        usersOffset += usersLimit

        // const moreUsers = await ofetch('/users/users/', {
        //   // baseURL: runtimeConfig.apiBaseUrl,
        //   params: {
        //     offset: usersOffset,
        //     limit: usersLimit,
        //     display_name: filter.value,
        //   },
        // })
        const moreUsers = {
            results: [],
            count: 0,
            next: 1,
        }

        users.value.results.push(...(moreUsers.results as []))
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
    step.value = 0
    Object.assign(data, getInitialData())
}

const submit = async () => {
    let photoId = null

    if (data.photo) {
        //   const createdImage = await $imagesApi.create(data.photo)
        //   photoId = createdImage.id
    }

    const chatData = {
        photo_id: photoId,
        title: data.title,
        member_ids: data.members.map((member: any) => member.id),
    }

    if (props.extraData) {
        Object.assign(chatData, props.extraData)
    }

    // store.createGroup(chatData)

    reset()
}
</script>

<template>
    <form class="h-full" action="#" method="POST" @submit.prevent="submit()">
        <div v-show="step === 0" class="h-full">
            <div>
                <div class="text-h5 md:text-h4 text-primary-brand-primary">Select Users</div>
            </div>

            <div class="h-[calc(100%-30px-16px)] mt-[16px]">
                <div>
                    <Input
                        placeholder="Whom would you like to invite?"
                        type="search"
                        size="l"
                        variant="outlined"
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

                <div
                    ref="scrollableContainer"
                    class="h-[calc(100%-44px+100px)] pb-[100px] overflow-y-auto scrollbar pt-[16px]"
                >
                    <div v-if="false" class="flex justify-center items-center">
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
                                :disabled="isOptionDisabled(user)"
                                v-model="data.members"
                            />
                        </div>
                        <div v-if="users.next" class="flex justify-center items-center mt-[16px]">
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

            <div
                class="absolute right-[16px] md:right-[32px] left-[16px] md:left-[32px] bottom-[32px] flex justify-center items-center pointer-events-none"
            >
                <Button
                    class="min-w-[288px] pointer-events-auto"
                    type="button"
                    size="l"
                    variant="primary"
                    :disabled="!isFirstStepComplete"
                    round
                    @click="++step"
                >
                    <span> Next </span>
                </Button>
            </div>
        </div>

        <div v-show="step === 1" class="h-full">
            <div>
                <div class="text-h5 md:text-h4 text-primary-brand-primary">New Group</div>
            </div>

            <div class="h-[calc(100%-30px-16px)] mt-[16px]">
                <div class="flex flex-col md:flex-row gap-x-[32px] gap-y-[12px]">
                    <div class="shrink-0 flex items-center w-[120px]">
                        <span class="block text-m-14 text-primary-seattle-100"> Avatar </span>
                    </div>

                    <div class="grow">
                        <GhostButton
                            class="max-w-max !justify-start !gap-x-[20px]"
                            type="file"
                            size="l-regular"
                            variant="link"
                            @change="($event: any) => (data.photo = $event.target.files[0])"
                        >
                            <Avatar class="w-[88px] h-[88px] rounded-full" :src="photoUrl" />

                            <span> Upload an avatar </span>
                        </GhostButton>
                    </div>
                </div>

                <div class="flex flex-col md:flex-row gap-x-[32px] gap-y-[4px] mt-[32px]">
                    <div class="shrink-0 flex items-center w-[120px]">
                        <div class="text-m-14 text-primary-seattle-100">Group name</div>
                    </div>

                    <div class="grow">
                        <div>
                            <Input
                                type="text"
                                size="l"
                                variant="outlined"
                                placeholder="Group name"
                                v-model="data.title"
                            />
                        </div>
                    </div>
                </div>

                <Divider class="mt-[16px]" />

                <div
                    class="h-[calc(100%-120px-32px-68px-16px-1px+100px)] md:h-[calc(100%-88px-32px-44px-16px-1px+100px)] pt-[16px] pb-[100px] overflow-y-auto scrollable"
                >
                    <div v-for="user in data.members" class="flex items-center gap-x-[9px] p-[9px]">
                        <div class="shrink-0">
                            <RouterLink class="block" to="/chats">
                                <Avatar class="w-[44px] h-[44px] rounded-full" />
                            </RouterLink>
                        </div>

                        <div class="grow">
                            <RouterLink
                                class="block max-w-max text-primary-brand-onPrimary text-l-long-16 font-medium"
                                to="/chats"
                            >
                                <span>
                                    {{ user.display_name }}
                                </span>
                            </RouterLink>
                        </div>
                    </div>
                </div>
            </div>

            <div
                class="absolute right-[16px] md:right-[32px] left-[16px] md:left-[32px] bottom-[32px] flex justify-center items-center pointer-events-none"
            >
                <Button
                    class="min-w-[288px] pointer-events-auto"
                    type="submit"
                    size="l"
                    variant="primary"
                    :disabled="!isSecondStepComplete"
                    round
                    data-popup-closer
                >
                    <span> Create </span>
                </Button>
            </div>
        </div>
    </form>
</template>
