.PHONY: init
init:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/incu6us/goimports-reviser/v3@latest
	go install mvdan.cc/gofumpt@latest

# 管理端前端产物不随仓库提交，构建后端之前先生成
.PHONY: web
web:
	cd web/admin && npm ci --no-audit --no-fund && npm run build

.PHONY: build
build: web
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server

# 只编后端：要求 web/admin/dist 里已经有构建好的产物
.PHONY: build-server
build-server:
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server

.PHONY: docker
docker:
	docker build -t webstack-go:v3 --build-arg APP_CONF=config/prod.yml --build-arg  APP_RELATIVE_PATH=./cmd/server  .
	docker run -itd -p 8000:8000 --name webstack-go webstack-go:v3

.PHONY: swag
swag:
	swag init  -g cmd/server/main.go -o ./docs --parseDependency

.PHONY: fmt
fmt:
	goimports-reviser -rm-unused -set-alias -format ./...
	find . -name '*.go' -not -name "*.pb.go" -not -name "*.gen.go" | xargs gofumpt -w -extra

.PHONY: run
run: build
	./bin/server -conf=config/prod.yml
