<script setup lang="ts">
import { NButton, NCard, NCheckbox, NForm, NFormItem, NInput, NRadio, NRadioGroup, NSelect, useMessage } from 'naive-ui'
import { computed, onMounted, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'

import { listCategories, type Category } from '@/api/category'
import { errorMessage } from '@/api/client'
import { createSites, type SiteCreateResult } from '@/api/site'

const message = useMessage()
const router = useRouter()

const categories = ref<Category[]>([])
const submitting = ref(false)
const result = ref<SiteCreateResult | null>(null)

const form = reactive({
  categoryId: null as number | null,
  url: '',
  isUsed: true,
  failSwitch: false,
})

/** 一级分类 + 其二级分类，二级用缩进表示层级（与旧页面一致） */
const categoryOptions = computed(() => {
  const roots = categories.value.filter((item) => !item.pid)
  const options: { label: string; value: number }[] = []
  roots.forEach((root) => {
    options.push({ label: root.name, value: root.id })
    categories.value
      .filter((item) => item.pid === root.id)
      .forEach((child) => options.push({ label: `  └ ${child.name}`, value: child.id }))
  })
  return options
})

async function loadCategories() {
  try {
    categories.value = await listCategories()
  } catch (error) {
    message.error(errorMessage(error))
  }
}

async function onSubmit() {
  if (!form.url.trim()) {
    message.warning('请输入网站链接')
    return
  }
  submitting.value = true
  result.value = null
  try {
    result.value = await createSites({
      category_id: form.categoryId,
      url: form.url.trim(),
      is_used: form.isUsed,
      fail_switch: form.failSwitch,
    })
    if (result.value.failCount === 0) {
      message.success(`成功创建 ${result.value.successCount} 条`)
      form.url = ''
    }
    await loadCategories()
  } catch (error) {
    message.error(errorMessage(error))
  } finally {
    submitting.value = false
  }
}

onMounted(loadCategories)
</script>

<template>
  <n-card size="small" title="新增网站" class="card">
    <n-form label-placement="left" label-width="120" @submit.prevent="onSubmit">
      <n-form-item label="分类">
        <n-select v-model:value="form.categoryId" :options="categoryOptions" clearable placeholder="请选择分类" />
      </n-form-item>
      <n-form-item label="网址">
        <n-input
          v-model:value="form.url"
          type="textarea"
          :autosize="{ minRows: 6, maxRows: 14 }"
          placeholder="必须包含 http:// 或 https://，一行一个"
        />
      </n-form-item>
      <n-form-item label="是否启用">
        <n-radio-group v-model:value="form.isUsed">
          <n-radio :value="true">启用</n-radio>
          <n-radio :value="false">禁用</n-radio>
        </n-radio-group>
      </n-form-item>
      <n-form-item label="失败处理">
        <n-checkbox v-model:checked="form.failSwitch">拉取失败时仅保存网址</n-checkbox>
      </n-form-item>
      <n-form-item label=" ">
        <div class="actions">
          <n-button type="primary" :loading="submitting" attr-type="submit">提交</n-button>
          <n-button quaternary @click="router.push('/admin/site')">返回列表</n-button>
        </div>
      </n-form-item>
    </n-form>

    <div v-if="result" class="result">
      <p class="result-title">
        成功创建 {{ result.successCount }} 条，失败 {{ result.failCount }} 条
      </p>
      <template v-if="result.failURLs && result.failURLs.length > 0">
        <p class="result-hint">以下 URL 创建失败：</p>
        <ul class="fail-list">
          <li v-for="url in result.failURLs" :key="url">{{ url }}</li>
        </ul>
      </template>
      <p v-else class="result-hint">无失败 URL。</p>
    </div>
  </n-card>
</template>

<style scoped>
.card {
  max-width: 720px;
}

.actions {
  display: flex;
  gap: 8px;
}

.result {
  margin-top: 8px;
  padding: 12px 16px;
  border: 1px solid var(--color-border);
  border-radius: 6px;
  background: #fafafa;
  font-size: 13px;
}

.result-title {
  margin: 0 0 8px;
  font-weight: 500;
}

.result-hint {
  margin: 0;
  color: var(--color-text-secondary);
}

.fail-list {
  margin: 6px 0 0;
  padding-left: 18px;
  color: #b23b3b;
  word-break: break-all;
}
</style>
