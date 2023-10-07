import { ref } from 'vue'
import { defineStore } from 'pinia'

export const usePopupsStore = defineStore('popups', () => {
    const popups = ref({})
  
    const isActive = (popupId) => {
        return popups.value.hasOwnProperty(popupId)
    }
  
    const getData = (popupId) => {
        return popups.value[popupId] ?? {}
    }
  
    const show = (popupId, data = {}) => {
        popups.value[popupId] = data
    }
  
    const hide = (popupId) => {
        delete popups.value[popupId]
    }
  
    return {
        isActive,
  
        getData,
        show,
        hide,
    }
})
  