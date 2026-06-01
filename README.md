# Blackroot 🛡️

**AI-Powered Secure Polyglot Sandbox Execution Platform**

Blackroot is a high-performance secure code execution platform built in Go. It functions as a lightweight secure online judge, sandbox runtime, and threat detection engine all in one. 

Built rapidly during a hackathon sprint. 🚀

## Core Architecture

Blackroot uses a two-tier execution strategy:
1. **Threat Detection Layer (Static Analysis)**: Before execution, the code is scanned against a set of language-specific regex rules to catch obvious malicious intents (`rm -rf`, fork bombs, subprocess abuses).
2. **Execution Sandbox**: We spin up isolated temporary workspaces. We have a `LocalExecutor` for raw speed (with strict context timeouts) and a `DockerExecutor` for heavy-duty containerized isolation (no network, restricted RAM/CPU).

## Features
- **Multi-language**: Python, Node.js, Go, C++
- **Threat Detection**: Blocks file system abuse, os.system, eval, and fork bombs.
- **Resource Constraints**: Strict timeouts and disposable environments.
- **Execution Analytics**: In-memory logging of execution speed, language stats, and security alerts.

## Setup & Run

### Prerequisites
- Go 1.21+
- Docker (optional, but recommended for full isolation)

### Quick Start

1. **Clone & Run**
```bash
make run
```
By default, this uses the LocalExecutor.

2. **Run with Docker Isolation**
```bash
make run-docker
```
*Note: Make sure Docker daemon is running.*

## API Endpoints

### `POST /run`
Executes code securely.
```json
// Request
{
  "language": "python",
  "code": "print('Hello World')"
}

// Response
{
  "status": "success",
  "output": "Hello World\n",
  "duration": 15400200
}
```

### `POST /scan`
Dry-run the threat detection scanner.
```json
// Request
{
  "language": "python",
  "code": "import os; os.system('rm -rf /')"
}

// Response
{
  "safe": false,
  "alerts": [
    "os.system calls are not allowed"
  ]
}
```

### `GET /logs`
Fetch analytics for recent executions.

## Future Improvements (Post-Hackathon)
- Implement AST parsing for deeper threat analysis instead of regex
- Swap in-memory logs for Redis / PostgreSQL
- WebSockets for streaming execution output
- More granular cgroups resource limits

---
*Built with ❤️ by a tired systems engineer at 3 AM.*
