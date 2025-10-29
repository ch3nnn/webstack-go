#!/bin/bash

# 飞书OAuth环境变量配置
export FEISHU_APP_ID=cli_a8723d9ef275d00e
export FEISHU_APP_SECRET=gQwAupMXlTaI5dy47DVwLgLNuwIRTUy1
export FEISHU_REDIRECT_URL=http://localhost:8000/api/auth/feishu/callback
export FEISHU_TENANT_KEY=""

# 启动服务器
cd /Users/han/cy/webstack-go
./bin/server -conf=config/prod.yml
