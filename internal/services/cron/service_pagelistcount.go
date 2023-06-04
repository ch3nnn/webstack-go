package cron

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

func (s *service) PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error) {
	iCronTaskDo := query.CronTask.WithContext(ctx.RequestContext())
	if searchData.Name != "" {
		iCronTaskDo = iCronTaskDo.Where(query.CronTask.Name.Eq(searchData.Name))
	}
	if searchData.Protocol != 0 {
		iCronTaskDo = iCronTaskDo.Where(query.CronTask.Protocol.Eq(searchData.Protocol))
	}
	if searchData.IsUsed != 0 {
		iCronTaskDo = iCronTaskDo.Where(query.CronTask.IsUsed.Eq(searchData.IsUsed))

	}

	total, err = iCronTaskDo.Count()
	if err != nil {
		return 0, err
	}

	return
}
