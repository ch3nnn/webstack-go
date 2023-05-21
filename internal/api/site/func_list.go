package site

import (
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	siteMd "github.com/ch3nnn/webstack-go/internal/repository/mysql/site"
	"github.com/ch3nnn/webstack-go/internal/services/site"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

type listRequest struct {
	Page              int    `form:"page"`               // 第几页
	PageSize          int    `form:"page_size"`          // 每页显示条数
	BusinessKey       string `form:"business_key"`       // 调用方key
	BusinessSecret    string `form:"business_secret"`    // 调用方secret
	BusinessDeveloper string `form:"business_developer"` // 调用方对接人
	Remark            string `form:"remark"`             // 备注
}
type listData struct {
	Id          int                 `json:"id"`          // ID
	Thumb       string              `json:"thumb"`       // 网站 logo
	Title       string              `json:"title"`       // 名称简介
	Url         string              `json:"url"`         // 链接
	Category    string              `json:"category"`    // 分类
	CategoryId  int32               `json:"category_id"` // 分类id
	Description string              `json:"description"` // 描述
	IsUsed      siteMd.IsUsedStatus `json:"is_used"`     // 是否启用
	CreatedAt   string              `json:"created_at"`  // 创建时间
	UpdatedAt   string              `json:"updated_at"`  // 更新时间
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int `json:"total"`
		CurrentPage  int `json:"current_page"`
		PerPageCount int `json:"per_page_count"`
	} `json:"pagination"`
}

// List 网站列表
// @Summary 网站列表
// @Description 网站列表
// @Tags API.site
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/site [get]
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listRequest)
		res := new(listResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		page := req.Page
		if page == 0 {
			page = 1
		}

		pageSize := req.PageSize
		if pageSize == 0 {
			pageSize = 10
		}

		searchData := new(site.SearchData)
		searchData.Page = page
		searchData.PageSize = pageSize
		searchData.BusinessKey = req.BusinessKey
		searchData.BusinessSecret = req.BusinessSecret
		searchData.Remark = req.Remark
		searchData.Search = strings.TrimSpace(c.Query("search"))

		resListData, err := h.siteService.PageList(c, searchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedListError,
				code.Text(code.AuthorizedListError)).WithError(err),
			)
			return
		}

		resCountData, err := h.siteService.PageListCount(c, searchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SiteListError,
				code.Text(code.SiteListError)).WithError(err),
			)
			return
		}
		res.Pagination.Total = cast.ToInt(resCountData)
		res.Pagination.PerPageCount = pageSize
		res.Pagination.CurrentPage = page
		res.List = make([]listData, len(resListData))

		for k, v := range resListData {
			_, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			data := listData{
				Id:          cast.ToInt(v.Id),
				Thumb:       v.Thumb,
				Title:       v.Title,
				Url:         v.Url,
				Category:    v.Category.Title,
				CategoryId:  v.Category.Id,
				Description: v.Description,
				IsUsed:      v.IsUsed,
				CreatedAt:   v.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:   v.UpdatedAt.Format("2006-01-02 15:04:05"),
			}

			res.List[k] = data
		}

		c.Payload(res)
	}
}
