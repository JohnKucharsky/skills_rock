# Getting started

The `.env` file is intentionally not added to `.gitignore` because it doesn't contain sensitive information.

## To Run

### Install goose

[goose](https://github.com/pressly/goose)

```bash

go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Run docker database

```bash

docker compose up -d
```

### Migrate db

```bash

goose -dir ./migrations postgres postgres://postgres:pass@localhost:5432/data up
```

### Start app

```bash

cd cmd && go run .
```

### Docs

[localhost:8080/docs](http://localhost:8080/docs)

## For development

### Install air

[air](https://github.com/air-verse/air)

```bash

go install github.com/air-verse/air@latest
```

### Install swag

[swag](https://github.com/swaggo/swag)

```bash

go install github.com/swaggo/swag/cmd/swag@latest
```


### Make Targets

- **make dev-db**: Starts database for dev
- **make migrate**: Migrations, goose
- **make start**: Starts the app
- **make dev**: Starts the app in dev mode, users air
