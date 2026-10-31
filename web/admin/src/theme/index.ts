import type { GlobalThemeOverrides } from 'naive-ui'

/**
 * 管理端主题基线。
 * 只用 token 定调，禁止在页面里写零散的局部样式覆盖。
 */
export const themeOverrides: GlobalThemeOverrides = {
  common: {
    primaryColor: '#2f6feb',
    primaryColorHover: '#1f5fd8',
    primaryColorPressed: '#1a52bb',
    primaryColorSuppl: '#1f5fd8',
    borderRadius: '6px',
    borderRadiusSmall: '4px',
    fontSize: '13px',
    fontSizeMedium: '13px',
    heightMedium: '32px',
    // 对齐 emil-design-eng：内置 ease 曲线太弱，改用更强的自定义曲线
    cubicBezierEaseOut: 'cubic-bezier(0.23, 1, 0.32, 1)',
    cubicBezierEaseInOut: 'cubic-bezier(0.77, 0, 0.175, 1)',
  },
  DataTable: {
    thPaddingMedium: '0 12px',
    tdPaddingMedium: '0 12px',
    thColor: '#fafafa',
    thTextColor: '#525252',
    thFontWeight: '500',
    tdColorHover: '#f7f8fa',
  },
  Card: {
    paddingMedium: '16px 20px',
  },
}

/** 表格统一行高，列表页与抽屉共用 */
export const TABLE_ROW_HEIGHT = 44
