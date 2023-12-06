run:
	go run cmd/main.go

test:
	go test ./... --cover

migrateup:
	migrate -path database/postgres/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path database/postgres/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

docker-build:
	docker build -t login_service:latest .

.PHONY: run test migrateup migratedown docker-build new_migration
