# forge-sentinel

## Overview
forge-sentinel is the monitoring and logging service for the QuantForge platform. It provides centralized logging, monitoring, and alerting capabilities for all microservices.

## Key Features
- Centralized log aggregation
- Real-time monitoring of system health
- Customizable alerting system
- Performance metrics collection and visualization
- Automated incident response

## Technology Stack
- ELK Stack (Elasticsearch, Logstash, Kibana)
- Prometheus for metrics collection
- Grafana for metrics visualization
- Node.js for custom monitoring scripts

## Setup
1. Clone the repository:
   ```
   git clone https://github.com/quantforge/forge-sentinel.git
   cd forge-sentinel
   ```
2. Install dependencies:
   ```
   npm install
   ```
3. Set up environment variables:
   ```
   cp .env.example .env
   # Edit .env with your configuration
   ```
4. Start the ELK stack (requires Docker):
   ```
   docker-compose up -d
   ```
5. Start the Node.js monitoring service:
   ```
   npm start
   ```

## Accessing Dashboards
- Kibana: http://localhost:5601
- Grafana: http://localhost:3000

## Configuring Alerts
Edit `config/alerts.yaml` to set up custom alerts.

## Testing
```
npm test
```

## Contributing
Please read CONTRIBUTING.md for details on our code of conduct and the process for submitting pull requests.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
