package cron

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/validation"
)

type updateUsedRequest struct {
	Id   string `form:"id"`   // 主键ID
	Used int64  `form:"used"` // 是否启用 1:是 -1:否
}

type updateUsedResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// UpdateUsed 更新任务为启用/禁用
// @Summary 更新任务为启用/禁用
// @Description 更新任务为启用/禁用
// @Tags API.cron
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "hashID"
// @Param used formData int true "是否启用 1:是 -1:否"
// @Success 200 {object} updateUsedResponse
// @Failure 400 {object} code.Failure
// @Router /api/cron/used [patch]
// @Security LoginToken
func (h *handler) UpdateUsed() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(updateUsedRequest)
		res := new(updateUsedResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
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

		id := int64(ids[0])

		err = h.cronService.UpdateUsed(ctx, id, req.Used)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminUpdateError,
				code.Text(code.AdminUpdateError)).WithError(err),
			)
			return
		}

		res.Id = id
		ctx.Payload(res)
	}
}
