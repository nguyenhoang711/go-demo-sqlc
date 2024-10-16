# server name
APP_NAME = demosqlc

# Define the environment (default is 'dev')
# ENV ?= local

# these are the default values
GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING ?= "root:123456@tcp(127.0.0.1:3306)/dev_go"
GOOSE_MIGRATION_DIR ?= sql/schema

docker-build:
	docker-compose up -d --build
	docker-compose ps

docker-stop:
	docker-compose down

docker-up:
	docker compose up -d

run-local:
	go run ./cmd/demosqlc
# go run ./cmd/$(APP_NAME)

# @echo "Running with environment: $(ENV)"
# ENV=$(ENV) go run ./cmd/$(APP_NAME)
gen-proto:
	cd proto && sh gen.sh
gen-swagger:
	swag init -g ./cmd/demosqlc/main.go ./cmd/swag/docs

up-gen:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

down-gen:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

reset-gen:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlcgen:
	sqlc generate

.PHONY: run-local down-gen up-gen reset-gen docker-build docker-stop docker-up

.PHONY: air