# Gateway Integration Service

## Overview

The Gateway Integration Service is a core component of the ScopeAPI platform that provides centralized management and monitoring capabilities for multiple API gateways. It enables organizations to manage Kong, NGINX, Traefik, Envoy, and HAProxy gateways from a unified interface.

## Architecture Integration

This service is part of the **Core Services Layer** in the ScopeAPI architecture:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                        Core Services Layer                                  │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐             │
│  │  Endpoint       │  │  Threat         │  │  Attack         │             │
│  │  Discovery      │  │  Detection      │  │  Blocking       │             │
│  │  Service        │  │  Engine         │  │  Engine         │             │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘             │
│                                                                             │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐             │
│  │  Sensitive      │  │  Security       │  │  Gateway        │             │
│  │  Data Scanner   │  │  Testing        │  │  Integration    │             │
│  │                 │  │  Engine         │  │  Service        │             │
│  │ • PII detection │  │ • Automated     │  │                 │             │
│  │ • Data classify │  │ • Vuln scanning │  │ • Kong/NGINX    │             │
│  │ • Compliance    │  │ • Pen testing   │  │ • Traefik/Envoy │             │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘             │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### **Integration Points**

- **API Gateway Layer**: Manages Kong, NGINX, Traefik, Envoy, HAProxy
- **Data Storage Layer**: Uses PostgreSQL for integration metadata
- **Message Queue**: Publishes events to Kafka for real-time updates
- **Frontend**: Provides Angular components for gateway management UI

## Features

### 🔗 **Multi-Gateway Support**
- **Kong**: Cloud-native API gateway and platform
- **NGINX**: High-performance HTTP server and reverse proxy
- **Traefik**: Modern HTTP reverse proxy and load balancer
- **Envoy**: High-performance C++ distributed proxy
- **HAProxy**: Reliable, high-performance TCP/HTTP load balancer

### 🛠️ **Core Functionality**
- **Integration Management**: Create, update, and delete gateway integrations
- **Configuration Sync**: Synchronize configurations across gateways
- **Health Monitoring**: Real-time health checks and status monitoring
- **Credential Management**: Secure handling of authentication credentials
- **Event Processing**: Process gateway events and security events
- **RESTful API**: Complete REST API for programmatic access

### 🔒 **Security Features**
- **Credential Encryption**: Secure storage of authentication credentials
- **JWT Authentication**: Protected API endpoints
- **Role-based Access**: Integration with ScopeAPI's authentication system
- **Audit Logging**: Comprehensive logging of all operations

## Architecture

### Service Structure
```
gateway-integration/
├── cmd/
│   └── main.go                 # Service entry point
├── internal/
│   ├── handlers/               # HTTP request handlers
│   ├── models/                 # Data models and structures
│   ├── repository/             # Database operations
│   └── services/               # Business logic
├── go.mod                      # Go module dependencies
└── README.md                   # This file
```

### Key Components

#### 1. **Integration Service**
- Manages gateway integration lifecycle
- Handles validation and testing
- Processes events and notifications

#### 2. **Gateway Clients**
- **KongClient**: Kong-specific operations
- **NginxClient**: NGINX configuration management
- **TraefikClient**: Traefik middleware and routing
- **EnvoyClient**: Envoy cluster and listener management
- **HAProxyClient**: HAProxy frontend/backend configuration

#### 3. **Configuration Management**
- Version-controlled configuration storage
- Configuration validation and deployment
- Rollback capabilities

## API Endpoints

### Integration Management
```
GET    /api/v1/integrations           # List all integrations
GET    /api/v1/integrations/:id       # Get integration details
POST   /api/v1/integrations           # Create new integration
PUT    /api/v1/integrations/:id       # Update integration
DELETE /api/v1/integrations/:id       # Delete integration
POST   /api/v1/integrations/:id/test  # Test integration connection
POST   /api/v1/integrations/:id/sync  # Sync configuration
```

### Configuration Management
```
GET    /api/v1/configs               # List configurations
GET    /api/v1/configs/:id           # Get configuration details
POST   /api/v1/configs               # Create configuration
PUT    /api/v1/configs/:id           # Update configuration
DELETE /api/v1/configs/:id           # Delete configuration
POST   /api/v1/configs/:id/validate  # Validate configuration
POST   /api/v1/configs/:id/deploy    # Deploy configuration
```

