<script lang="ts" setup>
import { ref, reactive, computed } from 'vue'

import { useChatsStore } from '@/stores/chats'
import { usePopupsStore } from '@/stores/popups'
import Avatar from '@/components/Avatar.vue'
import Button from '@/components/Button.vue'
import GhostButton from '@/components/GhostButton.vue'
import CopyText from '@/components/CopyText.vue'
import ChatFormRow from '@/components/ChatFormRow.vue'
import ChatMembership from '@/components/ChatMembership.vue'
import ChatSection from '@/components/ChatSection.vue'
import ChatTopBar from '@/components/ChatTopBar.vue'
import Divider from '@/components/Divider.vue'
import Textarea from '@/components/Textarea.vue'
import Switcher from '@/components/Switcher.vue'
import Radio from '@/components/Radio.vue'
import Input from '@/components/Input.vue'
import PopupConfirm from '@/components/PopupConfirm.vue'
import PopupShare from '@/components/PopupShare.vue'
import ArrowLeftSmallIcon from '@/components/icons/ArrowLeftSmall.vue'
import DotsMenuIcon from '@/components/icons/DotsMenu.vue'

const props = defineProps<{
    chat: any
}>()

const store = useChatsStore()
const popupsStore = usePopupsStore()

const staffMemberships: any = ref({
    results: [
        {
            user: {
                id: 1,
                display_name: 'Ruslan Iskandarov',
            },
            role: 'admin',
        },
    ],
})

const data = reactive({
    photo: null,
    title: props.chat.title,
    description: props.chat.description || '',
    accessibility: props.chat.accessibility,
    joinByRequest: props.chat.join_by_request,
    inviteHash: props.chat.invite_hash,
    ownerAndModerators: [],
})

const photoUrl = computed(() => {
    if (data.photo) {
        return URL.createObjectURL(data.photo)
    } else if (props.chat.photo) {
        //   return $media(props.chat.photo.url)
    }

    return undefined
})

const inviteLink = computed(() => {
    return `domain/+${props.chat.invite_hash}`
})

const submit = async () => {
    // let photoId = props.chat.photo ? props.chat.photo.id : null
    // if (data.photo) {
    // //   const createdImage = await $imagesApi.create(data.photo)
    // //   photoId = createdImage.id
    // }
    // store.updateGroup(props.chat.id, {
    //   photo_id: photoId,
    //   title: data.title,
    //   description: data.description || null,
    //   accessibility: data.accessibility,
    //   join_by_request: data.joinByRequest,
    // })
}

const onChatDelete = () => {
    // store.deleteChat(props.chat.id)
}
</script>

