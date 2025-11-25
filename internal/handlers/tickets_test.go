package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListTickets(t *testing.T) {
	rec := httptest.NewRecorder()
	ListTickets(rec, httptest.NewRequest("GET", "/api/tickets", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
