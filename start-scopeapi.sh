#!/bin/bash

# ScopeAPI Startup Script
# This script helps you start all ScopeAPI components

set -e

echo "🚀 Starting ScopeAPI Platform..."
echo "=========================================="
echo "🚀 ScopeAPI Platform Startup"
echo "=========================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}[SUCCESS]${NC} $2"
    else
        echo -e "${YELLOW}[WARNING]${NC} $2"
    fi
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Check if required tools are installed
check_dependencies() {
    echo "[INFO] Checking dependencies..."
    
    # Check Go
    if command_exists go; then
        print_status 0 "Go is installed"
    else
        print_status 1 "Go is not installed"
        echo "[INFO] Please install Go from https://golang.org/dl/"
        exit 1
    fi
    
    # Check Node.js
    if command_exists node; then
        print_status 0 "Node.js is installed"
    else
        print_status 1 "Node.js is not installed"
        echo "[INFO] Please install Node.js from https://nodejs.org/"
    fi
    
    # Check npm
    if command_exists npm; then
        print_status 0 "npm is installed"
    else
        print_status 1 "npm is not installed"
    fi
    
    echo "[SUCCESS] All dependencies are installed"
}

# Setup environment variables
setup_environment() {
    echo "[INFO] Setting up environment variables..."
    
    # Database configuration
    export DB_HOST=${DB_HOST:-localhost}
    export DB_PORT=${DB_PORT:-5432}
    export DB_USER=${DB_USER:-postgres}
    export DB_PASSWORD=${DB_PASSWORD:-password}
    export DB_NAME=${DB_NAME:-scopeapi}
    
    # Kafka configuration
    export KAFKA_BROKERS=${KAFKA_BROKERS:-localhost:9092}
    export KAFKA_TOPIC_PREFIX=${KAFKA_TOPIC_PREFIX:-scopeapi}
    
    # Server ports
    export DATA_INGESTION_PORT=${DATA_INGESTION_PORT:-8080}
    export API_DISCOVERY_PORT=${API_DISCOVERY_PORT:-8081}
    export THREAT_DETECTION_PORT=${THREAT_DETECTION_PORT:-8082}
    
    export GO111MODULE=on
    export GOPATH=$HOME/go
    
    print_status 0 "Environment variables configured"
}

# Check if PostgreSQL is running
check_postgresql() {
    echo "[INFO] Checking PostgreSQL connection..."
    
    if command_exists psql; then
        if pg_isready -h $DB_HOST -p $DB_PORT -U $DB_USER &> /dev/null; then
            print_status 0 "PostgreSQL is running"
            return 0
        else
            print_status 1 "PostgreSQL is not running or not accessible"
            print_warning "Please start PostgreSQL and ensure it's accessible at $DB_HOST:$DB_PORT"
            print_warning "You can start PostgreSQL with: sudo systemctl start postgresql"
            return 1
        fi
    else
        print_status 1 "PostgreSQL client not found"
        return 1
    fi
}

# Check if Kafka is running
check_kafka() {
    echo "[INFO] Checking Kafka connection..."
    
    if command_exists docker; then
        # Check if user has permission to run docker
        if docker ps &>/dev/null; then
            if docker ps | grep -q kafka; then
                print_status 0 "Kafka is running in Docker"
                return 0
            else
                print_status 1 "Kafka is not running or not accessible"
                print_warning "Please start Kafka and ensure it's accessible at localhost:9092"
                print_warning "You can start Kafka with Docker:"
                print_warning "docker run -d --name kafka -p 9092:9092 confluentinc/cp-kafka:latest"
                return 1
            fi
        else
            print_status 1 "Docker permission denied - add user to docker group or run with sudo"
            print_warning "Run: sudo usermod -aG docker $USER && newgrp docker"
            return 1
        fi
    else
        print_status 1 "Docker not found"
        return 1
    fi
}

