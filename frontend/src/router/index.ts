import { createRouter, createWebHistory } from 'vue-router'

import RegistrationView from '@/views/RegistrationView.vue'
import LoginView from '@/views/LoginView.vue'
import ChatsView from '@/views/ChatsView.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/registration',
            name: 'registration',
            component: RegistrationView,
        },
        {
            path: '/login',
            name: 'login',
            component: LoginView,
        },
        {
            path: '/chats',
            name: 'chats',
            component: ChatsView,
        },
    ],
})

export default router
