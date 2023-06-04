package cron

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

func (s *service) Execute(ctx core.Context, id int64) (err error) {

	cronTask, err := query.CronTask.WithContext(ctx.RequestContext()).Where(query.CronTask.ID.Eq(id)).First()
	if err != nil {
		return err
	}

	cronTask.Spec = "手动执行"
	go s.cronServer.AddJob(cronTask)()

	return nil
}
