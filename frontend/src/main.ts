import { adjustTextarea } from '@/services/form'
import './assets/styles/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.directive('autoresize', {
    mounted(el: HTMLTextAreaElement): void {
        adjustTextarea(el)

        el.addEventListener('input', (event: Event) => {
            const textarea = event.target as HTMLTextAreaElement

            adjustTextarea(textarea)
        })
    },
})

app.directive('focus', {
    mounted(el) {
        const options = {
            preventScroll: true,
        }

        if (el.matches('.textarea')) {
            el.querySelector('.textarea__field').focus(options)
            return
        }

        el.focus(options)
    },
})

app.mount('#app')
