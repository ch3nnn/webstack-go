<script setup lang="ts">
import {
  AppsOutline,
  DocumentTextOutline,
  GridOutline,
  KeyOutline,
  LinkOutline,
  LogOutOutline,
  PricetagsOutline,
  SettingsOutline,
} from '@vicons/ionicons5'
import {
  NDropdown,
  NIcon,
  NLayout,
  NLayoutContent,
  NLayoutHeader,
  NLayoutSider,
  NMenu,
  type MenuOption,
  useMessage,
} from 'naive-ui'
import { computed, h, ref, type Component } from 'vue'
import { RouterView, useRoute, useRouter } from 'vue-router'

import { errorMessage } from '@/api/client'
import { useAuthStore } from '@/stores/auth'

const route = useRoute()
const router = useRouter()
const auth = useAuthStore()
const message = useMessage()

const collapsed = ref(false)

const ICONS: Record<string, Component> = {
  dashboard: GridOutline,
  site: LinkOutline,
  category: PricetagsOutline,
  config: SettingsOutline,
  about_page: DocumentTextOutline,
  modify_password: KeyOutline,
}

function iconFor(link: string): Component {
  const key = link.split('/').filter(Boolean).pop() ?? ''
  return ICONS[key] ?? AppsOutline
}

function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) })
}

const menuOptions = computed<MenuOption[]>(() => {
  const options: MenuOption[] = [
    { label: '仪表盘', key: '/admin/dashboard', icon: renderIcon(GridOutline) },
  ]
  for (const root of auth.menu.filter((item) => !item.pid)) {
    const children = auth.menu.filter((item) => item.pid === root.id)
    if (children.length === 0) {
      // 顶级菜单必须有链接才可点击
      if (root.link) {
        options.push({ label: root.name, key: root.link, icon: renderIcon(iconFor(root.link)) })
      }
      continue
    }
    options.push({
      label: root.name,
      key: `group-${root.id}`,
      icon: renderIcon(iconFor(root.link)),
      children: children.map((child) => ({ label: child.name, key: child.link })),
    })
  }
  return options
})

const userOptions: MenuOption[] = [
  { label: '修改密码', key: 'modify_password', icon: renderIcon(KeyOutline) },
  { label: '退出登录', key: 'logout', icon: renderIcon(LogOutOutline) },
]

function go(key: string) {
  if (key !== route.path) router.push(key)
}

async function onUserSelect(key: string) {
  if (key === 'modify_password') {
    router.push('/admin/modify_password')
    return
  }
  try {
    await auth.logout()
    router.replace('/login')
  } catch (error) {
    message.error(errorMessage(error))
  }
}
</script>

<template>
  <n-layout has-sider class="layout">
    <n-layout-sider
      bordered
      collapse-mode="width"
      :collapsed-width="64"
      :width="208"
      :collapsed="collapsed"
      show-trigger="bar"
      @collapse="collapsed = true"
      @expand="collapsed = false"
    >
      <div class="brand">
        <span class="brand-dot" />
        <span v-show="!collapsed" class="brand-text">WebStack</span>
      </div>
      <n-menu
        :value="route.path"
        :options="menuOptions"
        :collapsed="collapsed"
        :collapsed-width="64"
        :collapsed-icon-size="20"
        @update:value="go"
      />
    </n-layout-sider>

    <n-layout>
      <n-layout-header bordered class="header">
        <span class="header-title">{{ route.meta.title ?? '' }}</span>
        <n-dropdown :options="userOptions" trigger="click" @select="onUserSelect">
          <button class="user pressable" type="button">{{ auth.username || '管理员' }}</button>
        </n-dropdown>
      </n-layout-header>

      <n-layout-content class="content" content-style="padding: 16px 20px 24px;">
        <router-view v-slot="{ Component: Page }">
          <keep-alive>
            <component :is="Page" />
          </keep-alive>
        </router-view>
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>

<style scoped>
.layout {
  height: 100vh;
}

.brand {
  display: flex;
  align-items: center;
  gap: 8px;
  height: 52px;
  padding: 0 18px;
  border-bottom: 1px solid var(--color-border);
  font-weight: 600;
  color: #1f2329;
}

.brand-dot {
  width: 8px;
  height: 8px;
  border-radius: 2px;
  background: #2f6feb;
}

.brand-text {
  font-size: 14px;
  letter-spacing: 0.02em;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 52px;
  padding: 0 20px;
  background: #fff;
}

.header-title {
  font-size: 14px;
  font-weight: 500;
  color: #1f2329;
}

.user {
  padding: 4px 10px;
  border: 1px solid var(--color-border);
  border-radius: 6px;
  background: #fff;
  color: var(--color-text-secondary);
  font-size: 13px;
  cursor: pointer;
}

.content {
  background: #f5f6f8;
}
</style>
