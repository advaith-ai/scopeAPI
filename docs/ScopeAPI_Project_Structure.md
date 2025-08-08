# ScopeAPI Project Structure

This document provides a comprehensive overview of the ScopeAPI project structure, including all directories, files, and their purposes.

## Root Directory Structure

```
scopeAPI/
├── LICENSE                          # MIT License
├── README.md                        # Main project overview and quick start
├── README-RUN.md                    # Detailed manual startup instructions
├── docker-compose.yml               # Multi-service Docker environment
├── .gitignore                       # Git ignore rules
├── go.mod                           # Go workspace module
├── go.work                          # Go workspace configuration
├── go.work.sum                      # Go workspace checksums
├── scopeAPI.code-workspace          # VS Code workspace configuration
├── start-scopeapi.sh                # Automated startup script
├── scripts/                         # Utility scripts
│   ├── setup-database-integration.sh
│   └── setup-database.sh
├── backend/                         # Backend microservices
├── adminConsole/                    # Angular admin console application
└── docs/                           # Project documentation
```

## Backend Structure

```
backend/
├── bin/                            # Compiled binaries
│   ├── api-discovery
│   └── data-ingestion
├── config/                         # Configuration files
│   ├── api-discovery.yaml
│   ├── data-ingestion.yaml
│   └── threat-detection.yaml
├── go.mod                          # Go module file
├── go.sum                          # Go dependencies checksums
├── Makefile                        # Build and deployment automation
├── services/                       # Microservices
│   ├── api-discovery/              # API Discovery Service
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── internal/
│   │       ├── handlers/
│   │       │   ├── discovery_handler.go
│   │       │   ├── endpoint_handler.go
│   │       │   ├── handlers.go
│   │       │   └── inventory_handler.go
│   │       ├── models/
│   │       │   ├── api_spec.go
│   │       │   ├── endpoint.go
│   │       │   └── metadata.go
│   │       ├── repository/
│   │       │   ├── discovery_repository.go
│   │       │   ├── inventory_repository.go
│   │       │   └── repository.go
│   │       └── services/
│   │           ├── discovery_service.go
│   │           ├── inventory_service.go
│   │           ├── metadata_service.go
│   │           └── services.go
│   ├── attack-blocking/             # Attack Blocking Service
│   │   ├── internal/
│   │   │   ├── handlers/
│   │   │   │   ├── agent_handler.go
│   │   │   │   ├── blocking_handler.go
│   │   │   │   └── policy_handler.go
│   │   │   ├── models/
│   │   │   │   ├── agent.go
│   │   │   │   ├── blocking_rule.go
│   │   │   │   └── policy.go
│   │   │   └── repository/
│   │   │       ├── blocking_repository.go
│   │   │       └── policy_repository.go
│   │   ├── models/
│   │   │   ├── agent.go
│   │   │   ├── blocking_rule-README.md
│   │   │   ├── blocking_rule.go
│   │   │   ├── policy-README.md
│   │   │   └── policy.go
│   │   └── services/
│   │       ├── agent_management_service-README.md
│   │       ├── agent_management_service.go
│   │       ├── attack_blocking_service.go
│   │       ├── cloud_intelligence_service-README.md
│   │       ├── cloud_intelligence_service.go
│   │       └── policy_enforcement_service.go
│   ├── data-ingestion/              # Data Ingestion Service
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── config/
│   │   │   └── data-ingestion.yaml
│   │   ├── data-ingestion
│   │   ├── go.mod
│   │   └── internal/
│   │       ├── config/
│   │       │   └── config.go
│   │       ├── handlers/
│   │       │   └── handlers.go
│   │       ├── models/
│   │       │   └── traffic.go
│   │       └── services/
│   │           ├── data_ingestion_service.go
│   │           ├── data_normalizer_service.go
│   │           ├── data_parser_service.go
│   │           └── queue_service.go
│   ├── data-protection/             # Data Protection Service
│   │   └── internal/
│   │       ├── models/
│   │       │   ├── classification_rule.go
│   │       │   ├── compliance_report.go
│   │       │   └── pii_data.go
│   │       ├── repository/
│   │       │   └── repository.go
│   │       └── services/
│   │           ├── compliance_service.go
│   │           ├── data_classification_service.go
│   │           ├── pii_detection_service.go
│   │           └── risk_scoring_service.go
│   ├── gateway-integration/         # Gateway Integration Service
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── README.md
│   │   └── internal/
│   │       ├── handlers/
│   │       │   ├── integration_handler.go
│   │       │   └── integration_handler_test.go
│   │       ├── models/
│   │       │   └── integration.go
│   │       ├── repository/
│   │       │   ├── integration_repository.go
│   │       │   └── integration_repository_test.go
│   │       ├── services/
│   │       │   ├── integration_service.go
│   │       │   ├── integration_service_test.go
│   │       │   ├── kong_integration_service.go
│   │       │   ├── kong_integration_service_test.go
│   │       │   ├── nginx_integration_service.go
│   │       │   ├── traefik_integration_service.go
│   │       │   ├── envoy_integration_service.go
│   │       │   └── haproxy_integration_service.go
│   │       └── integration_test.go
│   └── threat-detection/            # Threat Detection Service
│       ├── cmd/
│       │   └── main.go
│       ├── go.mod
│       ├── go.sum
│       └── internal/
│           ├── handlers/
│           │   └── handlers.go
│           ├── models/
│           │   ├── anomaly.go
│           │   ├── behavior_pattern.go
│           │   ├── signature.go
│           │   └── threat.go
│           ├── repository/
│           │   └── repository.go
│           └── services/
│               ├── anomaly_detection_service.go
│               ├── behavioral_analysis_service.go
│               ├── signature_detection_service.go
│               └── threat_detection_service.go
└── shared/                         # Shared libraries and utilities
    ├── auth/
    │   └── jwt/
    │       └── jwt.go
    ├── database/
    │   └── postgresql/
    │       ├── migrations/
    │       │   ├── 001_initial_schema.down.sql
    │       │   ├── 001_initial_schema.up.sql
    │       │   ├── 002_api_discovery.down.sql
    │       │   ├── 002_api_discovery.up.sql
    │       │   ├── 003_gateway_integration.down.sql
    │       │   └── 003_gateway_integration.up.sql
    │       ├── migrator.go
    │       └── postgresql.go
    ├── go.mod
    ├── go.sum
    ├── logging/
    │   └── logging.go
    ├── messaging/
    │   └── kafka/
    │       └── kafka.go
    ├── monitoring/
    │   ├── health/
    │   │   └── health.go
    │   └── metrics/
    │       └── metrics.go
    └── utils/
        └── config/
            └── config.go
```

