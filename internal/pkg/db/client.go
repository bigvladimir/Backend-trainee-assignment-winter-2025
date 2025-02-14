package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDb(ctx context.Context, dsn string) (*Database, error) {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		err = fmt.Errorf("Ошибка при подключении к базе данных: %w", err)
		return nil, err
	}
	return newDatabase(pool), nil
}
