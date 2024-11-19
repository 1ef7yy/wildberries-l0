package v1

import (
	"net/http"
	"wildberries/l0/internal/view"
)

type Router struct {
	View view.View
}

func NewRouter(view view.View) *Router {
	return &Router{
		View: view,
	}
}

func (v *Router) Ping() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	return http.StripPrefix("/v1/", mux)
}
