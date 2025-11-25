package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jacksondprice/delta-course-etl/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/tickets", handlers.ListTickets)
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, mux))
}
