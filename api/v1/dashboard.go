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
