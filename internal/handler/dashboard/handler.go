/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:46
 */

package index

import (
	"math/big"
	"net/http"
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
	"github.com/ch3nnn/webstack-go/internal/handler"
)

type Handler struct {
	*handler.Handler
}

func NewHandler(handler *handler.Handler) *Handler {
	return &Handler{Handler: handler}
}

func (h *Handler) memory() (m mem.VirtualMemoryStat) {
	info, err := mem.VirtualMemory()
	if err != nil {
		return m
	}

	return *info
}

func (h *Handler) disk() (d disk.UsageStat) {
	info, err := disk.Usage("/")
	if err != nil {
		return d
	}

	return *info
}

func (h *Handler) cpu() (c cpu.InfoStat) {
	info, err := cpu.Info()
	if err != nil {
		return c
	}

	if len(info) > 0 {
		return info[0]
	}

	return c
}

func (h *Handler) Dashboard(ctx *gin.Context) {
	memoryInfo := h.memory()
	diskInfo := h.disk()
	cpuInfo := h.cpu()

	dir, _ := os.Getwd()

	var cpuPercent float64
	cpuPercents, _ := cpu.Percent(time.Second, false)
	if len(cpuPercents) > 0 {
		cpuPercent = mathutil.RoundToFloat(cpuPercents[0], 2)
	}

	ctx.HTML(http.StatusOK, "dashboard.html", v1.DashboardResp{
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
		CpuName:         cpuInfo.ModelName,
		CpuCores:        cpuInfo.Cores,
		CpuUsedPercent:  cpuPercent,
	})
}
