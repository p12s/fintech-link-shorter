.PHONY:
.SILENT:
.DEFAULT_GOAL := run

s:
	go run cmd/main.go

build:
	go mod download && go build cmd/main.go

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage

test.coverage:
	go tool cover -func=cover.out

swag:
	swag init -g cmd/main.go

lint:
	golangci-lint run