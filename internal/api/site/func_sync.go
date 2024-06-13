/**
 * @Author: chentong
 * @Date: 2024/05/19 下午7:12
 */

package site

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
)

type syncRequest struct {
	Id int64 `uri:"id"` // 主键ID
}

type syncResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// SyncSite 一键同步
// @Summary 一键同步
// @Description 一键同步
// @Tags API.site
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body syncRequest true "请求信息"
// @Success 200 {object} syncResponse
// @Failure 400 {object} code.Failure
// @Router /api/site/sync [patch]
func (h *handler) SyncSite() core.HandlerFunc {
	return func(c core.Context) {
		req := new(syncRequest)
		res := new(syncResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		if err := h.siteService.Sync(c, req.Id); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SiteUpdateError,
				code.Text(code.SiteUpdateError)).WithError(err),
			)
			return
		}

		res.Id = req.Id
		c.Payload(res)
	}
}
