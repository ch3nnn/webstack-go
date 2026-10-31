<script setup lang="ts">
import { NButton, NCard, NForm, NFormItem, NInput, NSpin, NSwitch, useMessage } from 'naive-ui'
import { onMounted, reactive, ref } from 'vue'

import { fetchConfig, updateSiteSettings, type SiteSettingsPayload } from '@/api/config'
import { errorMessage } from '@/api/client'
import { imgSrc } from '@/utils/img'

const message = useMessage()

const loading = ref(false)
const savingSite = ref(false)

/**
 * 站点设置改成改一项存一项：串行排队，避免连点造成并发写。
 * 失败时重拉服务端值，把输入框回滚到真正存下的内容。
 */
let saveQueue: Promise<void> = Promise.resolve()

function persistSite(patch: SiteSettingsPayload) {
  saveQueue = saveQueue
    .then(async () => {
      savingSite.value = true
      try {
        await updateSiteSettings(patch)
        message.success('站点设置已保存')
      } catch (error) {
        message.error(errorMessage(error))
        await load()
      } finally {
        savingSite.value = false
      }
    })
    .catch(() => {})
}

const site = reactive({
  isSearch: true,
  siteTitle: '',
  siteKeyword: '',
  siteDesc: '',
  siteRecord: '',
  siteURL: '',
  siteLogo: '',
  siteFavicon: '',
})

const logoFile = ref<File | null>(null)
const faviconFile = ref<File | null>(null)
const logoInput = ref<HTMLInputElement | null>(null)
const faviconInput = ref<HTMLInputElement | null>(null)

async function load() {
  loading.value = true
  try {
    const data = await fetchConfig()
    site.isSearch = data.is_search
    site.siteTitle = data.site_title
    site.siteKeyword = data.site_keyword
    site.siteDesc = data.site_desc
    site.siteRecord = data.site_record
    site.siteURL = data.site_url
    site.siteLogo = data.site_logo
    site.siteFavicon = data.site_favicon
  } catch (error) {
    message.error(errorMessage(error))
  } finally {
    loading.value = false
  }
}

function pickFile(event: Event, target: 'logo' | 'favicon') {
  const input = event.target as HTMLInputElement
  const file = input.files?.[0] ?? null
  if (!file) return
  if (target === 'logo') {
    logoFile.value = file
    persistSite({ logo: file })
  } else {
    faviconFile.value = file
    persistSite({ favicon: file })
  }
}

onMounted(load)
</script>

<template>
  <n-spin :show="loading">
    <n-card size="small" title="站点设置" class="card">
      <p class="auto-save-hint">
        {{ savingSite ? '正在保存…' : '每项修改后自动保存，无需手动提交。' }}
      </p>
      <n-form label-placement="left" label-width="88">
        <n-form-item label="首页搜索框">
          <div class="switch-field">
            <n-switch v-model:value="site.isSearch" @update:value="(v: boolean) => persistSite({ is_search: v })" />
            <span class="hint">关闭后首页不显示搜索框</span>
          </div>
        </n-form-item>
        <n-form-item label="网站标题">
          <n-input
            v-model:value="site.siteTitle"
            maxlength="20"
            placeholder="请输入网站标题"
            @change="persistSite({ site_title: site.siteTitle })"
          />
        </n-form-item>
        <n-form-item label="网站关键词">
          <n-input
            v-model:value="site.siteKeyword"
            maxlength="20"
            placeholder="请输入站点关键词"
            @change="persistSite({ site_keyword: site.siteKeyword })"
          />
        </n-form-item>
        <n-form-item label="站点描述">
          <n-input
            v-model:value="site.siteDesc"
            type="textarea"
            :autosize="{ minRows: 2, maxRows: 4 }"
            @change="persistSite({ site_desc: site.siteDesc })"
          />
        </n-form-item>
        <n-form-item label="备案号">
          <n-input
            v-model:value="site.siteRecord"
            placeholder="请输入备案号"
            @change="persistSite({ site_record: site.siteRecord })"
          />
        </n-form-item>
        <n-form-item label="备案链接">
          <n-input
            v-model:value="site.siteURL"
            placeholder="请输入 ICP/IP/域名备案管理地址"
            @change="persistSite({ site_url: site.siteURL })"
          />
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
      </n-form>
    </n-card>
  </n-spin>
</template>

<style scoped>
.card {
  margin-bottom: 12px;
}

.switch-field {
  display: flex;
  gap: 10px;
  align-items: center;
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

.auto-save-hint {
  margin: 0 0 12px;
  color: var(--color-text-secondary);
  font-size: 12px;
}

.hidden-input {
  display: none;
}
</style>
