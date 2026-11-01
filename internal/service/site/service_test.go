/**
 * @Author: chentong
 * @Date: 2025/02/09 11:03
 */

package site

import (
	"flag"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/duke-git/lancet/v2/pointer"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	s "github.com/ch3nnn/webstack-go/internal/service"
	"github.com/ch3nnn/webstack-go/pkg/config"
	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/ch3nnn/webstack-go/pkg/log"
)

var (
	logger *log.Logger
	j      *jwt.JWT
)

func setupRepository(t *testing.T) (*repository.Repository, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm connection: %v", err)
	}

	query.SetDefault(db)

	return repository.NewRepository(logger, db), mock
}

func TestMain(m *testing.M) {
	err := os.Setenv("APP_CONF", "../../../config/local.yml")
	if err != nil {
		panic(err)
	}

	envConf := flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger = log.NewLog(conf)
	j = jwt.NewJwt(conf)

	code := m.Run()
	os.Exit(code)
}

func TestSiteService_Delete(t *testing.T) {
	req := &v1.SiteDeleteReq{
		ID: 1,
	}

	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	siteDao := repository.NewMockIStSiteDao(ctrl)
	categoryDao := repository.NewMockIStCategoryDao(ctrl)

	siteFunc := repository.NewMockiCustomGenStSiteFunc(ctrl)
	siteDao.EXPECT().WithContext(gomock.Any()).Return(siteFunc)
	siteDao.EXPECT().WhereByID(gomock.Any()).Return(func(dao gen.Dao) gen.Dao {
		return dao.Where(query.SysUser.ID.Eq(req.ID))
	})
	siteFunc.EXPECT().Delete(gomock.Any()).Return(nil)

	siteService := NewService(srv, siteDao, categoryDao)
	resp, err := siteService.Delete(ctx, req)

	// 断言
	assert.NoError(t, err)
	assert.Equal(t, &v1.SiteDeleteResp{ID: req.ID}, resp)
}

func TestSiteService_BatchCreate(t *testing.T) {
	req := &v1.SiteCreateReq{
		CategoryID: 1,
		Url:        "www.123.com",
		IsUsed:     true,
		FailSwitch: true,
	}

	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	siteDao := repository.NewMockIStSiteDao(ctrl)
	categoryDao := repository.NewMockIStCategoryDao(ctrl)

	siteFunc := repository.NewMockiCustomGenStSiteFunc(ctrl)
	siteDao.EXPECT().WithContext(gomock.Any()).Return(siteFunc)
	siteFunc.EXPECT().Create(gomock.Any()).Return(nil, nil)

	siteService := NewService(srv, siteDao, categoryDao)
	resp, err := siteService.BatchCreate(ctx, req)

	// 断言
	assert.NoError(t, err)
	assert.Equal(t, &v1.SiteCreateResp{
		SuccessCount: 1,
		FailCount:    0,
		FailURLs:     nil,
	}, resp)
}

func TestSiteService_Export(t *testing.T) {
	req := &v1.SiteExportReq{
		Search:     "webstack-go",
		CategoryID: 1,
	}
	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	siteDao := repository.NewMockIStSiteDao(ctrl)
	categoryDao := repository.NewMockIStCategoryDao(ctrl)

	siteFunc := repository.NewMockiCustomGenStSiteFunc(ctrl)
	siteDao.EXPECT().WithContext(gomock.Any()).Return(siteFunc)
	siteDao.EXPECT().WhereByCategoryID(gomock.Any()).Return(func(dao gen.Dao) gen.Dao {
		return dao
	})
	siteDao.EXPECT().LikeInByTitleOrDescOrURL(gomock.Any()).Return(func(dao gen.Dao) gen.Dao {
		return dao
	})

	var siteCategories []repository.SiteCategory
	siteFunc.EXPECT().FindSiteCategoryWithPage(gomock.Any(), gomock.Any(), &siteCategories, gomock.Any(), gomock.Any()).Return(int64(1), nil)

	siteService := NewService(srv, siteDao, categoryDao)
	_, err := siteService.Export(ctx, req)
	assert.NoError(t, err)
}

func TestSiteService_List(t *testing.T) {
	req := &v1.SiteListReq{
		Page:       1,
		PageSize:   1,
		Search:     "webstack-go",
		CategoryID: 1,
	}
	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	siteDao := repository.NewMockIStSiteDao(ctrl)
	categoryDao := repository.NewMockIStCategoryDao(ctrl)

	siteFunc := repository.NewMockiCustomGenStSiteFunc(ctrl)
	siteDao.EXPECT().WithContext(gomock.Any()).Return(siteFunc)
	siteDao.EXPECT().WhereByCategoryID(gomock.Any()).Return(func(dao gen.Dao) gen.Dao {
		return dao
	})
	siteDao.EXPECT().LikeInByTitleOrDescOrURL(gomock.Any()).Return(func(dao gen.Dao) gen.Dao {
		return dao
	})
	var siteCategories []repository.SiteCategory
	siteFunc.EXPECT().FindSiteCategoryWithPage(gomock.Any(), gomock.Any(), &siteCategories, gomock.Any(), gomock.Any()).Return(int64(1), nil)

	siteService := NewService(srv, siteDao, categoryDao)
	_, err := siteService.List(ctx, req)
	assert.NoError(t, err)
}

