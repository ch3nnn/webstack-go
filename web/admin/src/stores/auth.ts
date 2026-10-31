import { defineStore } from 'pinia'
import { ref } from 'vue'

import { fetchInfo, login as loginApi, logout as logoutApi, type MenuItem } from '@/api/user'

const TOKEN_COOKIE = '_login_token_'
const TOKEN_MAX_AGE = 24 * 60 * 60

export const useAuthStore = defineStore('auth', () => {
  const username = ref('')
  const menu = ref<MenuItem[]>([])
  const loaded = ref(false)

  async function ensureUser(): Promise<void> {
    if (loaded.value) return
    const info = await fetchInfo()
    username.value = info.username
    menu.value = info.menu ?? []
    loaded.value = true
  }

  async function login(name: string, password: string): Promise<void> {
    const token = await loginApi(name, password)
    document.cookie = `${TOKEN_COOKIE}=${token}; path=/; max-age=${TOKEN_MAX_AGE}; SameSite=Lax`
    loaded.value = false
    await ensureUser()
  }

  async function logout(): Promise<void> {
    try {
      await logoutApi()
    } finally {
      document.cookie = `${TOKEN_COOKIE}=; path=/; max-age=0`
      username.value = ''
      menu.value = []
      loaded.value = false
    }
  }

  return { username, menu, loaded, ensureUser, login, logout }
})
