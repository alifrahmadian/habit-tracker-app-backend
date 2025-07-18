include .env

APP_NAME=habit_tracker_app
MAIN=./cmd

# DB Migrations
MIGRATION_PATH = db/migrations
DATABASE_URL = "postgres://${USERNAME}:${PASSWORD}@${HOST}:${PORT}/${DB_NAME}?sslmode=disable"

.PHONY: run migrate-create migrate-up migrate-down migrate-version

run:
	go run $(MAIN)

migrate-create:
ifndef NAME
	$(error NAME is undefined. Usage: make migrate-create NAME=create_table)
endif
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(NAME)

migrate-up:
	migrate -path $(MIGRATION_PATH) -database $(DATABASE_URL) up

migrate-down:
	migrate -path $(MIGRATION_PATH) -database $(DATABASE_URL) down

migrate-version:
		migrate -path $(MIGRATION_PATH) -database $(DATABASE_URL) version

migrate-force-version:
		migrate -path $(MIGRATION_PATH) -database $(DATABASE_URL) force $(VERSION_NUMBER)