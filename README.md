# ping-cli

A fast and concurrent HTTP endpoint monitoring tool built with Go.

## Overview

`ping-cli` is a command-line utility that allows you to monitor the health and response times of multiple HTTP endpoints simultaneously. It uses Go's concurrency features to ping multiple URLs in parallel, providing quick status reports with visual indicators.

## Features

- 🚀 **Concurrent pinging** - Monitor multiple endpoints simultaneously
- ⚡ **Fast response times** - Get results in milliseconds
- 📊 **Visual status indicators** - Clear emoji-based status reporting
- 🛠️ **Simple CLI interface** - Easy to use command-line tool
- 🔄 **HTTP status monitoring** - Detects UP, DOWN, and UNHEALTHY states

## Installation

### Prerequisites

- Go 1.24.4 or later

### Build from source

```bash
git clone https://github.com/fearfactor3/ping-cli-project.git
cd ping-cli-project
go build -o ping-cli
```

### Install directly

```bash
go install github.com/fearfactor3/ping-cli-project@latest
```

## Usage

### Basic Usage

```bash
./ping-cli <URL1> [URL2] [URL3] ...
```

### Examples

**Monitor a single endpoint:**
```bash
./ping-cli https://google.com
```

**Monitor multiple endpoints:**
```bash
./ping-cli https://google.com https://github.com https://stackoverflow.com
```

**Monitor local services:**

```bash
./ping-cli http://localhost:3000 http://localhost:8080
```

## Output Format

The tool provides clear status indicators for each endpoint:

- ✅ **UP** - Endpoint is healthy (HTTP 2xx status)
- ❌ **DOWN** - Endpoint is unreachable (connection error)
- ⚠️ **UNHEALTHY** - Endpoint responded with non-2xx status code

### Example Output

```text
✅ https://google.com - UP (45.23ms)
✅ https://github.com - UP (120.45ms)
❌ https://nonexistent-site.com - DOWN (no such host)
⚠️ https://httpstat.us/404 - UNHEALTHY (HTTP 404)
```