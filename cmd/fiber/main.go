package main

import (
	"github.com/rafgugi/angsle/internal/http"
	"github.com/rafgugi/angsle/internal/http/handler"
)

func main() {
	genericHandler := handler.NewGenericHandler()

	router := http.NewRouter(genericHandler)
	router.Listen(":3000")
}
