#!/bin/bash

# ScopeAPI Database Integration Setup
# This script sets up complete database integration for all ScopeAPI services

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    local status=$1
    local message=$2
    if [ "$status" -eq 0 ]; then
        echo -e "${GREEN}[SUCCESS]${NC} $message"
    else
        echo -e "${RED}[ERROR]${NC} $message"
    fi
}

# Function to print info
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

# Function to print warning
print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# Default configuration
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-password}
DB_NAME=${DB_NAME:-scopeapi}
DB_SSL_MODE=${DB_SSL_MODE:-disable}

print_info "ScopeAPI Database Integration Setup"
print_info "=================================="

# Check if Docker is available for PostgreSQL
check_docker() {
    print_info "Checking Docker availability..."
    
    if command -v docker >/dev/null 2>&1; then
        print_status 0 "Docker found"
        return 0
    else
        print_status 1 "Docker not found"
        print_warning "Docker is recommended for easy PostgreSQL setup"
        return 1
    fi
}

# Start PostgreSQL with Docker if available
start_postgresql_docker() {
    print_info "Starting PostgreSQL with Docker..."
    
    # Check if PostgreSQL container is already running
    if docker ps --format "table {{.Names}}" | grep -q "scopeapi-postgres"; then
        print_status 0 "PostgreSQL container already running"
        return 0
    fi
    
    # Start PostgreSQL container
    if docker run --name scopeapi-postgres \
        -e POSTGRES_PASSWORD="$DB_PASSWORD" \
        -e POSTGRES_DB="$DB_NAME" \
        -e POSTGRES_USER="$DB_USER" \
        -p "$DB_PORT:5432" \
        -d postgres:13 >/dev/null 2>&1; then
        print_status 0 "PostgreSQL container started successfully"
        
        # Wait for PostgreSQL to be ready
        print_info "Waiting for PostgreSQL to be ready..."
        sleep 5
        
        # Test connection
        if pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" >/dev/null 2>&1; then
            print_status 0 "PostgreSQL is ready"
            return 0
        else
            print_status 1 "PostgreSQL is not ready yet"
            return 1
        fi
    else
        print_status 1 "Failed to start PostgreSQL container"
        return 1
    fi
}

# Check if PostgreSQL is installed locally
check_postgresql_local() {
    print_info "Checking local PostgreSQL installation..."
    
    if command -v psql >/dev/null 2>&1; then
        print_status 0 "PostgreSQL client found"
        
        # Check if PostgreSQL server is running
        if pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" >/dev/null 2>&1; then
            print_status 0 "PostgreSQL server is running"
            return 0
        else
            print_warning "PostgreSQL server is not running"
            print_warning "Please start PostgreSQL server:"
            print_warning "  Ubuntu/Debian: sudo systemctl start postgresql"
            print_warning "  CentOS/RHEL: sudo systemctl start postgresql"
            print_warning "  macOS: brew services start postgresql"
            return 1
        fi
    else
        print_status 1 "PostgreSQL client not found"
        return 1
    fi
}

# Create database and run migrations
setup_database() {
    print_info "Setting up database and running migrations..."
    
    # Change to project root directory
    cd "$(dirname "$0")/.."
    
    # Create bin directory if it doesn't exist
    mkdir -p bin
    
    # Create migration runner
    cat > bin/migrate.go << 'EOF'
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"scopeapi.local/backend/shared/database/postgresql"
)

func main() {
	var (
		host     = flag.String("host", "localhost", "Database host")
		port     = flag.String("port", "5432", "Database port")
		user     = flag.String("user", "postgres", "Database user")
		password = flag.String("password", "password", "Database password")
		dbname   = flag.String("dbname", "scopeapi", "Database name")
		sslmode  = flag.String("sslmode", "disable", "SSL mode")
		action   = flag.String("action", "migrate", "Action: migrate, rollback, status")
	)
	flag.Parse()

	config := postgresql.Config{
		Host:     *host,
		Port:     *port,
		User:     *user,
		Password: *password,
		DBName:   *dbname,
		SSLMode:  *sslmode,
	}

	migrator, err := postgresql.NewMigrator(config)
	if err != nil {
		log.Fatalf("Failed to create migrator: %v", err)
	}
	defer migrator.Close()

	migrationsDir := "backend/shared/database/postgresql/migrations"

	switch *action {
	case "migrate":
		if err := migrator.Migrate(migrationsDir); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		fmt.Println("Migrations completed successfully")
	case "rollback":
		if err := migrator.Rollback(); err != nil {
			log.Fatalf("Rollback failed: %v", err)
		}
		fmt.Println("Rollback completed successfully")
	case "status":
		if err := migrator.Status(migrationsDir); err != nil {
			log.Fatalf("Status check failed: %v", err)
		}
	default:
		fmt.Printf("Unknown action: %s\n", *action)
		fmt.Println("Available actions: migrate, rollback, status")
		os.Exit(1)
	}
}
EOF

    print_status 0 "Migration runner created"
    
    # Build the migration tool
    print_info "Building migration tool..."
    if go build -o bin/migrate bin/migrate.go; then
        print_status 0 "Migration tool built successfully"
    else
        print_status 1 "Failed to build migration tool"
        return 1
    fi
    
    # Set password for psql
    export PGPASSWORD="$DB_PASSWORD"
    
    # Create database if it doesn't exist
    print_info "Creating database '$DB_NAME' if it doesn't exist..."
    if psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -lqt | cut -d \| -f 1 | grep -qw "$DB_NAME"; then
        print_status 0 "Database '$DB_NAME' already exists"
    else
        if psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -c "CREATE DATABASE $DB_NAME;" >/dev/null 2>&1; then
            print_status 0 "Database '$DB_NAME' created successfully"
        else
            print_status 1 "Failed to create database '$DB_NAME'"
            return 1
        fi
    fi
    
    # Run migrations
    print_info "Applying database migrations..."
    if ./bin/migrate -host="$DB_HOST" -port="$DB_PORT" -user="$DB_USER" -password="$DB_PASSWORD" -dbname="$DB_NAME" -sslmode="$DB_SSL_MODE" -action=migrate; then
        print_status 0 "Database migrations applied successfully"
    else
        print_status 1 "Failed to apply database migrations"
        return 1
    fi
    
    # Check migration status
    print_info "Checking migration status..."
    if ./bin/migrate -host="$DB_HOST" -port="$DB_PORT" -user="$DB_USER" -password="$DB_PASSWORD" -dbname="$DB_NAME" -sslmode="$DB_SSL_MODE" -action=status; then
        print_status 0 "Migration status checked successfully"
    else
        print_status 1 "Failed to check migration status"
        return 1
    fi
}

