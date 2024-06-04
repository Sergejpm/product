package repository

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

func Open(ctx context.Context, creds Creds) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "pgx", creds.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("connect error: %s", err)
	}

	if creds.MaxOpenConns > 0 {
		db.SetMaxOpenConns(creds.MaxOpenConns)
	}
	if creds.MaxIdleConns > 0 {
		db.SetMaxIdleConns(creds.MaxIdleConns)
	}
	if creds.MaxConnLifetimeSec > 0 {
		db.SetConnMaxLifetime(creds.MaxConnLifetimeSec * time.Second)
	}

	return db, nil
}
