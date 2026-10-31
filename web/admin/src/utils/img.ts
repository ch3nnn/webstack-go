/**
 * 把后端存的图片值变成可直接放进 <img src> 的值。
 * 库里可能是：外链、本站相对路径、SVG 的 data URI，或位图的纯 base64。
 * 与后端 pkg/tools.ImgSrc 保持同一套规则。
 */
export function imgSrc(value: string): string {
  if (!value) return ''
  if (value.startsWith('data:') || value.startsWith('http') || value.startsWith('/')) return value
  return `data:image/png;base64,${value}`
}
