package db

import (
	"wildberries/l0/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	Logger logger.Logger
	db     *pgxpool.Pool
}

type Postgres interface {
}

func NewStorage(logger logger.Logger) Postgres {
	return &postgres{
		Logger: logger,
	}
}
