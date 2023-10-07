import { usePopupsStore } from '@/stores/popups'

export default {
    install: (app, options) => {
        app.config.globalProperties.$showPopup = (popupId, data = {}) => {
            const store = usePopupsStore()

            store.show(popupId, data)
        }
    }
}
