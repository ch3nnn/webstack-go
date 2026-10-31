import { client } from './client'

export interface SiteConfig {
  id: number
  about_site: string
  about_author: string
  is_about: boolean
  /** 是否在首页显示搜索框 */
  is_search: boolean
  site_title: string
  site_keyword: string
  site_desc: string
  site_record: string
  site_url: string
  site_logo: string
  site_favicon: string
}

export async function fetchConfig(): Promise<SiteConfig> {
  const { data } = await client.get<SiteConfig>('/admin/config')
  return data
}

/** 每项都是可选的：只提交改动过的那一项，未传的字段后端按 nil 跳过 */
export interface SiteSettingsPayload {
  is_search?: boolean
  site_title?: string
  site_keyword?: string
  site_desc?: string
  site_record?: string
  site_url?: string
  logo?: File | null
  favicon?: File | null
}

/** 只提交传进来的字段，其余留在库里不动 */
export async function updateSiteSettings(payload: SiteSettingsPayload): Promise<void> {
  const form = new FormData()
  if (payload.is_search !== undefined) form.append('is_search', String(payload.is_search))
  if (payload.site_title !== undefined) form.append('site_title', payload.site_title)
  if (payload.site_keyword !== undefined) form.append('site_keyword', payload.site_keyword)
  if (payload.site_desc !== undefined) form.append('site_desc', payload.site_desc)
  if (payload.site_record !== undefined) form.append('site_record', payload.site_record)
  if (payload.site_url !== undefined) form.append('site_url', payload.site_url)
  if (payload.logo) form.append('logo', payload.logo)
  if (payload.favicon) form.append('favicon', payload.favicon)

  await client.put('/admin/config', form)
}

export interface AboutPayload {
  about_site: string
  about_author: string
  is_about: boolean
}

export async function updateAbout(payload: AboutPayload): Promise<void> {
  await client.put('/admin/config', payload)
}
