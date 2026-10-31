# webstack-go 网址导航后台系统

基于 Golang 开源的网址导航网站项目，具备完整的前后台，您可以拿来制作自己平日收藏的网址导航。
> v1: 使用 mysql 和 redis 组件, 丰富的后端功能。 v2: 简化版无需额外组件, 使用轻量级 sqlite 数据库。

- 前端模板: 公共站 [WebStackPage](https://github.com/WebStackPage/WebStackPage.github.io)（服务端渲染）；管理端 Vue 3 + TypeScript + Vite + Naive UI（SPA，构建产物内嵌进二进制）
- 后端框架: 基于 [go-nunu](https://github.com/go-nunu/nunu) 脚手架搭建

功能：
- [x] 新增 webstack - 导航首页
- [x] 新增 仪表盘 (SSE）
- [x] 新增 网站管理 - 网站分类
- [x] 新增 网站管理 - 网站列表
- [x] 新增 系统管理 - 自定义导航基本信息 (Logo、favicon、备案信息等）
- [x] 新增 支持批量添加 (自动获取标题、Logo、网站描述）
- [x] 新增 一键同步、导出功能
- [x] 新增 由 [gorm-gen](https://github.com/go-gorm/gen) 代码生成提供支持的友好且更安全的 GORM
- [x] 杂项 与仓库保持同步 [Docker Hub](https://hub.docker.com/r/ch3nnn/webstack-go/tags)

## 快速开始

### 一、运行环境

- Golang 1.23+（`go.mod` 的 go 指令是 1.23.0，1.22 无法编译）
- SQLite
- Node.js 20+（**只有从源码构建时需要**，用来编译管理端前端）

> 管理端 SPA 的构建产物不随仓库提交：`web/admin/dist` 里只保留一个 `.gitkeep` 占位，
> 真正的产物在构建时生成并被打进二进制。因此：
>
> - 直接下载 Releases 二进制或使用 Docker 镜像 —— 不需要 Node.js，产物已内置；
> - 从源码 `make build` —— 需要 Node.js，构建脚本会先编译前端再编译后端。

### 二、启动服务

**1、二进制文件**

你可以直接从[ Releases ](https://github.com/ch3nnn/webstack-go/releases)下载预先编译好的二进制文件，解压后执行:

```bash
./webstack-go -conf config/prod.yml 
```

> [!NOTE]
> MacOS 在执行二进制文件时会提示：未打开“webstack-go”，因为 Apple 无法检查其是否包含恶意软件。
> 
> 可在“系统设置 > 隐私与安全性 > 安全性”中点击“仍然允许”，然后再次尝试执行二进制文件。


**2、源码运行服务 (需要 Golang + Node.js 环境)**
1. 执行 `make build`：先 `npm ci && npm run build` 生成管理端产物，再 `go build` 把产物内嵌进 `./bin/server`
2. 执行 `./bin/server -conf=config/prod.yml` 首次启动程序之后，会生成 SQLite 数据库，并自动创建表结构

只想重新编译后端（`web/admin/dist` 里已有产物）时用 `make build-server`。

> `//go:embed all:admin/dist` 要求该目录非空，所以仓库里保留了一个 `dist/.gitkeep`。
> 如果产物缺失，`go build` 仍能通过，但启动日志会告警、`/admin` 返回 503 并提示先执行 `make web`。


**3、Docker 运行服务**
#### 下载镜像
1. docker run 运行
```bash
docker run -i -t --restart always -p 8000:8000 --name webstack-go -v ./data/storage:/data/app/storage ch3nnn/webstack-go:latest
```

2. docker compose (推荐)
```yaml
services:
  webstack-go:
    stdin_open: true
    tty: true
    restart: always
    ports:
      - 8000:8000
    container_name: webstack-go
    image: ch3nnn/webstack-go:latest
    volumes:
      - ./data/storage:/data/app/storage
```

#### 本地编译
1. 目录下执行 `make docker` 等待启动
   ```shell
   CONTAINER ID   IMAGE            COMMAND      CREATED         STATUS         PORTS                    NAMES
   5cb641ff3950 webstack-go:v2   "./server"   5 seconds ago Up 5 seconds 0.0.0.0:8000->8000/tcp webstack-go
   ```
2. docker container 正常运行后, 在浏览器中打开界面，链接地址：http://127.0.0.1:8000

## 三、管理端前端开发

管理端位于 `web/admin`，Vue 3 + TypeScript + Vite + Naive UI。产物 `web/admin/dist` 不提交入库，由 `make web`（或 CI / Docker 构建阶段）生成，再由 `web/assets.go` 的 `//go:embed admin/dist` 打进二进制；静态资源路径为 `/ui`，SPA 路由挂在 `/admin/*`，公共站（`/`、`/about`）仍是 Go 模板。

```bash
cd web/admin
npm ci
npm run dev      # http://127.0.0.1:5173，/api 与 /assets 代理到 127.0.0.1:8000
```

改完前端后跑一遍类型检查、lint 与构建（或直接在仓库根目录 `make web`）：

```bash
cd web/admin
npm run typecheck && npm run lint && npm run build
```

产物不需要提交：`web/admin/dist/*` 已在 `.gitignore` 里，只有 `.gitkeep` 占位入库。
CI（`.github/workflows/web.yml`）校验类型、lint 与构建；Docker 镜像与 release 二进制在各自的构建流程里现场生成前端产物。

### 回滚

管理端前端已内嵌进二进制，**前端无法单独回滚**：回滚 = 换回上一个镜像 tag / 二进制版本。

## 效果图

> **首页**

![](.github/image/index.png)

> **网站分类**

![](.github/image/category.png)

> **新增网站**

![](.github/image/add_site.png)

> **网站信息**

![](.github/image/site.png)

> **网站配置**

![](.github/image/config.png)

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=ch3nnn/webstack-go&type=Date)](https://star-history.com/#ch3nnn/webstack-go&Date)