# Start Data Ingestion Service
start_data_ingestion() {
    echo "[INFO] Starting Data Ingestion Service..."
    
    # Build services using centralized Makefile
    echo "[INFO] Building backend services..."
    cd backend
    if make data-ingestion; then
        echo "[SUCCESS] Data Ingestion Service built successfully"
        
        # Start the service from root bin directory
        echo "[INFO] Starting Data Ingestion Service on port $DATA_INGESTION_PORT..."
        cd ..
        ./bin/data-ingestion &
        DATA_INGESTION_PID=$!
        echo "[SUCCESS] Data Ingestion Service started with PID: $DATA_INGESTION_PID"
    else
        echo "[ERROR] Failed to build Data Ingestion Service"
        exit 1
    fi
}

# Start API Discovery Service
start_api_discovery() {
    echo "[INFO] Starting API Discovery Service..."
    
    # Build API Discovery Service
    cd backend
    if make api-discovery; then
        echo "[SUCCESS] API Discovery Service built successfully"
        
        # Start the service from root bin directory
        echo "[INFO] Starting API Discovery Service on port $API_DISCOVERY_PORT..."
        cd ..
        ./bin/api-discovery &
        API_DISCOVERY_PID=$!
        echo "[SUCCESS] API Discovery Service started with PID: $API_DISCOVERY_PID"
    else
        echo "[ERROR] Failed to build API Discovery Service"
        return 1
    fi
}

# Start Threat Detection Service
start_threat_detection() {
    echo "[INFO] Starting Threat Detection Service..."
    
    # Build Threat Detection Service (if compilation issues are fixed)
    cd backend
    if make threat-detection 2>/dev/null; then
        echo "[SUCCESS] Threat Detection Service built successfully"
        
        # Start the service from root bin directory
        echo "[INFO] Starting Threat Detection Service on port $THREAT_DETECTION_PORT..."
        cd ..
        ./bin/threat-detection &
        THREAT_DETECTION_PID=$!
        echo "[SUCCESS] Threat Detection Service started with PID: $THREAT_DETECTION_PID"
    else
        echo "[WARNING] Threat Detection Service has compilation issues, skipping..."
        return 0
    fi
}

# Start Admin Console
start_admin_console() {
    echo "[INFO] Starting Admin Console..."
    
    cd adminConsole
    
    # Check if admin console dependencies are installed
    if [ -d "node_modules" ]; then
        echo "[INFO] Admin Console dependencies are already installed"
    else
        echo "[INFO] Installing Admin Console dependencies..."
        npm install
    fi
    
    echo "[INFO] Starting Angular development server..."
    npm start &
    ADMIN_CONSOLE_PID=$!
    echo "[SUCCESS] Admin Console started with PID: $ADMIN_CONSOLE_PID"
    
    cd ..
}

# Main execution
main() {
    echo "=========================================="
    echo "🎉 ScopeAPI Platform is starting up!"
    echo "=========================================="
    
    # Check dependencies
    check_dependencies
    
    # Setup environment
    setup_environment
    
    # Check infrastructure
    check_postgresql || print_warning "Continuing without PostgreSQL..."
    check_kafka || print_warning "Continuing without Kafka..."
    
    # Start services
    start_data_ingestion
    start_api_discovery
    start_threat_detection
    start_admin_console
    
    echo ""
    echo "📊 Data Ingestion Service: http://localhost:$DATA_INGESTION_PORT"
    echo "🔍 API Discovery Service: http://localhost:$API_DISCOVERY_PORT"
    echo "🛡️ Threat Detection Service: http://localhost:$THREAT_DETECTION_PORT"
    echo "🌐 Admin Console: http://localhost:4200 (if started)"
    echo ""
    echo "📈 Health Checks:"
    echo "  • Data Ingestion: http://localhost:$DATA_INGESTION_PORT/health"
    echo "  • API Discovery: http://localhost:$API_DISCOVERY_PORT/health"
    echo "  • Threat Detection: http://localhost:$THREAT_DETECTION_PORT/health"
    echo "📊 Metrics: http://localhost:$DATA_INGESTION_PORT/metrics"
    echo ""
    echo "Press Ctrl+C to stop all services"
    
    # Wait for interrupt
    trap 'echo ""; echo "🛑 Shutting down ScopeAPI Platform..."; kill $DATA_INGESTION_PID $API_DISCOVERY_PID $THREAT_DETECTION_PID $ADMIN_CONSOLE_PID 2>/dev/null; exit 0' INT
    wait
}

# Run main function
main "$@" 