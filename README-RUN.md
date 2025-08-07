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

**Frontend:**
```bash
cd frontend
npm install
npm start
```
- **URL**: http://localhost:4200

## 📊 Available Services

### **✅ Ready to Run**

| Service | Status | Port | Description |
|---------|--------|------|-------------|
| **Data Ingestion** | ✅ Complete | 8080 | Captures and processes API traffic |
| **API Discovery** | ✅ Complete | 8081 | Discovers and catalogs API endpoints |
| **Threat Detection** | ✅ Complete | 8082 | Detects malicious API requests |
| **Frontend** | ✅ Complete | 4200 | Angular UI dashboard |

### **⚠️ Needs Implementation**

| Service | Status | Description |
|---------|--------|-------------|
| Attack Blocking | 🚧 Partial | Real-time request blocking |
| Data Protection | 🚧 Partial | PII detection and classification |
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
```

### **Configuration Files**

The services use YAML configuration files:

- **Data Ingestion**: `backend/services/data-ingestion/config/data-ingestion.yaml`
- **API Discovery**: Uses environment variables and defaults
- **Threat Detection**: Uses environment variables and defaults

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

## 📈 Health Checks

Monitor service health:

```bash
# Data Ingestion
curl http://localhost:8080/health

# API Discovery
curl http://localhost:8081/health

# Threat Detection
curl http://localhost:8082/health

# Frontend
curl http://localhost:4200
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

## 🎯 Next Steps

1. **Test all services** using the provided examples
2. **Explore the frontend** at http://localhost:4200
3. **Implement missing components** (Attack Blocking, Data Protection)
4. **Add authentication** and security features
5. **Scale the services** for production use

## 📞 Support

If you encounter issues:

1. Check the troubleshooting section above
2. Review service logs for error messages
3. Verify all dependencies are installed
4. Ensure infrastructure (DB, Kafka) is running

---

**🎉 Congratulations! You now have a working ScopeAPI platform!** 
