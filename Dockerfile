FROM golang:1.22-alpine AS builder
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

ARG APP_RELATIVE_PATH

COPY . /data/app
WORKDIR /data/app

RUN rm -rf /data/app/bin/
RUN export GOPROXY=https://goproxy.cn,direct && go mod tidy && go build -ldflags="-s -w" -o ./bin/server ${APP_RELATIVE_PATH}
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
