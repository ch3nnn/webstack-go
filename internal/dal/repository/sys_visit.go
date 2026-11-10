package repository

import (
	"context"

	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

var _ iCustomGenSysVisitFunc = (*customSysVisitDao)(nil)

type (
	// ISysVisitDao not edit interface name
	ISysVisitDao interface {
		iWhereSysVisitFunc
		iGenSysVisitFunc
		WithContext(ctx context.Context) iCustomGenSysVisitFunc
	}

	// not edit interface name
	iCustomGenSysVisitFunc interface {
		iGenSysVisitFunc

		// Record 记录一次首页访问
		Record(visit *model.SysVisit) error
		// FindPVUVByDays 按天汇总指定日期集合的 PV / UV
		FindPVUVByDays(days []string) ([]DayVisitCount, error)
	}

	// not edit interface name
	customSysVisitDao struct {
		sysVisitDao
	}
)

// DayVisitCount 某一天的访问量：PV 为访问次数，UV 为独立 IP 数
type DayVisitCount struct {
	Day string
	PV  int64
	UV  int64
}

// VisitDayFormat 访问统计按天分桶使用的日期格式
const VisitDayFormat = "2006-01-02"

func NewSysVisitDao() ISysVisitDao {
	return &customSysVisitDao{
		sysVisitDao{
			sysVisitDo: query.SysVisit.WithContext(context.Background()),
		},
	}
}

func (d *customSysVisitDao) WithContext(ctx context.Context) iCustomGenSysVisitFunc {
	d.sysVisitDo = d.sysVisitDo.WithContext(ctx)
	return d
}

// Record 记录一次首页访问。ip 由调用方做哈希后再传进来，库里不留明文。
func (d *customSysVisitDao) Record(visit *model.SysVisit) error {
	return d.sysVisitDo.Create(visit)
}

// FindPVUVByDays 一次分组查询拿到每天的 PV 与 UV，没数据的日期不会出现在结果里
func (d *customSysVisitDao) FindPVUVByDays(days []string) ([]DayVisitCount, error) {
	if len(days) == 0 {
		return nil, nil
	}

	var rows []DayVisitCount
	err := d.sysVisitDo.
		Select(
			query.SysVisit.Day.As("day"),
			query.SysVisit.ID.Count().As("pv"),
			query.SysVisit.IP.Distinct().Count().As("uv"),
		).
		Where(query.SysVisit.Day.In(days...)).
		Group(query.SysVisit.Day).
		Scan(&rows)

	return rows, err
}
