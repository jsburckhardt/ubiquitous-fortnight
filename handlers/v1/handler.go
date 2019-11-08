package handler

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

// GetPing is the handler for ping endpoint
func GetPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// GetV1Home is the handler for V1 home endpoint
func GetV1Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

// GetV1Status is the handler for V1 status endpoint
func GetV1Status(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
