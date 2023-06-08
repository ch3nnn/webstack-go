package site

import (
	"github.com/ch3nnn/webstack-go/configs"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/hash"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
	"github.com/ch3nnn/webstack-go/internal/services/site"

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

	// CategoryList 新增网站-分类目录列表
	// @Tags API.site
	// @Router /api/site/category_list [get]
	CategoryList() core.HandlerFunc

	// UpdateSite 编辑网站
	// @Tags API.site
	// @Router /api/site [put]
	UpdateSite() core.HandlerFunc
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
