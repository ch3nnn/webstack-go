package category

import (
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"net/http"
)

type detailRequest struct {
	Id int64 `uri:"id"` // 主键 id
}

type detailResponse struct {
	Id   int64  `json:"id"`   // 主键ID
	Pid  int64  `json:"pid"`  // 父类ID
	Name string `json:"name"` // 分类名称
	Icon string `json:"icon"` // 图标
}

// Detail 分类详情
// @Summary 分类详情
// @Description 分类详情
// @Tags API.category
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/category/{id} [get]
// @Security LoginToken
func (h *handler) Detail() core.HandlerFunc {
	return func(c core.Context) {
		req := new(detailRequest)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		cat, err := h.categoryService.Detail(c, req.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CategoryDetailError,
				code.Text(code.CategoryDetailError)).WithError(err),
			)
			return
		}

		c.Payload(detailResponse{
			Id:   cat.ID,
			Pid:  cat.ParentID,
			Name: cat.Title,
			Icon: cat.Icon,
		})
	}
}
