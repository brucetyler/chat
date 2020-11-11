import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: ()=>import("../views/login/login.vue")
    },
    {
      path: '/chat',
      component: ()=>import("../views/chat/chat.vue")
    }
  ]
})
export default router