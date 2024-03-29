import { ref } from 'vue'
import { defineStore } from 'pinia'
import { ofetch } from 'ofetch'

interface User {
    id: number
}

interface Ticket {
    id: string
}

interface Credentials {
    email: string
    password: string
}

export const useAuthStore = defineStore('auth', () => {
    const currentUser = ref<User | null>(null)
    const ticket = ref<Ticket | null>(null)

    const register = async (regData: any) => {
        return ofetch('/auth/register', {
            baseURL: 'http://localhost:8080',
            method: 'POST',
            body: regData,
        })
    }

    const login = async (creds: Credentials) => {
        return ofetch('/auth/login', {
            baseURL: 'http://localhost:8080',
            method: 'POST',
            body: creds,
            credentials: 'include',
        })
    }

    const logout = async () => {
        return ofetch('/auth/logout', {
            baseURL: 'http://localhost:8080',
            method: 'POST',
            credentials: 'include',
        })
    }

    const fetchCurrentUser = async () => {
        currentUser.value = await ofetch('/auth/me', {
            baseURL: 'http://localhost:8080',
            credentials: 'include',
        })
    }

    const createTicket = async () => {
        ticket.value = await ofetch('/auth/ticket', {
            baseURL: 'http://localhost:8080',
            method: 'POST',
            credentials: 'include',
        })
    }

    return {
        currentUser,
        ticket,

        register,
        login,
        logout,
        fetchCurrentUser,
        createTicket,
    }
})
