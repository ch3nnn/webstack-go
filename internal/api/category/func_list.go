package category

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/services/category"

	"github.com/spf13/cast"
)

type listData struct {
	Id     int32  `json:"id"`      // ID
	HashID string `json:"hashid"`  // hashid
	Pid    int32  `json:"pid"`     // 父类ID
	Name   string `json:"name"`    // 菜单名称
	Link   string `json:"link"`    // 链接地址
	Icon   string `json:"icon"`    // 图标
	IsUsed int32  `json:"is_used"` // 是否启用 1=启用 -1=禁用
	Sort   int32  `json:"sort"`    // 排序
	Level  int32  `json:"level"`   // 分类等级 1 一级分类  2 二级分类
}

type listRequest struct{}

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
		resListData, err := h.categoryService.List(c, new(category.SearchData))
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CategoryListError,
				code.Text(code.CategoryListError)).WithError(err),
			)
			return
		}

		res.List = make([]listData, len(resListData))

		for k, v := range resListData {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			data := listData{
				Id:     v.Id,
				HashID: hashId,
				Pid:    v.ParentId,
				Name:   v.Title,
				Icon:   v.Icon,
				IsUsed: v.IsUsed,
				Sort:   v.Sort,
				Level:  v.Level,
			}

			res.List[k] = data
		}

		c.Payload(res)
	}
}
