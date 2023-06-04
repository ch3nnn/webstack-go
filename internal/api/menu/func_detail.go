package menu

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/services/menu"
)

type detailRequest struct {
	Id string `uri:"id"` // HashID
}

type detailResponse struct {
	Id   int64  `json:"id"`   // 主键ID
	Pid  int64  `json:"pid"`  // 父类ID
	Name string `json:"name"` // 菜单名称
	Link string `json:"link"` // 链接地址
	Icon string `json:"icon"` // 图标
}

// Detail 菜单详情
// @Summary 菜单详情
// @Description 菜单详情
// @Tags API.menu
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu/{id} [get]
// @Security LoginToken
func (h *handler) Detail() core.HandlerFunc {
	return func(c core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithError(err),
			)
			return
		}

		searchOneData := new(menu.SearchOneData)
		searchOneData.Id = int64(ids[0])

		info, err := h.menuService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuDetailError,
				code.Text(code.MenuDetailError)).WithError(err),
			)
			return
		}

		res.Id = info.ID
		res.Pid = info.Pid
		res.Name = info.Name
		res.Link = info.Link
		res.Icon = info.Icon
		c.Payload(res)
	}
}
