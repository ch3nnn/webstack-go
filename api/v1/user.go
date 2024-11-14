package v1

type (
	LoginReq struct {
		Username string `form:"username" json:"username" binding:"required" example:"admin"`
		Password string `form:"password,default=value" json:"password,default=qweqwe" example:"123456"`
	}

	LoginResp struct {
		Token string `json:"token"`
	}
)

type (
	UpdatePasswordReq struct {
		OldPassword string `form:"old_password"` // 旧密码
		NewPassword string `form:"new_password"` // 新密码
	}

	UpdatePasswordResp struct{}
)

type (
	Menu struct {
		Id   int    `json:"id"`   // ID
		Pid  int    `json:"pid"`  // 父类ID
		Name string `json:"name"` // 菜单名称
		Link string `json:"link"` // 链接地址
		Icon string `json:"icon"` // 图标
	}

	InfoReq struct{}

	InfoResp struct {
		Username string `json:"username"` // 用户名
		Menus    []Menu `json:"menu"`     // 菜单栏
	}
)
