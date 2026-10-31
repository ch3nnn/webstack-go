<script setup lang="ts">
import { NButton, NCard, NForm, NFormItem, NInput, useMessage } from 'naive-ui'
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'

import { errorMessage } from '@/api/client'
import { updatePassword } from '@/api/user'
import { useAuthStore } from '@/stores/auth'

const message = useMessage()
const router = useRouter()
const auth = useAuthStore()

const form = reactive({ oldPassword: '', newPassword: '', confirmPassword: '' })
const submitting = ref(false)

async function onSubmit() {
  if (!form.oldPassword || !form.newPassword) {
    message.warning('请填写原密码与新密码')
    return
  }
  if (form.newPassword !== form.confirmPassword) {
    message.warning('两次输入的新密码不一致')
    return
  }
  submitting.value = true
  try {
    await updatePassword(form.oldPassword, form.newPassword)
    message.success('密码修改成功，请重新登录')
    await auth.logout()
    await router.replace('/login')
  } catch (error) {
    message.error(errorMessage(error))
  } finally {
    submitting.value = false
  }
}
</script>

<template>
  <n-card size="small" title="修改密码" class="card">
    <n-form label-placement="left" label-width="88" @submit.prevent="onSubmit">
      <n-form-item label="原密码">
        <n-input v-model:value="form.oldPassword" type="password" show-password-on="click" />
      </n-form-item>
      <n-form-item label="新密码">
        <n-input v-model:value="form.newPassword" type="password" show-password-on="click" />
      </n-form-item>
      <n-form-item label="确认新密码">
        <n-input v-model:value="form.confirmPassword" type="password" show-password-on="click" />
      </n-form-item>
      <n-form-item label=" ">
        <n-button type="primary" :loading="submitting" attr-type="submit">提交</n-button>
      </n-form-item>
    </n-form>
  </n-card>
</template>

<style scoped>
.card {
  max-width: 480px;
}
</style>
