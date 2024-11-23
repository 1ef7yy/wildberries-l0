package domain

import (
	"context"
	"os"
	"wildberries/l0/internal/models"
	"wildberries/l0/internal/storage/cache"
	"wildberries/l0/internal/storage/db"
	"wildberries/l0/pkg/logger"
)

type domain struct {
	Logger logger.Logger
	pg     db.Postgres
	cache  cache.Cache
}

type Domain interface {
	GetData(id string) (models.Schema, error)
	RestoreCache() error
}

func NewDomain(logger logger.Logger) Domain {
	return &domain{
		Logger: logger,
		pg:     *db.NewPostgres(context.Background(), os.Getenv("POSTGRES_CONN"), logger),
		cache:  *cache.NewCache(),
	}

}
