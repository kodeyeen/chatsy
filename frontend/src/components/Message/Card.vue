<script setup>
import Avatar from '@/components/Avatar.vue'
</script>

<template>
    <div
    ref="root"
    class="block px-[6px] 2md:px-[22px] py-[2px] transition-colors [&.selected]:bg-[rgba(45,164,48,0.20)]"
    :class="{
    //   'selected': isSelected,
    //   'cursor-pointer': selectable,
    }"
    xclick="onSelect"
  >
    <div
      class="relative flex [.own_&]:justify-end max-w-[720px] mx-auto 2md:pr-[52px]"
      :class="{
        // 'pl-[53px]': own && showUserInfo,
      }"
    >
      <input
        class="l absolute top-1/2 right-[calc(100%+20px)] -translate-y-1/2 opacity-0 border-primary-seattle-100 hover:border-primary-brand-primary checked:border-primary-brand-primary transition-opacity duration-300"
        :class="{
        //   'opacity-100': selectable,
        }"
        type="checkbox"
        xchecked="isSelected"
      >

      <div
        class="relative flex items-end max-w-max ml-0"
        :class="{
        //   'pl-[44px]': !own && showUserInfo,
        }"
      >
        <NuxtLink
          xv-if="!own && lastInGroup && showUserInfo"
          class="absolute bottom-0 left-0"
          xto="$userLink(message.author)"
        >
          <Avatar
            class="w-[44px] h-[44px] rounded-full"
            xsrc="message.author.photo ? $media(message.author.photo.url) : null"
          />
        </NuxtLink>

        <div class="flex [.own_&]:flex-row-reverse items-end max-w-[426px]">
          <svg
            class="shrink-0 text-primary-brand-white [.own_&]:text-[rgba(241,254,225,1)] [.own_&]:-scale-x-100"
            :class="{
            //   'invisible': !lastInGroup,
            }"
            width="9"
            height="16"
          >
            <use xlink:href="#cornerShape"></use>
          </svg>

          <div
            class="relative grow flex flex-col items-start max-w-[417px] px-[12px] py-[6px] rounded-[12px] [.own_&]:rounded-bl-[12px] bg-primary-brand-white [.own_&]:bg-[rgba(241,254,225,1)]"
            :class="{
            //   'rounded-bl-none [.own_&]:rounded-br-none': lastInGroup,
            }"
          >
            <div
              xv-if="(firstInGroup && showUserInfo) || message.forward_original"
              class="flex items-center gap-x-[4px]"
              :class="{
                // '[.own_&]:hidden': !message.forward_original,
              }"
            >
              <span class="flex items-center gap-x-[4px]">
                <span xv-if="message.forward_original" class="text-m-14 font-medium text-primary-atlantic-140">
                  Переслано от:
                </span>

                <span class="button-ghost m-medium link font-bold">
                  <span>
                    Ruslan Iskandarov
                    <!-- {{ author.display_name }} -->
                  </span>
                </span>
              </span>
            </div>

            <div>
              <!-- <ChatMessageParent
                v-if="message.reply_to"
                :message="message.reply_to"
                @click="onRepliedMessageClick"
              /> -->

              <div class="text-[16px] leading-[21px] text-primary-brand-onPrimary">
                <pre class="inline font-[inherit] whitespace-pre-wrap" xv-html="formattedText">Test message card template</pre>
                <span class="inline-block w-[44px] [.own_&]:w-[61px]"></span>
              </div>

              <!-- <div class="flex flex-col items-start mt-[8px]">
                <ChatMessageAttachment/>
              </div> -->
            </div>

            <span class="absolute bottom-[6px] right-[12px] flex items-center gap-x-[3px]">
              <time class="text-s-12 text-primary-seattle-100" xdatetime="message.sent_at">
                15:59
                <!-- {{ new Date(message.sent_at).toLocaleTimeString('ru', { hour: '2-digit', minute: '2-digit' }) }} -->
              </time>

              <svg class="hidden [.own_&]:block text-utilitarian-geneva-100" width="17" height="17">
                <use xlink:href="#readIcon"></use>
                <!-- <use v-if="message.is_seen" xlink:href="#readIcon"></use> -->
                <!-- <use v-else xlink:href="#checkIcon"></use> -->
              </svg>
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>