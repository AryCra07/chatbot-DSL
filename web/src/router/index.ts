import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'
import ChatBox from "@/components/ChatBox.vue";

const routes: Array<RouteRecordRaw> = [
    {
        path: '/chat',
        name: 'ChatBox',
        component: ChatBox
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
