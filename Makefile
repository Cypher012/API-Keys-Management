# Load DATABASE_URL from .env (only for make commands)
DATABASE_URL := $(shell grep -E '^DATABASE_URL=' .env | cut -d '=' -f2-)

GOOSE_DIR := internal/db/migrations
DB_DRIVER := postgres

.PHONY: migrate-up migrate-down migrate-status migrate-create

migrate-up:
	goose -dir $(GOOSE_DIR) $(DB_DRIVER) "$(DATABASE_URL)" up

migrate-down:
	goose -dir $(GOOSE_DIR) $(DB_DRIVER) "$(DATABASE_URL)" down

migrate-status:
	goose -dir $(GOOSE_DIR) $(DB_DRIVER) "$(DATABASE_URL)" status

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Error: name is required. Usage: make migrate-create name=your_migration_name"; \
		exit 1; \
	fi
	goose -dir $(GOOSE_DIR) create $(name) sql
