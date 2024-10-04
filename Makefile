# server name
APP_NAME = demosqlc

# Define the environment (default is 'dev')
ENV ?= local

run-local:
	@echo "Running with environment: $(ENV)"
	ENV=$(ENV) go run cmd/demosqlc/main.go
gen-proto:
	cd proto && sh gen.sh
gen-swagger:
	swag init -g ./cmd/demosqlc/main.go ./cmd/swag/docs