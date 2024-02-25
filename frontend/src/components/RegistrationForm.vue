<script lang="ts" setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import { useAuthStore } from '@/stores/auth'

const data = ref({
    username: 'scanderoff',
    email: 'scanderoff@gmail.com',
    firstName: 'Ruslan',
    lastName: 'Iskandarov',
    password: 'svetabaran',
})

const router = useRouter()
const auth = useAuthStore()

const onSubmit = async () => {
    try {
        await auth.register(data.value)
        // await router.push({ name: 'chats' })
        // window.location.reload()
    } catch {
        console.log('CATCHED')
    }
}
</script>

<template>
    <form action="#" @submit.prevent="onSubmit">
        <span>Username</span><br />
        <input
            class="border"
            type="text"
            placeholder="username"
            v-model="data.username"
            required
        /><br />
        <span>Email</span><br />
        <input
            class="border"
            type="email"
            placeholder="email"
            v-model="data.email"
            required
        /><br />
        <span>First Name</span><br />
        <input
            class="border"
            type="text"
            placeholder="first name"
            v-model="data.firstName"
            required
        /><br />
        <span>Last Name</span><br />
        <input
            class="border"
            type="text"
            placeholder="last name"
            v-model="data.lastName"
            required
        /><br />
        <span>Password</span><br />
        <input
            class="border"
            type="password"
            placeholder="password"
            v-model="data.password"
            required
        /><br />
        <button type="submit">
            <span> Register </span>
        </button>
    </form>
</template>
