package site

import (
	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"net/http"
)

type updateUsedRequest struct {
	Id   int32 `form:"id"`   // 主键ID
	Used int32 `form:"used"` // 是否启用 1:是 -1:否
}

type updateUsedResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// UpdateUsed 更新网站为启用/禁用
// @Summary 更新网站为启用/禁用
// @Description 更新网站为启用/禁用
// @Tags API.site
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateUsedRequest true "请求信息"
// @Success 200 {object} updateUsedResponse
// @Failure 400 {object} code.Failure
// @Router /api/site/used [patch]
func (h *handler) UpdateUsed() core.HandlerFunc {
	return func(c core.Context) {
		req := new(updateUsedRequest)
		res := new(updateUsedResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		if err := h.siteService.UpdateUsed(c, req.Id, req.Used); err != nil {
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
