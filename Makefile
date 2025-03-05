dev:
	cd ./cmd; air
dev-db:
	docker compose up --build
migrate:
	cd ./migrations; goose postgres postgres://postgres:pass@localhost:5432/data up
swag:
	swag init -g cmd/main.go -o cmd/docs