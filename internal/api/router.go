package api

import (
	"net/http"
)

// SetupRouter wires up all our API endpoints
func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	h := NewHandler()

	mux.HandleFunc("/health", h.HealthCheck)
	mux.HandleFunc("/run", h.RunCode)
	mux.HandleFunc("/scan", h.ScanCode)
	mux.HandleFunc("/logs", h.GetLogs)

	return mux
}
