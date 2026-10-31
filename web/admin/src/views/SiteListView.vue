<script setup lang="ts">
import {
  NButton,
  NCard,
  NDataTable,
  NForm,
  NFormItem,
  NInput,
  NInputNumber,
  NModal,
  NPagination,
  NPopconfirm,
  NSelect,
  NSwitch,
  NTooltip,
  NUpload,
  useMessage,
  type DataTableColumns,
  type UploadFileInfo,
} from 'naive-ui'
import { computed, h, onMounted, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'

import { listCategories, type Category } from '@/api/category'
import { errorMessage } from '@/api/client'
import {
  deleteSite,
  exportUrl,
  listSites,
  syncSite,
  updateSite,
  type Site,
} from '@/api/site'
import EmptyState from '@/components/EmptyState.vue'
import ErrorState from '@/components/ErrorState.vue'

const message = useMessage()
const router = useRouter()

const rows = ref<Site[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
const loading = ref(false)
const failed = ref(false)
const search = ref('')
const categoryId = ref<number | null>(null)
const categories = ref<Category[]>([])

const detailVisible = ref(false)
const detailRow = ref<Site | null>(null)

const editVisible = ref(false)
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
})

const categoryOptions = computed(() => {
  const options: { label: string; value: number }[] = []
  categories.value
    .filter((item) => !item.pid)
    .forEach((root) => {
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

async function load() {
  loading.value = true
  failed.value = false
  try {
    const data = await listSites({
      page: page.value,
      page_size: pageSize.value,
      search: search.value || undefined,
      category_id: categoryId.value,
    })
    rows.value = data.list
    total.value = data.pagination?.total ?? 0
  } catch (error) {
    failed.value = true
    message.error(errorMessage(error))
  } finally {
    loading.value = false
  }
}

function onSearch() {
  page.value = 1
  load()
}

function onExport() {
  window.location.href = exportUrl({ search: search.value, category_id: categoryId.value })
}

async function onSync(row: Site) {
  try {
    await syncSite(row.id)
    message.success(`编号 ${row.id} 已同步`)
    await load()
  } catch (error) {
    message.error(errorMessage(error))
  }
}

async function onUsedChange(row: Site, value: boolean) {
  try {
    await updateSite({ id: row.id, is_used: value })
    row.is_used = value
    message.success(`编号 ${row.id} 已${value ? '启用' : '禁用'}`)
  } catch (error) {
    message.error(errorMessage(error))
    await load()
  }
}

async function onSortChange(row: Site, value: number) {
  try {
    await updateSite({ id: row.id, sort: value })
    row.sort = value
  } catch (error) {
    message.error(errorMessage(error))
    await load()
  }
}

async function onDelete(row: Site) {
  try {
    await deleteSite(row.id)
    message.success(`编号 ${row.id} 已删除`)
    await load()
  } catch (error) {
    message.error(errorMessage(error))
  }
}

function onDetail(row: Site) {
  detailRow.value = row
  detailVisible.value = true
}

function onEdit(row: Site) {
  edit.id = row.id
  edit.title = row.title
  edit.url = row.url
  edit.description = row.description
  edit.categoryId = row.category_id || null
  edit.sort = row.sort
  edit.thumb = ''
  edit.file = null
  editVisible.value = true
}

function onFileChange(options: { file: UploadFileInfo }) {
  edit.file = options.file.file ?? null
}

async function onSave() {
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
    editVisible.value = false
    await load()
  } catch (error) {
    message.error(errorMessage(error))
  } finally {
    saving.value = false
  }
}

function logoUrl(row: Site): string {
  if (!row.icon) return ''
  return row.icon.startsWith('data:') ? row.icon : `data:image/png;base64,${row.icon}`
}

const columns: DataTableColumns<Site> = [
  { title: 'ID', key: 'id', width: 66, fixed: 'left' },
  {
    title: 'Logo',
    key: 'icon',
    width: 62,
    render: (row) => h('img', { src: logoUrl(row), class: 'logo', alt: '' }),
  },
  {
    title: '名称简介',
    key: 'title',
    width: 220,
    fixed: 'left',
    render: (row) =>
      h(NTooltip, null, {
        trigger: () => h('span', { class: 'truncate' }, row.title),
        default: () => row.description || row.title,
      }),
  },
  {
    title: '链接',
    key: 'url',
    width: 220,
    render: (row) =>
      h(
        'a',
        { href: row.url, target: '_blank', rel: 'noreferrer', class: 'truncate link' },
        row.url,
      ),
  },
  { title: '分类', key: 'category', width: 110 },
  { title: '创建日期', key: 'created_at', width: 160 },
  { title: '更新日期', key: 'updated_at', width: 160 },
  {
    title: '状态',
    key: 'is_used',
    width: 76,
    align: 'center',
    render: (row) =>
      h(NSwitch, {
        size: 'small',
        value: row.is_used,
        'onUpdate:value': (value: boolean) => onUsedChange(row, value),
      }),
  },
  {
    title: '排序',
    key: 'sort',
    width: 92,
    align: 'center',
    render: (row) =>
      h(NInputNumber, {
        value: row.sort,
        size: 'small',
        showButton: false,
        min: 0,
        onUpdateValue: (value: number | null) => onSortChange(row, value ?? 0),
      }),
  },
  {
    title: '操作',
    key: 'actions',
    width: 236,
    fixed: 'right',
    render: (row) =>
      h('div', { class: 'actions' }, [
        h(NButton, { size: 'tiny', quaternary: true, onClick: () => onSync(row) }, { default: () => '同步' }),
        h(NButton, { size: 'tiny', quaternary: true, onClick: () => onDetail(row) }, { default: () => '详情' }),
        h(NButton, { size: 'tiny', quaternary: true, onClick: () => onEdit(row) }, { default: () => '编辑' }),
        h(
          NPopconfirm,
          { onPositiveClick: () => onDelete(row) },
          {
            trigger: () =>
              h(NButton, { size: 'tiny', quaternary: true, type: 'error' }, { default: () => '删除' }),
            default: () => `确认删除「${row.title}」？`,
          },
        ),
      ]),
  },
]

function rowProps(row: Site) {
  // 拉取失败时后端用 url 兜底成 title，旧页面用高亮提示需要手工修正
  return row.url === row.title ? { class: 'row-warn' } : {}
}

onMounted(async () => {
  await loadCategories()
  await load()
})
</script>

<template>
  <n-card size="small">
    <div class="toolbar">
      <div class="toolbar-left">
        <n-button type="primary" @click="router.push('/admin/site/add')">新增</n-button>
        <n-button attr-type="button" @click="onExport">导出</n-button>
      </div>
      <div class="toolbar-right">
        <n-input
          v-model:value="search"
          placeholder="请输入网址标题..."
          clearable
          class="search"
          @keyup.enter="onSearch"
        />
        <n-select
          v-model:value="categoryId"
          :options="categoryOptions"
          clearable
          placeholder="全部分类"
          class="filter"
          @update:value="onSearch"
        />
        <n-button secondary @click="onSearch">搜索</n-button>
      </div>
    </div>

    <p class="notice">
      注：勾选「拉取失败时仅保存网址」创建的行，背景会有高亮，记得点编辑修正。
    </p>

    <ErrorState v-if="failed" @retry="load" />
    <EmptyState
      v-else-if="!loading && rows.length === 0"
      description="暂无网站数据"
      action-text="新增网站"
      @action="router.push('/admin/site/add')"
    />
    <template v-else>
      <n-data-table
        :columns="columns"
        :data="rows"
        :loading="loading"
        :row-key="(row: Site) => row.id"
        :row-props="rowProps"
        :bordered="false"
        size="small"
        :scroll-x="1400"
        :single-line="false"
      />
      <div class="pager">
        <n-pagination
          v-model:page="page"
          v-model:page-size="pageSize"
          :item-count="total"
          :page-sizes="[10, 20, 50]"
          show-size-picker
          @update:page="load"
          @update:page-size="onSearch"
        />
      </div>
    </template>
  </n-card>

  <n-modal v-model:show="detailVisible" preset="card" title="网站详情" class="modal">
    <p class="detail-row"><span>名称</span>{{ detailRow?.title }}</p>
    <p class="detail-row"><span>链接</span>{{ detailRow?.url }}</p>
    <p class="detail-row"><span>分类</span>{{ detailRow?.category }}</p>
    <p class="detail-row"><span>描述</span>{{ detailRow?.description }}</p>
  </n-modal>

  <n-modal v-model:show="editVisible" preset="card" title="编辑网站" class="modal">
    <n-form label-placement="left" label-width="76">
      <n-form-item label="分类">
        <n-select v-model:value="edit.categoryId" :options="categoryOptions" clearable />
      </n-form-item>
      <n-form-item label="Logo">
        <div class="logo-field">
          <n-input v-model:value="edit.thumb" placeholder="请输入 Logo 链接" />
          <n-upload :default-upload="false" :show-file-list="false" accept=".jpg,.png" @change="onFileChange">
            <n-button attr-type="button">选择文件</n-button>
          </n-upload>
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
    <template #footer>
      <div class="modal-footer">
        <n-button quaternary @click="editVisible = false">取消</n-button>
        <n-button type="primary" :loading="saving" @click="onSave">确认</n-button>
      </div>
    </template>
  </n-modal>
</template>

<style scoped>
.toolbar {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 12px;
}

.toolbar-left,
.toolbar-right {
  display: flex;
  gap: 8px;
  align-items: center;
}

.search {
  width: 200px;
}

.filter {
  width: 160px;
}

.notice {
  margin: 0 0 12px;
  padding: 8px 12px;
  border-radius: 6px;
  background: #f2f6ff;
  color: #3d5b9e;
  font-size: 12px;
}

.logo {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  object-fit: cover;
}

.truncate {
  display: inline-block;
  overflow: hidden;
  max-width: 100%;
  text-overflow: ellipsis;
  white-space: nowrap;
  vertical-align: bottom;
}

.link {
  color: #2f6feb;
}

.actions {
  display: flex;
  gap: 2px;
}

.pager {
  display: flex;
  justify-content: flex-end;
  margin-top: 12px;
}

.modal {
  width: 560px;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
}

.logo-field {
  display: flex;
  gap: 8px;
  width: 100%;
}

.detail-row {
  display: flex;
  gap: 12px;
  margin: 0 0 10px;
  font-size: 13px;
  word-break: break-all;
}

.detail-row span {
  flex: 0 0 48px;
  color: var(--color-text-secondary);
}

:deep(.row-warn td) {
  background: #fffaf0;
}
</style>
