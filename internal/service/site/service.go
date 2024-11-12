/**
 * @Author: chentong
 * @Date: 2024/05/27 下午5:56
 */

package site

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/duke-git/lancet/v2/random"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/mat/besticon/besticon"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	s "github.com/ch3nnn/webstack-go/internal/service"
	"github.com/ch3nnn/webstack-go/pkg/tools"
)

// Service 层
var _ Service = (*service)(nil)

type Service interface {
	i()

	List(ctx context.Context, req *v1.SiteListReq) (resp *v1.SiteListResp, err error)
	Delete(ctx context.Context, req *v1.SiteDeleteReq) (resp *v1.SiteDeleteResp, err error)
	Update(ctx *gin.Context, req *v1.SiteUpdateReq) (resp *v1.SiteUpdateResp, err error)
	BatchCreate(ctx context.Context, req *v1.SiteCreateReq) (resp *v1.SiteCreateResp, err error)
	Sync(ctx *gin.Context, req *v1.SiteSyncReq) (resp *v1.SiteSyncResp, err error)
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

const defaultIcon = "/upload/favicon.png"

func getWebLogoIconByURL(url string) string {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	b := besticon.New(
		besticon.WithLogger(besticon.NewDefaultLogger(ioutil.Discard)), // disable verbose logging
		besticon.WithHTTPClient(client),
	)

	icons, err := b.NewIconFinder().FetchIcons(url)
	if err != nil || len(icons) == 0 {
		return defaultIcon
	}

	// get picture format
	var format string
	if ext := filepath.Ext(icons[0].URL); ext != "" {
		format = ext[1:]
	}

	// image save static resource directory
	dst := path.Join("upload", fmt.Sprintf("%s.%s", random.RandNumeralOrLetter(6), format))
	file, err := os.Create(filepath.Join("web", dst))
	if err != nil {
		return defaultIcon
	}
	defer file.Close()

	response, err := client.Get(icons[0].URL)
	if err != nil {
		return defaultIcon
	}
	defer response.Body.Close()

	if _, err := io.Copy(file, response.Body); err != nil {
		return defaultIcon
	}

	return filepath.Join("/", dst)
}

func getWebTitle(url string) (title string) {
	c := tools.NewColly()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title += e.Text
	})
	if err := c.Visit(url); err != nil {
		return
	}

	return
}

func getWebDescription(url string) (doc string) {
	c := tools.NewColly()
	c.OnXML("//meta[@name='description']/@content|//meta[@name='Description']/@content|//meta[@name='DESCRIPTION']",
		func(e *colly.XMLElement) {
			doc += e.Text
		},
	)
	if err := c.Visit(url); err != nil {
		return
	}

	return
}
