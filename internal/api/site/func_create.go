package site

import (
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/services/site"
	"net/http"
)

type createRequest struct {
	CategoryId int32  `form:"category_id"`
	Url        string `form:"url"`
}

type createResponse struct {
	Id int32 `json:"id"`
}

// Create 创建网站
// @Summary 创建网站
// @Description 创建网站
// @Tags API.site
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/site [post]
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		createData := new(site.CreateSiteData)
		createData.CategoryId = req.CategoryId
		createData.Url = req.Url

		id, err := h.siteService.Create(c, createData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SiteCreateError,
				code.Text(code.SiteCreateError)).WithError(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
