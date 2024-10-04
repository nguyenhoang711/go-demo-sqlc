# server name
APP_NAME = demosqlc

# Define the environment (default is 'dev')
ENV ?= dev

run-local:
	@echo "Running with environment: $(ENV)"
	ENV=$(ENV) go run cmd/demosqlc/main.go
gen-proto:
	cd proto && sh gen.sh
gen-swagger:
	swag init