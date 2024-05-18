<script lang="ts" setup>
import { ref, computed, watch } from 'vue'
import { Tippy } from 'vue-tippy'

import { useChatsStore } from '@/stores/chats'

import Avatar from '@/components/Avatar.vue'
import Popup from '@/components/Popup.vue'
import ArrowDownSmallIcon from '@/components/icons/ArrowDownSmall.vue'
import ArrowLeftIcon from '@/components/icons/ArrowLeft.vue'
import DoneCheckIcon from '@/components/icons/DoneCheck.vue'
import DotsMenuIcon from '@/components/icons/DotsMenu.vue'
import RotateRefreshLoadingIcon from '@/components/icons/RotateRefreshLoading.vue'
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

const chats = useChatsStore()

const filter = ref('')

watch(filter, (newFilter) => {
    model.value = newFilter
})
</script>

<template>
    <div>
        <div class="flex justify-between px-[16px] py-[12px] bg-primary-brand-wildSand">
            <div>
                <button
                    class="button m clear gap-x-[6px] px-[12px]"
                    type="button"
                    click="navigateTo(chats.back)"
                >
                    <ArrowLeftIcon width="20" height="20" />

                    <span> Go Back </span>
                </button>
            </div>

            <div class="flex">
                <Tippy
                    v-if="false"
                    class="mr-[16px]"
                    theme="none"
                    placement="bottom-end"
                    trigger="click"
                    :arrow="false"
                    :duration="[0, 0]"
                    :offset="[0, 4]"
                    :zIndex="35"
                    interactive
                >
                    <button
                        class="button-ghost l-medium ghost-30 gap-x-[4px]"
                        type="button"
                        click="$showPopup('publisherMenuPopup')"
                    >
                        <Avatar class="w-[36px] h-[36px] rounded-full" />

                        <ArrowDownSmallIcon width="24" height="24" />
                    </button>

                    <template #content="{ hide }">
                        <div class="options-menu hidden md:block w-[288px] max-h-max">
                            <div>
                                <div class="group-heading">Написать от имени...</div>

                                <ul>
                                    <li @click="hide">
                                        <button
                                            class="option-button items-start gap-x-[6px]"
                                            type="button"
                                        >
                                            <Avatar
                                                class="shrink-0 w-[20px] h-[20px] rounded-full"
                                            />

                                            <span class="grow"> Иван Сизых </span>

                                            <DoneCheckIcon
                                                v-show="true"
                                                class="shrink-0"
                                                width="20"
                                                height="20"
                                            />
                                        </button>
                                    </li>
                                </ul>

                                <div class="p-[10px] pb-[9px]">
                                    <hr class="divider" />
                                </div>

                                <ul>
                                    <li @click="hide">
                                        <button
                                            class="option-button items-start gap-x-[6px]"
                                            type="button"
                                        >
                                            <Avatar
                                                class="shrink-0 w-[20px] h-[20px] rounded-[7px]"
                                            />

                                            <span class="grow">
                                                Клиника профессора Калинченко
                                            </span>

                                            <DoneCheckIcon
                                                v-show="false"
                                                class="shrink-0"
                                                width="20"
                                                height="20"
                                            />
                                        </button>
                                    </li>
                                </ul>
                            </div>
                        </div>
                    </template>
                </Tippy>

                <button v-if="false" class="button m accent round only-icon mr-[6px]" type="button">
                    <svg width="20" height="20">
                        <use xlink:href="#pencilIcon"></use>
                    </svg>
                </button>

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
                    <button
                        class="button m ghost only-icon"
                        type="button"
                        click="$showPopup('sidebarMenuPopup')"
                    >
                        <DotsMenuIcon width="20" height="20" />
                    </button>

                    <template #content="{ hide }">
                        <div class="options-menu hidden md:block min-w-[260px] max-h-max">
                            <ul>
                                <li @click="hide">
                                    <button
                                        class="option-button gap-x-[6px]"
                                        type="button"
                                        click="$showPopup('groupCreationPopup')"
                                    >
                                        <UsersIcon width="20" height="20" />

                                        <span> Создать группу </span>
                                    </button>
                                </li>
                                <li v-if="false" @click="hide">
                                    <button class="option-button gap-x-[6px]" type="button">
                                        <RotateRefreshLoadingIcon width="20" height="20" />

                                        <span> Сменить роль </span>
                                    </button>
                                </li>
                                <li @click="hide">
                                    <RouterLink class="option-button gap-x-[6px]" to="/account/">
                                        <SettingsEditFilterGearIcon width="20" height="20" />

                                        <span> Настройки </span>
                                    </RouterLink>
                                </li>
                            </ul>
                        </div>
                    </template>
                </Tippy>
            </div>
        </div>

        <div class="px-[16px] py-[8px] border-b border-b-primary-london-110 bg-primary-brand-white">
            <div class="relative">
                <input
                    class="input m filled !pl-[40px]"
                    type="search"
                    placeholder="Search"
                    xv-model="filter"
                />

                <SearchIcon
                    class="absolute top-[8px] left-[12px] text-primary-seattle-100 pointer-events-none"
                    width="20"
                    height="20"
                />
            </div>
        </div>

        <Teleport to="#popupsContainer">
            <Popup id="sidebarMenuPopup" class="md:hidden" contentClass="pt-[10px] px-[6px]">
                <div class="options-menu max-h-max p-0 shadow-none">
                    <ul>
                        <li>
                            <button
                                class="option-button gap-x-[6px]"
                                type="button"
                                data-popup-closer
                                click="$showPopup('groupCreationPopup')"
                            >
                                <UsersIcon width="20" height="20" />

                                <span> Создать группу </span>
                            </button>
                        </li>
                        <li v-if="false">
                            <button
                                class="option-button gap-x-[6px]"
                                type="button"
                                data-popup-closer
                            >
                                <RotateRefreshLoadingIcon width="20" height="20" />

                                <span> Сменить роль </span>
                            </button>
                        </li>
                        <li>
                            <RouterLink class="option-button gap-x-[6px]" to="/account/">
                                <SettingsEditFilterGearIcon width="20" height="20" />

                                <span> Настройки </span>
                            </RouterLink>
                        </li>
                    </ul>
                </div>

                <div class="mt-[16px] px-[10px]">
                    <button class="button m clear round w-full" type="button" data-popup-closer>
                        <span> Закрыть </span>
                    </button>
                </div>
            </Popup>

            <Popup id="publisherMenuPopup" class="md:hidden">
                <div>
                    <div class="text-h5 text-primary-brand-primary">Написать от имени...</div>
                </div>

                <div class="mt-[12px]">
                    <button class="option-button px-0" type="button" data-popup-closer>
                        <Avatar class="shrink-0 w-[36px] h-[36px] rounded-full" />

                        <span class="grow ml-[8px] mr-[6px]"> Иван Сизых </span>

                        <DoneCheckIcon v-show="true" class="shrink-0" width="20" height="20" />
                    </button>

                    <div class="pt-[20px] pb-[12px]">
                        <div class="text-caption-14 font-medium uppercase text-primary-seattle-60">
                            Организации
                        </div>
                    </div>

                    <button class="option-button px-0" type="button" data-popup-closer>
                        <Avatar class="shrink-0 w-[36px] h-[36px] rounded-[7px]" />

                        <span class="grow ml-[8px] mr-[6px]"> Клиника профессора Калинченко </span>

                        <DoneCheckIcon v-show="false" class="shrink-0" width="20" height="20" />
                    </button>
                </div>

                <div class="mt-[24px]">
                    <button class="button l clear round w-full" type="button" data-popup-closer>
                        <span> Закрыть </span>
                    </button>
                </div>
            </Popup>

            <Popup
                id="groupCreationPopup"
                dialogClass="md:max-w-[600px] md:max-h-[658px]"
                contentClass="pb-[100px] overflow-y-hidden"
            >
                <!-- <ChatGroupCreationForm/> -->
            </Popup>
        </Teleport>
    </div>
</template>
