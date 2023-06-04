package menu

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/services/menu"
)

type createActionRequest struct {
	Id     string `form:"id"`     // HashID
	Method string `form:"method"` // 请求方法
	API    string `form:"api"`    // 请求地址
}

type createActionResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// CreateAction 创建功能权限
// @Summary 创建功能权限
// @Description 创建功能权限
// @Tags API.menu
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "HashID"
// @Param method formData string true "请求方法"
// @Param api formData string true "请求地址"
// @Success 200 {object} createActionResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu_action [post]
// @Security LoginToken
func (h *handler) CreateAction() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createActionRequest)
		res := new(createActionResponse)
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

		searchOneData := new(menu.SearchOneData)
		searchOneData.Id = int64(ids[0])
		menuInfo, err := h.menuService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuDetailError,
				code.Text(code.MenuDetailError)).WithError(err),
			)
			return
		}

		createActionData := new(menu.CreateMenuActionData)
		createActionData.MenuId = menuInfo.ID
		createActionData.Method = req.Method
		createActionData.API = req.API

		createId, err := h.menuService.CreateAction(c, createActionData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuCreateActionError,
				code.Text(code.MenuCreateActionError)).WithError(err),
			)
			return
		}

		res.Id = int64(createId)
		c.Payload(res)
	}
}
