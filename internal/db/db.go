package db

import (
	"context"
	"fmt"
	"github.com/induzo/gocom/database/pginit/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func New(dbURI string) (*pgxpool.Pool, error) {
	pgi, err := pginit.New(
		dbURI,
		pginit.WithDecimalType(),
		pginit.WithGoogleUUIDType(),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to initialize pginit: %w", err)
	}

	pool, err := pgi.ConnPool(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	log.Println("Successfully connected to the database")
	return pool, nil
}
