

go:
	go run cmd/main.go

swag-init:
	swag init -g api/api.go -o api/docs

migration-up:
	migrate -path ./migrations/postgres/ -database 'admin//postgres:2711@localhost:5432/korzinka?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres/ -database 'admin://postgres:2711@localhost:5432/korzinka?sslmode=disable' down
