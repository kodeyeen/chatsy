<script lang="ts" setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { Tippy } from 'vue-tippy'
import { useVModel } from '@vueuse/core'

import { useChatsStore } from '@/stores/chats'
import { usePopupsStore } from '@/stores/popups'
import { useFetch } from '@/services/fetch'

import type { Page } from '@/models/api'
import type { Chat } from '@/models/chat'

import Button from '@/components/Button.vue'
import ChatCard from '@/components/ChatCard.vue'
import ChatCreateForm from '@/components/ChatCreateForm.vue'
import Input from '@/components/Input.vue'
import OptionButton from '@/components/OptionButton.vue'
import OptionsMenu from '@/components/OptionsMenu.vue'
import OptionItem from '@/components/OptionItem.vue'
import Popup from '@/components/Popup.vue'
import SearchIcon from '@/components/icons/Search.vue'
import SpinnerIcon from '@/components/icons/Spinner.vue'
import DotsMenuIcon from '@/components/icons/DotsMenu.vue'
import SettingsEditFilterGearIcon from '@/components/icons/SettingsEditFilterGear.vue'
import UsersIcon from '@/components/icons/Users.vue'

const props = defineProps<{
    activeChat: Chat | null,
}>()

const emit = defineEmits(['update:activeChat'])
const activeChat = useVModel(props, 'activeChat', emit)

const chatsStore = useChatsStore()
const popupsStore = usePopupsStore()
const router = useRouter()

const { data: chats } = useFetch('/chats/mine?limit=10&offset=0').get().json<Page<Chat>>()

const title = ref('')

const onChatSelect = (chat: Chat): void => {
    if (activeChat.value && activeChat.value.id === chat.id) {
        return
    }

    chatsStore.currentState = chatsStore.State.CHAT_MESSAGES

    router.push({
        hash: `#${chat.id}`,
        query: {},
    })

    activeChat.value = chat
}

// const chats = ref<any>({
//     items: [
//         {
//             id: 1,
//             title: 'First Chat',
//             description: 'descr',
//             type: 'group',
//             areNotificationsEnabled: false,
//             participantCount: 3,
//             isJoined: true,
//             owner_id: 1,
//             join_request_count: 12,
//             lastMessage: {
//                 text: "i'm the last",
//                 sentAt: new Date().toString(),
//             },
//         },
//         {
//             id: 2,
//             title: 'Second Chat',
//             description: 'Kek',
//             type: 'group',
//             areNotificationsEnabled: false,
//             participantCount: 5,
//             isJoined: true,
//             owner_id: 1,
//             join_request_count: 0,
//             lastMessage: {
//                 text: "i'm the second",
//                 sentAt: new Date().toString(),
//             },
//         },
//     ],
// })
// const activeChat = ref<any>(chats.value.items[0])
</script>

<template>
    <div class="flex flex-col h-full">
        <div class="shrink-0 flex justify-between px-[16px] py-[12px] bg-primary-brand-wildSand">
            <div class="flex ml-auto">
                <Tippy
                    theme="none"
                    placement="bottom-end"
                    trigger="click"
                    :arrow="false"
                    :duration="[0, 0]"
                    :offset="[0, 4]"
                    :zIndex="35"
                    interactive
                >
                    <Button
                        class="w-[36px] h-[36px] !p-0"
                        type="button"
                        size="m"
                        variant="ghost"
                        @click="popupsStore.show('sidebarMenuPopup')"
                    >
                        <DotsMenuIcon :width="20" :height="20" />
                    </Button>

                    <template #content="{ hide }">
                        <OptionsMenu class="hidden md:block min-w-[260px] max-h-max">
                            <OptionItem @click="hide()">
                                <OptionButton
                                    class="gap-x-[6px]"
                                    text="Create group"
                                    type="button"
                                    @click="popupsStore.show('groupCreationPopup')"
                                >
                                    <template #start>
                                        <UsersIcon :width="20" :height="20" />
                                    </template>
                                </OptionButton>
                            </OptionItem>

                            <OptionItem>
                                <OptionButton
                                    class="gap-x-[6px]"
                                    text="Settings"
                                    type="link"
                                    to="/chats"
                                >
                                    <template #start>
                                        <SettingsEditFilterGearIcon :width="20" :height="20" />
                                    </template>
                                </OptionButton>
                            </OptionItem>
                        </OptionsMenu>
                    </template>
                </Tippy>
            </div>
        </div>

        <div class="shrink-0 px-[16px] pt-[8px] pb-[10px] border-b border-b-primary-london-110 bg-primary-brand-white">
            <Input type="search" size="m" variant="filled" placeholder="Search" xv-model="filter">
                <template #start>
                    <SearchIcon
                        class="absolute top-[8px] left-[12px] text-primary-seattle-100 pointer-events-none"
                        :width="20"
                        :height="20"
                    />
                </template>
            </Input>
        </div>

        <div class="relative grow bg-primary-brand-white">
            <div
                ref="scrollableContainer"
                class="absolute top-0 left-0 w-full h-full overflow-x-hidden overflow-y-auto scrollbar p-[8px]"
            >
                <div class="h-full">
                    <div v-if="chats">
                        <div>
                            <ChatCard
                                v-for="chat in chats.items"
                                :key="chat.id"
                                :chat="chat"
                                :active="activeChat !== null && activeChat.id === chat.id"
                                @click="onChatSelect(chat)"
                            />
                        </div>

                        <div v-if="chats.offset + chats.limit < chats.count" class="flex justify-center items-center my-[16px]">
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

        <Teleport to="#popupsContainer">
            <Popup id="sidebarMenuPopup" class="md:!hidden" contentClass="pt-[10px] px-[6px]">
                <OptionsMenu class="max-h-max p-0 shadow-none">
                    <OptionItem>
                        <OptionButton
                            class="gap-x-[6px]"
                            text="Create group"
                            type="button"
                            @click="
                                popupsStore.hide('sidebarMenuPopup');
                                popupsStore.show('groupCreationPopup')
                            "
                        >
                            <template #start>
                                <UsersIcon :width="20" :height="20" />
                            </template>
                        </OptionButton>
                    </OptionItem>

                    <OptionItem>
                        <OptionButton class="gap-x-[6px]" text="Settings" type="link" to="/chats">
                            <template #start>
                                <SettingsEditFilterGearIcon :width="20" :height="20" />
                            </template>
                        </OptionButton>
                    </OptionItem>
                </OptionsMenu>

                <div class="mt-[16px] px-[10px]">
                    <Button
                        class="w-full rounded-full"
                        type="button"
                        size="m"
                        variant="clear"
                        @click="popupsStore.hide('sidebarMenuPopup')"
                    >
                        <span> Close </span>
                    </Button>
                </div>
            </Popup>

            <Popup
                id="groupCreationPopup"
                dialogClass="md:max-w-[600px] md:max-h-[658px]"
                contentClass="pb-[100px] overflow-y-hidden"
            >
                <ChatCreateForm />
            </Popup>
        </Teleport>
    </div>
</template>
