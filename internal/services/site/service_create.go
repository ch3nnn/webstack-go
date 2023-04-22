package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/tools"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/site"
	"github.com/gocolly/colly"
	"github.com/mat/besticon/besticon"
	"io/ioutil"
)

type CreateSiteData struct {
	CategoryId int32  `json:"category_id"`
	Url        string `json:"Url"`
}

// 获取网站 logo
func getWebLogoIconUrlByUrl(siteData *CreateSiteData) string {
	b := besticon.New(besticon.WithLogger(besticon.NewDefaultLogger(ioutil.Discard))) // Disable verbose logging
	icons, err := b.NewIconFinder().FetchIcons(siteData.Url)
	if err != nil || len(icons) == 0 {
		return ""
	}
	return icons[0].URL

}

// getWebTitle 获取网站标题
func getWebTitle(siteData *CreateSiteData) string {
	var title string
	c := tools.NewColly()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title += e.Text
	})
	if err := c.Visit(siteData.Url); err != nil {
		return ""
	}
	return title

}

// getWebDescription 获取网站描述
func getWebDescription(siteData *CreateSiteData) string {
	var description string
	c := tools.NewColly()
	c.OnXML("//meta[@name='description']/@content|//meta[@name='Description']/@content|//meta[@name='DESCRIPTION']",
		func(e *colly.XMLElement) {
			description += e.Text
		},
	)
	if err := c.Visit(siteData.Url); err != nil {
		return ""
	}
	return description
}

func (s *service) Create(ctx core.Context, siteData *CreateSiteData) (id int32, err error) {

	model := site.NewModel()
	model.IsUsed = -1
	model.CategoryId = siteData.CategoryId
	model.Url = siteData.Url
	model.Thumb = getWebLogoIconUrlByUrl(siteData)
	model.Title = getWebTitle(siteData)
	model.Description = getWebDescription(siteData)

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
