package db

import (
	"context"
	"sync"
	"wildberries/l0/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Logger   logger.Logger
	Database *pgxpool.Pool
}

func NewPostgres(ctx context.Context, dsn string, log logger.Logger) *Postgres {
	var (
		pgInstance *Postgres
		pgOnce     sync.Once
	)

	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, dsn)
		if err != nil {
			log.Error("Unable to connect to database: " + err.Error())
		}

		pgInstance = &Postgres{
			Logger:   log,
			Database: db,
		}
	})
	return pgInstance
}

func (pg *Postgres) Close() {
	pg.Database.Close()
}
