<script setup lang="ts">
import {
  NButton,
  NCard,
  NDataTable,
  NForm,
  NFormItem,
  NGrid,
  NGridItem,
  NInput,
  NInputNumber,
  NPopconfirm,
  NSelect,
  NSwitch,
  useMessage,
  type DataTableColumns,
} from 'naive-ui'
import { computed, h, onMounted, reactive, ref } from 'vue'
import type { SelectOption } from 'naive-ui'

import {
  createCategory,
  deleteCategory,
  getCategory,
  listCategories,
  updateCategory,
  type Category,
} from '@/api/category'
import { errorMessage } from '@/api/client'
import { LINE_ICONS } from '@/constants/icons'
import EmptyState from '@/components/EmptyState.vue'
import ErrorState from '@/components/ErrorState.vue'

interface CategoryRow extends Category {
  children?: CategoryRow[]
  /** 本次编辑前的排序值，用于判断是否需要提交 */
  originSort?: number
}

const message = useMessage()

const rows = ref<CategoryRow[]>([])
const loading = ref(false)
const failed = ref(false)
const submitting = ref(false)
const editingId = ref<number | null>(null)

const form = reactive({
  parentId: 0,
  name: '',
  icon: 'linecons-globe',
  sortId: 0,
  isUsed: true,
})

const roots = computed(() => rows.value)

const parentOptions = computed(() => [
  { label: '一级分类', value: 0 },
  ...roots.value.map((item) => ({ label: item.name, value: item.id })),
])

const iconOptions = LINE_ICONS.map((icon) => ({ label: icon, value: icon }))

function renderIconLabel(option: SelectOption) {
  return h('span', { class: 'icon-option' }, [
    h('i', { class: String(option.value) }),
    String(option.label),
  ])
}

const isChildForm = computed(() => form.parentId !== 0)

async function load() {
  loading.value = true
  failed.value = false
  try {
    const flat = await listCategories()
    const byId = new Map<number, CategoryRow>()
    flat.forEach((item) => byId.set(item.id, { ...item, originSort: item.sort }))
    const tree: CategoryRow[] = []
    byId.forEach((item) => {
      if (item.pid && byId.has(item.pid)) {
        const parent = byId.get(item.pid)!
        parent.children = [...(parent.children ?? []), item]
      } else {
        tree.push(item)
      }
    })
    tree.forEach((item) => {
      if (item.children) item.children.sort((a, b) => a.sort - b.sort)
    })
    rows.value = tree.sort((a, b) => a.sort - b.sort)
  } catch (error) {
    failed.value = true
    message.error(errorMessage(error))
  } finally {
    loading.value = false
  }
}

function resetForm() {
  editingId.value = null
  form.parentId = 0
  form.name = ''
  form.icon = 'linecons-globe'
  form.sortId = 0
  form.isUsed = true
}

async function onSubmit() {
  if (!form.name.trim()) {
    message.warning('请输入分类名称')
    return
  }
  submitting.value = true
  try {
    if (editingId.value !== null) {
      await updateCategory({
        id: editingId.value,
        parent_id: form.parentId,
        name: form.name.trim(),
        icon: form.icon,
        used: form.isUsed,
        sort_id: form.sortId,
      })
      message.success('修改成功')
    } else {
      await createCategory({
        parent_id: form.parentId,
        level: form.parentId === 0 ? 1 : 2,
        name: form.name.trim(),
        icon: form.icon,
        is_used: form.isUsed,
        sort_id: form.sortId,
      })
      message.success('新增成功')
    }
    resetForm()
    await load()
  } catch (error) {
    message.error(errorMessage(error))
  } finally {
    submitting.value = false
  }
}

async function onAddChild(row: CategoryRow) {
  try {
    const detail = await getCategory(row.id)
    if (!detail.is_add) {
      message.warning('二级分类不可创建分类')
      return
    }
    resetForm()
    form.parentId = detail.id
  } catch (error) {
    message.error(errorMessage(error))
  }
}

async function onEdit(row: CategoryRow) {
  try {
    const detail = await getCategory(row.id)
    editingId.value = detail.id
    form.parentId = detail.pid
    form.name = detail.name
    form.icon = detail.icon || 'linecons-globe'
    form.sortId = detail.sort_id
    form.isUsed = row.is_used
  } catch (error) {
    message.error(errorMessage(error))
  }
}

async function onDelete(row: CategoryRow) {
  try {
    await deleteCategory(row.id)
    message.success('已删除')
    if (editingId.value === row.id) resetForm()
    await load()
  } catch (error) {
    message.error(errorMessage(error))
  }
}

