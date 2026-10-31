# 管理端前端产物不入库，镜像构建时现场生成
FROM node:20-alpine AS web
WORKDIR /web
COPY web/admin/package.json web/admin/package-lock.json ./
RUN npm ci --no-audit --no-fund
COPY web/admin/ ./
RUN npm run build


FROM golang:1.23-alpine AS builder
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

ARG APP_RELATIVE_PATH

COPY . /data/app
# 覆盖仓库里的占位 dist（只有 .gitkeep），让 go:embed 拿到真正的产物
COPY --from=web /web/dist /data/app/web/admin/dist
WORKDIR /data/app

RUN rm -rf /data/app/bin/
#RUN export GOPROXY=https://goproxy.cn,direct && go mod tidy && go build -ldflags="-s -w" -o ./bin/server ${APP_RELATIVE_PATH}
RUN go mod tidy && go build -ldflags="-s -w" -o ./bin/server ${APP_RELATIVE_PATH}
RUN mv config /data/app/bin/


FROM alpine:3.14
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories


RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata


ARG APP_ENV
ENV APP_ENV=${APP_ENV}

WORKDIR /data/app
COPY --from=builder /data/app/bin /data/app
COPY --from=builder /data/app/web/upload /data/app/web/upload/
RUN mkdir -p /data/app/storage/

EXPOSE 8000
ENTRYPOINT [ "./server" ]
