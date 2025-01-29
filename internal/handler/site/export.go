/**
 * @Author: chentong
 * @Date: 2025/01/29 19:06
 */

package site

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (h *Handler) Export(ctx *gin.Context) {
	var rep v1.SiteExportReq
	if err := ctx.BindQuery(&rep); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	resp, err := h.siteService.Export(ctx, &rep)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	// 设置响应头 Excel 文件
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename=sites.xlsx")
	ctx.Header("Content-Transfer-Encoding", "binary")

	if err = resp.File.Write(ctx.Writer); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	return
}
