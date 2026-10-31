<script setup lang="ts">
import {
  NButton,
  NForm,
  NFormItem,
  NInput,
  NInputNumber,
  NModal,
  NSelect,
  NUpload,
  useMessage,
  type UploadFileInfo,
} from 'naive-ui'
import { computed, reactive, ref, watch } from 'vue'

import { errorMessage } from '@/api/client'
import { updateSite, type Site } from '@/api/site'
import { imgSrc } from '@/utils/img'

const props = defineProps<{
  show: boolean
  site: Site | null
  categoryOptions: { label: string; value: number }[]
}>()

const emit = defineEmits<{
  'update:show': [value: boolean]
  /** 保存成功，父组件需要重拉列表 */
  saved: []
}>()

const message = useMessage()

const saving = ref(false)
const edit = reactive({
  id: 0,
  title: '',
  url: '',
  description: '',
  categoryId: null as number | null,
  sort: 0,
  thumb: '',
  file: null as File | null,
  /** 当前行的 logo（base64 / data URI），用于先展示已有图标 */
  icon: '',
  /** 选中文件后的本地预览地址（object URL，用完要 revoke） */
  fileUrl: '',
  fileName: '',
})

/** 预览优先级：刚选的文件 > 正在输入的链接 > 当前已有图标 */
const logoPreview = computed(() => {
  if (edit.fileUrl) return edit.fileUrl
  const typed = edit.thumb.trim()
  if (typed) return typed
  return imgSrc(edit.icon)
})

function resetLogoPreview() {
  if (edit.fileUrl) URL.revokeObjectURL(edit.fileUrl)
  edit.fileUrl = ''
  edit.fileName = ''
}

// 打开时用传入的行填充表单；关闭时（含点 X / 遮罩）回收 object URL
watch(
  () => props.show,
  (visible) => {
    if (!visible) {
      resetLogoPreview()
      edit.file = null
      return
    }
    const row = props.site
    if (!row) return
    edit.id = row.id
    edit.title = row.title
    edit.url = row.url
    edit.description = row.description
    edit.categoryId = row.category_id || null
    edit.sort = row.sort
    edit.thumb = ''
    edit.file = null
    edit.icon = row.icon
    resetLogoPreview()
  },
)

function close() {
  emit('update:show', false)
}

function onFileChange(options: { file: UploadFileInfo }) {
  const file = options.file.file ?? null
  edit.file = file
  resetLogoPreview()
  if (!file) return
  edit.fileUrl = URL.createObjectURL(file)
  edit.fileName = file.name
}

async function save() {
  if (!edit.title.trim()) {
    message.warning('请输入网站名称')
    return
  }
  if (!edit.url.trim()) {
    message.warning('请输入网站链接')
    return
  }
  if (!edit.description.trim()) {
    message.warning('请输入网站描述')
    return
  }

  saving.value = true
  try {
    await updateSite({
      id: edit.id,
      title: edit.title.trim(),
      url: edit.url.trim(),
      description: edit.description.trim(),
      category_id: edit.categoryId ?? undefined,
      sort: edit.sort,
      thumb: edit.thumb.trim() || undefined,
      file: edit.file ?? undefined,
    })
    message.success(`编号 ${edit.id} 修改完成`)
    close()
    emit('saved')
  } catch (error) {
    message.error(errorMessage(error))
  } finally {
    saving.value = false
  }
}
</script>

<template>
  <n-modal
    :show="show"
    preset="card"
    title="编辑网站"
    :style="{ width: '560px' }"
    @update:show="emit('update:show', $event)"
  >
    <div class="modal-body">
      <n-form label-placement="left" label-width="76">
        <n-form-item label="分类">
          <n-select v-model:value="edit.categoryId" :options="categoryOptions" clearable />
        </n-form-item>
        <n-form-item label="Logo">
          <div class="logo-field">
            <img v-if="logoPreview" :src="logoPreview" class="logo-preview" alt="" />
            <span v-else class="logo-preview logo-preview--empty" />
            <div class="logo-inputs">
              <n-input v-model:value="edit.thumb" placeholder="请输入 Logo 链接" />
              <div class="logo-actions">
                <n-upload
                  :default-upload="false"
                  :show-file-list="false"
                  accept=".jpg,.jpeg,.png,.svg"
                  @change="onFileChange"
                >
                  <n-button attr-type="button">选择文件</n-button>
                </n-upload>
                <span class="logo-hint">{{ edit.fileName || '未选择新文件（PNG / JPG / SVG）' }}</span>
              </div>
            </div>
          </div>
        </n-form-item>
        <n-form-item label="名称">
          <n-input v-model:value="edit.title" placeholder="请输入网站名称" />
        </n-form-item>
        <n-form-item label="链接">
          <n-input v-model:value="edit.url" placeholder="请输入网站地址" />
        </n-form-item>
        <n-form-item label="描述">
          <n-input v-model:value="edit.description" type="textarea" :autosize="{ minRows: 2, maxRows: 4 }" />
        </n-form-item>
        <n-form-item label="排序">
          <n-input-number v-model:value="edit.sort" :min="0" />
        </n-form-item>
      </n-form>
    </div>
    <template #footer>
      <div class="modal-footer">
        <n-button quaternary @click="close">取消</n-button>
        <n-button type="primary" :loading="saving" @click="save">确认</n-button>
      </div>
    </template>
  </n-modal>
</template>

<style scoped>
.logo-field {
  display: flex;
  gap: 10px;
  width: 100%;
  align-items: flex-start;
}

.logo-preview {
  flex: 0 0 auto;
  width: 44px;
  height: 44px;
  border-radius: 8px;
  border: 1px solid var(--color-border);
  object-fit: contain;
  background: #fafafa;
}

.logo-preview--empty {
  display: block;
}

.logo-inputs {
  display: flex;
  flex-direction: column;
  gap: 6px;
  flex: 1 1 auto;
  min-width: 0;
}

.logo-actions {
  display: flex;
  gap: 8px;
  align-items: center;
}

/* n-upload 默认撑满一行，会把文件名挤到最右边 */
.logo-actions :deep(.n-upload) {
  width: auto;
}

.logo-hint {
  color: var(--color-text-secondary);
  font-size: 12px;
  max-width: 220px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
