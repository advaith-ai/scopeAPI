# ScopeAPI - Comprehensive API Security Platform

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![Angular Version](https://img.shields.io/badge/Angular-17+-red.svg)](https://angular.io/)

## 🚀 Overview

ScopeAPI is a comprehensive, enterprise-grade API security platform designed to provide complete visibility, protection, and testing capabilities for modern API ecosystems. Built with a distributed microservices architecture, ScopeAPI offers real-time threat detection, automated security testing, and intelligent attack prevention powered by machine learning and cloud intelligence.

## ✨ Key Features

- 🔍 **Endpoint Discovery** - Automatic discovery and cataloging of API endpoints
- 🛡️ **Sensitive Data Scanning** - PII detection and risk scoring
- ⚡ **Attack Detection** - ML-powered anomaly and behavioral analysis
- 🔍 **Attack Context** - Full context around attacks for quick vulnerability fixes
- 🛡️ **Attack Blocking** - Real-time threat blocking with cloud intelligence
- 🧪 **API Security Testing** - Automated OWASP API Top 10 vulnerability testing
- 🔗 **CI/CD Integration** - Seamless integration with development pipelines
- 📊 **Intelligent Analytics** - Interactive dashboards and comprehensive reporting

## 🏗️ Architecture

ScopeAPI employs a distributed, microservices-based architecture designed for scalability, resilience, and maintainability.

### **Core Services**
- **API Discovery Service**: Automatic endpoint discovery and cataloging
- **Threat Detection Engine**: ML-powered anomaly and behavioral analysis
- **Attack Blocking Engine**: Real-time threat blocking with cloud intelligence
- **Data Protection Service**: PII detection and compliance management
- **Security Testing Engine**: Automated vulnerability testing
- **Gateway Integration Service**: Multi-gateway management (Kong, NGINX, Traefik, Envoy, HAProxy)
- **Data Ingestion Service**: High-volume traffic processing

### **Technology Stack**
- **Backend**: Go microservices with Gin framework
- **Frontend**: Angular 17+ with TypeScript
- **Databases**: PostgreSQL (relational), MongoDB (document), Neo4j (graph), Redis (cache)
- **Message Queues**: Apache Kafka for event-driven communication
- **ML/AI**: TensorFlow, PyTorch, Apache Spark
- **Container Orchestration**: Kubernetes, Docker
- **Monitoring**: Prometheus, Grafana, ELK Stack

### **Key Architectural Principles**
- **Microservices**: Independent, scalable services
- **Event-Driven**: Asynchronous communication via Kafka
- **Polyglot Persistence**: Multi-database strategy for optimal performance
- **Cloud-Native**: Containerized deployment
- **Security-First**: Zero-trust architecture
- **Observability**: Full-stack monitoring and tracing

For detailed architecture information, see [Technical Architecture](ScopeAPI_Technical_Architecture.md).

## 🚀 Quick Start

### Prerequisites

- Go 1.21+
- Node.js 18+
- Angular CLI 17+
- Docker & Docker Compose
- PostgreSQL 14+

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-org/scopeapi.git
   cd scopeapi
   ```

2. **Start the backend services**
   ```bash
   # Using the provided script
   ./start-scopeapi.sh
   
   # Or manually with Docker Compose
   docker-compose up -d
   ```

3. **Start the frontend**
   ```bash
   cd frontend
   npm install
   npm start
   ```

4. **Access the application**
   - Frontend: http://localhost:4200
   - API Gateway: http://localhost:8080

### Development Setup

For detailed development setup instructions, see [README-RUN.md](README-RUN.md).

## 📚 Documentation

- **[Product Overview](ScopeAPI_Product_Overview.md)** - High-level product overview and features
- **[Technical Architecture](ScopeAPI_Technical_Architecture.md)** - Detailed system architecture and design
- **[Project Structure](ScopeAPI_Project_Structure.md)** - Complete project structure and organization

## 🛠️ Development

### Backend Development

```bash
cd backend
go mod tidy
go run ./services/api-discovery/cmd/main.go
```

### Frontend Development

```bash
cd frontend
npm install
ng serve
```

### Running Tests

```bash
# Backend tests
cd backend
go test ./...

# Frontend tests
cd frontend
npm test
```

## 🔧 Configuration

Configuration files are located in:
- Backend services: `backend/config/`
- Frontend: `frontend/src/environments/`

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

- **Documentation**: Check the [docs](docs/) directory
- **Issues**: Report bugs and feature requests via [GitHub Issues](https://github.com/your-org/scopeapi/issues)
- **Discussions**: Join the conversation in [GitHub Discussions](https://github.com/your-org/scopeapi/discussions)

## 🏢 Enterprise

For enterprise support, custom deployments, and professional services, please contact us at enterprise@scopeapi.com.

---

**ScopeAPI** - Securing APIs with Intelligence 