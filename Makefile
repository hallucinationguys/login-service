run:
	go run cmd/serverd/main.go
test:
	go test ./... --cover

docker-build:
	docker build -t login_service:latest .