import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// 全局样式
// import '@/assets/css/global.scss';

import 'element-plus/lib/theme-chalk/index.css'

// 引入 ElementUI
import ElementPlus from 'element-plus'

const app = createApp(App)
app.use(ElementPlus)
app.use(store).use(router).mount('#app')
