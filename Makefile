run:
	go run cmd/main.go

test:
	go test ./... --cover

migrateup:
	migrate -path database/postgres/migration -database "postgres://postgres:datnguyennnx@database_login:5432/login_service?sslmode=disable" -verbose up

migratedown:
	migrate -path database/postgres/migration -database "postgres://postgres:datnguyennnx@database_login:5432/login_service?sslmode=disable" -verbose down

docker-build:
	docker build -t login_service:latest .

.PHONY: run test migrateup migratedown docker-build