# Test database connection for all services
test_database_connection() {
    print_info "Testing database connection for all services..."
    
    # Test basic connection
    if psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c "SELECT 1;" >/dev/null 2>&1; then
        print_status 0 "Basic database connection successful"
    else
        print_status 1 "Basic database connection failed"
        return 1
    fi
    
    # Test schema and tables
    print_info "Testing database schema..."
    if psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c "SELECT COUNT(*) FROM scopeapi.endpoints;" >/dev/null 2>&1; then
        print_status 0 "Database schema is valid"
    else
        print_status 1 "Database schema validation failed"
        return 1
    fi
    
    # Test system configuration
    print_info "Testing system configuration..."
    if psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c "SELECT COUNT(*) FROM scopeapi.system_config;" >/dev/null 2>&1; then
        print_status 0 "System configuration table is accessible"
    else
        print_status 1 "System configuration table access failed"
        return 1
    fi
}

# Create test data
create_test_data() {
    print_info "Creating test data..."
    
    # Insert test endpoints
    psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c "
    INSERT INTO scopeapi.endpoints (path, method, parameters, response_structure, risk_score) VALUES
    ('/api/v1/users', 'GET', '{\"query\": {\"page\": \"integer\", \"limit\": \"integer\"}}', '{\"users\": \"array\", \"total\": \"integer\"}', 0.0),
    ('/api/v1/users', 'POST', '{\"body\": {\"name\": \"string\", \"email\": \"string\"}}', '{\"user_id\": \"uuid\", \"created_at\": \"timestamp\"}', 0.0),
    ('/api/v1/users/{id}', 'GET', '{\"path\": {\"id\": \"uuid\"}}', '{\"user\": \"object\"}', 0.0),
    ('/api/v1/auth/login', 'POST', '{\"body\": {\"username\": \"string\", \"password\": \"string\"}}', '{\"token\": \"string\", \"expires\": \"timestamp\"}', 0.0)
    ON CONFLICT (path, method) DO NOTHING;
    " >/dev/null 2>&1
    
    if [ $? -eq 0 ]; then
        print_status 0 "Test data created successfully"
    else
        print_status 1 "Failed to create test data"
        return 1
    fi
}

# Test service database integration
test_service_integration() {
    print_info "Testing service database integration..."
    
    # Test Data Ingestion service
    print_info "Testing Data Ingestion service..."
    cd backend/services/data-ingestion
    if go build -o ../../../bin/data-ingestion cmd/main.go; then
        print_status 0 "Data Ingestion service builds successfully"
    else
        print_status 1 "Data Ingestion service build failed"
    fi
    
    # Test API Discovery service
    print_info "Testing API Discovery service..."
    cd ../api-discovery
    if go build -o ../../../bin/api-discovery cmd/main.go; then
        print_status 0 "API Discovery service builds successfully"
    else
        print_status 1 "API Discovery service build failed"
    fi
    
    # Test Threat Detection service
    print_info "Testing Threat Detection service..."
    cd ../threat-detection
    if go build -o ../../../bin/threat-detection cmd/main.go; then
        print_status 0 "Threat Detection service builds successfully"
    else
        print_status 1 "Threat Detection service build failed"
    fi
    
    cd ../../..
}

# Main execution
main() {
    print_info "Starting database integration setup..."
    
    # Check Docker availability
    if check_docker; then
        # Try to start PostgreSQL with Docker
        if start_postgresql_docker; then
            print_status 0 "PostgreSQL started with Docker"
        else
            print_warning "Failed to start PostgreSQL with Docker, trying local installation"
            if ! check_postgresql_local; then
                print_status 1 "No PostgreSQL available"
                exit 1
            fi
        fi
    else
        # Try local PostgreSQL
        if ! check_postgresql_local; then
            print_status 1 "No PostgreSQL available"
            print_warning "Please install and start PostgreSQL:"
            print_warning "  Ubuntu/Debian: sudo apt-get install postgresql postgresql-contrib"
            print_warning "  CentOS/RHEL: sudo yum install postgresql postgresql-server"
            print_warning "  macOS: brew install postgresql"
            exit 1
        fi
    fi
    
    # Setup database and run migrations
    if ! setup_database; then
        exit 1
    fi
    
    # Test database connection
    if ! test_database_connection; then
        exit 1
    fi
    
    # Create test data
    create_test_data
    
    # Test service integration
    test_service_integration
    
    print_info "Database integration setup completed successfully!"
    print_info "================================================"
    print_info "Database Configuration:"
    print_info "  Host: $DB_HOST"
    print_info "  Port: $DB_PORT"
    print_info "  Database: $DB_NAME"
    print_info "  User: $DB_USER"
    print_info ""
    print_info "You can now start the ScopeAPI services with database integration."
    print_info "Services will automatically connect to the database."
}

# Run main function
main "$@" 