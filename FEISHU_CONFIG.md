# 飞书OAuth配置环境变量说明

## 环境变量配置

在部署时，需要设置以下环境变量：

```bash
# 飞书应用配置
FEISHU_APP_ID=your_app_id
FEISHU_APP_SECRET=your_app_secret
FEISHU_REDIRECT_URL=https://your-domain.com/api/auth/feishu/callback
FEISHU_TENANT_KEY=your_tenant_key  # 可选
```

## K8s部署配置示例

### ConfigMap
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: webstack-config
data:
  FEISHU_REDIRECT_URL: "https://your-domain.com/api/auth/feishu/callback"
  FEISHU_TENANT_KEY: ""
```

### Secret
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: webstack-secrets
type: Opaque
data:
  FEISHU_APP_ID: <base64-encoded-app-id>
  FEISHU_APP_SECRET: <base64-encoded-app-secret>
```

### Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: webstack-go
spec:
  template:
    spec:
      containers:
      - name: webstack-go
        image: webstack-go:latest
        env:
        - name: FEISHU_APP_ID
          valueFrom:
            secretKeyRef:
              name: webstack-secrets
              key: FEISHU_APP_ID
        - name: FEISHU_APP_SECRET
          valueFrom:
            secretKeyRef:
              name: webstack-secrets
              key: FEISHU_APP_SECRET
        envFrom:
        - configMapRef:
            name: webstack-config
```

## 本地开发

创建 `.env` 文件：
```bash
FEISHU_APP_ID=cli_a8723d9ef275d00e
FEISHU_APP_SECRET=gQwAupMXlTaI5dy47DVwLgLNuwIRTUy1
FEISHU_REDIRECT_URL=http://localhost:8000/api/auth/feishu/callback
FEISHU_TENANT_KEY=
```

然后使用 `source .env` 或 `export` 命令设置环境变量。
