package v1

import (
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
