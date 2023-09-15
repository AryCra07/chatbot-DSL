import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import { createPinia } from 'pinia'
import i18n from '@/lang'

const pinia = createPinia();

createApp(App).use(router).use(ElementPlus).use(pinia).use(i18n).mount('#app')
