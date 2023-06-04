package cron

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchData struct {
	Page     int64  // 第几页
	PageSize int64  // 每页显示条数
	Name     string // 任务名称
	Protocol int64  // 执行方式
	IsUsed   int64  // 是否启用
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (cronTasks []*model.CronTask, err error) {
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
	cronTasks, _, err = iCronTaskDo.Order(query.CronTask.ID.Desc()).FindByPage(int((searchData.Page-1)*searchData.PageSize), int(searchData.PageSize))
	if err != nil {
		return nil, err
	}

	return
}
