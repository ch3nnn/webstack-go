<script setup lang="ts">
import {
  NButton,
  NCard,
  NDataTable,
  NInput,
  NInputNumber,
  NPagination,
  NPopconfirm,
  NSelect,
  NSwitch,
  NTooltip,
  useMessage,
  type DataTableColumns,
} from 'naive-ui'
import { computed, h, onActivated, onMounted, ref } from 'vue'

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
import SiteCreateModal from '@/components/site/SiteCreateModal.vue'
import SiteDetailModal from '@/components/site/SiteDetailModal.vue'
import SiteEditModal from '@/components/site/SiteEditModal.vue'
import { imgSrc } from '@/utils/img'

const message = useMessage()

const rows = ref<SiteRow[]>([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
const loading = ref(false)
const failed = ref(false)
const search = ref('')
/** 列表按创建时间倒序；只有带分类过滤时后端才按 sort 升序，那种情况下重拉才看得到新位置 */
type SiteRow = Site & { originSort: number }
const categoryId = ref<number | null>(null)
const categories = ref<Category[]>([])
let categoryRequest = 0
let initialActivation = true

const createVisible = ref(false)

const detailVisible = ref(false)
const detailRow = ref<Site | null>(null)

const editVisible = ref(false)
const editRow = ref<Site | null>(null)

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
  const request = ++categoryRequest
  try {
    const result = await listCategories()
    if (request === categoryRequest) categories.value = result
  } catch (error) {
    if (request === categoryRequest) message.error(errorMessage(error))
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
    rows.value = data.list.map((item) => ({ ...item, originSort: item.sort }))
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

async function commitSort(row: SiteRow) {
  const next = row.sort
  if (next === row.originSort) return
  try {
    await updateSite({ id: row.id, sort: next })
    row.originSort = next
  } catch (error) {
    message.error(errorMessage(error))
    await load()
  }
  if (categoryId.value) await load()
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
  editRow.value = row
  editVisible.value = true
}

const LOGO_SIZE = 24
/** 表格里没有 data-v，尺寸只能内联；样式恒定，提到模块级避免逐行拼接 */
const LOGO_STYLE =
  `width:${LOGO_SIZE}px;height:${LOGO_SIZE}px;border-radius:50%;object-fit:contain;` +
  `display:block;margin:0 auto;background:#f2f3f5`

const columns: DataTableColumns<SiteRow> = [
  { title: 'ID', key: 'id', width: 66, fixed: 'left' },
  {
    title: 'Logo',
    key: 'icon',
    width: 62,
    align: 'center',
    render: (row) => {
      // 列渲染函数在 NDataTable 的渲染上下文里执行，拿不到本组件的 data-v，
      // 作用域样式对它无效，尺寸必须内联
      const src = imgSrc(row.icon)
      return src ? h('img', { src, style: LOGO_STYLE, alt: '' }) : h('span', { style: LOGO_STYLE })
    },
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
        onUpdateValue: (value: number | null) => {
          row.sort = value ?? 0
        },
        onBlur: () => void commitSort(row),
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

/**
 * 高亮 = 「拉取失败的自动兜底」且「用户还没动过」。
 *
 * 抓取失败时后端把 title 兜底成 url（batchcreate.go 的 condition.Ternary），
 * 所以 url == title 是「标题仍是原始网址」的信号；但只看它的话，用户改完标题
 * 之外的字段（描述、Logo、分类…）高亮不会消失。加上 updated_at == created_at，
 * 意味着这行自创建后没被编辑过——保存过一次编辑就会让它消失。
 */
function rowProps(row: Site) {
  const untouched = row.updated_at === row.created_at
  return row.url === row.title && untouched ? { class: 'row-warn' } : {}
}

onActivated(() => {
  if (initialActivation) {
    initialActivation = false
    return
  }
  void loadCategories()
})
onMounted(async () => {
  await loadCategories()
  await load()
})
</script>

<template>
  <n-card size="small">
    <div class="toolbar">
      <div class="toolbar-left">
        <n-button type="primary" @click="createVisible = true">新增</n-button>
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
      注：抓取失败的行（标题仍是原始网址）会有高亮，保存一次编辑后自动消失。
    </p>

    <ErrorState v-if="failed" @retry="load" />
    <EmptyState
      v-else-if="!loading && rows.length === 0"
      description="暂无网站数据"
      action-text="新增网站"
      @action="createVisible = true"
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

  <SiteCreateModal v-model:show="createVisible" :category-options="categoryOptions" @done="load" />
  <SiteDetailModal v-model:show="detailVisible" :row="detailRow" />

  <!-- class 不会落到 naive 的弹框根节点上，宽度必须走 style（否则会铺满整个窗口） -->
  <SiteEditModal
    v-model:show="editVisible"
    :site="editRow"
    :category-options="categoryOptions"
    @saved="load"
  />
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

:deep(.truncate) {
  display: inline-block;
  overflow: hidden;
  max-width: 100%;
  text-overflow: ellipsis;
  white-space: nowrap;
  vertical-align: bottom;
}

:deep(.link) {
  color: #2f6feb;
}

.pager {
  display: flex;
  justify-content: flex-end;
  margin-top: 12px;
}

</style>
