<script setup lang="ts">
import {
  EyeOutline,
  LinkOutline,
  PeopleOutline,
  PersonOutline,
} from '@vicons/ionicons5'
import { NCard, NGrid, NGridItem, NProgress, NSkeleton } from 'naive-ui'
import { computed, onActivated, onDeactivated, onMounted, onUnmounted, ref } from 'vue'

import { errorMessage } from '@/api/client'
import { fetchStat, type StatData } from '@/api/stat'
import BaseChart from '@/components/chart/BaseChart.vue'
import { baseChartOption, type EChartsOption } from '@/components/chart/echarts'
import ErrorState from '@/components/ErrorState.vue'
import StatCard from '@/components/StatCard.vue'

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

onMounted(loadStat)
connect()
// 页面被 keep-alive 缓存：切走只触发 deactivated，不关连接服务端会一直采样
onActivated(() => {
  if (!source) connect()
  loadStat()
})
onDeactivated(() => source?.close())
onUnmounted(() => source?.close())

function percent(value: number | undefined): number {
  if (typeof value !== 'number' || Number.isNaN(value)) return 0
  return Math.min(100, Math.max(0, Number(value.toFixed(1))))
}

/* ---- 统计卡片与近 7 天趋势 ---- */
const stat = ref<StatData | null>(null)
const statFailed = ref(false)
const statLoading = ref(false)

async function loadStat() {
  statLoading.value = true
  statFailed.value = false
  try {
    stat.value = await fetchStat()
  } catch (error) {
    statFailed.value = true
    console.error(errorMessage(error))
  } finally {
    statLoading.value = false
  }
}

const cards = computed(() => [
  { label: '链接数量', value: stat.value?.site_count ?? 0, icon: LinkOutline, from: '#4f8ef7', to: '#2563eb' },
  { label: '今日浏览量', value: stat.value?.today_pv ?? 0, icon: PersonOutline, from: '#f26b6b', to: '#dc2626' },
  { label: '今日独立 IP', value: stat.value?.today_uv ?? 0, icon: PeopleOutline, from: '#34c05e', to: '#16a34a' },
  { label: '累计浏览量', value: stat.value?.total_pv ?? 0, icon: EyeOutline, from: '#a97bf5', to: '#7c3aed' },
])

const dayLabels = computed(() => stat.value?.days.map((day) => day.label) ?? [])

/** 近 7 天独立 IP：柱状图 */
const uvOption = computed<EChartsOption>(() => ({
  ...baseChartOption,
  legend: { data: ['独立 IP'], top: 0, right: 0, itemWidth: 12, itemHeight: 8, textStyle: { color: '#6b7280' } },
  xAxis: { type: 'category', data: dayLabels.value, axisLine: { lineStyle: { color: '#e5e7eb' } }, axisTick: { show: false } },
  yAxis: { type: 'value', minInterval: 1, splitLine: { lineStyle: { color: '#f0f1f4' } } },
  series: [
    {
      name: '独立 IP',
      type: 'bar',
      barMaxWidth: 36,
      data: stat.value?.days.map((day) => day.uv) ?? [],
      itemStyle: { color: '#7aa7f8', borderRadius: [4, 4, 0, 0] },
    },
  ],
}))

/** 近 7 天浏览量：折线图 */
const pvOption = computed<EChartsOption>(() => ({
  ...baseChartOption,
  legend: { data: ['浏览量(PV)'], top: 0, right: 0, itemWidth: 12, itemHeight: 8, textStyle: { color: '#6b7280' } },
  xAxis: { type: 'category', data: dayLabels.value, boundaryGap: false, axisLine: { lineStyle: { color: '#e5e7eb' } }, axisTick: { show: false } },
  yAxis: { type: 'value', minInterval: 1, splitLine: { lineStyle: { color: '#f0f1f4' } } },
  series: [
    {
      name: '浏览量(PV)',
      type: 'line',
      smooth: false,
      symbolSize: 6,
      data: stat.value?.days.map((day) => day.pv) ?? [],
      lineStyle: { color: '#7aa7f8', width: 2 },
      itemStyle: { color: '#7aa7f8' },
    },
  ],
}))

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
    <!-- 统计卡片 -->
    <n-grid :cols="4" :x-gap="12" :y-gap="12" responsive="screen" item-responsive>
      <n-grid-item v-for="card in cards" :key="card.label" span="4 s:2 m:1">
        <stat-card
          :label="card.label"
          :value="card.value"
          :icon="card.icon"
          :from="card.from"
          :to="card.to"
        />
      </n-grid-item>
    </n-grid>

    <!-- 近 7 天趋势 -->
    <n-grid class="charts" :cols="2" :x-gap="12" :y-gap="12" responsive="screen" item-responsive>
      <n-grid-item span="2 m:1">
        <n-card size="small" title="近7天独立 IP 统计">
          <n-skeleton v-if="statLoading && !stat" text :repeat="6" />
          <base-chart v-else :option="uvOption" />
          <p v-if="statFailed" class="stat-failed">统计加载失败</p>
        </n-card>
      </n-grid-item>
      <n-grid-item span="2 m:1">
        <n-card size="small" title="近7天浏览量统计">
          <n-skeleton v-if="statLoading && !stat" text :repeat="6" />
          <base-chart v-else :option="pvOption" />
          <p v-if="statFailed" class="stat-failed">统计加载失败</p>
        </n-card>
      </n-grid-item>
    </n-grid>

    <n-grid class="metrics" :cols="3" :x-gap="12" :y-gap="12" responsive="screen" item-responsive>
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
/* 卡片 → 图表 → 资源占用 → 系统信息，区块之间保持同一条纵向间距 */
.charts,
.metrics,
.system {
  margin-top: 16px;
}

.stat-failed {
  margin: 0;
  color: #d03050;
  font-size: 12px;
}

.metric {
  margin: 10px 0 0;
  font-size: 13px;
  color: var(--color-text-secondary);
  font-variant-numeric: tabular-nums;
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
