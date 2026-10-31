<script setup lang="ts">
import { NButton, NCard, NForm, NFormItem, NSpin, NSwitch, useMessage } from 'naive-ui'
import { MdEditor } from 'md-editor-v3'
import { onMounted, reactive, ref } from 'vue'

import { fetchConfig, updateAbout } from '@/api/config'
import { errorMessage } from '@/api/client'

import 'md-editor-v3/lib/style.css'

const message = useMessage()

const loading = ref(false)
const saving = ref(false)

const about = reactive({
  aboutSite: '',
  aboutAuthor: '',
  isAbout: false,
})

async function load() {
  loading.value = true
  try {
    const data = await fetchConfig()
    about.aboutSite = data.about_site
    about.aboutAuthor = data.about_author
    about.isAbout = data.is_about
  } catch (error) {
    message.error(errorMessage(error))
  } finally {
    loading.value = false
  }
}

async function onSave() {
  saving.value = true
  try {
    await updateAbout({
      about_site: about.aboutSite,
      about_author: about.aboutAuthor,
      is_about: about.isAbout,
    })
    message.success('关于页面已保存')
  } catch (error) {
    message.error(errorMessage(error))
  } finally {
    saving.value = false
  }
}

onMounted(load)
</script>

<template>
  <n-spin :show="loading">
    <n-card size="small" title="关于页面" class="card">
      <n-form label-placement="top">
        <n-form-item label="关于网站（Markdown）">
          <MdEditor v-model="about.aboutSite" :toolbars-exclude="['github']" preview-theme="github" />
        </n-form-item>
        <n-form-item label="关于站长（Markdown）">
          <MdEditor v-model="about.aboutAuthor" :toolbars-exclude="['github']" preview-theme="github" />
        </n-form-item>
        <n-form-item label="是否显示关于信息">
          <n-switch v-model:value="about.isAbout" />
        </n-form-item>
        <n-form-item label=" ">
          <n-button type="primary" :loading="saving" @click="onSave">保存关于页面</n-button>
        </n-form-item>
      </n-form>
    </n-card>
  </n-spin>
</template>

<style scoped>
.card {
  margin-bottom: 12px;
}
</style>
