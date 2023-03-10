package site

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/services/site"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 创建/编辑网站列表
	// @Tags API.site
	// @Router /api/site [post]
	Create() core.HandlerFunc

	// Delete 删除网站列表
	// @Tags API.site
	// @Router /api/site/{id} [delete]
	Delete() core.HandlerFunc

	// UpdateUsed 更新网站为启用/禁用
	// @Tags API.site
	// @Router /api/site/used [patch]
	UpdateUsed() core.HandlerFunc

	// List 网站列表
	// @Tags API.site
	// @Router /api/site [get]
	List() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	cache       redis.Repo
	hashids     hash.Hash
	siteService site.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		hashids:     hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		siteService: site.New(db, cache),
	}
}

func (h *handler) i() {}
