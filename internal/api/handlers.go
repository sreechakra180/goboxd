package api

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/user/blackroot/internal/executor"
	"github.com/user/blackroot/internal/logger"
	"github.com/user/blackroot/internal/models"
	"github.com/user/blackroot/internal/scanner"
)

type Handler struct {
	exec    executor.Executor
	scanner *scanner.Scanner
	
	// in-memory logs for hackathon speed (avoids DB setup)
	logs   []models.LogEntry
	logsMu sync.Mutex
}

func NewHandler() *Handler {
	// quick docker support based on env var
	var exec executor.Executor
	if os.Getenv("USE_DOCKER") == "true" {
		logger.InfoLog.Println("Using DockerExecutor for isolated sandboxing")
		exec = executor.NewDockerExecutor()
	} else {
		logger.InfoLog.Println("Using LocalExecutor (fallback)")
		exec = executor.NewLocalExecutor()
	}

	return &Handler{
		exec:    exec,
		scanner: scanner.NewScanner(),
		logs:    make([]models.LogEntry, 0),
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"up"}`))
}

func (h *Handler) ScanCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.RunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	alerts := h.scanner.ScanCode(req.Language, req.Code)
	
	resp := map[string]interface{}{
		"safe": len(alerts) == 0,
		"alerts": alerts,
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) RunCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.RunRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	startTime := time.Now()

	// 1. Threat Detection Layer
	alerts := h.scanner.ScanCode(req.Language, req.Code)
	if len(alerts) > 0 {
		h.saveLog(req.Language, "blocked")
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.RunResponse{
			Status:   "blocked",
			Output:   "Execution blocked by security scanner.",
			Duration: time.Since(startTime),
			Alerts:   alerts,
		})
		return
	}

	// 2. Execution Layer
	out, err := h.exec.Run(req.Language, req.Code)
	
	status := "success"
	if err != nil {
		status = "error"
		logger.ErrorLog.Printf("Execution error: %v", err)
	}

	duration := time.Since(startTime)
	h.saveLog(req.Language, status)

	resp := models.RunResponse{
		Status:   status,
		Output:   out,
		Duration: duration,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetLogs(w http.ResponseWriter, r *http.Request) {
	h.logsMu.Lock()
	defer h.logsMu.Unlock()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.logs)
}

func (h *Handler) saveLog(language, status string) {
	h.logsMu.Lock()
	defer h.logsMu.Unlock()
	
	h.logs = append(h.logs, models.LogEntry{
		Timestamp: time.Now(),
		Language:  language,
		Status:    status,
	})
}
