ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST = postgres://admin:admin@localhost:5432/avitointro
endif

MIGRATION_FOLDER = migrations
LOCAL_BIN:=$(CURDIR)/bin

add-linters:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
add-bins:
	GOBIN=$(LOCAL_BIN) go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

generate-api:
	$(LOCAL_BIN)/oapi-codegen -generate chi-server  -package api api/openapi.yaml > internal/infrastructure/http/chi-server.gen.go
	$(LOCAL_BIN)/oapi-codegen -generate types -package api api/openapi.yaml > internal/infrastructure/http/types.gen.go

migration-create:
	bin/goose   -dir "$(MIGRATION_FOLDER)" create "$(name)" sql
migration-up:
	bin/goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up
migration-down:
	bin/goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

lint: 
	bin/golangci-lint run  
	bin/fieldalignment -fix ./...

build-infrastructure:
	docker-compose up postgres -d 