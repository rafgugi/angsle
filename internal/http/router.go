package http

import (
	"github.com/gofiber/fiber/v2"
)

// Registration is the interface for handler, so
// it can be registered to the router
type Registration interface {
	Register(router *Router) error
}

// Router extends router.Router
type Router struct {
	r *fiber.App
}

// NewRouter initializes new Router instance, then registers the regs
func NewRouter(regs ...Registration) *Router {
	router := &Router{fiber.New()}

	for _, reg := range regs {
		reg.Register(router)
	}

	return router
}

func (r *Router) Handle(method, path string, handler fiber.Handler) {
	r.r.Add(method, path, handler)
}

func (r *Router) Listen(path string) {
	r.r.Listen(path)
}

