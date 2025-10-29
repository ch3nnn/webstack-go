# WebStack-Go Docker 部署指南

## 概述

本项目已成功集成飞书OAuth登录功能，并支持Docker容器化部署。本文档介绍如何使用Docker部署WebStack-Go应用。

## 功能特性

- ✅ 飞书OAuth 2.0登录认证
- ✅ 首页访问控制（需要飞书登录）
- ✅ 管理后台独立认证（不受影响）
- ✅ 环境变量配置支持
- ✅ Docker容器化部署
- ✅ Kubernetes部署支持

## 快速开始

### 1. 构建Docker镜像

```bash
# 编译Linux版本二进制文件
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/server-linux ./cmd/server

# 构建Docker镜像
docker build -f Dockerfile.simple -t webstack-go:latest .
```

### 2. 运行容器

```bash
# 使用环境变量运行
docker run -d \
  --name webstack-go \
  -p 8000:8000 \
  -e FEISHU_APP_ID=your_app_id \
  -e FEISHU_APP_SECRET=your_app_secret \
  -e FEISHU_REDIRECT_URL=http://localhost:8000/api/auth/feishu/callback \
  -e FEISHU_TENANT_KEY=your_tenant_key \
  webstack-go:latest
```

### 3. 使用Docker Compose

```bash
# 修改docker-compose.yml中的环境变量
# 然后运行
docker-compose up -d
```

## 环境变量配置

| 变量名 | 描述 | 示例值 | 必需 |
|--------|------|--------|------|
| `FEISHU_APP_ID` | 飞书应用ID | `cli_a8723d9ef275d00e` | ✅ |
| `FEISHU_APP_SECRET` | 飞书应用密钥 | `gQwAupMXlTaI5dy47DVwLgLNuwIRTUy1` | ✅ |
| `FEISHU_REDIRECT_URL` | OAuth回调地址 | `http://localhost:8000/api/auth/feishu/callback` | ✅ |
| `FEISHU_TENANT_KEY` | 企业租户Key | `your_tenant_key` | ❌ |

## 飞书应用配置

### 1. 创建飞书应用

1. 访问 [飞书开放平台](https://open.feishu.cn/)
2. 创建企业自建应用
3. 获取 `App ID` 和 `App Secret`

### 2. 配置OAuth重定向URI

在飞书应用管理后台添加以下重定向URI：
- 开发环境：`http://localhost:8000/api/auth/feishu/callback`
- 生产环境：`https://yourdomain.com/api/auth/feishu/callback`

### 3. 配置权限

确保应用具有以下权限：
- `user:read` - 读取用户基本信息
- `user:read:email` - 读取用户邮箱

## 部署方式

### 方式1：Docker运行

```bash
# 停止现有容器
docker stop webstack-go || true
docker rm webstack-go || true

# 运行新容器
docker run -d \
  --name webstack-go \
  --restart unless-stopped \
  -p 8000:8000 \
  -v $(pwd)/storage:/app/storage \
  -e FEISHU_APP_ID=your_app_id \
  -e FEISHU_APP_SECRET=your_app_secret \
  -e FEISHU_REDIRECT_URL=http://yourdomain.com/api/auth/feishu/callback \
  -e FEISHU_TENANT_KEY=your_tenant_key \
  webstack-go:latest
```

### 方式2：Docker Compose

```yaml
version: '3.8'

services:
  webstack-go:
    image: webstack-go:latest
    container_name: webstack-go
    restart: unless-stopped
    ports:
      - "8000:8000"
    environment:
      - FEISHU_APP_ID=your_app_id
      - FEISHU_APP_SECRET=your_app_secret
      - FEISHU_REDIRECT_URL=http://yourdomain.com/api/auth/feishu/callback
      - FEISHU_TENANT_KEY=your_tenant_key
    volumes:
      - ./storage:/app/storage
    healthcheck:
      test: ["CMD", "wget", "--quiet", "--tries=1", "--spider", "http://localhost:8000/"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
```

### 方式3：Kubernetes部署

参考 `FEISHU_CONFIG.md` 文件中的Kubernetes配置示例。

## 访问应用

### 1. 首页访问

- 访问 `http://localhost:8000/`
- 自动重定向到飞书登录页面
- 完成飞书授权后返回首页

### 2. 管理后台

- 访问 `http://localhost:8000/login`
- 使用原有的管理员账号登录
- 不受飞书登录影响

### 3. API文档

- 访问 `http://localhost:8000/swagger/index.html`
- 查看完整的API文档

## 故障排除

### 1. 容器启动失败

```bash
# 查看容器日志
docker logs webstack-go

# 检查容器状态
docker ps -a | grep webstack
```

### 2. 飞书登录失败

1. 检查环境变量是否正确设置
2. 确认飞书应用配置中的重定向URI
3. 检查网络连接和防火墙设置

### 3. 权限问题

```bash
# 检查文件权限
docker exec webstack-go ls -la /app/

# 修复权限
docker exec webstack-go chmod +x /app/server
```

## 监控和维护

### 1. 健康检查

```bash
# 检查应用状态
curl http://localhost:8000/api/about

# 检查飞书登录
curl http://localhost:8000/api/auth/feishu/login
```

### 2. 日志查看

```bash
# 实时查看日志
docker logs -f webstack-go

# 查看最近100行日志
docker logs --tail 100 webstack-go
```

### 3. 数据备份

```bash
# 备份数据目录
docker cp webstack-go:/app/storage ./backup-$(date +%Y%m%d)
```

## 安全建议

1. **环境变量安全**：使用Kubernetes Secret或Docker Secret管理敏感信息
2. **网络安全**：配置适当的防火墙规则
3. **HTTPS**：生产环境建议使用HTTPS
4. **定期更新**：定期更新Docker镜像和依赖

## 更新日志

- **v1.0.0** - 集成飞书OAuth登录功能
- **v1.0.1** - 支持环境变量配置
- **v1.0.2** - 优化Docker镜像构建

## 技术支持

如有问题，请查看：
1. 项目README.md
2. 飞书开放平台文档
3. Docker官方文档
4. Kubernetes官方文档
