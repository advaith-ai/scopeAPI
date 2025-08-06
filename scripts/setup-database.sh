#!/bin/bash

# ScopeAPI Database Setup Script
# This script sets up PostgreSQL database and runs migrations

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

print_info "ScopeAPI Database Setup"
print_info "======================="

# Check if PostgreSQL is installed
check_postgresql() {
    print_info "Checking PostgreSQL installation..."
    
    if command -v psql >/dev/null 2>&1; then
        print_status 0 "PostgreSQL client found"
        return 0
    else
        print_status 1 "PostgreSQL client not found"
        print_warning "Please install PostgreSQL client:"
        print_warning "  Ubuntu/Debian: sudo apt-get install postgresql-client"
        print_warning "  CentOS/RHEL: sudo yum install postgresql"
        print_warning "  macOS: brew install postgresql"
        return 1
    fi
}

# Check if PostgreSQL server is running
check_postgresql_server() {
    print_info "Checking PostgreSQL server connection..."
    
    if pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" >/dev/null 2>&1; then
        print_status 0 "PostgreSQL server is running and accessible"
        return 0
    else
        print_status 1 "PostgreSQL server is not accessible"
        print_warning "Please ensure PostgreSQL server is running:"
        print_warning "  Ubuntu/Debian: sudo systemctl start postgresql"
        print_warning "  CentOS/RHEL: sudo systemctl start postgresql"
        print_warning "  macOS: brew services start postgresql"
        print_warning "  Or start with Docker: docker run --name postgres -e POSTGRES_PASSWORD=$DB_PASSWORD -p 5432:5432 -d postgres:13"
        return 1
    fi
}

# Create database if it doesn't exist
create_database() {
    print_info "Creating database '$DB_NAME' if it doesn't exist..."
    
    # Set password for psql
    export PGPASSWORD="$DB_PASSWORD"
    
    # Check if database exists
    if psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -lqt | cut -d \| -f 1 | grep -qw "$DB_NAME"; then
        print_status 0 "Database '$DB_NAME' already exists"
    else
        # Create database
        if psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -c "CREATE DATABASE $DB_NAME;" >/dev/null 2>&1; then
            print_status 0 "Database '$DB_NAME' created successfully"
        else
            print_status 1 "Failed to create database '$DB_NAME'"
            return 1
        fi
    fi
}

# Build and run migrations
run_migrations() {
    print_info "Running database migrations..."
    
    # Change to project root directory
    cd "$(dirname "$0")/.."
    
    # Build the migration tool
    print_info "Building migration tool..."
    if go build -o bin/migrate backend/shared/database/postgresql/migrator.go; then
        print_status 0 "Migration tool built successfully"
    else
        print_status 1 "Failed to build migration tool"
        return 1
    fi
    
    # Run migrations
    print_info "Applying database migrations..."
    if ./bin/migrate migrate; then
        print_status 0 "Database migrations applied successfully"
    else
        print_status 1 "Failed to apply database migrations"
        return 1
    fi
}

# Create a simple migration runner
create_migration_runner() {
    print_info "Creating migration runner..."
    
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
}

# Main execution
main() {
    print_info "Starting database setup..."
    
    # Check prerequisites
    if ! check_postgresql; then
        exit 1
    fi
    
    if ! check_postgresql_server; then
        exit 1
    fi
    
    # Create database
    if ! create_database; then
        exit 1
    fi
    
    # Create migration runner
    create_migration_runner
    
    # Run migrations
    if ! run_migrations; then
        exit 1
    fi
    
    print_info "Database setup completed successfully!"
    print_info "You can now start the ScopeAPI services."
}

# Run main function
main "$@" 