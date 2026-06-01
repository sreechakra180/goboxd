package models

import "time"

// RunRequest is what the user sends to /run
type RunRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

// RunResponse is what we send back
type RunResponse struct {
	Status   string        `json:"status"`   // "success", "error", "timeout", "blocked"
	Output   string        `json:"output"`   // stdout/stderr combined for simplicity
	Duration time.Duration `json:"duration"` // execution time
	Alerts   []string      `json:"alerts,omitempty"` // if blocked by scanner
}

// LogEntry for our quick in-memory analytics
type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Language  string    `json:"language"`
	Status    string    `json:"status"`
	MemoryUse int64     `json:"memory_usage,omitempty"` // hard to get perfectly accurate in hackathon, maybe skip
}
