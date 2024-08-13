# Sentinel: Enhanced Logging and Performance Monitoring Microservice

Sentinel is a robust microservice designed for monitoring Docker containers, collecting logs, parsing them for insights, and providing performance metrics and alerting. It's particularly suited for low-latency environments such as quantitative trading systems.

## Features

- Concurrent monitoring of multiple Docker containers
- Sophisticated log parsing and analysis
- Prometheus metrics for log lines, processing time, errors, and latency
- Configurable alerting system
- Extensible architecture for easy addition of new features

## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go 1.21 or later

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/sentinel.git
   cd sentinel
   ```

2. Build the Docker image:
   ```
   docker build -t sentinel .
   ```

3. Set up environment variables:
   Create a `.env` file in the project root and add the following variables:
   ```
   PORT=8080
   DB_DATABASE=sentinel
   DB_USERNAME=user
   DB_PASSWORD=password
   DB_PORT=5432
   ```

4. Start the services using Docker Compose:
   ```
   docker-compose up -d
   ```

## Usage

Once the services are up and running, you can access:

- Metrics: `http://localhost:8080/metrics`
- Prometheus: `http://localhost:9090`
- Grafana: `http://localhost:3000`

## Configuration

The service can be configured using environment variables:

- `METRICS_ADDR`: Address to serve metrics (default: ":8080")
- `DOCKER_HOST`: Docker daemon socket (default: "unix:///var/run/docker.sock")
- `SMTP_HOST`: SMTP server for sending alerts
- `SMTP_PORT`: SMTP server port
- `SMTP_FROM`: Email address to send alerts from
- `SMTP_PASSWORD`: Password for the email account

## Development

### Running Tests

To run unit tests:
```
go test ./...
```

To run integration tests (requires Docker):
```
go test ./test/integration -tags=integration
```

### Linting

We use golangci-lint for linting. To run the linter:
```
golangci-lint run
```

### Pre-commit Hooks

To set up pre-commit hooks:
```
./scripts/setup_pre_commit.sh
```

## Extending the Project

To add more features:

1. Create new packages in the `internal/` directory
2. Add new metrics in `internal/metrics/metrics.go`
3. Extend the log parsing logic in `internal/logparser/parser.go`
4. Add new alerts in `internal/alerting/alerting.go`
5. Update the monitor logic in `internal/monitor/monitor.go`
6. Update the main function in `cmd/monitor/main.go` as needed

## CI/CD

The project uses GitHub Actions for CI/CD. The workflow includes:

- Running tests
- Linting
- Building and pushing Docker images
- Deploying to Kubernetes (including canary deployments)
- Security scanning with CodeQL and Trivy

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
