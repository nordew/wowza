.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: build
build:
	go build -o wowza cmd/api/main.go

.PHONY: migrate-up
migrate-up:
	migrate -path migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" down

.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  run             Run the application"
	@echo "  build           Build the application binary"
	@echo "  migrate-up      Apply all up migrations"
	@echo "  migrate-down    Apply all down migrations"
	@echo "  migrate-create  Create a new migration with name=\`name\`" 