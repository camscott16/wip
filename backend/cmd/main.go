package main

import (
	"log"
	"net/http"
	"os"

	"github.com/camscott16/wip/pkg/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	// health check endpoint
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "ok"}`))
	})

	http.HandleFunc("/api/salary/estimate", api.SalaryEstimateHandler)

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}