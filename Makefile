run:
	go run main.go

gen:
	sqlc generate

create-db:
	psql -U postgres -c "CREATE DATABASE $(DB_NAME)"

create-schema:
	psql -f ./db/schema.sql -U postgres -d $(DB_NAME)

new-setup: create-db create-schema
	@echo "Set up database $(DB_NAME)..."

