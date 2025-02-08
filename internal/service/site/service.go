/**
 * @Author: chentong
 * @Date: 2024/05/27 下午5:56
 */

package site

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/mat/besticon/besticon"
	"github.com/pkg/errors"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	s "github.com/ch3nnn/webstack-go/internal/service"
	"github.com/ch3nnn/webstack-go/pkg/tools"
)

// Service 层
var _ Service = (*service)(nil)

type Service interface {
	i()

	// List 站点列表
	List(ctx context.Context, req *v1.SiteListReq) (resp *v1.SiteListResp, err error)
	// Delete 删除站点
	Delete(ctx context.Context, req *v1.SiteDeleteReq) (resp *v1.SiteDeleteResp, err error)
	// Update 更新站点
	Update(ctx *gin.Context, req *v1.SiteUpdateReq) (resp *v1.SiteUpdateResp, err error)
	// BatchCreate 批量创建站点
	BatchCreate(ctx context.Context, req *v1.SiteCreateReq) (resp *v1.SiteCreateResp, err error)
	// Sync 同步站点信息
	Sync(ctx *gin.Context, req *v1.SiteSyncReq) (resp *v1.SiteSyncResp, err error)
	// Export 导出站点信息
	Export(ctx *gin.Context, req *v1.SiteExportReq) (resp *v1.SiteExportResp, err error)
}

type service struct {
	*s.Service
	siteRepository     repository.IStSiteDao
	categoryRepository repository.IStCategoryDao
}

func NewService(s *s.Service) Service {
	return &service{
		Service:            s,
		siteRepository:     repository.NewStSiteDao(),
		categoryRepository: repository.NewStCategoryDao(),
	}
}

func (s *service) i() {}

// -----------------------------------------------------------------------------------------------------------------------------------------

func getWebLogoIconBase64(url string) (string, error) {
	client := &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	b := besticon.New(
		besticon.WithHTTPClient(client),
		besticon.WithLogger(besticon.NewDefaultLogger(io.Discard)), // disable verbose logging
	)

	finder := b.NewIconFinder()
	if _, err := finder.FetchIcons(url); err != nil {
		return repository.DefaultFaviconBase64, err
	}

	icon := finder.IconInSizeRange(besticon.SizeRange{Min: 16, Perfect: 64, Max: 512})
	if icon == nil {
		return repository.DefaultFaviconBase64, errors.New("failed to fetch icons")
	}

	return base64.StdEncoding.EncodeToString(icon.ImageData), nil
}

func getWebTitle(url string) (title string, err error) {
	c := tools.NewColly()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title += e.Text
	})
	if err := c.Visit(url); err != nil {
		return "", err
	}

	return
}

func getWebDescription(url string) (doc string, err error) {
	c := tools.NewColly()
	c.OnXML("//meta[@name='description']/@content|//meta[@name='Description']/@content|//meta[@name='DESCRIPTION']",
		func(e *colly.XMLElement) {
			doc += e.Text
		},
	)
	if err := c.Visit(url); err != nil {
		return "", err
	}

	return
}
