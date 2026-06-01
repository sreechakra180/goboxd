package executor

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/user/blackroot/internal/logger"
)

// Executor defines how we run code
type Executor interface {
	Run(language, code string) (output string, err error)
}

// Generate a random ID for the temp workspace
func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// createTempWorkspace makes an isolated temp dir for this run
func createTempWorkspace(code, filename string) (string, error) {
	dir := filepath.Join(os.TempDir(), "blackroot_"+generateID())
	
	// TODO: improve sandbox isolation later
	// for now just standard perms
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	filePath := filepath.Join(dir, filename)
	if err := os.WriteFile(filePath, []byte(code), 0644); err != nil {
		// cleanup if write fails
		os.RemoveAll(dir)
		return "", err
	}

	logger.InfoLog.Printf("Created workspace %s", dir)
	return dir, nil
}
