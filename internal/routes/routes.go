package routes

import (
	"net/http"
	v1 "wildberries/l0/internal/routes/v1"
	"wildberries/l0/internal/view"
)

func InitRouter(view view.View) *http.ServeMux {
	mux := http.NewServeMux()
	v1 := v1.NewRouter(view)

	mux.Handle("/api/", v1.Endpoints())

	return mux
}
