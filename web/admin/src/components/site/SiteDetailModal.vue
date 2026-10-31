<script setup lang="ts">
import { NModal } from 'naive-ui'

import type { Site } from '@/api/site'

defineProps<{
  show: boolean
  row: Site | null
}>()

const emit = defineEmits<{ 'update:show': [value: boolean] }>()
</script>

<template>
  <n-modal
    :show="show"
    preset="card"
    title="网站详情"
    :style="{ width: '520px' }"
    @update:show="emit('update:show', $event)"
  >
    <p class="detail-row"><span>名称</span>{{ row?.title }}</p>
    <p class="detail-row"><span>链接</span>{{ row?.url }}</p>
    <p class="detail-row"><span>分类</span>{{ row?.category }}</p>
    <p class="detail-row"><span>描述</span>{{ row?.description }}</p>
  </n-modal>
</template>

<style scoped>
.detail-row {
  display: flex;
  gap: 12px;
  margin: 0 0 10px;
  font-size: 13px;
  word-break: break-all;
}

.detail-row span {
  flex: 0 0 48px;
  color: var(--color-text-secondary);
}
</style>
