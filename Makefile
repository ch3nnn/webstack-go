.PHONY: init
init:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/incu6us/goimports-reviser/v3@latest
	go install mvdan.cc/gofumpt@latest

.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server

.PHONY: docker
docker:
	docker build -t webstack-go:v2 --build-arg APP_CONF=config/prod.yml --build-arg  APP_RELATIVE_PATH=./cmd/server  .
	docker run -itd -p 8000:8000 --name webstack-go webstack-go:v2

.PHONY: swag
swag:
	swag init  -g cmd/server/main.go -o ./docs --parseDependency

.PHONY: fmt
fmt:
	goimports-reviser -rm-unused -set-alias -format ./...
	find . -name '*.go' -not -name "*.pb.go" -not -name "*.gen.go" | xargs gofumpt -w -extra

.PHONY: run
run:
	go mod tidy
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server
	./bin/server -conf=config/prod.yml
