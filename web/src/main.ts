import { createApp } from 'vue';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import './assets/css/index.css';
import App from './App.vue';
import router from './router';
import { createPinia } from 'pinia';
import i18n from '@/lang';

const pinia = createPinia();

createApp(App).use(router).use(ElementPlus).use(i18n).use(pinia).mount('#app');
