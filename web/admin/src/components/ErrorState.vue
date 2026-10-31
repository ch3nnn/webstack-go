<script setup lang="ts">
import { NButton, NResult } from 'naive-ui'

withDefaults(defineProps<{ description?: string }>(), {
  description: '请稍后重试',
})

const emit = defineEmits<{ retry: [] }>()
</script>

<template>
  <div class="error">
    <n-result status="error" title="加载失败" :description="description">
      <template #footer>
        <n-button size="small" @click="emit('retry')">重试</n-button>
      </template>
    </n-result>
  </div>
</template>

<style scoped>
.error {
  padding: 24px 0;
}
</style>
