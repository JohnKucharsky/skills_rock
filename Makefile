dev:
	cd cmd && air
start:
	cd cmd && go run .
dev-db:
	docker compose up -d
migrate:
	goose -dir ./migrations postgres postgres://postgres:pass@localhost:5432/data up
swag:
	swag init -g cmd/main.go -o cmd/docs