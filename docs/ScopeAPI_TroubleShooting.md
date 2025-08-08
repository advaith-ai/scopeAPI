# ScopeAPI Troubleshooting Guide

This guide helps you resolve common issues when starting the ScopeAPI platform.

## Common Issues and Solutions

### 1. Script Syntax Error: "too many arguments"

**Error:**
```
./start-scopeapi.sh: line 22: [: too many arguments
```

**Solution:**
This has been fixed in the latest version. If you're still seeing this error, make sure you're using the updated `start-scopeapi.sh` script.

### 2. Docker Permission Denied

**Error:**
```
permission denied while trying to connect to the Docker daemon socket
```

**Solution:**
Run the Docker permissions fix script:
```bash
./fix-docker-permissions.sh
```

Or manually:
```bash
# Add user to docker group
sudo usermod -aG docker $USER

# Apply changes (choose one):
# Option 1: Log out and log back in
# Option 2: Run this command:
newgrp docker
```

### 3. Kafka Connection Failures

**Error:**
```
Failed to consume Kafka messages: connection refused
```

**Solution:**
Kafka is optional for development. The service will continue to work without Kafka, but you'll see these error messages. To eliminate them:

**Option 1: Start Kafka with Docker**
```bash
# Start Kafka container
docker run -d --name kafka -p 9092:9092 confluentinc/cp-kafka:latest

# Or use docker-compose (if available)
docker-compose up -d kafka
```

**Option 2: Disable Kafka (Recommended for development)**
The service will work without Kafka. The error messages will be reduced with exponential backoff.

### 4. Admin Console Compilation Errors

**Error:**
Multiple TypeScript compilation errors in Angular admin console.

**Solution:**
These are development-time errors that don't prevent the service from running. The admin console will still be accessible at `http://localhost:4200`.

To fix the compilation errors:
```bash
cd adminConsole
npm install
npm audit fix
```

### 5. PostgreSQL Connection Issues

**Error:**
```
PostgreSQL is not running or not accessible
```

**Solution:**
```bash
# Start PostgreSQL service
sudo systemctl start postgresql

# Or use Docker
docker run -d --name postgres -e POSTGRES_PASSWORD=password -p 5432:5432 postgres:13
```

### 6. Port Already in Use

**Error:**
```
bind: address already in use
```

**Solution:**
```bash
# Find process using the port
sudo lsof -i :8080

# Kill the process
sudo kill -9 <PID>

# Or use different ports
export DATA_INGESTION_PORT=8081
./start-scopeapi.sh
```

## Service Status Check

After starting the platform, verify all services are running:

```bash
# Check if services are responding
curl http://localhost:8080/health
curl http://localhost:8080/metrics

# Check admin console
curl http://localhost:4200
```

## Environment Variables

You can customize the startup by setting environment variables:

```bash
# Database configuration
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=password
export DB_NAME=scopeapi

# Service ports
export DATA_INGESTION_PORT=8080
export API_DISCOVERY_PORT=8081
export THREAT_DETECTION_PORT=8082

# Kafka configuration
export KAFKA_BROKERS=localhost:9092
export KAFKA_TOPIC_PREFIX=scopeapi
```

## Development vs Production

### Development Mode
- Kafka is optional
- PostgreSQL can be local or Docker
- Admin Console runs in development mode
- Detailed logging enabled

### Production Mode
- All services required
- Proper database setup
- Kafka cluster required
- Admin Console built for production

## Getting Help

If you're still experiencing issues:

1. Check the logs in the terminal output
2. Verify all dependencies are installed
3. Ensure ports are not in use
4. Check firewall settings
5. Review the error messages for specific guidance

## Quick Start for Development

For a quick development setup without Kafka:

```bash
# 1. Fix Docker permissions (if needed)
./fix-docker-permissions.sh

# 2. Start PostgreSQL (if not running)
sudo systemctl start postgresql

# 3. Start ScopeAPI
./start-scopeapi.sh
```

The platform will start with:
- Data Ingestion Service: http://localhost:8080
- Admin Console: http://localhost:4200
- Health Check: http://localhost:8080/health
- Metrics: http://localhost:8080/metrics 