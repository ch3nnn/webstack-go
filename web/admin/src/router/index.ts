import { createRouter, createWebHistory } from 'vue-router'

import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', redirect: '/admin' },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue'),
      meta: { public: true },
    },
    {
      path: '/admin',
      component: () => import('@/layouts/AdminLayout.vue'),
      children: [
        { path: '', redirect: '/admin/dashboard' },
        {
          path: 'dashboard',
          name: 'dashboard',
          component: () => import('@/views/DashboardView.vue'),
          meta: { title: '仪表盘' },
        },
      ],
    },
    { path: '/:pathMatch(.*)*', redirect: '/admin' },
  ],
})

router.beforeEach(async (to) => {
  if (to.meta.public) return true
  try {
    await useAuthStore().ensureUser()
    return true
  } catch {
    return { path: '/login', query: { redirect: to.fullPath } }
  }
})

export default router
