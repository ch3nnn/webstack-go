import { client } from './client'

export interface Category {
  id: number
  pid: number
  name: string
  link: string
  icon: string
  is_used: boolean
  sort: number
  level: number
}

export async function listCategories(): Promise<Category[]> {
  const { data } = await client.get<{ list: Category[] }>('/admin/category')
  return data.list ?? []
}

export interface CategoryDetail {
  id: number
  pid: number
  name: string
  icon: string
  /** 是否允许在该节点下新增子分类（仅一级分类为 true） */
  is_add: boolean
  sort_id: number
}

export async function getCategory(id: number): Promise<CategoryDetail> {
  const { data } = await client.get<CategoryDetail>(`/admin/category/${id}`)
  return data
}

export interface CategoryCreatePayload {
  parent_id: number
  level: number
  name: string
  icon: string
  is_used: boolean
  sort_id: number
}

export async function createCategory(payload: CategoryCreatePayload): Promise<void> {
  await client.post('/admin/category', payload)
}

export interface CategoryUpdatePayload {
  id: number
  parent_id?: number
  name?: string
  icon?: string
  used?: boolean
  sort_id?: number
}

export async function updateCategory(payload: CategoryUpdatePayload): Promise<void> {
  await client.put('/admin/category/update', payload)
}

export async function deleteCategory(id: number): Promise<void> {
  await client.delete(`/admin/category/${id}`)
}
