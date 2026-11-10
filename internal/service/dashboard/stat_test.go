package dashboard

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRecentDays(t *testing.T) {
	today := time.Date(2026, 11, 1, 10, 30, 0, 0, time.Local)
	days := recentDays(today, 7)

	assert.Len(t, days, 7)
	// 时间正序，最后一个是今天
	assert.Equal(t, "2026-10-26", days[0].Format("2006-01-02"))
	assert.Equal(t, "2026-11-01", days[6].Format("2006-01-02"))
	assert.Equal(t, today.Format("2006-01-02"), days[6].Format("2006-01-02"))
}

func TestDayLabel(t *testing.T) {
	today := time.Date(2026, 11, 1, 10, 0, 0, 0, time.Local)

	cases := map[string]string{
		"2026-11-01": "今天",
		"2026-10-31": "昨天",
		"2026-10-30": "前天",
		"2026-10-29": "10-29",
	}
	for date, want := range cases {
		day, err := time.ParseInLocation("2006-01-02", date, time.Local)
		assert.NoError(t, err)
		assert.Equal(t, want, dayLabel(day, today), date)
	}
}

func TestRecentDays_跨月跨年(t *testing.T) {
	today := time.Date(2027, 1, 2, 0, 0, 0, 0, time.Local)
	days := recentDays(today, 7)

	assert.Equal(t, "2026-12-27", days[0].Format("2006-01-02"))
	assert.Equal(t, "2027-01-02", days[6].Format("2006-01-02"))
	// 跨年不影响口语标签：01-01 是昨天，12-30 才回落到 月-日
	assert.Equal(t, "昨天", dayLabel(days[5], today))
	assert.Equal(t, "12-30", dayLabel(days[3], today))
}
