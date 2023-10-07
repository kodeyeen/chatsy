import '@/assets/styles/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from '@/App.vue'
import utilsPlugin from '@/plugins/utils'

const app = createApp(App)

app.use(createPinia())
app.use(utilsPlugin)

app.mount('#app')
