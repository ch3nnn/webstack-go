package v1

var (
	ErrSuccess             = newError(0, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")
)

var (
	ErrorUserNameAndPassword = newError(100, "用户名和密码错误")
	ErrorUserOldPassword     = newError(100, "原密码错误")
	ErrorTokenGeneration     = newError(101, "令牌生成错误")

	ErrorUnableToGetFile      = newError(200, "无法获取文件")
	ErrorFileSizeExceedsLimit = newError(201, "文件大小超过限制")
)
