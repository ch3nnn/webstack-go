<script setup lang="ts">
import { NButton, NEmpty } from 'naive-ui'

withDefaults(defineProps<{ description?: string; actionText?: string }>(), {
  description: '暂无数据',
  actionText: '',
})

const emit = defineEmits<{ action: [] }>()
</script>

<template>
  <div class="empty">
    <n-empty :description="description" size="large">
      <template v-if="actionText" #extra>
        <n-button size="small" secondary @click="emit('action')">{{ actionText }}</n-button>
      </template>
    </n-empty>
  </div>
</template>

<style scoped>
.empty {
  padding: 48px 0;
}
</style>
