package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	log.Printf("server is running on port%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
