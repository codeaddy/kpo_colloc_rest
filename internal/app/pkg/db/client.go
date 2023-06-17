package db

import (
	"colloc_rest/internal/app/pkg/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDB(ctx context.Context) (*Database, error) {
	dsn := generateDsn()
	//dsn := "postgresql://test:test@postgres/postgres?sslmode=disable"
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return NewDatabase(pool), nil
}

func generateDsn() string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Dbname)
	//return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	config.Host, config.Port, config.User, config.Password, config.Dbname)
}
