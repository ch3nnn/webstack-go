import { BarChart, LineChart } from 'echarts/charts'
import { GridComponent, TooltipComponent } from 'echarts/components'
import * as echarts from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import type { EChartsOption } from 'echarts'

// 按需注册：只引柱状图与折线图，避免把整包 echarts 塞进产物
echarts.use([BarChart, LineChart, GridComponent, TooltipComponent, CanvasRenderer])

export { echarts }
export type { EChartsOption }

/** 折线/柱状图共用的基础配置：细网格、灰字、去掉动画以贴合后台的低干扰风格 */
export const baseChartOption: EChartsOption = {
  // 后台图表每 1s 可能被刷新，动画会让数字“追”得比数据还慢，直接关掉
  animation: false,
  grid: { left: 8, right: 16, top: 32, bottom: 8, containLabel: true },
  tooltip: { trigger: 'axis' },
  textStyle: { fontSize: 12, color: '#6b7280' },
}
