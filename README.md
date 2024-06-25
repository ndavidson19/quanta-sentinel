# quanta-sentinel

## Overview
quanta-sentinel is the monitoring and logging service for the QuantForge platform. It provides centralized logging, monitoring, and alerting capabilities for all microservices.

## Key Features
- Centralized log aggregation
- Real-time monitoring of system health
- Customizable alerting system
- Performance metrics collection and visualization
- Automated incident response

## Technology Stack
- Go
- Prometheus for metrics collection
- Grafana for metrics visualization
- Loki for log aggregation
- Alertmanager for alert management

## Setup
1. Clone the repository:
   ```
   git clone https://github.com/quantforge/quanta-sentinel.git
   cd quanta-sentinel
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Set up environment variables:
   ```
   cp .env.example .env
   # Edit .env with your configuration
   ```
4. Build and run the service:
   ```
   go build
   ./quanta-sentinel
   ```

## Accessing Dashboards
- Grafana: http://localhost:3000
- Prometheus: http://localhost:9090

## Configuring Alerts
Edit `config/alerts.yaml` to set up custom alerts.

## Testing
```
go test ./...
```

## Contributing
Please read CONTRIBUTING.md for details on our code of conduct and the process for submitting pull requests.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
