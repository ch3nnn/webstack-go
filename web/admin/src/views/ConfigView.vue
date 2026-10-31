<script setup lang="ts">
import { NButton, NCard, NForm, NFormItem, NInput, NSpin, NSwitch, useMessage } from 'naive-ui'
import { MdEditor } from 'md-editor-v3'
import { onMounted, reactive, ref } from 'vue'

import { fetchConfig, updateAbout, updateSiteSettings } from '@/api/config'
import { errorMessage } from '@/api/client'

import 'md-editor-v3/lib/style.css'

const message = useMessage()

const loading = ref(false)
const savingSite = ref(false)
const savingAbout = ref(false)

const site = reactive({
  siteTitle: '',
  siteKeyword: '',
  siteDesc: '',
  siteRecord: '',
  siteURL: '',
  siteLogo: '',
  siteFavicon: '',
})

const about = reactive({
  aboutSite: '',
  aboutAuthor: '',
  isAbout: false,
})

const logoFile = ref<File | null>(null)
const faviconFile = ref<File | null>(null)
const logoInput = ref<HTMLInputElement | null>(null)
const faviconInput = ref<HTMLInputElement | null>(null)

async function load() {
  loading.value = true
  try {
    const data = await fetchConfig()
    site.siteTitle = data.site_title
    site.siteKeyword = data.site_keyword
    site.siteDesc = data.site_desc
    site.siteRecord = data.site_record
    site.siteURL = data.site_url
    site.siteLogo = data.site_logo
    site.siteFavicon = data.site_favicon
    about.aboutSite = data.about_site
    about.aboutAuthor = data.about_author
    about.isAbout = data.is_about
  } catch (error) {
    message.error(errorMessage(error))
  } finally {
    loading.value = false
  }
}

/** 后端把 logo / favicon 存成纯 base64（无 data: 前缀） */
function imgSrc(value: string): string {
  if (!value) return ''
  if (value.startsWith('data:') || value.startsWith('http') || value.startsWith('/')) return value
  return `data:image/png;base64,${value}`
}

function pickFile(event: Event, target: 'logo' | 'favicon') {
  const input = event.target as HTMLInputElement
  const file = input.files?.[0] ?? null
  if (target === 'logo') logoFile.value = file
  else faviconFile.value = file
}

async function onSubmitSite() {
  savingSite.value = true
  try {
    await updateSiteSettings({
      site_title: site.siteTitle,
      site_keyword: site.siteKeyword,
      site_desc: site.siteDesc,
      site_record: site.siteRecord,
      site_url: site.siteURL,
      logo: logoFile.value,
      favicon: faviconFile.value,
    })
    message.success('站点设置已保存')
    await load()
  } catch (error) {
    message.error(errorMessage(error))
  } finally {
    savingSite.value = false
  }
}

async function onSubmitAbout() {
  savingAbout.value = true
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
    savingAbout.value = false
  }
}

onMounted(load)
</script>

<template>
  <n-spin :show="loading">
    <n-card size="small" title="站点设置" class="card">
      <n-form label-placement="left" label-width="88">
        <n-form-item label="网站标题">
          <n-input v-model:value="site.siteTitle" maxlength="20" placeholder="请输入网站标题" />
        </n-form-item>
        <n-form-item label="网站关键词">
          <n-input v-model:value="site.siteKeyword" maxlength="20" placeholder="请输入站点关键词" />
        </n-form-item>
        <n-form-item label="站点描述">
          <n-input v-model:value="site.siteDesc" type="textarea" :autosize="{ minRows: 2, maxRows: 4 }" />
        </n-form-item>
        <n-form-item label="备案号">
          <n-input v-model:value="site.siteRecord" placeholder="请输入备案号" />
        </n-form-item>
        <n-form-item label="备案链接">
          <n-input v-model:value="site.siteURL" placeholder="请输入 ICP/IP/域名备案管理地址" />
        </n-form-item>
        <n-form-item label="网站 Logo">
          <div class="upload-field">
            <img v-if="site.siteLogo" :src="imgSrc(site.siteLogo)" class="preview" alt="logo" />
            <n-button attr-type="button" @click="logoInput?.click()">选择文件</n-button>
            <input ref="logoInput" type="file" accept="image/*" class="hidden-input" @change="pickFile($event, 'logo')" />
            <span class="hint">{{ logoFile?.name ?? '未选择新文件' }}</span>
          </div>
        </n-form-item>
        <n-form-item label="网站 Favicon">
          <div class="upload-field">
            <img v-if="site.siteFavicon" :src="imgSrc(site.siteFavicon)" class="preview" alt="favicon" />
            <n-button attr-type="button" @click="faviconInput?.click()">选择文件</n-button>
            <input
              ref="faviconInput"
              type="file"
              accept="image/*"
              class="hidden-input"
              @change="pickFile($event, 'favicon')"
            />
            <span class="hint">{{ faviconFile?.name ?? '未选择新文件' }}</span>
          </div>
        </n-form-item>
        <n-form-item label=" ">
          <n-button type="primary" :loading="savingSite" @click="onSubmitSite">保存站点设置</n-button>
        </n-form-item>
      </n-form>
    </n-card>

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
          <n-button type="primary" :loading="savingAbout" @click="onSubmitAbout">保存关于页面</n-button>
        </n-form-item>
      </n-form>
    </n-card>
  </n-spin>
</template>

<style scoped>
.card {
  margin-bottom: 12px;
}

.upload-field {
  display: flex;
  gap: 10px;
  align-items: center;
}

.preview {
  width: 28px;
  height: 28px;
  object-fit: contain;
}

.hint {
  color: var(--color-text-secondary);
  font-size: 12px;
}

.hidden-input {
  display: none;
}
</style>
