<script setup lang="ts">
import { NButton, NForm, NFormItem, NInput, useMessage } from 'naive-ui'
import { reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import { errorMessage } from '@/api/client'
import { useAuthStore } from '@/stores/auth'

const auth = useAuthStore()
const router = useRouter()
const route = useRoute()
const message = useMessage()

const form = reactive({ username: '', password: '' })
const submitting = ref(false)

async function onSubmit() {
  if (!form.username || !form.password) {
    message.warning('请输入用户名和密码')
    return
  }
  submitting.value = true
  try {
    await auth.login(form.username, form.password)
    const redirect = typeof route.query.redirect === 'string' ? route.query.redirect : '/admin'
    await router.replace(redirect)
  } catch (error) {
    message.error(errorMessage(error))
  } finally {
    submitting.value = false
  }
}
</script>

<template>
  <div class="page">
    <div class="card">
      <h1 class="title">WebStack 管理后台</h1>
      <n-form :show-label="false" @submit.prevent="onSubmit">
        <n-form-item>
          <n-input v-model:value="form.username" placeholder="用户名" autofocus />
        </n-form-item>
        <n-form-item>
          <n-input
            v-model:value="form.password"
            type="password"
            show-password-on="click"
            placeholder="密码"
            @keyup.enter="onSubmit"
          />
        </n-form-item>
        <n-button type="primary" block :loading="submitting" attr-type="submit" @click="onSubmit">
          登录
        </n-button>
      </n-form>
    </div>
  </div>
</template>

<style scoped>
.page {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background: #f5f6f8;
}

.card {
  width: 340px;
  padding: 28px 24px 24px;
  border: 1px solid var(--color-border);
  border-radius: 10px;
  background: #fff;
}

.title {
  margin: 0 0 20px;
  font-size: 16px;
  font-weight: 600;
  color: #1f2329;
}
</style>
