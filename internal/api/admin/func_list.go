package admin

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/configs"
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/password"
	"github.com/ch3nnn/webstack-go/internal/services/admin"
	"github.com/xinliangnote/go-gin-api/pkg/timeutil"

	"github.com/spf13/cast"
)

type listRequest struct {
	Page     int64  `form:"page,default=1"`       // 第几页
	PageSize int64  `form:"page_size,default=10"` // 每页显示条数
	Username string `form:"username"`             // 用户名
	Nickname string `form:"nickname"`             // 昵称
	Mobile   string `form:"mobile"`               // 手机号
}

type listData struct {
	Id          int64  `json:"id"`           // ID
	HashID      string `json:"hashid"`       // hashid
	Username    string `json:"username"`     // 用户名
	Nickname    string `json:"nickname"`     // 昵称
	Mobile      string `json:"mobile"`       // 手机号
	IsUsed      int64  `json:"is_used"`      // 是否启用 1:是 -1:否
	IsOnline    int64  `json:"is_online"`    // 是否在线 1:是 -1:否
	CreatedAt   string `json:"created_at"`   // 创建时间
	CreatedUser string `json:"created_user"` // 创建人
	UpdatedAt   string `json:"updated_at"`   // 更新时间
	UpdatedUser string `json:"updated_user"` // 更新人
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int64 `json:"total"`
		CurrentPage  int64 `json:"current_page"`
		PerPageCount int64 `json:"per_page_count"`
	} `json:"pagination"`
}

// List 管理员列表
// @Summary 管理员列表
// @Description 管理员列表
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param page query int true "第几页" default(1)
// @Param page_size query int true "每页显示条数" default(10)
// @Param username query string false "用户名"
// @Param nickname query string false "昵称"
// @Param mobile query string false "手机号"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [get]
// @Security LoginToken
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listRequest)
		res := new(listResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		searchData := new(admin.SearchData)
		searchData.Page = req.Page
		searchData.PageSize = req.PageSize
		searchData.Username = req.Username
		searchData.Nickname = req.Nickname
		searchData.Mobile = req.Mobile

		admins, err := h.adminService.PageList(c, searchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminListError,
				code.Text(code.AdminListError)).WithError(err),
			)
			return
		}

		resCountData, err := h.adminService.PageListCount(c, searchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminListError,
				code.Text(code.AdminListError)).WithError(err),
			)
			return
		}
		res.Pagination.Total = resCountData
		res.Pagination.PerPageCount = req.PageSize
		res.Pagination.CurrentPage = req.Page
		res.List = make([]listData, len(admins))

		for k, v := range admins {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.ID)})
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			isOnline := -1
			if h.cache.Exists(configs.RedisKeyPrefixLoginUser + password.GenerateLoginToken(v.ID)) {
				isOnline = 1
			}

			res.List[k] = listData{
				Id:          v.ID,
				HashID:      hashId,
				Username:    v.Username,
				Nickname:    v.Nickname,
				Mobile:      v.Mobile,
				IsUsed:      v.IsUsed,
				IsOnline:    int64(isOnline),
				CreatedAt:   v.CreatedAt.Format(timeutil.CSTLayout),
				CreatedUser: v.CreatedUser,
				UpdatedAt:   v.UpdatedAt.Format(timeutil.CSTLayout),
				UpdatedUser: v.UpdatedUser,
			}
		}

		c.Payload(res)
	}
}
