package category

import (
	"github.com/ch3nnn/webstack-go/configs"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/hash"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
	"github.com/ch3nnn/webstack-go/internal/services/category"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 新增网站分类
	// @Tags API.category
	// @Router /api/category [post]
	Create() core.HandlerFunc

	// List  分类列表
	// @Tags API.category
	// @Router /api/category [get]
	List() core.HandlerFunc

	// Delete 删除分类
	// @Tags API.category
	// @Router /api/category/{id} [delete]
	Delete() core.HandlerFunc

	// UpdateUsed 更新分了启用/禁用
	// @Tags API.category
	// @Router /api/category/used [patch]
	UpdateUsed() core.HandlerFunc

	// Detail 获取分类详情数据
	// @Tags API.category
	// @Router /api/category/used [patch]
	Detail() core.HandlerFunc

	// UpdateSort 更新分类排序
	// @Tags API.category
	// @Router /api/category/used [patch]
	UpdateSort() core.HandlerFunc
}

type handler struct {
	logger          *zap.Logger
	cache           redis.Repo
	hashids         hash.Hash
	categoryService category.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:          logger,
		cache:           cache,
		hashids:         hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		categoryService: category.New(db, cache),
	}
}

func (h *handler) i() {}
