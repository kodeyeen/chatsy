<script setup>
import { ref } from 'vue'
import { Tippy } from 'vue-tippy'

import Avatar from '@/components/Avatar.vue'
import Popup from '@/components/Popup.vue'
import CopyButton from '@/components/CopyButton.vue'
import MessageCard from '@/components/Message/Card.vue'
import MessageForm from '@/components/Message/Form.vue'
import SpinnerIcon from '@/icons/Spinner.vue'
import DotsMenuIcon from '@/icons/DotsMenu.vue'
import CheckmarkBigCircleIcon from '@/icons/CheckmarkBigCircle.vue'
import VolumeDisabledIcon from '@/icons/VolumeDisabled.vue'
import ReplyIcon from '@/icons/Reply.vue'
import CopyIcon from '@/icons/Copy.vue'
import ShareArrowIcon from '@/icons/ShareArrow.vue'
import TrashDeleteBinIcon from '@/icons/TrashDeleteBin.vue'
import LoginEnterArrowIcon from '@/icons/LoginEnterArrow.vue'
import SearchIcon from '@/icons/Search.vue'
import PenEditWriteIcon from '@/icons/PenEditWrite.vue'
import InfoInformationCircleIcon from '@/icons/InfoInformationCircle.vue'
import DeleteDisabledBigIcon from '@/icons/DeleteDisabledBig.vue'
import LockIcon from '@/icons/Lock.vue'

const contextMenu = ref(null)
const contextMenuMessage = ref(null)

const onContextMenu = (event, message) => {
  event.preventDefault()

  contextMenuMessage.value = message

  contextMenu.value.setProps({
    getReferenceClientRect: () => ({
      width: 0,
      height: 0,
      top: event.clientY,
      bottom: event.clientY,
      left: event.clientX,
      right: event.clientX,
    })
  })

  contextMenu.value.show()
}
</script>

