package category

import (
	"github.com/xinliangnote/go-gin-api/internal/services/category"
	"net/http"
	"strconv"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type detailRequest struct {
	Id string `uri:"id"` // 主键 id
}

type detailResponse struct {
	Id   int32  `json:"id"`   // 主键ID
	Pid  int32  `json:"pid"`  // 父类ID
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

		id, err := strconv.Atoi(req.Id)
		if err != nil {
			return
		}

		searchOneData := new(category.SearchOneData)
		searchOneData.Id = int32(id)

		info, err := h.categoryService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CategoryDetailError,
				code.Text(code.CategoryDetailError)).WithError(err),
			)
			return
		}

		res := detailResponse{
			Id:   info.Id,
			Pid:  info.ParentId,
			Name: info.Title,
			Icon: info.Icon,
		}

		c.Payload(res)
	}
}
