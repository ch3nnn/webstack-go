# webstack-go 网址导航后台系统

基于 Golang 开源的网址导航网站项目，具备完整的前后台，您可以拿来制作自己平日收藏的网址导航。
- 图标库: [lineicons](https://lineicons.com/icons/)
- 前端模板: [WebStackPage](https://github.com/WebStackPage/WebStackPage.github.io)
- 后端 Gin 框架: 基于 [go-gin-api ](https://github.com/xinliangnote/go-gin-api)项目二次开发

原有后端项目基础上新增功能：
1. webstack - 导航首页
2. 系统管理员 - 网站分类 
3. 系统管理员 - 网站列表
4. 网站新增支持自动获取标题、Logo、网站描述
5. 新增 docker-compose.yml 一键安装各个组件

## 快速开始

### 一、运行环境
> 安装 Golang、Mysql、Redis

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

### 二、启动服务
> 两种方式运行 1. 源码启动服务 2. Docker启动服务

**一、源码运行服务**

 1. 目录下执行 `go mod tidy` 拉取项目依赖库
 2. 安装完依赖执行 `go run main.go` 首次启动程序之后，会在浏览器中自动打开安装界面，链接地址：http://127.0.0.1:9999/install
 3. 点击 `初始化项目` 会看到如下图所示, 如果提示重新运行服务说明项目初始化完成, 只需重新运行服务就 OK 了! 

   <img src="assets/bootstrap/images/init_project.png" width="600"/>

**二、Docker运行服务**
 1. 目录下执行 `docker-compose -f docker-compose.yml up -d` 等待组件启动
    ```shell
     $ docker-compose -f docker-compose.yml start   
    [+] Running 7/7
     ⠿ Container webstack-go-mysql       Healthy                                                                                                                                                                                                                                                    11.7s
     ⠿ Container webstack-go-redis       Healthy                                                                                                                                                                                                                                                    11.2s
     ⠿ Container webstack-go-service     Healthy                                                                                                                                                                                                                                                    12.2s
     ⠿ Container webstack-go-grafana     Started                                                                                                                                                                                                                                                     2.2s
     ⠿ Container webstack-go-loki        Started                                                                                                                                                                                                                                                     3.0s
     ⠿ Container webstack-go-prometheus  Started                                                                                                                                                                                                                                                     3.2s
     ⠿ Container webstack-go-promtail    Started 
    ```
 2. 执行`docker ps`服务正常运行如下
    ```shell
    CONTAINER ID   IMAGE                             COMMAND                  CREATED         STATUS                    PORTS                               NAMES
    698e64590652   grafana/promtail:2.7.3            "/usr/bin/promtail -…"   6 hours ago     Up 14 seconds                                                 webstack-go-promtail
    b30e56157328   grafana/loki:2.7.3                "/usr/bin/loki -conf…"   6 hours ago     Up 14 seconds             0.0.0.0:3100->3100/tcp              webstack-go-loki
    269d18273780   grafana/grafana-enterprise        "/run.sh"                6 hours ago     Up 14 seconds             0.0.0.0:3000->3000/tcp              webstack-go-grafana
    e8f1de150ef7   prom/prometheus:latest            "/bin/prometheus --c…"   6 hours ago     Up 13 seconds             0.0.0.0:9090->9090/tcp              webstack-go-prometheus
    5956a1de364a   webstack-go_webstack-go-service   "./webstack-go -env …"   6 hours ago     Up 26 seconds (healthy)   0.0.0.0:9999->9999/tcp              webstack-go-service
    fd7ccf68ebd2   mysql/mysql-server:5.7            "/entrypoint.sh mysq…"   6 hours ago     Up 37 seconds (healthy)   33060/tcp, 0.0.0.0:3305->3306/tcp   webstack-go-mysql
    eaaac671f0d5   redis:6.2.4                       "docker-entrypoint.s…"   6 hours ago     Up 37 seconds (healthy)   0.0.0.0:6378->6379/tcp              webstack-go-redis
    ```
 3. docker container 正常运行后, 在浏览器中打开安装界面，链接地址：http://127.0.0.1:9999/install
 4. 点击 `初始化项目` 会看到如下图所示, 如果提示重新运行服务说明项目初始化完成, 只需重新运行 `webstack-go-service` 容器服务就 OK 了!
 
    <img src="assets/bootstrap/images/init_project.png" width="600"/>


## 效果图

> **首页**

![](assets/bootstrap/images/index.png)

> **网站分类**

![](assets/bootstrap/images/category.png)

> **新增网站**

![](assets/bootstrap/images/add_site.png)

> **网站信息**

![](assets/bootstrap/images/site.png)

> **监控组件**

![grafana.png](assets/bootstrap/images/grafana.png)![]()
