package api

import (
	"encoding/json"
	"net/http"
)

// JobStatus represents a dummy job response
type JobStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// NewRouter returns a basic HTTP mux
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Dummy job status endpoint
	mux.HandleFunc("/jobs/", func(w http.ResponseWriter, r *http.Request) {
		resp := JobStatus{ID: "123", Status: "queued"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	return mux
}