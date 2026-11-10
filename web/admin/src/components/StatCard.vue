<script setup lang="ts">
import type { Component } from 'vue'

defineProps<{
  label: string
  value: number | string
  icon: Component
  /** 卡片渐变，两个色标 */
  from: string
  to: string
}>()
</script>

<template>
  <div class="stat" :style="{ background: `linear-gradient(135deg, ${from}, ${to})` }">
    <span class="stat-icon">
      <component :is="icon" />
    </span>
    <div class="stat-body">
      <span class="stat-label">{{ label }}</span>
      <span class="stat-value">{{ value }}</span>
    </div>
  </div>
</template>

<style scoped>
.stat {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 20px 22px;
  border-radius: 10px;
  color: #fff;
  box-shadow: 0 4px 14px rgb(0 0 0 / 8%);
}

.stat-icon {
  display: grid;
  flex: none;
  width: 46px;
  height: 46px;
  border-radius: 50%;
  background: rgb(255 255 255 / 22%);
  place-items: center;
}

.stat-icon :deep(svg) {
  width: 22px;
  height: 22px;
}

.stat-body {
  display: flex;
  flex: 1;
  flex-direction: column;
  gap: 2px;
  align-items: flex-end;
  min-width: 0;
}

.stat-label {
  font-size: 13px;
  opacity: 0.9;
}

.stat-value {
  font-size: 22px;
  font-weight: 600;
  font-variant-numeric: tabular-nums;
  line-height: 1.2;
}
</style>
