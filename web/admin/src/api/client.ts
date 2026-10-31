import axios from 'axios'

export interface ApiError {
  code: number
  message: string
  data?: unknown
}

export const client = axios.create({ baseURL: '/api', timeout: 30000 })

client.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error?.response?.status === 401) {
      document.cookie = '_login_token_=; path=/; max-age=0'
      if (location.pathname !== '/login') location.replace('/login')
    }
    return Promise.reject(error)
  },
)

/** 后端错误体统一为 { code, message, data } */
export function errorMessage(error: unknown): string {
  const body = (error as { response?: { data?: ApiError } })?.response?.data
  if (body?.message) return body.message
  const message = (error as Error)?.message
  return message || '请求失败'
}
