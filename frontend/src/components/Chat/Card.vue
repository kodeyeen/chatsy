<script setup>
import Avatar from '@/components/Avatar.vue'
</script>

<template>
    <div
    class="chat-card flex gap-x-[9px] py-[8px] pl-[9px] pr-[12px] rounded-[10px] hover:bg-primary-brand-wildSand [&.active]:bg-primary-brand-accent cursor-pointer transition-colors"
    :class="{
      // 'items-center': chat.last_message,
      'items-center': true,
    }"
  >
    <div class="shrink-0">
      <Avatar
        class="w-[56px] h-[56px] rounded-full"
        xsrc="chat.photo ? $media(chat.photo.url) : undefined"
      />
    </div>

    <div class="grow">
      <div class="flex justify-between items-center gap-x-[8px]">
        <span class="grow flex items-center gap-x-[4px]">
          <span class="text-l-long-16 font-bold text-primary-brand-onPrimary [.chat-card.active_&]:text-primary-brand-white line-clamp-1">
            <!-- {{ chat.title }} -->
            Сын шлюхи
          </span>

          <svg xv-if="!chat.are_notifications_enabled" class="text-primary-seattle-100 [.chat-card.active_&]:text-primary-brand-white" width="24" height="24">
            <use xlink:href="#muteIcon"></use>
          </svg>
        </span>

        <span xv-if="chat.last_message" class="shrink-0 flex items-center gap-x-[3px]">
          <svg xv-if="chat.last_message.author.id === currentUser.id" class="text-utilitarian-geneva-100 [.chat-card.active_&]:text-primary-brand-white" width="24" height="24">
            <use v-if="true" xlink:href="#readIcon"></use>
            <!-- <use v-if="chat.last_message.is_seen" xlink:href="#readIcon"></use> -->
            <use v-else xlink:href="#checkIcon"></use>
          </svg>

          <time
            xv-if="chat.last_message"
            class="text-s-12 text-primary-seattle-100 [.chat-card.active_&]:text-primary-brand-white"
            xdatetime="chat.last_message.sent_at"
          >
            15:59
            <!-- {{ new Date(chat.last_message.sent_at).toLocaleTimeString('ru', { hour: '2-digit', minute: '2-digit' }) }} -->
          </time>
        </span>
      </div>

      <div xv-if="chat.last_message" class="flex items-center gap-x-[8px]">
        <div class="grow flex items-center gap-x-[4px] mt-[2px] text-l-short-16">
          <span xv-if="chat.type === 'group' && chat.last_message.author.id !== currentUser.id" class="text-primary-brand-onPrimary [.chat-card.active_&]:text-primary-brand-white whitespace-nowrap">
            <!-- {{ chat.last_message.author.display_name }}: -->
            Иван Сизых
          </span>

          <span v-if="false" class="shrink-0 w-[24px] h-[24px] border border-primary-seattle-10 rounded-[4px] overflow-hidden">
            <img class="w-full h-full object-cover" src="https://telegra.ph/file/75085f39c9315c6eeefb6.jpg" alt="Attachment">
          </span>

          <span class="text-primary-seattle-100 [.chat-card.active_&]:text-primary-brand-white line-clamp-1">
            Last message text
            <!-- {{ chat.last_message.text }} -->
          </span>
        </div>

        <div xv-if="chat.unread_message_count > 0" class="shrink-0">
          <Counter class="accent [.chat-card.active_&]:hidden" :value="9"/>
          <!-- <Counter class="accent [.chat-card.active_&]:hidden" :value="chat.unread_message_count"/> -->
          <!-- <Counter class="neutral hidden [.chat-card.active_&]:block" :value="chat.unread_message_count"/> -->
        </div>
      </div>
    </div>
  </div>
</template>