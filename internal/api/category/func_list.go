package category

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/spf13/cast"
)

type listData struct {
	Id     int64  `json:"id"`      // ID
	HashID string `json:"hashid"`  // hashid
	Pid    int64  `json:"pid"`     // 父类ID
	Name   string `json:"name"`    // 菜单名称
	Link   string `json:"link"`    // 链接地址
	Icon   string `json:"icon"`    // 图标
	IsUsed int64  `json:"is_used"` // 是否启用 1=启用 -1=禁用
	Sort   int64  `json:"sort"`    // 排序
	Level  int64  `json:"level"`   // 分类等级 1 一级分类  2 二级分类
}

type listResponse struct {
	List []listData `json:"list"`
}

// List 分类列表
// @Summary 分类列表
// @Description 分类列表
// @Tags API.category
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/category [get]
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		res := new(listResponse)
		categories, err := h.categoryService.List(c)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CategoryListError,
				code.Text(code.CategoryListError)).WithError(err),
			)
			return
		}

		res.List = make([]listData, len(categories))

		for i, cat := range categories {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(cat.ID)})
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			res.List[i] = listData{
				Id:     cat.ID,
				HashID: hashId,
				Pid:    cat.ParentID,
				Name:   cat.Title,
				Icon:   cat.Icon,
				IsUsed: cat.IsUsed,
				Sort:   cat.Sort,
				Level:  cat.Level,
			}
		}

		c.Payload(res)
	}
}