## Admin Console Structure

```
adminConsole/
├── angular.json                     # Angular CLI configuration
├── package.json                     # Node.js dependencies
├── package-lock.json                # Locked dependency versions
├── README.md                        # Admin Console-specific documentation
├── tsconfig.json                    # TypeScript configuration
├── tsconfig.app.json                # App-specific TypeScript config
├── tsconfig.spec.json               # Test TypeScript configuration
├── .editorconfig                    # Editor configuration
├── .gitignore                       # Admin Console-specific git ignore
└── src/
    ├── app/
    │   ├── app-routing.module.ts    # Main routing configuration
    │   ├── app.component.html       # Root component template
    │   ├── app.component.scss       # Root component styles
    │   ├── app.component.spec.ts    # Root component tests
    │   ├── app.component.ts         # Root component logic
    │   ├── app.module.ts            # Main application module
    │   ├── core/                    # Core functionality
    │   │   ├── core-routing.module.ts
    │   │   ├── core.module.ts
    │   │   ├── guards/
    │   │   │   ├── auth.guard.spec.ts
    │   │   │   └── auth.guard.ts
    │   │   ├── interceptors/
    │   │   │   ├── auth.interceptor.spec.ts
    │   │   │   ├── auth.interceptor.ts
    │   │   │   ├── error.interceptor.spec.ts
    │   │   │   └── error.interceptor.ts
    │   │   ├── models/
    │   │   │   ├── api-endpoint.model.ts
    │   │   │   ├── threat.model.ts
    │   │   │   └── user.model.ts
    │   │   └── services/
    │   │       ├── api.service.spec.ts
    │   │       ├── api.service.ts
    │   │       ├── auth.service.spec.ts
    │   │       ├── auth.service.ts
    │   │       ├── notification.service.spec.ts
    │   │       ├── notification.service.ts
    │   │       ├── threat.service.spec.ts
    │   │       ├── threat.service.ts
    │   │       ├── websocket.service.spec.ts
    │   │       └── websocket.service.ts
    │   ├── features/                # Feature modules
    │   │   ├── api-discovery/       # API Discovery feature
    │   │   │   ├── api-discovery-routing.module.ts
    │   │   │   ├── api-discovery.module.ts
    │   │   │   └── components/
    │   │   │       └── api-discovery-overview/
    │   │   │           ├── api-discovery-overview.component.html
    │   │   │           ├── api-discovery-overview.component.scss
    │   │   │           ├── api-discovery-overview.component.spec.ts
    │   │   │           └── api-discovery-overview.component.ts
    │   │   ├── attack-protection/   # Attack Protection feature
    │   │   │   ├── attack-protection-routing.module.ts
    │   │   │   ├── attack-protection.module.ts
    │   │   │   └── components/
    │   │   │       └── attack-protection-overview/
    │   │   │           ├── attack-protection-overview.component.html
    │   │   │           ├── attack-protection-overview.component.scss
    │   │   │           ├── attack-protection-overview.component.spec.ts
    │   │   │           └── attack-protection-overview.component.ts
    │   │   ├── auth/                # Authentication feature
    │   │   │   ├── auth-routing.module.ts
    │   │   │   ├── auth.module.ts
    │   │   │   └── components/
    │   │   │       ├── access-policies/
    │   │   │       │   ├── access-policies.component.html
    │   │   │       │   ├── access-policies.component.scss
    │   │   │       │   ├── access-policies.component.spec.ts
    │   │   │       │   └── access-policies.component.ts
    │   │   │       ├── login/
    │   │   │       │   ├── login.component.html
    │   │   │       │   ├── login.component.scss
    │   │   │       │   ├── login.component.spec.ts
    │   │   │       │   └── login.component.ts
    │   │   │       └── user-management/
    │   │   │           ├── user-management.component.html
    │   │   │           ├── user-management.component.scss
    │   │   │           ├── user-management.component.spec.ts
    │   │   │           └── user-management.component.ts
    │   │   ├── dashboard/           # Dashboard feature
    │   │   │   ├── components/
    │   │   │   │   └── dashboard-overview/
    │   │   │   │       ├── dashboard-overview.component.html
    │   │   │   │       ├── dashboard-overview.component.scss
    │   │   │   │       ├── dashboard-overview.component.spec.ts
    │   │   │   │       └── dashboard-overview.component.ts
    │   │   │   ├── dashboard-routing.module.ts
    │   │   │   └── dashboard.module.ts
    │   │   ├── data-protection/     # Data Protection feature
    │   │   │   ├── components/
    │   │   │   │   └── data-protection-overview/
    │   │   │   │       ├── data-protection-overview.component.html
    │   │   │   │       ├── data-protection-overview.component.scss
    │   │   │   │       ├── data-protection-overview.component.spec.ts
    │   │   │   │       └── data-protection-overview.component.ts
    │   │   │   ├── data-protection-routing.module.ts
    │   │   │   └── data-protection.module.ts
    │   │   ├── gateway-integration/ # Gateway Integration feature
    │   │   │   ├── components/
    │   │   │   │   ├── gateway-integration-overview/
    │   │   │   │   │   ├── gateway-integration-overview.component.html
    │   │   │   │   │   ├── gateway-integration-overview.component.scss
    │   │   │   │   │   ├── gateway-integration-overview.component.spec.ts
    │   │   │   │   │   └── gateway-integration-overview.component.ts
    │   │   │   │   ├── integration-list/
    │   │   │   │   │   ├── integration-list.component.html
    │   │   │   │   │   ├── integration-list.component.scss
    │   │   │   │   │   ├── integration-list.component.spec.ts
    │   │   │   │   │   └── integration-list.component.ts
    │   │   │   │   ├── integration-form/
    │   │   │   │   │   ├── integration-form.component.html
    │   │   │   │   │   ├── integration-form.component.scss
    │   │   │   │   │   ├── integration-form.component.spec.ts
    │   │   │   │   │   └── integration-form.component.ts
    │   │   │   │   ├── integration-details/
    │   │   │   │   │   ├── integration-details.component.html
    │   │   │   │   │   ├── integration-details.component.scss
    │   │   │   │   │   ├── integration-details.component.spec.ts
    │   │   │   │   │   └── integration-details.component.ts
    │   │   │   │   ├── kong-integration/
    │   │   │   │   │   ├── kong-integration.component.html
    │   │   │   │   │   ├── kong-integration.component.scss
    │   │   │   │   │   ├── kong-integration.component.spec.ts
    │   │   │   │   │   └── kong-integration.component.ts
    │   │   │   │   ├── nginx-integration/
    │   │   │   │   │   ├── nginx-integration.component.html
    │   │   │   │   │   ├── nginx-integration.component.scss
    │   │   │   │   │   ├── nginx-integration.component.spec.ts
    │   │   │   │   │   └── nginx-integration.component.ts
    │   │   │   │   ├── traefik-integration/
    │   │   │   │   │   ├── traefik-integration.component.html
    │   │   │   │   │   ├── traefik-integration.component.scss
    │   │   │   │   │   ├── traefik-integration.component.spec.ts
    │   │   │   │   │   └── traefik-integration.component.ts
    │   │   │   │   ├── envoy-integration/
    │   │   │   │   │   ├── envoy-integration.component.html
    │   │   │   │   │   ├── envoy-integration.component.scss
    │   │   │   │   │   ├── envoy-integration.component.spec.ts
    │   │   │   │   │   └── envoy-integration.component.ts
    │   │   │   │   └── haproxy-integration/
    │   │   │   │       ├── haproxy-integration.component.html
    │   │   │   │       ├── haproxy-integration.component.scss
    │   │   │   │       ├── haproxy-integration.component.spec.ts
    │   │   │   │       └── haproxy-integration.component.ts
    │   │   │   ├── services/
    │   │   │   │   └── gateway-integration.service.ts
    │   │   │   ├── gateway-integration-routing.module.ts
    │   │   │   └── gateway-integration.module.ts
    │   │   └── threat-detection/    # Threat Detection feature
    │   │       ├── components/
    │   │       │   └── threat-detection-overview/
    │   │       │       ├── threat-detection-overview.component.html
    │   │       │       ├── threat-detection-overview.component.scss
    │   │       │       ├── threat-detection-overview.component.spec.ts
    │   │       │       └── threat-detection-overview.component.ts
    │   │       ├── threat-detection-routing.module.ts
    │   │       └── threat-detection.module.ts
    │   └── shared/                  # Shared components
    │       ├── components/
    │       │   ├── header/
    │       │   │   ├── header.component.html
    │       │   │   ├── header.component.scss
    │       │   │   ├── header.component.spec.ts
    │       │   │   └── header.component.ts
    │       │   ├── loading-spinner/
    │       │   │   ├── loading-spinner.component.html
    │       │   │   ├── loading-spinner.component.scss
    │       │   │   ├── loading-spinner.component.spec.ts
    │       │   │   └── loading-spinner.component.ts
    │       │   └── sidebar/
    │       │       ├── sidebar.component.html
    │       │       ├── sidebar.component.scss
    │       │       ├── sidebar.component.spec.ts
    │       │       └── sidebar.component.ts
    │       └── shared.module.ts
    ├── assets/                      # Static assets
    ├── favicon.ico                  # Application icon
    ├── index.html                   # Main HTML template
    ├── main.ts                      # Application entry point
    └── styles.scss                  # Global styles
```

