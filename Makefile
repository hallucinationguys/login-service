run:
	go run cmd/main.go

test:
	go test ./... --cover

migrateup:
	migrate -path database/postgres/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path database/postgres/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

docker-build:
	docker build -t login_service:latest .

.PHONY: run test migrateup migratedown docker-build
