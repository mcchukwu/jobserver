package main

import (
	"log"
	"net/http"

	"github.com/mcchukwu/jobserver/internal/api"
	"github.com/mcchukwu/jobserver/internal/service"
)

const port = ":8080"

func main() {
	jobService := service.NewJobService()
	jobHandler := api.NewJobHandler(jobService)

	http.HandleFunc("/job", jobHandler.CreateJob)

	log.Printf("server is running on port%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
