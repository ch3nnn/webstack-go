package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/site"
)

type UpdateSiteRequest struct {
	Id          int32  `json:"id"`
	CategoryId  int32  `json:"category_id"` // 网站分类id
	Title       string `json:"title"`       // 网站标题
	Thumb       string `json:"thumb"`       // 网站 logo
	Description string `json:"description"` // 网站描述
	Url         string `json:"url"`         // 网站地址
}

func (s *service) UpdateSite(ctx core.Context, updateSite *UpdateSiteRequest) (err error) {
	qb := site.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, updateSite.Id)
	data := map[string]any{
		"CategoryId":  updateSite.CategoryId,
		"Title":       updateSite.Title,
		"Thumb":       updateSite.Thumb,
		"Description": updateSite.Description,
		"Url":         updateSite.Url,
	}
	if err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data); err != nil {
		return err
	}
	return
}
