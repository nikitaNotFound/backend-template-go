install-tools:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
start-local-http:
	go run cmd/http-server/main.go
start-local-grpc:
	go run cmd/grpc-server/main.go
migrate-local:
	go run cmd/database-migration/main.go
swag:
	swag init -g ./cmd/http-server/main.go -o ./docs/