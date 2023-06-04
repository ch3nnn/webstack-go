package admin

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/validation"
	"github.com/ch3nnn/webstack-go/internal/services/admin"
)

type createRequest struct {
	Username string `form:"username" binding:"required"` // 用户名
	Nickname string `form:"nickname" binding:"required"` // 昵称
	Mobile   string `form:"mobile" binding:"required"`   // 手机号
	Password string `form:"password" binding:"required"` // MD5后的密码
}

type createResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// Create 新增管理员
// @Summary 新增管理员
// @Description 新增管理员
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "用户名"
// @Param nickname formData string true "昵称"
// @Param mobile formData string true "手机号"
// @Param password formData string true "MD5后的密码"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [post]
// @Security LoginToken
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		createData := new(admin.CreateAdminData)
		createData.Nickname = req.Nickname
		createData.Username = req.Username
		createData.Mobile = req.Mobile
		createData.Password = req.Password

		id, err := h.adminService.Create(c, createData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminCreateError,
				code.Text(code.AdminCreateError)).WithError(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
