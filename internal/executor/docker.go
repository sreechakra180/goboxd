package executor

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/user/blackroot/internal/logger"
)

// DockerExecutor runs code inside throwaway containers
type DockerExecutor struct {
	timeout time.Duration
}

func NewDockerExecutor() *DockerExecutor {
	return &DockerExecutor{
		timeout: 5 * time.Second, // slightly longer for container spin up
	}
}

func (e *DockerExecutor) Run(language, code string) (string, error) {
	var filename string
	var image string
	var runCmd string

	switch language {
	case "python":
		filename = "main.py"
		image = "python:3.9-alpine"
		runCmd = "python3 /workspace/main.py"
	case "javascript", "node":
		filename = "main.js"
		image = "node:18-alpine"
		runCmd = "node /workspace/main.js"
	case "go":
		filename = "main.go"
		image = "golang:1.21-alpine"
		runCmd = "go run /workspace/main.go"
	// hackathon version — optimize cpp later, compiling in alpine takes time
	default:
		return "", fmt.Errorf("unsupported language for docker executor: %s", language)
	}

	dir, err := createTempWorkspace(code, filename)
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(dir) // cleanup temp dir

	// using os/exec with docker run --rm is a quick hackathon way to get isolation
	// without dealing with the complex Docker Go SDK
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	// --network none prevents making external requests (e.g. reverse shells)
	// --memory 128m restricts RAM
	// --cpus 0.5 restricts CPU
	args := []string{
		"run", "--rm",
		"--network", "none",
		"--memory", "128m",
		"--cpus", "0.5",
		"-v", fmt.Sprintf("%s:/workspace", dir),
		image,
		"sh", "-c", runCmd,
	}

	cmd := exec.CommandContext(ctx, "docker", args...)
	
	out, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		logger.WarnLog.Printf("Docker execution timeout for %s", dir)
		return string(out) + "\n--- EXECUTION KILLED (TIMEOUT) ---", nil
	}

	return string(out), err
}
