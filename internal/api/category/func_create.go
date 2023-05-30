package category

import (
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/services/category"
	"net/http"
)

type createRequest struct {
	Id    int64  `form:"id"`    // ID
	Pid   int64  `form:"pid"`   // 父类ID
	Name  string `form:"name"`  // 分类名称
	Icon  string `form:"icon"`  // 图标
	Level int64  `form:"level"` // 分类等级
}

type createResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// Create 创建/编辑分类
// @Summary 创建/编辑分类
// @Description 创建/编辑分类
// @Tags API.category
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/category [post]
// @Security LoginToken
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

		if req.Id != 0 { // 编辑功能

			updateCategory := new(category.UpdateCategory)
			updateCategory.Name = req.Name
			updateCategory.Icon = req.Icon

			if err := h.categoryService.Modify(c, req.Id, updateCategory); err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.CategoryUpdateError,
					code.Text(code.CategoryUpdateError)).WithError(err),
				)
				return
			}

			res.Id = req.Id
			c.Payload(res)

		} else { // 新增功能
			pid := req.Level
			level := 2

			if req.Level == -1 {
				pid = 0
				level = 1
			}

			createData := new(category.CreateCategoryData)
			createData.Pid = pid
			createData.Name = req.Name
			createData.Icon = req.Icon
			createData.Level = int64(level)

			if err := h.categoryService.Create(c, createData); err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.CategoryCreateError,
					code.Text(code.CategoryCreateError)).WithError(err),
				)
				return
			}

			c.Payload(res)
		}
	}
}