## Documentation Files

```
docs/
├── ScopeAPI_Product_Overview.md     # Product overview and features
├── ScopeAPI_Technical_Architecture.md # Technical architecture details
└── ScopeAPI_Project_Structure.md   # This file
```

## Key Configuration Files

### Docker Compose
The `docker-compose.yml` file defines the complete development environment including:
- PostgreSQL, MongoDB, Neo4j, Redis databases
- Kafka and Zookeeper for messaging
- All 5 backend microservices
- NGINX API Gateway
- Admin Console Angular application
- Prometheus and Grafana for monitoring

### Go Workspace
The `go.work` file manages the Go workspace for the backend microservices, allowing them to share dependencies and be built together.

### Angular Configuration
The `angular.json` file configures the Angular CLI build process, including:
- Build targets for development and production
- Asset management
- Testing configuration
- Linting rules

## Development Guidelines

### Backend Development
1. Each microservice follows a clean architecture pattern with:
   - `cmd/` - Application entry points
   - `internal/` - Private application code
   - `models/` - Data structures
   - `services/` - Business logic
   - `repository/` - Data access layer
   - `handlers/` - HTTP request handlers

2. Shared utilities are in the `shared/` directory and imported by all services

3. Database migrations are managed through the shared PostgreSQL package

### Admin Console Development
1. Feature-based architecture with lazy-loaded modules
2. Core functionality in the `core/` directory
3. Shared components in the `shared/` directory
4. Each feature has its own routing and module configuration

### Testing
- Backend: Unit tests with Go's testing package and testify
- Admin Console: Unit tests with Jasmine and Karma
- Integration tests for API endpoints
- End-to-end tests for critical user flows

### Deployment
- Docker containers for all services
- Kubernetes manifests for production deployment
- CI/CD pipelines for automated testing and deployment
- Monitoring and logging infrastructure

## Getting Started

1. Clone the repository
2. Run `docker-compose up` to start the development environment
3. Access the admin console at `http://localhost:4200`
4. API documentation available at `http://localhost:8080/docs`

For detailed setup instructions, see `README-RUN.md`. 