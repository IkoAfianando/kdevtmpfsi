# Process Killer Service

A Go-based service that monitors and automatically terminates a specific process (`kdevtmpfsi`) at regular intervals.

## Overview

This service continuously monitors the system for a process named `kdevtmpfsi` and automatically kills it when detected. It runs as a daemon service with configurable check intervals.

## Features

- **Process Monitoring**: Continuously monitors for the target process
- **Automatic Termination**: Sends SIGKILL signal to terminate the process
- **Configurable Interval**: Check interval set to 10 seconds by default
- **Logging**: Comprehensive logging of all operations
- **Docker Support**: Can be run as a containerized service
- **Host PID Access**: Uses host PID namespace for process monitoring

## Requirements

- Go 1.23.5 or higher
- Docker and Docker Compose (for containerized deployment)
- Linux environment

## Installation

### Local Installation

1. Clone the repository:
```bash
git clone https://github.com/IkoAfianando/kdevtmpfsi
cd kdevtmpfsi
```

2. Build the application:
```bash
go build -o process-killer main.go
```

3. Run the service:
```bash
./process-killer
```

### Docker Installation

1. Build and run using Docker Compose:
```bash
docker-compose up -d
```

2. Check service status:
```bash
docker-compose ps
```

3. View logs:
```bash
docker-compose logs -f process-killer
```

## Configuration

### Environment Variables

The service uses the following constants that can be modified in the source code:

- `processName`: The name of the process to monitor (default: "kdevtmpfsi")
- `checkInterval`: Time interval between checks (default: 10 seconds)

### Docker Configuration

The Docker Compose configuration includes:

- **Container Name**: `kdevtmpfsi-killer`
- **PID Mode**: `host` (required for accessing host processes)
- **Restart Policy**: `always`

## How It Works

1. The service starts and immediately performs an initial check
2. It then runs continuously, checking every 10 seconds
3. For each check, it:
   - Executes `ps aux | grep kdevtmpfsi | grep -v grep`
   - Parses the output to extract the PID
   - Sends a SIGKILL signal (`kill -9`) to the process
   - Logs the operation results

## Logging

The service provides detailed logging for:
- Service startup
- Process detection
- Kill operations (success/failure)
- Error conditions

## Docker Components

### Dockerfile

Uses a multi-stage build:
- **Builder Stage**: Uses `golang:1.23.5-alpine` to compile the Go application
- **Runtime Stage**: Uses `alpine:latest` with `procps` package for process utilities

### Docker Compose

Configures the service with:
- Host PID namespace access
- Automatic restart policy
- Custom container name

## Usage Examples

### Running Locally
```bash
# Build and run
go run main.go

# Or build first, then run
go build -o process-killer main.go
./process-killer
```

### Running with Docker
```bash
# Start the service
docker-compose up -d

# Stop the service
docker-compose down

# Rebuild and restart
docker-compose up -d --build
```

## Security Considerations

- The service requires elevated privileges to kill processes
- When running in Docker, it uses host PID namespace
- SIGKILL is used for immediate termination (non-graceful)

## Troubleshooting

### Common Issues

1. **Process not found**: Ensure the target process name is correct
2. **Permission denied**: Service may need elevated privileges
3. **Docker access**: Ensure Docker daemon is running and accessible

### Logs

Check service logs for troubleshooting:
```bash
# Docker logs
docker-compose logs process-killer

# Local logs
# Logs are printed to stdout when running locally
```

## License

This project is provided as-is for educational and utility purposes.

## Contributing

Feel free to submit issues and enhancement requests.
