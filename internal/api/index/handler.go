package index

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/services/category"
	"github.com/xinliangnote/go-gin-api/internal/services/index"
	"github.com/xinliangnote/go-gin-api/internal/services/site"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Index 导航网站首页
	// @Tags API.admin
	// @Router / [get]
	Index() core.HandlerFunc
}

type handler struct {
	logger          *zap.Logger
	cache           redis.Repo
	hashids         hash.Hash
	indexService    index.Service
	categoryService category.Service
	siteService     site.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:          logger,
		cache:           cache,
		hashids:         hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		indexService:    index.New(db, cache),
		categoryService: category.New(db, cache),
		siteService:     site.New(db, cache),
	}
}

func (h *handler) i() {}
