package view

import (
	"net/http"
	"wildberries/l0/internal/domain"
	logger "wildberries/l0/pkg/logger"
)

type view struct {
	Logger logger.Logger
	domain domain.Domain
}

type View interface {
	GetData(w http.ResponseWriter, r *http.Request)
	RestoreCache() error
}

func NewView(logger logger.Logger, domain domain.Domain) View {
	return &view{
		Logger: logger,
		domain: domain,
	}
}
