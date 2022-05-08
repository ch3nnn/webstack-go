package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-web-site-admin/src/go/routes"
)

func main() {
	//注册路由
	r := routes.SetUserRouter()
	//启动端口为8080的项目
	r.Run(":8080")
}
