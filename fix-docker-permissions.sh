#!/bin/bash

# Fix Docker Permissions Script
# This script helps fix Docker permission issues

set -e

echo "🔧 Fixing Docker Permissions..."
echo "=========================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if user is root
if [ "$EUID" -eq 0 ]; then
    print_error "This script should not be run as root"
    print_warning "Please run as a regular user"
    exit 1
fi

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    print_error "Docker is not installed"
    print_warning "Please install Docker first: https://docs.docker.com/get-docker/"
    exit 1
fi

# Check if user is already in docker group
if groups $USER | grep -q docker; then
    print_success "User is already in docker group"
    print_warning "You may need to log out and log back in, or run: newgrp docker"
else
    print_warning "Adding user to docker group..."
    
    # Add user to docker group
    if sudo usermod -aG docker $USER; then
        print_success "User added to docker group"
        print_warning "You need to log out and log back in for changes to take effect"
        print_warning "Or run: newgrp docker"
    else
        print_error "Failed to add user to docker group"
        exit 1
    fi
fi

# Test Docker access
echo ""
echo "Testing Docker access..."
if docker ps &> /dev/null; then
    print_success "Docker access is working correctly"
else
    print_warning "Docker access still not working"
    print_warning "Try logging out and logging back in, or run: newgrp docker"
    print_warning "If the issue persists, restart your system"
fi

echo ""
echo "🔧 Docker permissions fix completed!"
echo "==========================================" 