func TestSiteService_Sync(t *testing.T) {
	req := &v1.SiteSyncReq{
		ID: 1,
	}
	site := &model.StSite{
		ID:          1,
		CategoryID:  1,
		Title:       "site",
		Icon:        "https://www.baidu.com",
		Description: "site",
		URL:         "https://www.baidu.com",
		IsUsed:      false,
		CreatedAt:   nil,
		UpdatedAt:   nil,
		DeletedAt:   nil,
		Sort:        1,
	}
	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	siteDao := repository.NewMockIStSiteDao(ctrl)
	categoryDao := repository.NewMockIStCategoryDao(ctrl)

	siteFunc := repository.NewMockiCustomGenStSiteFunc(ctrl)
	siteDao.EXPECT().WithContext(gomock.Any()).Return(siteFunc)
	siteDao.EXPECT().WhereByID(gomock.Any()).Return(func(dao gen.Dao) gen.Dao {
		return dao
	})

	siteFunc.EXPECT().FindOne(gomock.Any()).Return(site, nil)

	siteDao.EXPECT().WithContext(gomock.Any()).Return(siteFunc)
	siteDao.EXPECT().WhereByID(gomock.Any()).Return(func(dao gen.Dao) gen.Dao {
		return dao
	})
	siteFunc.EXPECT().Update(gomock.Any(), gomock.Any()).Return(int64(1), nil)

	siteService := NewService(srv, siteDao, categoryDao)
	_, err := siteService.Sync(ctx, req)
	assert.NoError(t, err)
}

// 页面里除文档级 <title> 外，内联 SVG 的 <title> 也会命中 colly 的 OnHTML("title")。
// 曾经用 title += e.Text 累加，导致标题变成整页所有标题的拼接（geogebra.org 抓到 368 字）。
func TestGetWebTitle(t *testing.T) {
	const longTitle = "GeoGebra - the world's favorite, free math tools used by over 100 million students and teachers"

	tests := []struct {
		name    string
		page    string
		want    string
		wantLen int
	}{
		{
			name: "只取首个 title，不拼接内联 SVG 的 title",
			page: `<html><head><title>GitHub · Change is constant.</title></head><body>
				<svg><title>Duolingo</title><title>FedEx</title></svg></body></html>`,
			want:    "GitHub · Change is constant.",
			wantLen: 28,
		},
		{
			name: "标题超长时按 rune 截断到 50",
			page: `<html><head><title>  ` + longTitle + `  </title></head><body>
				<svg><title>Download_on_the_App_Store_Badge_US-UK_RGB_blk_4SVG_092917</title></svg></body></html>`,
			want:    string([]rune(longTitle)[:50]),
			wantLen: 50,
		},
		{
			name:    "文档标题为空白时不退而取 SVG 标题",
			page:    `<html><head><title>   </title></head><body><svg><title>Download</title></svg></body></html>`,
			want:    "",
			wantLen: 0,
		},
		{
			name:    "没有文档标题时不取 SVG 标题",
			page:    `<html><head></head><body><svg><title>Duolingo</title></svg></body></html>`,
			want:    "",
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				_, _ = w.Write([]byte(tt.page))
			}))
			defer srv.Close()

			title, err := getWebTitle(srv.URL)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, title)
			assert.Equal(t, tt.wantLen, len([]rune(title)))
		})
	}
}

func TestSiteService_Update(t *testing.T) {
	req := &v1.SiteUpdateReq{
		Id:          1,
		Icon:        "site",
		Title:       "site",
		Url:         "site",
		CategoryId:  1,
		Description: "site",
		IsUsed:      pointer.Of(true),
		File:        &multipart.FileHeader{},
		Sort:        pointer.Of(1),
	}

	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	siteDao := repository.NewMockIStSiteDao(ctrl)
	categoryDao := repository.NewMockIStCategoryDao(ctrl)

	siteFunc := repository.NewMockiCustomGenStSiteFunc(ctrl)
	siteDao.EXPECT().WithContext(gomock.Any()).Return(siteFunc)
	siteDao.EXPECT().WhereByID(gomock.Any()).Return(func(dao gen.Dao) gen.Dao {
		return dao
	})
	siteFunc.EXPECT().Update(gomock.Any(), gomock.Any()).Return(int64(1), nil)

	siteService := NewService(srv, siteDao, categoryDao)
	_, err := siteService.Update(ctx, req)
	assert.NoError(t, err)
}
