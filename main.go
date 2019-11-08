package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// StatusResponse is a template for the status endpoint.
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

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World."))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// RESTy routes for "service" resource
	r.Route("/v1", func(r chi.Router) {
		r.Get("/", V1Home)
		r.Get("/status", V1Status)
	})

	err := http.ListenAndServe(":8001", r)
	if err != nil {
		log.Fatalf("Can't start the sever. Error: %+v", err)
	}
}

func V1Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func V1Status(w http.ResponseWriter, r *http.Request) {
	// read git hash from environment variable
	hash := os.Getenv("HASH")

	// read version from first line of metadata fiile
	file, err := os.Open("./metadata")
	if err != nil {
		log.Printf("Error: %+v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		log.Printf(" > Failed!: %v\n", err)
	}

	if line == "" {
		line = "NO METADATA"
	}

	payload := fmt.Sprintf(StatusResponse, line, hash)
	data := []byte(payload)
	w.Write(data)
}
