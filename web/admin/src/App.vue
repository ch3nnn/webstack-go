<script setup lang="ts">
import { NConfigProvider, NDialogProvider, NMessageProvider, dateZhCN, zhCN } from 'naive-ui'

import { themeOverrides } from '@/theme'
</script>

<template>
  <n-config-provider :theme-overrides="themeOverrides" :locale="zhCN" :date-locale="dateZhCN">
    <n-message-provider>
      <n-dialog-provider>
        <div class="boot">前端工程已就绪</div>
      </n-dialog-provider>
    </n-message-provider>
  </n-config-provider>
</template>

<style scoped>
.boot {
  padding: 32px;
  color: var(--color-text-secondary);
}
</style>
