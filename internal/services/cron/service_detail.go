package cron

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchOneData struct {
	Id int64 // 任务ID
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (cronTask *model.CronTask, err error) {
	iCronTaskDo := query.CronTask.WithContext(ctx.RequestContext())
	if searchOneData.Id != 0 {
		iCronTaskDo = iCronTaskDo.Where(query.CronTask.ID.Eq(searchOneData.Id))
	}

	cronTask, err = iCronTaskDo.First()
	if err != nil {
		return nil, err
	}

	return
}
