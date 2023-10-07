<script setup>
import Avatar from '@/components/Avatar.vue'
</script>

<template>
    <div ref="root" class="min-w-[288px] max-h-[240px] overflow-x-hidden overflow-y-auto scrollbar p-[4px] rounded-[6px] bg-primary-brand-white shadow-elevation-3">
    <ul>
      <li v-for="i in 5" :key="i">
        <button
          class="option-button gap-x-[8px] px-[10px] py-[8px]"
          :class="{
            // 'active': index === activeUserIndex,
          }"
          type="button"
          xclick="select(user)"
        >
          <span class="shrink-0">
            <Avatar
              class="w-[36px] h-[36px] rounded-full"
              xsrc="user.photo ? $media(user.photo.url) : undefined"
            />
          </span>

          <span class="grow flex flex-col items-start">
            <span class="text-m-14 font-medium text-primary-brand-primary">
              <!-- {{ user.display_name }} -->
              Данил Фролов
            </span>

            <span class="text-s-12 text-primary-seattle-120">
              @frolov_danil
            </span>
          </span>
        </button>
      </li>
    </ul>
  </div>
</template>