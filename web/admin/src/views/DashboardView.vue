<script setup lang="ts">
import { NCard, NGrid, NGridItem, NProgress, NSkeleton } from 'naive-ui'
import { onActivated, onDeactivated, onUnmounted, ref } from 'vue'

import ErrorState from '@/components/ErrorState.vue'

/** 后端 DashboardResp 未定义 json tag，字段名即 PascalCase */
interface DashboardData {
  ProjectVersion: string
  GoOS: string
  GoArch: string
  GoVersion: string
  ProjectPath: string
  Host: string
  Env: string
  MemTotal: string
  MemUsed: string
  MemUsedPercent: number
  DiskTotal: string
  DiskUsed: string
  DiskUsedPercent: number
  CpuName: string
  CpuCores: number
  CpuUsedPercent: number
}

const data = ref<DashboardData | null>(null)
const failed = ref(false)
let source: EventSource | null = null

function connect() {
  failed.value = false
  source?.close()
  const es = new EventSource('/api/admin/dashboard')
  source = es
  es.addEventListener('dashboard', (event) => {
    data.value = JSON.parse((event as MessageEvent<string>).data) as DashboardData
  })
  es.onerror = () => {
    es.close()
    source = null
    if (!data.value) failed.value = true
  }
}

connect()
// 页面被 keep-alive 缓存：切走只触发 deactivated，不关连接服务端会一直采样
onActivated(() => {
  if (!source) connect()
})
onDeactivated(() => source?.close())
onUnmounted(() => source?.close())

function percent(value: number | undefined): number {
  if (typeof value !== 'number' || Number.isNaN(value)) return 0
  return Math.min(100, Math.max(0, Number(value.toFixed(1))))
}

const systemRows = [
  ['项目版本', 'ProjectVersion'],
  ['运行环境', 'Env'],
  ['操作系统', 'GoOS'],
  ['架构', 'GoArch'],
  ['Go 版本', 'GoVersion'],
  ['主机名', 'Host'],
  ['CPU', 'CpuName'],
  ['核心数', 'CpuCores'],
  ['项目路径', 'ProjectPath'],
] as const
</script>

<template>
  <ErrorState v-if="failed" @retry="connect" />

  <template v-else>
    <n-grid :cols="3" :x-gap="12" :y-gap="12" responsive="screen" item-responsive>
      <n-grid-item span="3 m:1">
        <n-card size="small" title="内存">
          <n-skeleton v-if="!data" text :repeat="3" />
          <template v-else>
            <n-progress
              type="line"
              :percentage="percent(data.MemUsedPercent)"
              :height="6"
              :show-indicator="false"
            />
            <p class="metric">{{ data.MemUsed }} / {{ data.MemTotal }}</p>
          </template>
        </n-card>
      </n-grid-item>
      <n-grid-item span="3 m:1">
        <n-card size="small" title="磁盘">
          <n-skeleton v-if="!data" text :repeat="3" />
          <template v-else>
            <n-progress
              type="line"
              :percentage="percent(data.DiskUsedPercent)"
              :height="6"
              :show-indicator="false"
            />
            <p class="metric">{{ data.DiskUsed }} / {{ data.DiskTotal }}</p>
          </template>
        </n-card>
      </n-grid-item>
      <n-grid-item span="3 m:1">
        <n-card size="small" title="CPU">
          <n-skeleton v-if="!data" text :repeat="3" />
          <template v-else>
            <n-progress
              type="line"
              :percentage="percent(data.CpuUsedPercent)"
              :height="6"
              :show-indicator="false"
            />
            <p class="metric">{{ percent(data.CpuUsedPercent) }}% · {{ data.CpuCores }} 核</p>
          </template>
        </n-card>
      </n-grid-item>
    </n-grid>

    <n-card size="small" title="系统信息" class="system">
      <n-skeleton v-if="!data" text :repeat="6" />
      <dl v-else class="rows">
        <template v-for="[label, field] in systemRows" :key="field">
          <dt>{{ label }}</dt>
          <dd>{{ data[field] }}</dd>
        </template>
      </dl>
    </n-card>
  </template>
</template>

<style scoped>
.metric {
  margin: 10px 0 0;
  font-size: 13px;
  color: var(--color-text-secondary);
  font-variant-numeric: tabular-nums;
}

.system {
  margin-top: 12px;
}

.rows {
  display: grid;
  grid-template-columns: 120px 1fr;
  row-gap: 8px;
  margin: 0;
  font-size: 13px;
}

.rows dt {
  color: var(--color-text-secondary);
}

.rows dd {
  margin: 0;
  color: #1f2329;
  word-break: break-all;
}
</style>
