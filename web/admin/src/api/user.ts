import md5 from 'js-md5'

import { client } from './client'

export interface MenuItem {
  id: number
  pid: number
  name: string
  link: string
  icon: string
}

export interface UserInfo {
  username: string
  menu: MenuItem[]
}

/** 登录密码沿用旧前端行为：先 md5 再提交 */
export async function login(username: string, password: string): Promise<string> {
  const { data } = await client.post<{ token: string }>('/login', {
    username,
    password: md5(password),
  })
  return data.token
}

export async function fetchInfo(): Promise<UserInfo> {
  const { data } = await client.get<UserInfo>('/admin/info')
  return data
}

export async function logout(): Promise<void> {
  await client.post('/admin/logout')
}
