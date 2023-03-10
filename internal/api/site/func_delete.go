package site

import (
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"net/http"
)

type deleteRequest struct {
	Id int32 `uri:"id"` // 主键ID
}

type deleteResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Delete 删除网站信息
// @Summary 删除网站信息
// @Description 删除网站信息
// @Tags API.site
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteRequest true "请求信息"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /api/site/{id} [delete]
func (h *handler) Delete() core.HandlerFunc {
	return func(c core.Context) {
		req := new(deleteRequest)
		res := new(deleteResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		if err := h.siteService.Delete(c, req.Id); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SiteDeleteError,
				code.Text(code.SiteDeleteError)).WithError(err),
			)
			return
		}

		res.Id = req.Id
		c.Payload(res)
	}
}
