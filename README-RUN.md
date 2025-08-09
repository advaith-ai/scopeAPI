# ScopeAPI - Run Instructions

## 🚀 Quick Start

The ScopeAPI platform has several components that are **ready to run**. Here's how to get everything up and running:

### **Option 1: Automated Startup (Recommended)**

Use the provided startup script to run all services automatically:

```bash
./start-scopeapi.sh
```

This script will:
- ✅ Check dependencies (Go, Node.js, npm)
- ✅ Setup environment variables
- ✅ Check infrastructure (PostgreSQL, Kafka)
- ✅ Start all services with proper configuration
- ✅ Provide health check URLs

**⚠️ Troubleshooting:** If you encounter issues, see [TROUBLESHOOTING.md](./docs/ScopeAPI_TroubleShooting.md) for common solutions.

### **Option 2: Manual Startup**

If you prefer to start services individually:

#### **1. Prerequisites**

Install required tools:
```bash
# Install Go 1.22+
sudo apt-get install golang-go

# Install Node.js 18+
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs

# Install PostgreSQL
sudo apt-get install postgresql postgresql-contrib

# Install Kafka (using Docker)
docker run -d --name kafka \
  -p 9092:9092 \
  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \
  confluentinc/cp-kafka:latest
```

#### **2. Setup Database**

```bash
# Create database
sudo -u postgres createdb scopeapi

# Set environment variables
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=your_password
export DB_NAME=scopeapi
```

#### **3. Start Services**

**Data Ingestion Service:**
```bash
cd backend/services/data-ingestion
go mod tidy
go run cmd/main.go
```
- **URL**: http://localhost:8080
- **Health**: http://localhost:8080/health

**API Discovery Service:**
```bash
cd backend/services/api-discovery
go mod tidy
go run cmd/main.go
```
- **URL**: http://localhost:8080
- **Health**: http://localhost:8080/health

**Threat Detection Service:**
```bash
cd backend/services/threat-detection
go mod tidy
go run cmd/main.go
```
- **URL**: http://localhost:8080
- **Health**: http://localhost:8080/health

**Admin Console Microservice:**
```bash
# Option A: Full microservice build
cd backend/services/admin-console
make full-build
make run

# Option B: Angular development only
cd adminConsole
npm install
npm start
```
- **URL**: http://localhost:8080 (Microservice) or http://localhost:4200 (Angular dev)

## 📊 Available Services

### **✅ Ready to Run**

| Service | Status | Port | Description |
|---------|--------|------|-------------|
| **Data Ingestion** | ✅ Complete | 8080 | Captures and processes API traffic |
| **API Discovery** | ✅ Complete | 8081 | Discovers and catalogs API endpoints |
| **Threat Detection** | ✅ Complete | 8082 | Detects malicious API requests |
| **Admin Console** | ✅ Complete | 8086/8080 | Angular UI dashboard (Microservice) |
| **Attack Blocking** | ✅ Complete | 8084 | Real-time request blocking |
| **Data Protection** | ✅ Complete | 8083 | PII detection and classification |

### **⚠️ Needs Implementation**

| Service | Status | Description |
|---------|--------|-------------|
| Contextualization | ❌ Missing | Attack context enrichment |
| Cloud Intelligence | ❌ Missing | Centralized threat intelligence |

## 🔧 Configuration

### **Environment Variables**

Set these environment variables for proper operation:

```bash
# Database
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=your_password
export DB_NAME=scopeapi

# Kafka
export KAFKA_BROKERS=localhost:9092
export KAFKA_TOPIC_PREFIX=scopeapi

# Service Ports
export DATA_INGESTION_PORT=8080
export API_DISCOVERY_PORT=8081
export THREAT_DETECTION_PORT=8082
export ADMIN_CONSOLE_PORT=8080

# Environment
export ENVIRONMENT=development
```

### **Microservice Communication**

**Important**: In a microservices architecture, services communicate using service names, not localhost:

#### **Docker Compose (Production)**
```yaml
# ✅ Correct - Use service names
services:
  api_discovery:
    url: "http://api-discovery:8080"
  threat_detection:
    url: "http://threat-detection:8080"
```

#### **Local Development**
```yaml
# ✅ Correct - Use localhost for local dev
services:
  api_discovery:
    url: "http://localhost:8081"
  threat_detection:
    url: "http://localhost:8082"
```

#### **Environment Variables**
```bash
# ✅ Correct - Override via environment
export SERVICE_API_DISCOVERY_URL=http://api-discovery:8080
export SERVICE_THREAT_DETECTION_URL=http://threat-detection:8080
```

### **Configuration Files**

The services use YAML configuration files:

- **Data Ingestion**: `backend/services/data-ingestion/config/data-ingestion.yaml`
- **API Discovery**: Uses environment variables and defaults
- **Threat Detection**: Uses environment variables and defaults
- **Admin Console**: 
  - Development: `backend/services/admin-console/config/config.development.yaml`
  - Production: `backend/services/admin-console/config/config.production.yaml`

## 🧪 Testing the Services

### **Data Ingestion Service**

Test traffic ingestion:
```bash
curl -X POST http://localhost:8080/api/v1/ingestion/traffic \
  -H "Content-Type: application/json" \
  -d '{
    "timestamp": "2024-01-15T10:30:00Z",
    "method": "POST",
    "path": "/api/users",
    "status_code": 201,
    "request_size": 1024,
    "response_size": 512
  }'
```

### **API Discovery Service**

Start a discovery scan:
```bash
curl -X POST http://localhost:8081/api/v1/discovery/scan \
  -H "Content-Type: application/json" \
  -d '{
    "target_url": "https://api.example.com",
    "scan_type": "passive"
  }'
```

