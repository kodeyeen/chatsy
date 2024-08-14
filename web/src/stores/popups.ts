import { ref } from 'vue'
import { defineStore } from 'pinia'

export const usePopupsStore = defineStore('popups', () => {
    const popups = ref<{ [key: string]: boolean }>({})

    const isActive = (popupId: string): boolean => {
        return popups.value.hasOwnProperty(popupId)
    }

    const show = (popupId: string): void => {
        popups.value[popupId] = true
    }

    const hide = (popupId: string): void => {
        delete popups.value[popupId]
    }

    return {
        isActive,

        show,
        hide,
    }
})
