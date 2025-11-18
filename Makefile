.PHONY: up down migrate-up migrate-down migrate-version migrate-create migrate-force

# Docker commands
up:
	docker compose up api --build

down:
	docker compose down

# Migration commands
DB_URL := postgres://postgres:postgres@localhost:5432/pet_checkup?sslmode=disable
MIGRATIONS_PATH := sql/migrations

migrate-up:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_PATH) up

migrate-down:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_PATH) down 1

migrate-version:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_PATH) version

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $$name

migrate-force:
	@read -p "Enter version to force: " version; \
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_PATH) force $$version
