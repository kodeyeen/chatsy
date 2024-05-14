<script lang="ts" setup>
import { onMounted } from 'vue'
import { RouterView, useRouter } from 'vue-router'

import { useAuthStore } from '@/stores/auth'
import { useChatsStore } from '@/stores/chats'

const router = useRouter()
const auth = useAuthStore()
const chats = useChatsStore()

onMounted(() => {
    auth.fetchCurrentUser()
        .then(() => {
            return auth.createTicket()
        })
        .then(() => {
            return chats.init(auth.ticket)
        })
        .catch((e) => {
            // router.push({
            //     name: 'login',
            // })
            console.log(e)
        })
})
</script>

<template>
    <RouterView />
</template>