/**
 * 排序在失焦/提交时才落库：输入过程中每敲一个字符都发请求 + 重拉列表，
 * 会把正在编辑的输入框整个换掉。
 * 落库后重新拉列表——列表按 sort 升序排，不重拉就看不到新位置。
 */
async function commitSort(row: CategoryRow) {
  const next = row.sort
  if (next === (row.originSort ?? 0)) return
  try {
    await updateCategory({ id: row.id, sort_id: next })
    row.originSort = next
    message.success(`分类「${row.name}」排序已更新`)
  } catch (error) {
    message.error(errorMessage(error))
  }
  await load()
}

async function onUsedChange(row: CategoryRow, value: boolean) {
  try {
    await updateCategory({ id: row.id, used: value })
    row.is_used = value
  } catch (error) {
    message.error(errorMessage(error))
    await load()
  }
}

const columns: DataTableColumns<CategoryRow> = [
  { title: '名称', key: 'name', minWidth: 140 },
  {
    title: '排序',
    key: 'sort',
    width: 110,
    render: (row) =>
      h(NInputNumber, {
        value: row.sort,
        size: 'small',
        showButton: false,
        min: 0,
        onUpdateValue: (value: number | null) => {
          row.sort = value ?? 0
        },
        onBlur: () => void commitSort(row),
      }),
  },
  {
    title: '图标',
    key: 'icon',
    width: 90,
    render: (row) => h('i', { class: row.icon }),
  },
  {
    title: '启用',
    key: 'is_used',
    width: 80,
    align: 'center',
    render: (row) =>
      h(NSwitch, {
        size: 'small',
        value: row.is_used,
        'onUpdate:value': (value: boolean) => onUsedChange(row, value),
      }),
  },
  {
    title: '操作',
    key: 'actions',
    width: 190,
    render: (row) =>
      h('div', { class: 'actions' }, [
        h(
          NButton,
          { size: 'tiny', quaternary: true, onClick: () => onAddChild(row) },
          { default: () => '新增' },
        ),
        h(
          NButton,
          { size: 'tiny', quaternary: true, onClick: () => onEdit(row) },
          { default: () => '修改' },
        ),
        h(
          NPopconfirm,
          { onPositiveClick: () => onDelete(row) },
          {
            trigger: () => h(NButton, { size: 'tiny', quaternary: true, type: 'error' }, { default: () => '删除' }),
            default: () => `确认删除分类「${row.name}」？`,
          },
        ),
      ]),
  },
]

onMounted(load)
</script>

<template>
  <n-grid :cols="24" :x-gap="12" :y-gap="12" responsive="screen" item-responsive>
    <n-grid-item span="24 l:8">
      <n-card size="small" :title="editingId === null ? '新增分类' : '修改分类'">
        <n-form label-placement="left" label-width="76" @submit.prevent="onSubmit">
          <n-form-item label="上级分类">
            <n-select
              v-model:value="form.parentId"
              :options="parentOptions"
              :disabled="editingId !== null || isChildForm"
            />
          </n-form-item>
          <n-form-item label="分类名称">
            <n-input v-model:value="form.name" placeholder="请输入分类名称" />
          </n-form-item>
          <n-form-item label="分类图标">
            <n-select
              v-model:value="form.icon"
              :options="iconOptions"
              :render-label="renderIconLabel"
              filterable
            />
          </n-form-item>
          <n-form-item label="排序">
            <n-input-number v-model:value="form.sortId" :min="0" class="full" />
          </n-form-item>
          <n-form-item label="是否启用">
            <n-switch v-model:value="form.isUsed" />
          </n-form-item>
          <n-form-item label=" ">
            <div class="form-actions">
              <n-button type="primary" :loading="submitting" attr-type="submit">提交</n-button>
              <n-button v-if="editingId !== null" quaternary @click="resetForm">取消</n-button>
            </div>
          </n-form-item>
        </n-form>
      </n-card>
    </n-grid-item>

    <n-grid-item span="24 l:16">
      <n-card size="small" title="分类列表">
        <ErrorState v-if="failed" @retry="load" />
        <EmptyState v-else-if="!loading && rows.length === 0" description="还没有分类" action-text="新增分类" @action="resetForm" />
        <n-data-table
          v-else
          :columns="columns"
          :data="rows"
          :loading="loading"
          :row-key="(row: CategoryRow) => row.id"
          :bordered="false"
          size="small"
          default-expand-all
        />
      </n-card>
    </n-grid-item>
  </n-grid>
</template>

<style scoped>
.full {
  width: 100%;
}

.form-actions {
  display: flex;
  gap: 8px;
}

:deep(.icon-option) {
  display: inline-flex;
  align-items: center;
  gap: 6px;
}

:deep(.actions) {
  display: flex;
  gap: 2px;
}
</style>
