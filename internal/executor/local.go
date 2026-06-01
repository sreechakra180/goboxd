package executor

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/user/blackroot/internal/logger"
)

type LocalExecutor struct {
	timeout time.Duration
}

func NewLocalExecutor() *LocalExecutor {
	return &LocalExecutor{
		// keeping timeout low for now to avoid abuse (infinite loops etc)
		timeout: 3 * time.Second,
	}
}

func (e *LocalExecutor) Run(language, code string) (string, error) {
	var filename string
	var cmdArgs []string

	switch language {
	case "python":
		filename = "main.py"
		cmdArgs = []string{"python3", "main.py"} // assume python3 is installed
	case "javascript", "node":
		filename = "main.js"
		cmdArgs = []string{"node", "main.js"}
	case "go":
		filename = "main.go"
		cmdArgs = []string{"go", "run", "main.go"}
	case "cpp", "c++":
		filename = "main.cpp"
		// temp fix for execution edge case: c++ needs compile then run
		// we'll just use a quick bash script or compound command for local
		// actually, let's keep it simple and compile it first
		return e.runCPP(code)
	default:
		return "", fmt.Errorf("unsupported language: %s", language)
	}

	return e.executeCommand(code, filename, cmdArgs[0], cmdArgs[1:]...)
}

func (e *LocalExecutor) runCPP(code string) (string, error) {
	dir, err := createTempWorkspace(code, "main.cpp")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(dir) // automatic cleanup

	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	// 1. Compile
	compileCmd := exec.CommandContext(ctx, "g++", "main.cpp", "-o", "main")
	compileCmd.Dir = dir
	if out, err := compileCmd.CombinedOutput(); err != nil {
		return string(out) + "\nCompilation Error", nil
	}

	// 2. Run
	runCmd := exec.CommandContext(ctx, "./main")
	runCmd.Dir = dir
	out, err := runCmd.CombinedOutput()
	
	if ctx.Err() == context.DeadlineExceeded {
		return string(out) + "\nExecution Timeout", nil
	}

	return string(out), err
}

func (e *LocalExecutor) executeCommand(code, filename, cmdName string, args ...string) (string, error) {
	dir, err := createTempWorkspace(code, filename)
	if err != nil {
		return "", err
	}
	// clean up temp dirs when done!
	defer os.RemoveAll(dir)

	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, cmdName, args...)
	cmd.Dir = dir
	
	// combine stdout and stderr for easier API responses
	out, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		logger.WarnLog.Printf("Process killed on timeout in %s", dir)
		return string(out) + "\n--- EXECUTION KILLED (TIMEOUT) ---", nil
	}

	return string(out), err
}
