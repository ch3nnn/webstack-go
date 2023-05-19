package site

import (
	"crypto/tls"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/tools"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/site"
	"github.com/gocolly/colly"
	"github.com/mat/besticon/besticon"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type CreateSiteData struct {
	CategoryId int32  `json:"category_id"`
	Url        string `json:"Url"`
}

// 获取网站 logo
func getWebLogoIconUrlByUrl(site *site.Site) string {

	// https 不安全跳过验证
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	b := besticon.New(
		besticon.WithLogger(besticon.NewDefaultLogger(ioutil.Discard)), // 禁用详细日志记录
		besticon.WithHTTPClient(client),                                // 设置用于请求的http客户端
	)
	icons, err := b.NewIconFinder().FetchIcons(site.Url)
	if err != nil || len(icons) == 0 {
		return ""
	}
	// 获取图片格式
	var format string
	if ext := filepath.Ext(icons[0].URL); ext != "" {
		format = ext[1:]
	}
	// 图片保存静态资源目录
	dst := path.Join("/upload/" + site.Title + "." + format)
	file, err := os.Create(path.Join("assets", dst))
	if err != nil {
		return ""
	}
	defer file.Close()

	response, err := client.Get(icons[0].URL)
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	if _, err := io.Copy(file, response.Body); err != nil {
		return ""
	}

	return dst

}

// getWebTitle 获取网站标题
func getWebTitle(site *site.Site) string {
	var title string
	c := tools.NewColly()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title += e.Text
	})
	if err := c.Visit(site.Url); err != nil {
		return ""
	}
	return title

}

// getWebDescription 获取网站描述
func getWebDescription(site *site.Site) string {
	var description string
	c := tools.NewColly()
	c.OnXML("//meta[@name='description']/@content|//meta[@name='Description']/@content|//meta[@name='DESCRIPTION']",
		func(e *colly.XMLElement) {
			description += e.Text
		},
	)
	if err := c.Visit(site.Url); err != nil {
		return ""
	}
	return description
}

func (s *service) Create(ctx core.Context, siteData *CreateSiteData) (id int32, err error) {

	model := site.NewModel()
	model.IsUsed = -1
	model.CategoryId = siteData.CategoryId
	model.Url = siteData.Url
	model.Title = getWebTitle(model)
	model.Description = getWebDescription(model)
	model.Thumb = getWebLogoIconUrlByUrl(model)

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
