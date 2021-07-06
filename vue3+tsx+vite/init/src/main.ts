import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios'

const app = createApp(App)
app.config.globalProperties.$https = axios
app.mount('#app')
