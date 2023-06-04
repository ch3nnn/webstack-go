package cron

import (
	"fmt"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
	"math"

	"go.uber.org/zap"
)

func (s *server) Start() {
	s.cron.Start()
	go s.taskCount.Wait()

	totalNum, err := query.CronTask.Where(query.CronTask.IsUsed.Eq(1)).Count()
	if err != nil {
		s.logger.Fatal("cron initialize tasks count err", zap.Error(err))
	}

	pageSize := 50
	maxPage := int(math.Ceil(float64(totalNum) / float64(pageSize)))

	taskNum := 0
	s.logger.Info("开始初始化后台任务")

	for page := 1; page <= maxPage; page++ {
		cronTasks, err := query.CronTask.Where(query.CronTask.IsUsed.Eq(1)).
			Limit(pageSize).
			Offset((page - 1) * pageSize).
			Order(query.CronTask.ID).
			Find()

		if err != nil {
			s.logger.Fatal("cron initialize tasks list err", zap.Error(err))
		}

		for _, cronTask := range cronTasks {
			s.AddTask(cronTask)
			taskNum++
		}
	}

	s.logger.Info(fmt.Sprintf("后台任务初始化完成，总数量：%d", taskNum))
}
