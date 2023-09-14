import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import { createPinia } from 'pinia';
// import VueI18n from 'vue-i18n'

const pinia = createPinia();

createApp(App).use(router).use(ElementPlus).use(pinia).mount('#app')
