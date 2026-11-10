import { client } from './client'

export interface StatDay {
  date: string
  label: string
  pv: number
  uv: number
}

export interface StatData {
  site_count: number
  today_pv: number
  today_uv: number
  total_pv: number
  days: StatDay[]
}

/** 仪表盘统计：链接数量、今日 PV/UV、累计 PV 与近 7 天趋势 */
export async function fetchStat(): Promise<StatData> {
  const { data } = await client.get<StatData>('/admin/stat')
  return data
}
