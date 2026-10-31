<script setup lang="ts">
import { NButton, NCard, NForm, NFormItem, NSpin, NSwitch, useMessage } from 'naive-ui'
import { MdEditor } from 'md-editor-v3'
import { computed, onMounted, reactive, ref } from 'vue'

import { fetchConfig, updateAbout } from '@/api/config'
import { errorMessage } from '@/api/client'

import 'md-editor-v3/lib/style.css'

const message = useMessage()

const loading = ref(false)
const saving = ref(false)
const savingSwitch = ref(false)

const about = reactive({
  aboutSite: '',
  aboutAuthor: '',
  isAbout: false,
})

/** 前台关于页地址，方便点开确认效果 */
const aboutUrl = computed(() => `${location.origin}/about`)

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

/** 开关只表示「是否展示入口」，没有草稿状态，切换即落库 */
async function onToggleShow(value: boolean) {
  savingSwitch.value = true
  try {
    await updateAbout({ is_about: value })
    about.isAbout = value
    message.success('关于页面已保存')
  } catch (error) {
    message.error(errorMessage(error))
    await load()
  } finally {
    savingSwitch.value = false
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
      <div class="about-head">
        <div class="switch-field">
          <n-switch
            :value="about.isAbout"
            :loading="savingSwitch"
            @update:value="onToggleShow"
          />
          <span class="switch-label">是否展示关于页面</span>
          <span class="hint">
            {{ about.isAbout ? '当前：首页侧边栏显示「关于本站」入口' : '当前：首页侧边栏不显示「关于本站」入口' }}
          </span>
        </div>
        <div class="url-field">
          <span class="hint">关于页面地址：</span>
          <a :href="aboutUrl" target="_blank" rel="noreferrer">{{ aboutUrl }}</a>
        </div>
      </div>

      <n-form label-placement="top">
        <n-form-item label="关于网站（Markdown）">
          <MdEditor v-model="about.aboutSite" :toolbars-exclude="['github']" preview-theme="github" />
        </n-form-item>
        <n-form-item label="关于站长（Markdown）">
          <MdEditor v-model="about.aboutAuthor" :toolbars-exclude="['github']" preview-theme="github" />
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

.about-head {
  display: flex;
  flex-wrap: wrap;
  gap: 8px 24px;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--color-border);
}

.switch-field,
.url-field {
  display: flex;
  gap: 8px;
  align-items: center;
}

.switch-label {
  font-size: 13px;
  color: #1f2329;
}

.url-field a {
  color: #2f6feb;
  font-size: 12px;
  word-break: break-all;
}
</style>
