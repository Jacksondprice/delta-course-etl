package handlers

import (
	"encoding/json"
	"net/http"
)

type Ticket struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListTickets(w http.ResponseWriter, _ *http.Request) {
	items := []Ticket{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
