import { client } from './client'

export interface SiteConfig {
  id: number
  about_site: string
  about_author: string
  is_about: boolean
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

export interface SiteSettingsPayload {
  site_title: string
  site_keyword: string
  site_desc: string
  site_record: string
  site_url: string
  logo?: File | null
  favicon?: File | null
}

export async function updateSiteSettings(payload: SiteSettingsPayload): Promise<void> {
  const form = new FormData()
  form.append('site_title', payload.site_title)
  form.append('site_keyword', payload.site_keyword)
  form.append('site_desc', payload.site_desc)
  form.append('site_record', payload.site_record)
  form.append('site_url', payload.site_url)
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
