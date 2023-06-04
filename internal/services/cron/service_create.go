package cron

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type CreateCronTaskData struct {
	Name                string // 任务名称
	Spec                string // crontab 表达式
	Command             string // 执行命令
	Protocol            int64  // 执行方式 1:shell 2:http
	HttpMethod          int64  // http 请求方式 1:get 2:post
	Timeout             int64  // 超时时间(单位:秒)
	RetryTimes          int64  // 重试次数
	RetryInterval       int64  // 重试间隔(单位:秒)
	NotifyStatus        int64  // 执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知
	NotifyType          int64  // 通知类型 1:邮件 2:webhook
	NotifyReceiverEmail string // 通知者邮箱地址(多个用,分割)
	NotifyKeyword       string // 通知匹配关键字(多个用,分割)
	Remark              string // 备注
	IsUsed              int64  // 是否启用 1:是  -1:否
}

func (s *service) Create(ctx core.Context, createData *CreateCronTaskData) (id int64, err error) {
	task := model.CronTask{
		Name:                createData.Name,
		Spec:                createData.Spec,
		Command:             createData.Command,
		Protocol:            createData.Protocol,
		HTTPMethod:          createData.HttpMethod,
		Timeout:             createData.Timeout,
		RetryTimes:          createData.RetryTimes,
		RetryInterval:       createData.RetryInterval,
		NotifyStatus:        createData.NotifyStatus,
		NotifyType:          createData.NotifyType,
		NotifyReceiverEmail: createData.NotifyReceiverEmail,
		NotifyKeyword:       createData.NotifyKeyword,
		Remark:              createData.Remark,
		IsUsed:              createData.IsUsed,
		CreatedUser:         ctx.SessionUserInfo().UserName,
	}
	if err = query.CronTask.Create(&task); err != nil {
		return 0, err
	}

	s.cronServer.AddTask(&task)

	return
}
