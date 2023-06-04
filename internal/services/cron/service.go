package cron

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/cron"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, createData *CreateCronTaskData) (id int64, err error)
	Modify(ctx core.Context, id int64, modifyData *ModifyCronTaskData) (err error)
	PageList(ctx core.Context, searchData *SearchData) (cronTasks []*model.CronTask, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id, used int64) (err error)
	Execute(ctx core.Context, id int64) (err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (cronTask *model.CronTask, err error)
}

type service struct {
	db         mysql.Repo
	cache      redis.Repo
	cronServer cron.Server
}

func New(db mysql.Repo, cache redis.Repo, cron cron.Server) Service {
	return &service{
		db:         db,
		cache:      cache,
		cronServer: cron,
	}
}

func (s *service) i() {}
