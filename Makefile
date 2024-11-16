-include .env.dev
export

define setup_env
		$(eval ENV_FILE := $(1))
		@echo "--SETUP ENV: $(ENV_FILE)--"
		$(eval include $(1))
		$(eval export sed 's/=.*//' $(1))
endef

start-dev:
	go run cmd/api/main.go --env=dev

start-prod:
	go run cmd/api/main.go --env=prod

tidy:
	go mod tidy

migrate-up:
	$(call setup_env,$(project_file))
	migrate -path db/migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}" up

migrate-down:
	$(call setup_env,$(project_file))
	migrate -path db/migrations -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}" down 1