package handler

import (
	"github.com/rafgugi/angsle/internal/http"
	"github.com/valyala/fasthttp"
)

// GenericHandler handles healthz and not found
type GenericHandler struct{}

// NewGenericHandler create new GenericHandler
func NewGenericHandler() *GenericHandler {
	return &GenericHandler{}
}

// Healthz handles GET /healthz and returns "ok"
func (handle *GenericHandler) Healthz(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("ok")
}

// NotFound for handle undefined path
func (handle *GenericHandler) NotFound(ctx *fasthttp.RequestCtx) {
	ctx.Error(`{"message":"Path not found","meta":{"http_status":404}}`, fasthttp.StatusNotFound)
}

// Register is used for register routes with its handler
func (handle *GenericHandler) Register(router *http.Router) error {
	router.Handle("GET", "/healthz", handle.Healthz)
	router.NotFound = handle.NotFound

	return nil
}
