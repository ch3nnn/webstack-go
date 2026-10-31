/**
 * @Author: chentong
 * @Date: 2025/02/07 20:26
 */

package dashboard

import (
	"golang.org/x/sync/errgroup"
	"math/big"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/duke-git/lancet/v2/mathutil"
	humanize "github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (s *service) Dashboard(ctx *gin.Context) (*v1.DashboardResp, error) {
	var (
		g          errgroup.Group
		dir        string
		cpuPercent float64
		memoryInfo *mem.VirtualMemoryStat
		diskInfo   *disk.UsageStat
		cpuInfo    *cpu.InfoStat
	)

	g.Go(func() (err error) {
		memoryInfo, err = mem.VirtualMemoryWithContext(ctx)
		if err != nil {
			return err
		}
		return
	})
	g.Go(func() (err error) {
		diskInfo, err = disk.UsageWithContext(ctx, "/")
		if err != nil {
			return err
		}
		return
	})
	g.Go(func() (err error) {
		cpuInfos, err := cpu.InfoWithContext(ctx)
		// 部分环境（如 macOS）取不到 CPU 型号，不应让整页指标一起失败
		if err != nil || len(cpuInfos) == 0 {
			return nil
		}
		cpuInfo = &cpuInfos[0]
		return
	})
	g.Go(func() (err error) {
		cpuPercents, err := cpu.PercentWithContext(ctx, time.Second, false)
		if len(cpuPercents) > 0 {
			cpuPercent = mathutil.RoundToFloat(cpuPercents[0], 2)
		}
		return
	})
	g.Go(func() (err error) {
		dir, err = os.Getwd()
		return
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	resp := &v1.DashboardResp{
		ProjectVersion:  "2.0",
		GoOS:            runtime.GOOS,
		GoArch:          runtime.GOARCH,
		GoVersion:       runtime.Version(),
		ProjectPath:     strings.Replace(dir, "\\", "/", -1),
		MemTotal:        humanize.BigBytes(big.NewInt(int64(memoryInfo.Total))),
		MemUsed:         humanize.BigBytes(big.NewInt(int64(memoryInfo.Used))),
		MemUsedPercent:  mathutil.RoundToFloat(memoryInfo.UsedPercent, 2),
		DiskTotal:       humanize.BigBytes(big.NewInt(int64(diskInfo.Total))),
		DiskUsed:        humanize.BigBytes(big.NewInt(int64(diskInfo.Used))),
		DiskUsedPercent: mathutil.RoundToFloat(diskInfo.UsedPercent, 2),
		CpuName:         "unknown",
		CpuCores:        int32(runtime.NumCPU()),
		CpuUsedPercent:  cpuPercent,
	}

	if cpuInfo != nil {
		resp.CpuName = cpuInfo.ModelName
		resp.CpuCores = cpuInfo.Cores
	}

	return resp, nil
}
