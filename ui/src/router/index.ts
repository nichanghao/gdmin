import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: () => import('@/views/login/index.vue'),
    },
    {
      path: '/index',
      component: () => import('@/views/home/index.vue'),
      name: 'Index',
      meta: {
        title: '首页',
        icon: 'ep:home-filled',
        noCache: false,
        affix: true
      }
    }
  ]
});

export default router;