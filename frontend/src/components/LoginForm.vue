<script lang="ts" setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import { useAuthStore } from '@/stores/auth'

const data = ref({
    email: 'scanderoff@gmail.com',
    password: 'svetabaran',
})

const router = useRouter()
const auth = useAuthStore()

const onSubmit = async () => {
    try {
        await auth.login(data.value)
        await router.push({ name: 'chats' })
        window.location.reload()
    } catch {
        console.log('CATCHED')
    }
}

const logout = async () => {
    await auth.logout()
    await router.push({ name: 'login' })
    window.location.reload()
}
</script>

<template>
    <form action="#" @submit.prevent="onSubmit">
        <input
            type="email"
            placeholder="email"
            name="email"
            v-model="data.email"
            autocomplete="off"
            required
        /><br />
        <input
            type="password"
            placeholder="password"
            name="password"
            v-model="data.password"
            required
        /><br />
        <button type="submit">
            <span> Login </span>
        </button>
    </form>

    <button @click="logout">Logout</button>
</template>
