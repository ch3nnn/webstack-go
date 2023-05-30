package site

import (
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/services/site"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type createRequest struct {
	CategoryId int64  `form:"category_id"`
	Url        string `form:"url"`
}

type createResponse struct {
	SuccessCount int64 `json:"successCount"`
	FailCount    int64 `json:"failCount"`
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
		sites := make([]*site.CreateSiteData, 0, 10)
		for _, url := range strings.Split(req.Url, "\n") {
			// 校验网址格式
			if err := validator.New().Var(url, "http_url"); err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.ParamBindError,
					code.Text(code.ParamBindError)).WithError(err),
				)
				return
			}
			sites = append(sites, &site.CreateSiteData{CategoryId: req.CategoryId, Url: url})
		}
		res.SuccessCount, res.FailCount = h.siteService.Create(c, sites)

		c.Payload(res)
	}
}