### Gateway-Specific Endpoints

#### Kong
```
GET    /api/v1/kong/status           # Get Kong status
GET    /api/v1/kong/services         # List Kong services
GET    /api/v1/kong/routes           # List Kong routes
GET    /api/v1/kong/plugins          # List Kong plugins
POST   /api/v1/kong/plugins          # Create Kong plugin
PUT    /api/v1/kong/plugins/:id      # Update Kong plugin
DELETE /api/v1/kong/plugins/:id      # Delete Kong plugin
POST   /api/v1/kong/sync             # Sync Kong configuration
```

#### NGINX
```
GET    /api/v1/nginx/status          # Get NGINX status
GET    /api/v1/nginx/config          # Get NGINX configuration
POST   /api/v1/nginx/config          # Update NGINX configuration
POST   /api/v1/nginx/reload          # Reload NGINX configuration
GET    /api/v1/nginx/upstreams       # List NGINX upstreams
POST   /api/v1/nginx/upstreams       # Update NGINX upstream
POST   /api/v1/nginx/sync            # Sync NGINX configuration
```

#### Traefik
```
GET    /api/v1/traefik/status        # Get Traefik status
GET    /api/v1/traefik/providers     # List Traefik providers
GET    /api/v1/traefik/middlewares   # List Traefik middlewares
POST   /api/v1/traefik/middlewares   # Create Traefik middleware
PUT    /api/v1/traefik/middlewares/:id # Update Traefik middleware
DELETE /api/v1/traefik/middlewares/:id # Delete Traefik middleware
POST   /api/v1/traefik/sync          # Sync Traefik configuration
```

#### Envoy
```
GET    /api/v1/envoy/status          # Get Envoy status
GET    /api/v1/envoy/clusters        # List Envoy clusters
GET    /api/v1/envoy/listeners       # List Envoy listeners
GET    /api/v1/envoy/filters         # List Envoy filters
POST   /api/v1/envoy/filters         # Create Envoy filter
PUT    /api/v1/envoy/filters/:id     # Update Envoy filter
DELETE /api/v1/envoy/filters/:id     # Delete Envoy filter
POST   /api/v1/envoy/sync            # Sync Envoy configuration
```

#### HAProxy
```
GET    /api/v1/haproxy/status        # Get HAProxy status
GET    /api/v1/haproxy/config        # Get HAProxy configuration
POST   /api/v1/haproxy/config        # Update HAProxy configuration
POST   /api/v1/haproxy/reload        # Reload HAProxy configuration
GET    /api/v1/haproxy/backends      # List HAProxy backends
POST   /api/v1/haproxy/backends      # Update HAProxy backend
POST   /api/v1/haproxy/sync          # Sync HAProxy configuration
```

## Data Models

### Integration
```go
type Integration struct {
    ID          string                 `json:"id"`
    Name        string                 `json:"name"`
    Type        GatewayType            `json:"type"`
    Status      IntegrationStatus      `json:"status"`
    Config      map[string]interface{} `json:"config"`
    Credentials *Credentials           `json:"credentials,omitempty"`
    Endpoints   []Endpoint             `json:"endpoints"`
    Health      *HealthStatus          `json:"health"`
    CreatedAt   time.Time              `json:"created_at"`
    UpdatedAt   time.Time              `json:"updated_at"`
    LastSync    *time.Time             `json:"last_sync"`
}
```

### Gateway Types
```go
type GatewayType string

const (
    GatewayTypeKong    GatewayType = "kong"
    GatewayTypeNginx   GatewayType = "nginx"
    GatewayTypeTraefik GatewayType = "traefik"
    GatewayTypeEnvoy   GatewayType = "envoy"
    GatewayTypeHAProxy GatewayType = "haproxy"
)
```

### Credentials
```go
type Credentials struct {
    Type     CredentialType `json:"type"`
    Username string         `json:"username,omitempty"`
    Password string         `json:"password,omitempty"`
    Token    string         `json:"token,omitempty"`
    APIKey   string         `json:"api_key,omitempty"`
    CertFile string         `json:"cert_file,omitempty"`
    KeyFile  string         `json:"key_file,omitempty"`
}
```

