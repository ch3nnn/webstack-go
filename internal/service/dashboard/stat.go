package dashboard

import (
	"context"
	"math"
	"time"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
)

// statDays 统计窗口天数（含今天）
const statDays = 7

// recentDays 返回以 today 结尾、连续 n 天的时间列表（时间正序）
func recentDays(today time.Time, n int) []time.Time {
	days := make([]time.Time, 0, n)
	for i := n - 1; i >= 0; i-- {
		days = append(days, today.AddDate(0, 0, -i))
	}
	return days
}

// dayLabel 图表横轴文案：最近三天用口语，更早的用 月-日。
// 用四舍五入算天数，避免两边的时钟时间不同（如 today 是 10:00、day 是 00:00）导致差一天。
func dayLabel(day, today time.Time) string {
	switch int(math.Round(today.Sub(day).Hours() / 24)) {
	case 0:
		return "今天"
	case 1:
		return "昨天"
	case 2:
		return "前天"
	default:
		return day.Format("01-02")
	}
}

// Stat 仪表盘的统计卡片与近 7 天趋势
func (s *service) Stat(ctx context.Context) (*v1.StatResp, error) {
	now := time.Now()
	window := recentDays(now, statDays)

	dates := make([]string, 0, len(window))
	for _, day := range window {
		dates = append(dates, day.Format(repository.VisitDayFormat))
	}

	rows, err := s.visitRepo.WithContext(ctx).FindPVUVByDays(dates)
	if err != nil {
		return nil, err
	}

	counts := make(map[string]repository.DayVisitCount, len(rows))
	for _, row := range rows {
		counts[row.Day] = row
	}

	siteCount, err := s.siteRepo.WithContext(ctx).FindCount()
	if err != nil {
		return nil, err
	}

	totalPV, err := s.visitRepo.WithContext(ctx).FindCount()
	if err != nil {
		return nil, err
	}

	days := make([]v1.StatDay, 0, len(window))
	for i, day := range window {
		c := counts[dates[i]]
		days = append(days, v1.StatDay{
			Date:  dates[i],
			Label: dayLabel(day, now),
			PV:    c.PV,
			UV:    c.UV,
		})
	}

	today := counts[dates[len(dates)-1]]

	return &v1.StatResp{
		SiteCount: siteCount,
		TodayPV:   today.PV,
		TodayUV:   today.UV,
		TotalPV:   totalPV,
		Days:      days,
	}, nil
}
