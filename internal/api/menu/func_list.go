package menu

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/services/menu"

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
}

type listResponse struct {
	List []listData `json:"list"`
}

// List 菜单列表
// @Summary 菜单列表
// @Description 菜单列表
// @Tags API.menu
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu [get]
// @Security LoginToken
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		res := new(listResponse)
		menus, err := h.menuService.List(c, new(menu.SearchData))
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuListError,
				code.Text(code.MenuListError)).WithError(err),
			)
			return
		}

		res.List = make([]listData, len(menus))

		for i, m := range menus {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(m.ID)})
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			res.List[i] = listData{
				Id:     m.ID,
				HashID: hashId,
				Pid:    m.Pid,
				Name:   m.Name,
				Link:   m.Link,
				Icon:   m.Icon,
				IsUsed: m.IsUsed,
				Sort:   m.Sort,
			}
		}

		c.Payload(res)
	}
}
