package cron

import (
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/constant"
	"github.com/spf13/cast"
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/timeutil"
	"github.com/ch3nnn/webstack-go/internal/pkg/validation"
	"github.com/ch3nnn/webstack-go/internal/services/cron"
)

type listRequest struct {
	Page     int64  `form:"page,default=1"`       // 第几页
	PageSize int64  `form:"page_size,default=10"` // 每页显示条数
	Name     string `form:"name"`                 // 任务名称
	Protocol int64  `form:"protocol"`             // 执行方式 1:shell 2:http
	IsUsed   int64  `form:"is_used"`              // 是否启用 1:是  -1:否
}

type listData struct {
	Id               int64  `json:"id"`                 // ID
	HashID           string `json:"hashid"`             // hashid
	Name             string `json:"name"`               // 任务名称
	Protocol         int64  `json:"protocol"`           // 执行方式 1:shell 2:http
	ProtocolText     string `json:"protocol_text"`      // 执行方式
	Spec             string `json:"spec"`               // crontab 表达式
	Command          string `json:"command"`            // 执行命令
	HttpMethod       int64  `json:"http_method"`        // http 请求方式 1:get 2:post
	HttpMethodText   string `json:"http_method_text"`   // http 请求方式
	Timeout          int64  `json:"timeout"`            // 超时时间(单位:秒)
	RetryTimes       int64  `json:"retry_times"`        // 重试次数
	RetryInterval    int64  `json:"retry_interval"`     // 重试间隔(单位:秒)
	NotifyStatus     int64  `json:"notify_status"`      // 执行结束是否通知 1:不通知 2:失败通知 3:结束通知 4:结果关键字匹配通知
	NotifyStatusText string `json:"notify_status_text"` // 执行结束是否通知
	IsUsed           int64  `json:"is_used"`            // 是否启用 1=启用 2=禁用
	IsUsedText       string `json:"is_used_text"`       // 是否启用
	CreatedAt        string `json:"created_at"`         // 创建时间
	CreatedUser      string `json:"created_user"`       // 创建人
	UpdatedAt        string `json:"updated_at"`         // 更新时间
	UpdatedUser      string `json:"updated_user"`       // 更新人
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int64 `json:"total"`
		CurrentPage  int64 `json:"current_page"`
		PerPageCount int64 `json:"per_page_count"`
	} `json:"pagination"`
}

// List 任务列表
// @Summary 任务列表
// @Description 任务列表
// @Tags API.cron
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param page query int true "第几页" default(1)
// @Param page_size query int true "每页显示条数" default(10)
// @Param name query string false "任务名称"
// @Param protocol query int false "执行方式 1:shell 2:http"
// @Param is_used query int false "是否启用 1:是  -1:否"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/cron [get]
// @Security LoginToken
func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(listRequest)
		res := new(listResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		searchData := new(cron.SearchData)
		searchData.Page = req.Page
		searchData.PageSize = req.PageSize
		searchData.Name = req.Name
		searchData.Protocol = req.Protocol
		searchData.IsUsed = req.IsUsed

		resListData, err := h.cronService.PageList(ctx, searchData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CronListError,
				code.Text(code.CronListError)).WithError(err),
			)
			return
		}

		resCountData, err := h.cronService.PageListCount(ctx, searchData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CronListError,
				code.Text(code.CronListError)).WithError(err),
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
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			res.List[k] = listData{
				Id:               v.ID,
				HashID:           hashId,
				Name:             v.Name,
				Protocol:         v.Protocol,
				ProtocolText:     constant.ProtocolText[v.Protocol],
				Spec:             v.Spec,
				Command:          v.Command,
				HttpMethod:       v.HTTPMethod,
				HttpMethodText:   constant.HttpMethodText[v.HTTPMethod],
				Timeout:          v.Timeout,
				RetryTimes:       v.RetryTimes,
				RetryInterval:    v.RetryInterval,
				NotifyStatus:     v.NotifyStatus,
				NotifyStatusText: constant.NotifyStatusText[v.NotifyStatus],
				IsUsed:           v.IsUsed,
				IsUsedText:       constant.IsUsedText[v.IsUsed],
				CreatedAt:        v.CreatedAt.Format(timeutil.CSTLayout),
				CreatedUser:      v.CreatedUser,
				UpdatedAt:        v.UpdatedAt.Format(timeutil.CSTLayout),
				UpdatedUser:      v.UpdatedUser,
			}
		}

		ctx.Payload(res)
	}
}
