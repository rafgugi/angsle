package http

import (
	"github.com/fasthttp/router"
)

// Registration is the interface for handler, so
// it can be registered to the router
type Registration interface {
	Register(router *Router) error
}

// Router extends router.Router
type Router struct {
	*router.Router
}

// NewRouter initialize new Router instance
func NewRouter(regs ...Registration) (*Router, error) {
	router := &Router{router.New()}
	router.HandleMethodNotAllowed = false

	for _, reg := range regs {
		reg.Register(router)
	}

	return router, nil
}
