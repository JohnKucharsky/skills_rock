## Getting started

### Install swag

[swag](https://github.com/swaggo/swag)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Install goose

[goose](https://github.com/pressly/goose)

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Install air

[air](https://github.com/air-verse/air)

```bash
go install github.com/cosmtrek/air@latest
```

### Make Targets

- **make dev-db**: Starts database for dev
- **make migrate**: Migrations, goose
- **make dev**: Starts app in dev mode, fresh
- **make prod**: docker compose up --build
