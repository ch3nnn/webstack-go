package site

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/site"
)

type CreateSiteData struct {
	CategoryId  int32  `json:"category_id"`
	Thumb       string `json:"thumb"`
	Title       string `json:"title"`
	Url         string `json:"Url"`
	Description string `json:"description"`
}

func (s *service) Create(ctx core.Context, siteData *CreateSiteData) (id int32, err error) {

	model := site.NewModel()
	model.CategoryId = siteData.CategoryId
	model.Thumb = siteData.Thumb
	model.Title = siteData.Title
	model.Url = siteData.Url
	model.Description = siteData.Description
	model.IsUsed = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
