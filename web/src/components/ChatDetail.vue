<script lang="ts" setup>
import { ref } from 'vue'

import { useAuthStore } from '@/stores/auth'
import { useChatsStore } from '@/stores/chats'
import { usePopupsStore } from '@/stores/popups'
import Avatar from '@/components/Avatar.vue'
import Button from '@/components/Button.vue'
import GhostButton from '@/components/GhostButton.vue'
import ChatMembership from '@/components/ChatMembership.vue'
import ChatSection from '@/components/ChatSection.vue'
import ChatTopBar from '@/components/ChatTopBar.vue'
import ChatAddMembersForm from '@/components/ChatAddMembersForm.vue'
import Popup from '@/components/Popup.vue'
import ArrowLeftSmallIcon from '@/components/icons/ArrowLeftSmall.vue'
import LoginEnterArrowIcon from '@/components/icons/LoginEnterArrow.vue'
import SpinnerIcon from '@/components/icons/Spinner.vue'
import SingleUserAddPlusFillIcon from '@/components/icons/SingleUserAddPlusFill.vue'
import MuteIcon from '@/components/icons/Mute.vue'

const props = defineProps<{
    chat: any
}>()

const auth = useAuthStore()
const store = useChatsStore()
const popupsStore = usePopupsStore()

const memberships = ref<any>({
    results: [
        {
            user: {
                display_name: 'Ruslan Iskandarov',
            },
            role: 'admin',
        },
        {
            user: {
                display_name: 'Henry Cavill',
            },
            role: 'member',
        },
        {
            user: {
                display_name: 'Brad Pitt',
            },
            role: 'member',
        },
    ],
})
const isFetchingMemberships = ref(false)
const spinner = ref(null)
</script>

<template>
    <div class="flex flex-col h-full bg-primary-brand-wildSand">
        <ChatTopBar class="shrink-0 flex justify-between px-[20px] py-[16px]">
            <div>
                <GhostButton
                    type="button"
                    size="l-medium"
                    variant="link"
                    @click="store.currentState = store.State.CHAT_MESSAGES"
                >
                    <ArrowLeftSmallIcon :width="24" :height="24" />

                    <span> Back </span>
                </GhostButton>
            </div>

            <div>
                <span class="text-l-long-16 font-medium text-primary-brand-onPrimary">
                    Details
                </span>
            </div>

            <div>
                <GhostButton
                    xv-if="auth.currentUser"
                    :class="{
                        // 'invisible': auth.currentUser.id !== chat.owner_id,
                    }"
                    type="button"
                    size="l-medium"
                    variant="link"
                    @click="store.currentState = store.State.CHAT_EDIT_FORM"
                >
                    <span> Edit </span>
                </GhostButton>
            </div>
        </ChatTopBar>

        <div class="relative grow">
            <div class="absolute top-0 left-0 w-full h-full overflow-y-auto scrollbar">
                <div class="max-w-[720px] mx-auto my-[40px]">
                    <div class="flex flex-col items-center text-center">
                        <Avatar class="w-[150px] h-[150px] rounded-full" />

                        <div class="mt-[17px]">
                            <div class="text-primary-brand-onPrimary text-l-long-16 font-bold">
                                {{ chat.title }}
                            </div>

                            <div class="text-primary-seattle-100 text-[14px] leading-[18px]">
                                <!-- {{ $tc('member', chat.member_count) }} -->
                                5 participants
                            </div>
                        </div>
                    </div>

                    <div class="flex flex-col md:flex-row gap-x-[12px] gap-y-[6px] mt-[40px]">
                        <Button
                            xv-if="auth.currentUser"
                            class="grow shrink-0 gap-x-[8px]"
                            type="button"
                            size="l"
                            variant="clear"
                            xdisabled="chat.owner_id !== auth.currentUser?.id"
                            @click="popupsStore.show('addMembersPopup')"
                        >
                            <SingleUserAddPlusFillIcon :width="24" :height="24" />

                            <span> Add </span>
                        </Button>

                        <Button
                            class="grow shrink-0 gap-x-[8px]"
                            type="button"
                            size="l"
                            variant="clear"
                            click="onNotificationsToggle"
                        >
                            <MuteIcon
                                v-if="!chat.areNotificationsEnabled"
                                :width="24"
                                :height="24"
                            />

                            <span>
                                {{
                                    chat.areNotificationsEnabled
                                        ? 'Sound enabled'
                                        : 'Sound disabled'
                                }}
                            </span>
                        </Button>

                        <Button
                            xv-if="auth.currentUser"
                            class="grow shrink-0 gap-x-[8px]"
                            type="button"
                            size="l"
                            variant="clear"
                            :disabled="auth.currentUser?.id === chat.owner_id"
                            click="onGroupLeave"
                        >
                            <LoginEnterArrowIcon :width="24" :height="24" />

                            <span> Leave </span>
                        </Button>
                    </div>

                    <ChatSection v-if="chat.description" class="mt-[24px]" title="Details">
                        <div class="text-m-14 text-primary-brand-primary">
                            {{ chat.description }}
                        </div>
                    </ChatSection>

                    <ChatSection class="mt-[24px]" title="Members">
                        <div
                            v-if="isFetchingMemberships"
                            class="flex justify-center items-center h-full"
                        >
                            <SpinnerIcon
                                class="text-primary-brand-accent animate-spin"
                                :width="32"
                                :height="32"
                            />
                        </div>
                        <div v-else-if="memberships">
                            <div>
                                <ChatMembership
                                    v-for="membership in memberships.results"
                                    :key="membership.id"
                                    :membership="membership"
                                    readonly
                                />
                            </div>
                            <!-- <div v-if="memberships.has_more" class="flex justify-center items-center my-[16px]">
                  <i ref="spinner" class="icon dd-Spiner-l text-[32px] text-primary-brand-accent animate-spin"></i>
                </div> -->
                        </div>
                    </ChatSection>
                </div>
            </div>
        </div>

        <Teleport to="#popupsContainer">
            <Popup
                id="addMembersPopup"
                contentClass="overflow-hidden"
                dialogClass="md:max-w-[600px] min-h-[658px]"
            >
                <ChatAddMembersForm v-if="memberships" :chat="chat" :memberships="memberships" />
            </Popup>
        </Teleport>
    </div>
</template>
