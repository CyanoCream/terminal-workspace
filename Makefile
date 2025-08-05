.PHONY: configure build run migrate

configure:
	go work sync
	find . -name 'go.mod' -execdir go mod tidy \;

build:
	go build -o bin/api ./apps/api/cmd

run: build
	./bin/api

migrate:
	@echo "Waiting for PostgreSQL to be ready..."
	@./scripts/wait-for-postgres.sh
	@echo "Running migrations..."
	@migrate -path ./migrations -database "postgres://$$POSTGRES_USER:$$POSTGRES_PASSWORD@$$POSTGRES_HOST:$$POSTGRES_PORT/$$POSTGRES_DB?sslmode=disable" up

test:
	go test ./...