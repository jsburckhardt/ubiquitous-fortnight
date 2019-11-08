package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	handler "github.com/jsburckhardt/ubiquitous-fortnight/handlers/v1"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// StatusResponse is a template for the response of status endpoint.
const StatusResponse = `{
  "myapplication": [
    {
      "version": "%s",
      "description": "pre-interview technical test",
      "lastcommitsha": "%s"
    }
  ]
}
`

// ServePort is the port use by the API to serve.
const ServePort = 8001

func main() {
	flag.Parse()

	r := chi.NewRouter()

	// useful middlewares for logging.
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Ping route to verify service is running.
	r.Get("/ping", handler.GetPing)

	// V1 routes for the service. Having versions allows modifications in APIs without affecting customers.
	r.Route("/v1", func(r chi.Router) {
		r.Get("/", handler.GetV1Home)
		r.Get("/status", handler.GetV1Status)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%v", ServePort), r)
	if err != nil {
		log.Fatalf("Can't start the sever. Error: %+v", err)
	}
}