<template>
    <div class="relative flex h-full">
        <div class="relative bg-[#EEEAE3] grow">
            <div class="absolute top-0 left-0 w-full h-full bg-[url('./assets/images/bg.png')] bg-[length:670px] bg-[0_0] opacity-[.25] pointer-events-none"></div>

            <div class="relative flex flex-col h-full">
                <div id="messageSearchBar" class="2md:hidden" xclass="{ 'hidden': !isSearching }"></div>

                <div class="shrink-0" xclass="{ 'hidden 2md:block': isSearching }">
                    <div class="flex items-center gap-x-[16px] px-[12px] 2md:px-[20px] py-[6px] bg-primary-brand-white border-b border-b-primary-london-110">
                        <div class="shrink-0">
                            <span xv-if="chat.type === 'personal'">
                                <Avatar
                                    class="w-[44px] h-[44px] rounded-full"
                                    xsrc="chat.photo ? $media(chat.photo.url) : undefined"
                                />
                            </span>
                            <!-- <button xv-else-if="chat.type === 'group'" type="button" @click="store.currentState = store.State.GROUP_DETAIL">
                                <Avatar
                                class="w-[44px] h-[44px] rounded-full"
                                :src="chat.photo ? $media(chat.photo.url) : undefined"
                                />
                            </button> -->
                        </div>

                        <div class="grow flex flex-col items-start">
                            <div class="flex items-center gap-x-[4px]">
                                <span xv-if="chat.type === 'personal'" class="whitespace-nowrap">
                                    <span>
                                        Сын шлюхи
                                        <!-- {{ chat.title }} -->
                                    </span>
                                </span>
                                <!-- <button v-else class="whitespace-nowrap" type="button" @click="store.currentState = store.State.GROUP_DETAIL">
                                    <span>
                                        {{ chat.title }}
                                    </span>
                                </button> -->

                                <svg xv-if="!chat.are_notifications_enabled" class="text-primary-seattle-100" width="24" height="24">
                                    <use xlink:href="#muteIcon"></use>
                                </svg>
                            </div>

                            <button
                                xv-if="chat.type === 'group'"
                                class="text-primary-seattle-100 text-[14px] leading-[18px]"
                                type="button"
                                xclick="store.currentState = store.State.GROUP_DETAIL"
                            >
                                32 участника
                                <!-- {{ $tc('member', chat.member_count) }} -->
                            </button>
                        </div>

                        <div class="shrink-0 flex items-center gap-x-[8px]">
                            <button
                                class="button l ghost only-icon hidden 2md:flex"
                                type="button"
                                xclick="isSearching = true"
                            >
                                <SearchIcon class="w-[24px] h-[24px]"/>
                            </button>

                            <Tippy
                                xv-if="chat.is_joined"
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
                                    class="button m 2md:l ghost only-icon"
                                    type="button"
                                    xdisabled="isSelecting"
                                    @click="$showPopup('topBarMenuPopup')"
                                >
                                    <DotsMenuIcon class="w-[20px] h-[20px] 2md:w-[24px] 2md:h-[24px] rotate-90"/>
                                </button>

                                <template #content="{ hide }">
                                    <div
                                        v-if="true"
                                        xv-if="chat.type === 'personal'"
                                        class="options-menu hidden md:block min-w-[260px] max-h-max"
                                    >
                                        <ul>
                                            <li @click="hide">
                                                <button class="option-button gap-x-[6px]" type="button" xclick="isSelecting = true">
                                                    <CheckmarkBigCircleIcon class="w-[20px] h-[20px]"/>

                                                    <span>
                                                        Выбрать сообщения
                                                    </span>
                                                </button>
                                            </li>
                                            <li @click="hide">
                                                <button class="option-button gap-x-[6px]" type="button" xclick="onNotificationsToggle">
                                                    <VolumeDisabledIcon xv-if="!chat.are_notifications_enabled" class="w-[20px] h-[20px]"/>

                                                    <span xv-if="chat.are_notifications_enabled">
                                                        Выключить уведомления
                                                    </span>
                                                    <!-- <span v-else>
                                                        Включить уведомления
                                                    </span> -->
                                                </button>
                                            </li>
                                            <li v-if="false" @click="hide">
                                                <button class="option-button gap-x-[6px]" type="button">
                                                    <LockIcon class="w-[20px] h-[20px]"/>

                                                    <span>
                                                        Заблокировать пользователя
                                                    </span>
                                                </button>
                                            </li>
                                            <li @click="hide">
                                                <button
                                                    class="option-button gap-x-[6px] !text-utilitarian-moscow-100"
                                                    type="button"
                                                    xclick="$showPopup('confirmChatDeletionPopup')"
                                                >
                                                    <TrashDeleteBinIcon class="w-[20px] h-[20px]"/>

                                                    <span>
                                                        Удалить чат
                                                    </span>
                                                </button>
                                            </li>
                                        </ul>
                                    </div>
                                    <!-- <div v-else-if="chat.type === 'group'" class="options-menu hidden md:block min-w-[260px] max-h-max">
                                        <ul>
                                        <li v-if="chat.owner_id === currentUser.id" @click="hide">
                                            <button class="option-button gap-x-[6px]" type="button" @click="store.currentState = store.State.GROUP_EDIT_FORM">
                                            <PenEditWriteIcon class="w-[20px] h-[20px]"/>

                                            <span>
                                                Редактировать
                                            </span>
                                            </button>
                                        </li>
                                        <li @click="hide">
                                            <button class="option-button gap-x-[6px]" type="button" @click="store.currentState = store.State.GROUP_DETAIL">
                                            <InfoInformationCircleIcon class="w-[20px] h-[20px]/>

                                            <span>
                                                Информация
                                            </span>
                                            </button>
                                        </li>
                                        <li @click="hide">
                                            <button class="option-button gap-x-[6px]" type="button" @click="onNotificationsToggle">
                                            <VolumeDisabledIcon xv-if="!chat.are_notifications_enabled" class="w-[20px] h-[20px]"/>

                                            <span v-if="chat.are_notifications_enabled">
                                                Выключить уведомления
                                            </span>
                                            <span v-else>
                                                Включить уведомления
                                            </span>
                                            </button>
                                        </li>
                                        <li @click="hide">
                                            <button
                                            v-if="currentUser.id === chat.owner_id"
                                            class="option-button gap-x-[6px] text-utilitarian-moscow-100"
                                            type="button"
                                            @click="onChatDelete"
                                            >
                                            <TrashDeleteBinIcon class="w-[20px] h-[20px]"/>

                                            <span>
                                                Удалить и покинуть
                                            </span>
                                            </button>
                                            <button
                                            v-else
                                            class="option-button gap-x-[6px] text-utilitarian-moscow-100"
                                            type="button"
                                            @click="onGroupLeave"
                                            >
                                            <LoginEnterArrowIcon class="w-[20px] h-[20px]"/>

                                            <span>
                                                Покинуть группу
                                            </span>
                                            </button>
                                        </li>
                                        </ul>
                                    </div> -->
                                </template>
                            </Tippy>
                        </div>
                    </div>

                    <div xv-if="chat.type === 'group' && chat.join_request_count > 0" class="px-[20px] py-[6px] bg-primary-brand-white border-b border-b-primary-london-110">
                        <button class="button-ghost l-medium link w-full gap-x-[12px] justify-center" type="button">
                            <Avatar class="w-[28px] h-[28px] rounded-full"/>

                            <span>
                                129 заявок на вступление
                                <!-- {{ $tc('membershipRequest', chat.join_request_count) }} на вступление -->
                            </span>
                        </button>
                    </div>
                </div>

                <div class="relative grow">
                    <div
                        ref="scrollableContainer"
                        class="absolute top-0 left-0 w-full h-full flex flex-col-reverse overflow-x-hidden overflow-y-auto scrollbar pb-[10px]"
                        xscroll="scrollTop = $event.currentTarget.scrollTop"
                    >
                        <!-- <div xv-if="isFetchingMessages" class="flex justify-center items-center h-full">
                            <span class="text-primary-london-120 animate-spin">
                                <SpinnerIcon/>
                            </span>
                        </div> -->
                        <div xv-else>
                            <div xv-if="offsetEnd + limit < count" class="flex justify-center items-center py-[16px]">
                                <span ref="topSpinner" class="text-primary-london-120 animate-spin">
                                    <SpinnerIcon/>
                                </span>
                            </div>

                            <div class="flex flex-col-reverse">
                                <div v-for="i in 4" :key="i">
                                    <div class="sticky top-[10px] z-[11] flex justify-center my-[10px] pointer-events-none">
                                        <span class="pointer-events-auto px-[8px] py-[2px] rounded-[6px] text-m-14 text-primary-seattle-120 bg-primary-brand-white shadow-[0px_1px_1px_0px_rgba(0,0,0,0.16)]">
                                        <!-- {{ date }} -->
                                        29 октября
                                        </span>
                                    </div>

                                    <div class="flex flex-col-reverse">
                                        <MessageCard
                                            v-for="j in 20"
                                            :key="j"
                                            ref="messageCards"
                                            :class="{
                                                // 'own': message.author.id === currentUser.id,
                                                // 'selected': contextMenuMessage && contextMenuMessage.id === message.id,
                                            }"
                                            xmessage="message"
                                            xselectable="isSelecting"
                                            xown="message.author.id === currentUser.id"
                                            xshowUserInfo="chat.type === 'group'"
                                            xfirstInGroup="store.isFirstInGroup(messages, message, index)"
                                            xlastInGroup="store.isLastInGroup(messages, message, index)"
                                            xv-model="selectedMessages"
                                            xdata-message-id="message.id"
                                            @contextmenu="onContextMenu($event, message)"
                                            xrepliedMessageClick="onRepliedMessageClick"
                                        />
                                    </div>
                                </div>
                            </div>

                            <div xv-if="offsetStart > 0" class="flex justify-center items-center py-[16px]">
                                <span ref="bottomSpinner" class="text-primary-london-120 animate-spin">
                                    <SpinnerIcon/>
                                </span>
                            </div>
                        </div>
                    </div>
                </div>

                <div id="messageSearchNav" class="2md:hidden" xclass="{ 'hidden': !isSearching }"></div>
                <div class="shrink-0 px-[6px] 2md:px-[22px] pb-[6px] 2md:pb-[22px]" xclass="{ 'hidden 2md:block': isSearching }">
                    <!-- <div xv-if="!chat.is_joined" class="max-w-[720px] mx-auto">
                        <button class="button xl primary w-full" type="button" xclick="onGroupJoin">
                        <span>
                            Вступить в группу
                        </span>
                        </button>
                    </div> -->
                    <!-- <div xv-else-if="isSelecting" class="flex justify-between items-center gap-x-[24px] max-w-[570px] mx-auto px-[12px] py-[14px] rounded-[12px] bg-primary-brand-white">
                        <button
                            class="button-ghost l-medium ghost-70"
                            type="button"
                            xclick="selectedMessages = []; isSelecting = false"
                            >
                            <DeleteDisabledBigIcon class="w-[24px] h-[24px]/>
                        </button>

                        <span class="grow text-l-long-16 font-medium text-primary-brand-onPrimary overflow-hidden text-ellipsis whitespace-nowrap">
                            {{ $tc('message', selectedMessages.length) }}
                            5 сообщений
                        </span>

                        <button
                            xv-if="selectedMessages.length > 0"
                            class="button-ghost l-medium ghost-70 gap-x-[8px]"
                            type="button"
                            xclick="$showPopup('forwardPopup')"
                        >
                            <ShareArrowIcon class="w-[24px] h-[24px]"/>

                            <span>
                                Переслать
                            </span>
                        </button>

                        <button
                            xv-if="selectedMessages.length > 0"
                            class="button-ghost l-medium primary gap-x-[8px]"
                            type="button"
                            xclick="$showPopup('deletionConfirmationPopup')"
                        >
                            <TrashDeleteBinIcon class="w-[24px] h-[24px]"/>

                            <span>
                                Удалить
                            </span>
                        </button>
                    </div> -->
                    <MessageForm
                        xv-else
                        ref="messageForm"
                        class="mx-auto"
                        xchat="chat"
                        xmessageToReply="messageToReply"
                        xcancelReplying="messageToReply = null"
                        xtoggleEmojiPicker="isPanelActive = !isPanelActive"
                        xcloseEmojiPicker="isPanelActive = false"
                    >
                        <template #actions>
                            <button
                                class="button xl white round only-icon translate-y-[calc(100%+12px)] pointer-events-auto"
                                :class="{
                                //    'translate-y-0': Math.abs(scrollTop) > 500 || (offsetEnd > 0 && Math.abs(scrollTop) > 500),
                                }"
                                type="button"
                                xclick="onGoDown"
                            >
                                <ArrowDownIcon class="w-[24px] h-[24px]"/>
                            </button>
                        </template>
                    </MessageForm>
                </div>

                <div
                    id="emojiPanel"
                    class="md:hidden h-0 overflow-hidden transition-[height]"
                    :class="{
                        // 'h-[275px]': isPanelActive,
                    }"
                ></div>
            </div>
        </div>

        <!-- <div
            class="shrink-0 hidden 2md:block absolute xl:static top-0 right-0 z-[12] w-full h-full max-w-0 transition-[max-width] duration-[350ms] border-t border-l border-t-primary-london-110 border-l-primary-london-110"
            :class="{
                '!max-w-[350px]': isSearching,
            }"
            >
            <ChatMessageSearchResults
                :chat="chat"
                @messageSelect="onSearchResultSelect"
                @close="isSearching = false"
            />
        </div> -->

        <Teleport to="#popupsContainer">
            <Tippy
                ref="contextMenu"
                placement="right-start"
                trigger="manual"
                interactive
                :arrow="false"
                :offset="[0, 0]"
                xhide="contextMenuMessage = null"
            >
                <template #content="{ hide }">
                    <div class="options-menu min-w-[180px] max-h-max">
                        <ul @click="hide">
                            <li>
                                <button
                                    class="option-button gap-x-[6px]"
                                    type="button"
                                    xclick="onReply"
                                >
                                    <ReplyIcon class="w-[20px] h-[20px]"/>

                                    <span>
                                        Ответить
                                    </span>
                                </button>
                            </li>
                            <li>
                                <CopyButton
                                    class="option-button gap-x-[6px]"
                                    :text="contextMenuMessage ? contextMenuMessage.text : ''"
                                >
                                    <CopyIcon class="w-[20px] h-[20px]"/>

                                    <span>
                                        Копировать текст
                                    </span>
                                </CopyButton>
                            </li>
                            <li>
                                <button
                                    class="option-button gap-x-[6px]"
                                    type="button"
                                    xpointerdown="messageToForward = contextMenuMessage"
                                    xclick="$showPopup('forwardPopup')"
                                >
                                    <ShareArrowIcon class="w-[20px] h-[20px]"/>

                                    <span>
                                        Переслать
                                    </span>
                                </button>
                            </li>
                            <li>
                                <button class="option-button gap-x-[6px]" type="button" xclick="onSelect">
                                    <CheckmarkBigCircleIcon class="w-[20px] h-[20px]"/>

                                    <span>
                                        Выбрать
                                    </span>
                                </button>
                            </li>
                            <li>
                                <button
                                    class="option-button gap-x-[6px] !text-utilitarian-moscow-100"
                                    type="button"
                                    xpointerdown="messageToDelete = contextMenuMessage"
                                    xclick="$showPopup('deletionConfirmationPopup')"
                                >
                                    <TrashDeleteBinIcon class="w-[20px] h-[20px]"/>

                                    <span>
                                        Удалить
                                    </span>
                                </button>
                            </li>
                        </ul>
                    </div>
                </template>
            </Tippy>

            <Popup id="topBarMenuPopup" class="md:!hidden" contentClass="pt-[10px] px-[6px]">
                <div xv-if="chat.type === 'personal'" class="options-menu max-h-max !p-0 !shadow-none">
                    <ul>
                        <li>
                            <button class="option-button gap-x-[6px]" type="button" data-popup-closer xclick="isSearching = true">
                                <SearchIcon class="w-[20px] h-[20px]"/>

                                <span>
                                    Поиск
                                </span>
                            </button>
                        </li>
                        <li>
                            <button class="option-button gap-x-[6px]" type="button" data-popup-closer xclick="isSelecting = true">
                                <CheckmarkBigCircleIcon class="w-[20px] h-[20px]"/>

                                <span>
                                    Выбрать сообщения
                                </span>
                            </button>
                        </li>
                        <li>
                            <button class="option-button gap-x-[6px]" type="button" data-popup-closer xclick="onNotificationsToggle">
                                <VolumeDisabledIcon xv-if="!chat.are_notifications_enabled" class="w-[20px] h-[20px]"/>

                                <span xv-if="chat.are_notifications_enabled">
                                    Выключить уведомления
                                </span>
                                <!-- <span v-else>
                                    Включить уведомления
                                </span> -->
                            </button>
                        </li>
                        <li v-if="false">
                            <button class="option-button gap-x-[6px]" type="button" data-popup-closer>
                                <LockIcon class="w-[20px] h-[20px]"/>

                                <span>
                                    Заблокировать пользователя
                                </span>
                            </button>
                        </li>
                        <li>
                            <button
                                class="option-button gap-x-[6px] !text-utilitarian-moscow-100"
                                type="button"
                                data-popup-closer
                                xclick="$showPopup('confirmChatDeletionPopup')"
                            >
                                <TrashDeleteBinIcon class="w-[20px] h-[20px]"/>

                                <span>
                                    Удалить чат
                                </span>
                            </button>
                        </li>
                    </ul>
                </div>
                <!-- <div v-else-if="chat.type === 'group'" class="options-menu max-h-max p-0 shadow-none">
                    <ul>
                        <li>
                        <button class="option-button gap-x-[6px]" type="button" data-popup-closer @click="isSearching = true">
                            <SearchIcon class="w-[20px] h-[20px]"/>

                            <span>
                            Поиск
                            </span>
                        </button>
                        </li>
                        <li v-if="chat.owner_id === currentUser.id">
                        <button
                            class="option-button gap-x-[6px]"
                            type="button"
                            data-popup-closer
                            @click="store.currentState = store.State.GROUP_EDIT_FORM"
                        >
                            <PenEditWriteIcon class="w-[20px] h-[20px]"/>

                            <span>
                            Редактировать
                            </span>
                        </button>
                        </li>
                        <li>
                        <button class="option-button gap-x-[6px]" type="button" data-popup-closer @click="store.currentState = store.State.GROUP_DETAIL">
                            <InfoInformationCircleIcon class="w-[20px] h-[20px]"/>

                            <span>
                            Информация
                            </span>
                        </button>
                        </li>
                        <li>
                        <button class="option-button gap-x-[6px]" type="button" data-popup-closer xclick="onNotificationsToggle">
                            <VolumeDisabledIcon xv-if="!chat.are_notifications_enabled" class="w-[20px] h-[20px]"/>

                            <span xv-if="chat.are_notifications_enabled">
                                Выключить уведомления
                            </span>
                            <span v-else>
                                Включить уведомления
                            </span>
                        </button>
                        </li>
                        <li>
                        <button
                            v-if="currentUser.id === chat.owner_id"
                            class="option-button gap-x-[6px] text-utilitarian-moscow-100"
                            type="button"
                            data-popup-closer
                            @click="onChatDelete"
                        >
                            <TrashDeleteBinIcon class="w-[20px] h-[20px]"/>

                            <span>
                            Удалить и покинуть
                            </span>
                        </button>
                        <button
                            v-else
                            class="option-button gap-x-[6px] text-utilitarian-moscow-100"
                            type="button"
                            data-popup-closer
                            @click="onGroupLeave"
                        >
                            <LoginEnterArrowIcon class="w-[20px] h-[20px]"/>

                            <span>
                            Покинуть группу
                            </span>
                        </button>
                        </li>
                    </ul>
                </div> -->

                <div class="mt-[16px] px-[10px]">
                    <button class="button m clear round w-full" type="button" data-popup-closer>
                        <span>
                            Закрыть
                        </span>
                    </button>
                </div>
            </Popup>
<!-- 
            <Popup
                id="deletionConfirmationPopup"
                dialogClass="md:max-w-[330px]"
                contentClass="md:px-[16px] md:py-[20px]"
                @hidden="messageToDelete = null"
            >
                <ChatMessageDeletionForm
                :messages="messageToDelete ? [messageToDelete] : selectedMessages"
                @submit="onMessagesDelete"
                />
            </Popup>

            <Popup id="forwardPopup">
                <ChatMessageForwardForm
                @submit="onMessagesForward"
                />
            </Popup>

            <PopupConfirm
                id="confirmChatDeletionPopup"
                confirmText="Удалить"
                cancelText="Отмена"
                title="Удалить чат"
                description="Вы уверены, что хотите удалить чат?"
                @confirm="onChatDelete"
            /> -->
        </Teleport>
    </div>
</template>