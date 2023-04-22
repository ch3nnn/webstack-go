#!/bin/bash
# shellcheck disable=SC2103

# 构建依赖
function build() {
  cd .. || exit
  go mod tidy
  echo "build server"
  go build -o .
  cd ..
}


# 运行服务
function start() {
  cd ../ || exit
  nohup ./webstack-go >/dev/null 2>&1 &
  echo "service running"
}

# 停止服务
function stop() {

  PID=""
  query_pid() {
    PID=$(ps -ef | grep webstack-go | grep -v grep | awk '{print $2}')
  }
  query_pid
  kill -TERM "$PID"


}

# 服务状态
function status() {
  PID=$(ps -ef | grep webstack-go | grep -v grep | wc -l)
  if [ $PID != 0 ]; then
    echo "webstack-go is running..."
  else
    echo "webstack-go is not running..."
  fi
}

# 重启服务
function restart() {
  stop
  sleep 3
  start
  echo "restart success .... "
}

case $1 in
build)
  build
  ;;
start)
  start
  ;;
stop)
  stop
  ;;
restart)
  restart
  ;;
status)
  status
  ;;
*)
  echo -e "\033[0;31m 输入操作名错误 \033[0m  \033[0;31m {build|start|stop|restart|status} \033[0m"
  ;;
esac
