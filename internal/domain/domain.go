package domain

import (
	"wildberries/l0/internal/storage/db"
	"wildberries/l0/pkg/logger"
)

type domain struct {
	Logger logger.Logger
	pg     db.Postgres
}

type Domain interface{}

func NewDomain(logger logger.Logger) Domain {
	return &domain{
		Logger: logger,
		pg:     db.NewStorage(logger),
	}

}