### **Threat Detection Service**

Analyze traffic for threats:
```bash
curl -X POST http://localhost:8082/api/v1/threats/analyze \
  -H "Content-Type: application/json" \
  -d '{
    "request_data": {
      "method": "POST",
      "path": "/api/login",
      "headers": {"Content-Type": "application/json"},
      "body": "{\"username\":\"admin\",\"password\":\"'\"'\"' OR 1=1--\"}"
    }
  }'
```

### **Admin Console Service**

Test the admin console APIs:
```bash
# Health check
curl http://localhost:8080/api/v1/health

# Dashboard statistics
curl http://localhost:8080/api/v1/dashboard/stats

# Service status (inter-service communication)
curl http://localhost:8080/api/v1/services/status

# User management
curl http://localhost:8080/api/v1/users
```

## 📈 Health Checks

Monitor service health:

```bash
# Data Ingestion
curl http://localhost:8080/health

# API Discovery
curl http://localhost:8081/health

# Threat Detection
curl http://localhost:8082/health

# Admin Console
curl http://localhost:8080/api/v1/health
```
```

## 🐛 Troubleshooting

### **Common Issues**

1. **Port Already in Use**
   ```bash
   # Find process using port
   sudo lsof -i :8080
   
   # Kill process
   sudo kill -9 <PID>
   ```

2. **Database Connection Failed**
   ```bash
   # Check PostgreSQL status
   sudo systemctl status postgresql
   
   # Start PostgreSQL
   sudo systemctl start postgresql
   ```

3. **Kafka Connection Failed**
   ```bash
   # Check if Kafka is running
   docker ps | grep kafka
   
   # Start Kafka
   docker run -d --name kafka -p 9092:9092 confluentinc/cp-kafka:latest
   ```

4. **Go Module Issues**
   ```bash
   # Clean and download modules
   go clean -modcache
   go mod tidy
   ```

5. **Node.js Dependencies**
   ```bash
   # Clear npm cache and reinstall
   npm cache clean --force
   rm -rf node_modules package-lock.json
   npm install
   ```

### **Logs**

Check service logs for errors:

```bash
# Data Ingestion logs
tail -f backend/services/data-ingestion/logs/*.log

# API Discovery logs
tail -f backend/services/api-discovery/logs/*.log

# Threat Detection logs
tail -f backend/services/threat-detection/logs/*.log

# Admin Console logs
tail -f backend/services/admin-console/logs/*.log
```
```

## 🏗️ Admin Console Microservice

The admin console is now a full-fledged microservice that serves both the Angular frontend and provides backend APIs.

### **Architecture**
- **Frontend**: Angular 17+ application served as static files
- **Backend**: Go microservice with Gin framework
- **APIs**: RESTful endpoints for user management, dashboard stats, system monitoring
- **Service Discovery**: Built-in inter-service communication

### **Key Features**
- **Static File Serving**: Serves the Angular application
- **API Gateway**: Centralized admin APIs
- **Service Monitoring**: Health checks for all microservices
- **User Management**: CRUD operations for users
- **Dashboard Statistics**: Real-time metrics and analytics

### **Development Workflow**

#### **Option 1: Full Microservice (Recommended)**
```bash
cd backend/services/admin-console

# Build everything (Angular + Go)
make full-build

# Run the service
make run

# Access at: http://localhost:8080
```

#### **Option 2: Angular Development Only**
```bash
cd adminConsole
npm install
npm start

# Access at: http://localhost:4200
```

#### **Option 3: Docker**
```bash
# Build and run with Docker Compose
docker-compose up admin-console

# Access at: http://localhost:8086
```

### **Configuration**
The service supports environment-based configuration:
- **Development**: `config/config.development.yaml` (uses localhost)
- **Production**: `config/config.production.yaml` (uses service names)

### **Inter-Service Communication**
The admin console can communicate with other microservices:
```bash
# Check service status
curl http://localhost:8080/api/v1/services/status

# This will show health of all microservices:
# - api-discovery
# - threat-detection
# - data-protection
# - gateway-integration
# - attack-blocking
```

## 🚀 Production Deployment

For production deployment, consider:

1. **Docker Containers**: Containerize each service
2. **Kubernetes**: Use K8s for orchestration
3. **Load Balancing**: Use nginx or HAProxy
4. **Monitoring**: Implement Prometheus + Grafana
5. **Logging**: Use ELK stack or similar
6. **Security**: Enable TLS, implement proper authentication

## 📚 API Documentation

Each service provides REST API endpoints:

- **Data Ingestion API**: http://localhost:8080/api/v1/
- **API Discovery API**: http://localhost:8081/api/v1/
- **Threat Detection API**: http://localhost:8082/api/v1/
- **Admin Console API**: http://localhost:8080/api/v1/
- **Attack Blocking API**: http://localhost:8084/api/v1/
- **Data Protection API**: http://localhost:8083/api/v1/

## 🎯 Next Steps

1. **Test all services** using the provided examples
2. **Explore the admin console** at http://localhost:8080 (Microservice) or http://localhost:4200 (Angular dev)
3. **Test inter-service communication** using the service status endpoint
4. **Add authentication** and security features
5. **Scale the services** for production use
6. **Implement missing components** (Contextualization, Cloud Intelligence)

## 📞 Support

If you encounter issues:

1. Check the troubleshooting section above
2. Review service logs for error messages
3. Verify all dependencies are installed
4. Ensure infrastructure (DB, Kafka) is running

---

**🎉 Congratulations! You now have a working ScopeAPI platform!** 
