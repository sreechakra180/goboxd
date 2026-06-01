<<<<<<< HEAD
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
=======
<div align="center">

# goboxd

**A Go HTTP service for executing untrusted code in isolated sandboxes.**

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.23-00ADD8.svg?logo=go&logoColor=white)](https://go.dev)
[![Docker](https://img.shields.io/badge/Docker-Required-2496ED.svg?logo=docker&logoColor=white)](https://www.docker.com)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/thesouldev/goboxd/pulls)

</div>

---

## Overview

goboxd is an HTTP service written in Go that compiles and runs untrusted code inside isolated sandboxes and returns the result. Optional test cases can be supplied to assert behaviour against expected output. It is built for safe execution of code across many languages, with strict isolation, bounded concurrency, and a plug and play language registry.

## Features

- Plug and play language registry driven by YAML
- Process isolation using Linux namespaces and cgroups
- Bounded concurrency with request queuing
- Fully containerised for local development and deployment
- Per request resource limits for time, memory, and processes
- Liveness and readiness probes for orchestration

## Getting started

### Prerequisites

- Docker with Compose v2

No Go toolchain or system dependencies are required on the host. Everything runs in containers.

### Installation

```sh
git clone https://github.com/thesouldev/goboxd.git
cd goboxd
make build
```

### Usage

```sh
make run          # start the service on :8080
make test         # run unit tests
make integration  # run end to end tests
make lint         # run static analysis
```

## Project structure

```
.
├── cmd/goboxd/   binary entry point
├── internal/     private application packages
├── docs/         api, languages, security, benchmarks, architecture
└── tests/        integration tests
```

## Contributing

Contributions are welcome. Open an issue to discuss substantial changes before sending a pull request.

## License

This project is distributed under the GNU General Public License v3.0. See [LICENSE](LICENSE) for the full text.
>>>>>>> ea5ab73544b662686ae9797ef36e03b1aced31bc
