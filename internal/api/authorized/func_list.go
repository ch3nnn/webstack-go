package authorized

import (
	"github.com/spf13/cast"
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/timeutil"
	"github.com/ch3nnn/webstack-go/internal/services/authorized"
)

type listRequest struct {
	Page              int64  `form:"page,default=1"`       // 第几页
	PageSize          int64  `form:"page_size,default=10"` // 每页显示条数
	BusinessKey       string `form:"business_key"`         // 调用方key
	BusinessSecret    string `form:"business_secret"`      // 调用方secret
	BusinessDeveloper string `form:"business_developer"`   // 调用方对接人
	Remark            string `form:"remark"`               // 备注
}

type listData struct {
	Id                int64  `json:"id"`                 // ID
	HashID            string `json:"hashid"`             // hashid
	BusinessKey       string `json:"business_key"`       // 调用方key
	BusinessSecret    string `json:"business_secret"`    // 调用方secret
	BusinessDeveloper string `json:"business_developer"` // 调用方对接人
	Remark            string `json:"remark"`             // 备注
	IsUsed            int64  `json:"is_used"`            // 是否启用 1:是 -1:否
	CreatedAt         string `json:"created_at"`         // 创建时间
	CreatedUser       string `json:"created_user"`       // 创建人
	UpdatedAt         string `json:"updated_at"`         // 更新时间
	UpdatedUser       string `json:"updated_user"`       // 更新人
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int64 `json:"total"`
		CurrentPage  int64 `json:"current_page"`
		PerPageCount int64 `json:"per_page_count"`
	} `json:"pagination"`
}

// List 调用方列表
// @Summary 调用方列表
// @Description 调用方列表
// @Tags API.authorized
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param page query int true "第几页" default(1)
// @Param page_size query int true "每页显示条数" default(10)
// @Param business_key query string false "调用方key"
// @Param business_secret query string false "调用方secret"
// @Param business_developer query string false "调用方对接人"
// @Param remark path string false "备注"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized [get]
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

		searchData := new(authorized.SearchData)
		searchData.Page = req.Page
		searchData.PageSize = req.PageSize
		searchData.BusinessKey = req.BusinessKey
		searchData.BusinessSecret = req.BusinessSecret
		searchData.Remark = req.Remark

		resListData, err := h.authorizedService.PageList(c, searchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedListError,
				code.Text(code.AuthorizedListError)).WithError(err),
			)
			return
		}

		resCountData, err := h.authorizedService.PageListCount(c, searchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedListError,
				code.Text(code.AuthorizedListError)).WithError(err),
			)
			return
		}
		res.Pagination.Total = resCountData
		res.Pagination.PerPageCount = req.PageSize
		res.Pagination.CurrentPage = req.Page
		res.List = make([]listData, len(resListData))

		for k, v := range resListData {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.ID)})
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			res.List[k] = listData{
				Id:                v.ID,
				HashID:            hashId,
				BusinessKey:       v.BusinessKey,
				BusinessSecret:    v.BusinessSecret,
				BusinessDeveloper: v.BusinessDeveloper,
				Remark:            v.Remark,
				IsUsed:            v.IsUsed,
				CreatedAt:         v.CreatedAt.Format(timeutil.CSTLayout),
				CreatedUser:       v.CreatedUser,
				UpdatedAt:         v.UpdatedAt.Format(timeutil.CSTLayout),
				UpdatedUser:       v.UpdatedUser,
			}
		}

		c.Payload(res)
	}
}