<template>
    <form
        class="flex flex-col h-full bg-primary-brand-wildSand"
        action="#"
        method="POST"
        @submit.prevent="submit()"
    >
        <ChatTopBar class="shrink-0 flex justify-between px-[20px] py-[16px]">
            <div>
                <GhostButton
                    type="button"
                    size="l-medium"
                    variant="link"
                    @click="store.currentState = store.State.CHAT_DETAIL"
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
                <GhostButton type="submit" size="l-medium" variant="link">
                    <span> Save </span>
                </GhostButton>
            </div>
        </ChatTopBar>

        <div class="relative grow">
            <div class="absolute top-0 left-0 w-full h-full overflow-y-auto scrollbar">
                <div class="max-w-[720px] mx-auto my-[32px]">
                    <ChatSection title="General">
                        <ChatFormRow>
                            <template #label>
                                <span class="block text-m-14 text-primary-seattle-100">
                                    Group Avatar
                                </span>
                            </template>

                            <template #entry>
                                <GhostButton
                                    class="max-w-max !justify-start !gap-x-[20px]"
                                    type="file"
                                    size="l-regular"
                                    variant="link"
                                    @change="($event: any) => (data.photo = $event.target.files[0])"
                                >
                                    <Avatar
                                        class="w-[88px] h-[88px] rounded-full"
                                        :src="photoUrl"
                                    />

                                    <span> Upload an avatar </span>
                                </GhostButton>
                            </template>
                        </ChatFormRow>

                        <ChatFormRow class="mt-[32px]">
                            <template #label>
                                <span class="block text-m-14 text-primary-seattle-100">
                                    Group Name
                                </span>
                            </template>

                            <template #entry>
                                <div>
                                    <Input
                                        type="text"
                                        size="l"
                                        variant="outlined"
                                        v-model="data.title"
                                    />
                                </div>
                            </template>
                        </ChatFormRow>

                        <div class="mt-[32px]">
                            <span class="block text-m-14 text-primary-seattle-100">
                                Description
                            </span>

                            <Textarea
                                class="xl outlined min-h-[92px] mt-[4px]"
                                :maxlength="150"
                                :rows="2"
                                v-model="data.description"
                            ></Textarea>

                            <span class="block text-m-14 text-primary-seattle-100">
                                You can add a group description
                            </span>
                        </div>
                    </ChatSection>

                    <ChatSection class="mt-[24px]" title="Group Type">
                        <div>
                            <div class="flex flex-col gap-y-[8px]">
                                <Radio
                                    text="Public"
                                    name="type"
                                    value="public"
                                    size="l"
                                    v-model="data.accessibility"
                                />

                                <Radio
                                    text="Private"
                                    name="type"
                                    value="private"
                                    size="l"
                                    v-model="data.accessibility"
                                />
                            </div>

                            <div class="mt-[16px] text-m-14 text-primary-seattle-100">
                                Public groups are visible when searching and anyone can join
                                them;<br />
                                Private groups are joinable only by invitations or invite links.
                            </div>
                        </div>

                        <Divider class="my-[16px]" />

                        <div>
                            <div class="flex flex-col gap-y-[8px]">
                                <Switcher
                                    text="Join by request"
                                    size="l"
                                    v-model="data.joinByRequest"
                                />
                            </div>

                            <div class="mt-[16px] text-m-14 text-primary-seattle-100">
                                If you enable this option new members will be able to join the group
                                only after owner or moderator of the group will approve the request
                            </div>
                        </div>
                    </ChatSection>

                    <ChatSection class="mt-[24px]" title="Permanent Link">
                        <div class="px-[16px] py-[12px] rounded-[10px] bg-primary-london-100">
                            <Input
                                type="text"
                                size="l"
                                variant="outlined"
                                :value="inviteLink"
                                readonly
                            >
                                <template #end>
                                    <button
                                        class="absolute top-1/2 right-[10px] -translate-y-1/2 text-primary-seattle-60"
                                        type="button"
                                    >
                                        <DotsMenuIcon :width="24" :height="24" />
                                    </button>
                                </template>
                            </Input>

                            <div class="flex gap-x-[16px] mt-[12px]">
                                <CopyText v-slot="{ copy }" :source="inviteLink">
                                    <Button
                                        class="grow"
                                        type="button"
                                        size="l"
                                        variant="primary"
                                        @click="copy()"
                                    >
                                        <span> Copy </span>
                                    </Button>
                                </CopyText>

                                <Button
                                    class="grow"
                                    size="l"
                                    variant="primary"
                                    @click="popupsStore.show('groupSharePopup')"
                                >
                                    <span> Share </span>
                                </Button>
                            </div>

                            <div class="flex justify-center items-center mt-[12px]">
                                <span class="text-primary-brand-primary text-m-14 font-medium">
                                    No one has joined yet
                                </span>
                            </div>
                        </div>

                        <div class="mt-[8px] text-m-14 text-primary-seattle-100">
                            Any user can join the group via this link. You can reset the link at any
                            time.
                        </div>
                    </ChatSection>

                    <ChatSection
                        v-if="staffMemberships"
                        class="mt-[24px]"
                        title="Owner and Moderators"
                    >
                        <div>
                            <ChatMembership
                                v-for="membership in staffMemberships.results"
                                :key="membership.id"
                                :membership="membership"
                            />
                        </div>

                        <div v-if="false" class="mt-[8px]">
                            <Button type="button" size="l" variant="primary">
                                <span> Add a moderator </span>
                            </Button>
                        </div>

                        <div v-if="false" class="mt-[8px] text-m-14 text-primary-seattle-100">
                            You can add a moderator which can help you manage the group.
                        </div>
                    </ChatSection>

                    <div class="mt-[24px]">
                        <Button
                            class="w-full"
                            type="button"
                            size="l"
                            variant="critical"
                            @click="popupsStore.show('confirmChatDeletionPopup')"
                        >
                            <span> Delete group </span>
                        </Button>
                    </div>
                </div>
            </div>
        </div>

        <Teleport to="#popupsContainer">
            <PopupShare
                id="groupSharePopup"
                :title="inviteLink"
                :text="inviteLink"
                :url="`/+${data.inviteHash}`"
            />

            <PopupConfirm
                id="confirmChatDeletionPopup"
                confirmText="Delete"
                cancelText="Cancel"
                title="Delete Group"
                description="Are you sure you want to delete this group?"
                @confirm="onChatDelete"
            />
        </Teleport>
    </form>
</template>