## Configuration

### Environment Variables
```bash
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=scopeapi_user
DB_PASSWORD=scopeapi_password
DB_NAME=scopeapi

# Kafka
KAFKA_BROKERS=localhost:9092
KAFKA_TOPIC_PREFIX=scopeapi

# Server
SERVER_PORT=8086
SERVER_HOST=0.0.0.0

# Authentication
JWT_SECRET=your-jwt-secret

# Logging
LOG_LEVEL=info
LOG_FORMAT=json
```

### Configuration File
```yaml
server:
  port: "8086"
  host: "0.0.0.0"
  read_timeout: 30s
  write_timeout: 30s
  idle_timeout: 60s

database:
  postgresql:
    host: "localhost"
    port: 5432
    user: "scopeapi_user"
    password: "scopeapi_password"
    database: "scopeapi"
    ssl_mode: "disable"
    max_conns: 10

messaging:
  kafka:
    brokers: ["localhost:9092"]
    topic_prefix: "scopeapi"

auth:
  jwt_secret: "your-jwt-secret"

logging:
  level: "info"
  format: "json"
```

## Usage Examples

### Creating a Kong Integration
```bash
curl -X POST http://localhost:8086/api/v1/integrations \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Production Kong Gateway",
    "type": "kong",
    "config": {
      "admin_url": "http://kong-admin:8001",
      "proxy_url": "http://kong-proxy:8000"
    },
    "credentials": {
      "type": "basic",
      "username": "admin",
      "password": "password"
    },
    "endpoints": [
      {
        "name": "Admin API",
        "url": "http://kong-admin:8001",
        "protocol": "http",
        "port": 8001,
        "timeout": 30000
      }
    ]
  }'
```

### Testing an Integration
```bash
curl -X POST http://localhost:8086/api/v1/integrations/1/test \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Syncing Configuration
```bash
curl -X POST http://localhost:8086/api/v1/integrations/1/sync \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Development

### Prerequisites
- Go 1.22+
- PostgreSQL 14+
- Apache Kafka
- Docker (optional)

### Building
```bash
cd backend/services/gateway-integration
go mod tidy
go build -o gateway-integration cmd/main.go
```

### Running
```bash
# Set environment variables
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=scopeapi_user
export DB_PASSWORD=scopeapi_password
export DB_NAME=scopeapi
export KAFKA_BROKERS=localhost:9092

# Run the service
./gateway-integration
```

### Testing
```bash
# Run unit tests
go test ./...

# Run integration tests
go test ./... -tags=integration

# Run with coverage
go test ./... -cover
```

## Docker

### Building Docker Image
```bash
docker build -t scopeapi-gateway-integration .
```

### Running with Docker
```bash
docker run -d \
  --name gateway-integration \
  -p 8086:8086 \
  -e DB_HOST=postgres \
  -e DB_PORT=5432 \
  -e DB_USER=scopeapi_user \
  -e DB_PASSWORD=scopeapi_password \
  -e DB_NAME=scopeapi \
  -e KAFKA_BROKERS=kafka:9092 \
  scopeapi-gateway-integration
```

## Monitoring

### Health Check
```bash
curl http://localhost:8086/health
```

### Metrics
```bash
curl http://localhost:8086/metrics
```

### Logs
The service uses structured logging with the following levels:
- `DEBUG`: Detailed debugging information
- `INFO`: General operational information
- `WARN`: Warning messages
- `ERROR`: Error conditions

## Security Considerations

1. **Credential Storage**: Credentials are encrypted at rest and never returned in API responses
2. **API Security**: All endpoints require JWT authentication
3. **Input Validation**: All inputs are validated and sanitized
4. **Rate Limiting**: Implement rate limiting for API endpoints
5. **Audit Logging**: All operations are logged for audit purposes

## Troubleshooting

### Common Issues

1. **Connection Failed**
   - Check gateway endpoint accessibility
   - Verify credentials
   - Check network connectivity

2. **Configuration Sync Failed**
   - Validate gateway configuration
   - Check gateway permissions
   - Review gateway logs

3. **Health Check Failed**
   - Verify gateway is running
   - Check gateway configuration
   - Review service logs

### Debug Mode
Enable debug logging:
```bash
export LOG_LEVEL=debug
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 