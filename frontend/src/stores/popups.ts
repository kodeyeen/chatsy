import { ref } from 'vue'
import { defineStore } from 'pinia'

export const usePopupsStore = defineStore('popups', () => {
    const popups = ref<{ [key: string]: any }>({})

    const isActive = (popupId: string): boolean => {
        return popups.value.hasOwnProperty(popupId)
    }

    const getData = (popupId: string): any | null => {
        return popups.value[popupId] ?? {}
    }

    const show = (popupId: string, data: any = {}): void => {
        popups.value[popupId] = data
    }

    const hide = (popupId: string): void => {
        delete popups.value[popupId]
    }

    return {
        isActive,

        getData,
        show,
        hide,
    }
})
