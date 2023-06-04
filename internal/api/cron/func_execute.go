package cron

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/validation"
)

type executeRequest struct {
	Id string `uri:"id"` // HashID
}

type executeResponse struct {
	Id int64 `json:"id"` // ID
}

// Execute 手动执行单条任务
// @Summary 手动执行单条任务
// @Description 手动执行单条任务
// @Tags API.cron
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/cron/exec/{id} [patch]
// @Security LoginToken
func (h *handler) Execute() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(executeRequest)
		res := new(executeResponse)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithError(err),
			)
			return
		}

		err = h.cronService.Execute(ctx, int64(ids[0]))
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CronExecuteError,
				code.Text(code.CronExecuteError)).WithError(err),
			)
			return
		}

		res.Id = int64(ids[0])
		ctx.Payload(res)
	}
}
