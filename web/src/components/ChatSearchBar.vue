<script lang="ts" setup>
import { ref, computed, watch } from 'vue'
import { Tippy } from 'vue-tippy'

import { useChatsStore } from '@/stores/chats'
import { usePopupsStore } from '@/stores/popups'

import Button from '@/components/Button.vue'
import OptionButton from '@/components/OptionButton.vue'
import OptionsMenu from '@/components/OptionsMenu.vue'
import OptionItem from '@/components/OptionItem.vue'
import Input from '@/components/Input.vue'
import Popup from '@/components/Popup.vue'
import ChatCreateForm from '@/components/ChatCreateForm.vue'
import DotsMenuIcon from '@/components/icons/DotsMenu.vue'
import SearchIcon from '@/components/icons/Search.vue'
import SettingsEditFilterGearIcon from '@/components/icons/SettingsEditFilterGear.vue'
import UsersIcon from '@/components/icons/Users.vue'

const props = defineProps<{
    filter: string
}>()

const emit = defineEmits(['update:filter'])

const model = computed({
    get() {
        return props.filter
    },

    set(newValue) {
        emit('update:filter', newValue)
    },
})

const popupsStore = usePopupsStore()

const filter = ref('')

watch(filter, (newFilter) => {
    model.value = newFilter
})
</script>

<template>
    <div>
        <div class="flex justify-between px-[16px] py-[12px] bg-primary-brand-wildSand">
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

        <div class="px-[16px] pt-[8px] pb-[10px] border-b border-b-primary-london-110 bg-primary-brand-white">
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
