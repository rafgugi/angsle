package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafgugi/angsle/internal/http"
	"rsc.io/quote"
)

// GenericHandler handles healthz and not found
type GenericHandler struct{}

// NewGenericHandler create new GenericHandler
func NewGenericHandler() *GenericHandler {
	return &GenericHandler{}
}

// Healthz handles GET /healthz and returns "ok"
func (handle *GenericHandler) Healthz(c *fiber.Ctx) error {
	c.WriteString("ok")
	return nil
}

// Quote handles GET /quote and returns random quote
func (handle *GenericHandler) Quote(c *fiber.Ctx) error {
	c.WriteString(quote.Go())
	return nil
}

// NotFound for handle undefined path
func (handle *GenericHandler) NotFound(c *fiber.Ctx) error {
	c.SendString(`{"message":"Path not found","meta":{"http_status":404}}`)
	c.SendStatus(404)
	return nil
}

// Register is used for register routes with its handler
func (handle *GenericHandler) Register(router *http.Router) error {
	router.Handle("GET", "/healthz", handle.Healthz)
	router.Handle("GET", "/quote", handle.Quote)
	// router.NotFound = handle.NotFound

	return nil
}
