<script setup lang="ts">
import { onBeforeUnmount, onMounted, ref, watch } from 'vue'

import { echarts, type EChartsOption } from './echarts'

const props = withDefaults(
  defineProps<{
    option: EChartsOption
    /** 图表高度，随卡片大小调整 */
    height?: number
  }>(),
  { height: 260 },
)

const el = ref<HTMLDivElement | null>(null)
let chart: ReturnType<typeof echarts.init> | null = null
let observer: ResizeObserver | null = null

onMounted(() => {
  if (!el.value) return
  chart = echarts.init(el.value)
  chart.setOption(props.option)
  // 容器宽度随窗口变化时重排，避免图表按初始宽度定格
  observer = new ResizeObserver(() => chart?.resize())
  observer.observe(el.value)
})

watch(
  () => props.option,
  (option) => chart?.setOption(option, true),
)

onBeforeUnmount(() => {
  observer?.disconnect()
  observer = null
  chart?.dispose()
  chart = null
})
</script>

<template>
  <div ref="el" class="chart" :style="{ height: `${height}px` }" />
</template>

<style scoped>
.chart {
  width: 100%;
}
</style>
