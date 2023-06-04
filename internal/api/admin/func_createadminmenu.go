package admin

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/services/admin"
)

type createAdminMenuRequest struct {
	Id      string `form:"id"`      // HashID
	Actions string `form:"actions"` // 功能权限ID,多个用,分割
}

type createAdminMenuResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// CreateAdminMenu 提交菜单授权
// @Summary 提交菜单授权
// @Description 提交菜单授权
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "Hashid"
// @Param actions formData string true "功能权限ID,多个用,分割"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/menu [post]
// @Security LoginToken
func (h *handler) CreateAdminMenu() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createAdminMenuRequest)
		res := new(createAdminMenuResponse)
		if err := c.ShouldBindForm(req); err != nil {
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

		createData := new(admin.CreateMenuData)
		createData.AdminId = int64(ids[0])
		createData.Actions = req.Actions

		err = h.adminService.CreateMenu(c, createData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminMenuCreateError,
				code.Text(code.AdminMenuCreateError)).WithError(err),
			)
			return
		}

		res.Id = int64(ids[0])
		c.Payload(res)
	}
}
