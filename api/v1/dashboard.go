/**
 * @Author: chentong
 * @Date: 2024/11/12 12:32
 */

package v1

type DashboardResp struct {
	ProjectVersion  string  // 项目版本
	GoOS            string  // 操作系统
	GoArch          string  // 架构
	GoVersion       string  // go版本
	ProjectPath     string  // 项目路径
	Host            string  // 主机名
	Env             string  // 环境
	MemTotal        string  // 内存
	MemUsed         string  // 已用内存
	MemUsedPercent  float64 // 已用内存百分比
	DiskTotal       string  // 磁盘
	DiskUsed        string  // 已用磁盘
	DiskUsedPercent float64 // 已用磁盘百分比
	CpuName         string  // cpu名称
	CpuCores        int32   // cpu核数
	CpuUsedPercent  float64 // cpu使用率
}

type (
	// StatResp 仪表盘统计卡片与近 7 天趋势
	StatResp struct {
		SiteCount int64     `json:"site_count"` // 链接数量（站点总数）
		TodayPV   int64     `json:"today_pv"`   // 今日浏览量
		TodayUV   int64     `json:"today_uv"`   // 今日独立 IP
		TotalPV   int64     `json:"total_pv"`   // 累计浏览量
		Days      []StatDay `json:"days"`       // 近 7 天，时间正序
	}

	// StatDay 某一天的访问数据
	StatDay struct {
		Date  string `json:"date"`  // 2026-11-01
		Label string `json:"label"` // 今天 / 昨天 / 前天 / 10-30
		PV    int64  `json:"pv"`    // 浏览量
		UV    int64  `json:"uv"`    // 独立 IP
	}
)
