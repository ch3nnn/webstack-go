<script setup lang="ts">
import {
  NButton,
  NCheckbox,
  NForm,
  NFormItem,
  NInput,
  NModal,
  NRadio,
  NRadioGroup,
  NSelect,
  useMessage,
} from 'naive-ui'
import { reactive, ref, watch } from 'vue'

import { errorMessage } from '@/api/client'
import { createSites, type SiteCreateResult } from '@/api/site'

const props = defineProps<{
  show: boolean
  categoryOptions: { label: string; value: number }[]
}>()

const emit = defineEmits<{
  'update:show': [value: boolean]
  /** 创建成功，父组件需要重拉列表 */
  done: []
}>()

const message = useMessage()

const creating = ref(false)
const result = ref<SiteCreateResult | null>(null)
const form = reactive({
  categoryId: null as number | null,
  url: '',
  isUsed: true,
  failSwitch: false,
})

function reset() {
  form.categoryId = null
  form.url = ''
  form.isUsed = true
  form.failSwitch = false
  result.value = null
}

// 每次打开都从空白表单开始
watch(
  () => props.show,
  (visible) => {
    if (visible) reset()
  },
)

function close() {
  emit('update:show', false)
  result.value = null
}

async function submit() {
  if (!form.url.trim()) {
    message.warning('请输入网站链接')
    return
  }

  creating.value = true
  result.value = null
  try {
    result.value = await createSites({
      category_id: form.categoryId,
      url: form.url.trim(),
      is_used: form.isUsed,
      fail_switch: form.failSwitch,
    })
    emit('done')
    if (result.value.failCount === 0) {
      message.success(`成功创建 ${result.value.successCount} 条`)
      form.url = ''
    }
  } catch (error) {
    message.error(errorMessage(error))
  } finally {
    creating.value = false
  }
}
</script>

<template>
  <n-modal
    :show="show"
    preset="card"
    title="新增网站"
    :style="{ width: '560px' }"
    :mask-closable="false"
    @update:show="emit('update:show', $event)"
  >
    <div class="modal-body">
      <n-form label-placement="left" label-width="96">
        <n-form-item label="分类">
          <n-select
            v-model:value="form.categoryId"
            :options="categoryOptions"
            clearable
            placeholder="请选择分类"
          />
        </n-form-item>
        <n-form-item label="网址">
          <n-input
            v-model:value="form.url"
            type="textarea"
            :autosize="{ minRows: 6, maxRows: 12 }"
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
      </n-form>
      <div v-if="result" class="create-result">
        <p class="create-result__title">
          成功创建 {{ result.successCount }} 条，失败 {{ result.failCount }} 条
        </p>
        <template v-if="result.failURLs?.length">
          <p class="create-result__hint">以下 URL 创建失败：</p>
          <ul class="create-result__list">
            <li v-for="url in result.failURLs" :key="url">{{ url }}</li>
          </ul>
        </template>
        <p v-else class="create-result__hint">无失败 URL。</p>
      </div>
    </div>
    <template #footer>
      <div class="modal-footer">
        <n-button quaternary @click="close">关闭</n-button>
        <n-button type="primary" :loading="creating" @click="submit">提交</n-button>
      </div>
    </template>
  </n-modal>
</template>

<style scoped>
.create-result {
  margin-top: 4px;
  padding: 10px 12px;
  border: 1px solid var(--color-border);
  border-radius: 6px;
  background: #fafafa;
  font-size: 13px;
}

.create-result__title {
  margin: 0 0 6px;
  font-weight: 500;
}

.create-result__hint {
  margin: 0;
  color: var(--color-text-secondary);
}

.create-result__list {
  margin: 4px 0 0;
  padding-left: 18px;
  color: #b23b3b;
  word-break: break-all;
}
</style>
