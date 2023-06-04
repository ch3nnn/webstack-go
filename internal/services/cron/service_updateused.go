package cron

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/constant"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

func (s *service) UpdateUsed(ctx core.Context, id, used int64) (err error) {
	_, err = query.CronTask.WithContext(ctx.RequestContext()).
		Where(query.CronTask.ID.Eq(id)).
		UpdateColumnSimple(
			query.CronTask.IsUsed.Value(used),
			query.CronTask.UpdatedUser.Value(ctx.SessionUserInfo().UserName),
		)
	if err != nil {
		return err
	}

	// region 操作定时任务 避免主从同步延迟，在这需要查询主库
	if used == constant.IsUsedNo {
		s.cronServer.RemoveTask(id)
	} else {
		cronTask, err := query.CronTask.WithContext(ctx.RequestContext()).Where(query.CronTask.ID.Eq(id)).First()
		if err != nil {
			return err
		}

		s.cronServer.RemoveTask(id)
		s.cronServer.AddTask(cronTask)

	}

	return
}
