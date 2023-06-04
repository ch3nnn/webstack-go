package cron

import "github.com/spf13/cast"

func (s *server) RemoveTask(taskId int64) {
	name := cast.ToString(taskId)
	s.cron.RemoveJob(name)
}
