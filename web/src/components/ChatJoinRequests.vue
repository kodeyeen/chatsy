<script lang="ts" setup>
import Button from '@/components/Button.vue'
import GhostButton from '@/components/GhostButton.vue'
import ChatSection from '@/components/ChatSection.vue'
import ChatTopBar from '@/components/ChatTopBar.vue'
import ChatJoinRequest from '@/components/ChatJoinRequest.vue'
import Divider from '@/components/Divider.vue'

import { useChatsStore } from '@/stores/chats'

const store = useChatsStore()
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
                    <span> Close </span>
                </GhostButton>
            </div>

            <div>
                <span class="text-l-long-16 font-medium text-primary-brand-onPrimary">
                    Join Requests
                </span>
            </div>

            <div class="invisible">Close</div>
        </ChatTopBar>

        <div class="relative grow">
            <div class="absolute top-0 left-0 w-full h-full overflow-y-auto scrollbar">
                <div class="max-w-[720px] mx-auto my-[32px]">
                    <ChatSection
                        v-if="true"
                        title="3 join requests"
                        subtitle="Join by request mode is on"
                    >
                        <div>
                            <ChatJoinRequest />
                            <ChatJoinRequest />
                            <ChatJoinRequest />
                            <ChatJoinRequest />
                        </div>

                        <Divider class="my-[16px]" />

                        <div class="flex gap-x-[8px] pl-[62px]">
                            <Button class="rounded-full" type="button" size="s" variant="primary">
                                <span> Approve all </span>
                            </Button>
                            <Button class="rounded-full" type="button" size="s" variant="clear">
                                <span> Refuse all </span>
                            </Button>
                        </div>
                    </ChatSection>
                    <div v-else class="flex flex-col gap-y-[16px] items-center text-center">
                        <img src="@/assets/images/thumbs-up.png" alt="Thumbs Up" />

                        <div class="text-primary-brand-primary text-h5">No requests</div>

                        <div class="text-m-14 text-primary-seattle-100">
                            There are no requests to join the group
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
