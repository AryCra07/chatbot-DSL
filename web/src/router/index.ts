import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';
import ChatBox from '@/components/ChatRoom.vue';
import LoginRegister from '@/views/LoginRegister.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/login',
  },
  {
    path: '/chat',
    name: 'ChatBox',
    component: ChatBox,
  },
  {
    path: '/login',
    name: 'LoginRegister',
    component: LoginRegister,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
