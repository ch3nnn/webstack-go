package cron

import (
	"github.com/ch3nnn/webstack-go/configs"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/hash"
	cronRepo "github.com/ch3nnn/webstack-go/internal/repository/cron"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
	"github.com/ch3nnn/webstack-go/internal/services/cron"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建任务
	// @Tags API.cron
	// @Router /api/cron [post]
	Create() core.HandlerFunc

	// Modify 编辑任务
	// @Tags API.cron
	// @Router /api/cron/{id} [post]
	Modify() core.HandlerFunc

	// List 任务列表
	// @Tags API.cron
	// @Router /api/cron [get]
	List() core.HandlerFunc

	// UpdateUsed 更新任务为启用/禁用
	// @Tags API.cron
	// @Router /api/cron/used [patch]
	UpdateUsed() core.HandlerFunc

	// Detail 获取单条任务详情
	// @Tags API.cron
	// @Router /api/cron/{id} [get]
	Detail() core.HandlerFunc

	// Execute 手动执行任务
	// @Tags API.cron
	// @Router /api/cron/exec/{id} [patch]
	Execute() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       redis.Repo
	hashids     hash.Hash
	cronService cron.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo, cronServer cronRepo.Server) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		cronService: cron.New(db, cache, cronServer),
	}
}

func (h *handler) i() {}
