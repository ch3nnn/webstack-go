# webstack-go 网址导航后台系统

基于 Golang 开源的网址导航网站项目，具备完整的前后台，您可以拿来制作自己平日收藏的网址导航。
- 前端模板: [WebStackPage](https://github.com/WebStackPage/WebStackPage.github.io)
- 图标库: [lineicons](https://lineicons.com/icons/)
- 后端框架: [go-gin-api](https://github.com/xinliangnote/go-gin-api)

## 快速开始

**运行环境**
- Golang 1.16+  因为使用了 //go:embed 特性；
- MySQL
  - 连接地址，例如：127.0.0.1:3306；
  - 数据库名，例如：webstack，会在此数据库下初始化数据表；
  - 用户名，不可为空；
  -  密码，不可为空；
- Redis
  - 连接地址，例如：127.0.0.1:6379；
  - 密码，可为空；
  - 连接DB，默认是 0 ；

**启动服务**

 1. 目录下执行 `go mod tidy`拉取项目依赖库
 2. 安装完依赖执行`go run main.go`首次启动程序之后，会在浏览器中自动打开安装界面，链接地址：http://127.0.0.1:9999/install
 3. 点击`初始化项目` 会看到如下图所示, 如果提示重新运行服务说明项目初始化完成, 只需重新运行服务就 OK 了! 

   <img src="assets/bootstrap/images/init_project.png" width="600"/>


## 效果图

> **首页**

![](assets/bootstrap/images/index.png)

> **网站分类**

![](assets/bootstrap/images/category.png)

> **网站信息**

![](assets/bootstrap/images/site.png)
