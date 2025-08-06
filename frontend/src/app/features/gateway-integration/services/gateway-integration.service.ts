import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../../../environments/environment';

export interface Integration {
  id: string;
  name: string;
  type: GatewayType;
  status: IntegrationStatus;
  config: any;
  credentials?: Credentials;
  endpoints: Endpoint[];
  health?: HealthStatus;
  createdAt: string;
  updatedAt: string;
  lastSync?: string;
}

export interface GatewayType {
  kong: 'kong';
  nginx: 'nginx';
  traefik: 'traefik';
  envoy: 'envoy';
  haproxy: 'haproxy';
}

export interface IntegrationStatus {
  active: 'active';
  inactive: 'inactive';
  error: 'error';
  pending: 'pending';
}

export interface Credentials {
  type: CredentialType;
  username?: string;
  password?: string;
  token?: string;
  apiKey?: string;
  certFile?: string;
  keyFile?: string;
}

export interface CredentialType {
  basic: 'basic';
  token: 'token';
  apiKey: 'api_key';
  tls: 'tls';
}

export interface Endpoint {
  id: string;
  name: string;
  url: string;
  protocol: string;
  port: number;
  headers?: { [key: string]: string };
  timeout: number;
}

export interface HealthStatus {
  status: string;
  message: string;
  lastCheck: string;
  latency: number;
}

export interface SyncResult {
  success: boolean;
  message: string;
  changes: Change[];
  errors: string[];
  timestamp: string;
  duration: number;
}

export interface Change {
  type: string;
  resource: string;
  action: string;
  details: string;
}

@Injectable({
  providedIn: 'root'
})
export class GatewayIntegrationService {
  private apiUrl = `${environment.apiUrl}/gateway-integration/api/v1`;

  constructor(private http: HttpClient) { }

  // Integration Management
  getIntegrations(filters?: any): Observable<Integration[]> {
    const params = filters ? { params: filters } : {};
    return this.http.get<Integration[]>(`${this.apiUrl}/integrations`, params);
  }

  getIntegration(id: string): Observable<Integration> {
    return this.http.get<Integration>(`${this.apiUrl}/integrations/${id}`);
  }

  createIntegration(integration: Partial<Integration>): Observable<Integration> {
    return this.http.post<Integration>(`${this.apiUrl}/integrations`, integration);
  }

  updateIntegration(id: string, integration: Partial<Integration>): Observable<Integration> {
    return this.http.put<Integration>(`${this.apiUrl}/integrations/${id}`, integration);
  }

  deleteIntegration(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/integrations/${id}`);
  }

  testIntegration(id: string): Observable<HealthStatus> {
    return this.http.post<HealthStatus>(`${this.apiUrl}/integrations/${id}/test`, {});
  }

  syncIntegration(id: string): Observable<SyncResult> {
    return this.http.post<SyncResult>(`${this.apiUrl}/integrations/${id}/sync`, {});
  }

  // Configuration Management
  getConfigs(): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/configs`);
  }

  getConfig(id: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/configs/${id}`);
  }

  createConfig(config: any): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/configs`, config);
  }

  updateConfig(id: string, config: any): Observable<any> {
    return this.http.put<any>(`${this.apiUrl}/configs/${id}`, config);
  }

  deleteConfig(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/configs/${id}`);
  }

  validateConfig(id: string): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/configs/${id}/validate`, {});
  }

  deployConfig(id: string): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/configs/${id}/deploy`, {});
  }

  // Gateway Type Helpers
  getGatewayTypes(): GatewayType[] {
    return [
      { kong: 'kong' },
      { nginx: 'nginx' },
      { traefik: 'traefik' },
      { envoy: 'envoy' },
      { haproxy: 'haproxy' }
    ];
  }

  getGatewayTypeDisplayName(type: string): string {
    const displayNames: { [key: string]: string } = {
      'kong': 'Kong',
      'nginx': 'NGINX',
      'traefik': 'Traefik',
      'envoy': 'Envoy',
      'haproxy': 'HAProxy'
    };
    return displayNames[type] || type;
  }

  getGatewayTypeDescription(type: string): string {
    const descriptions: { [key: string]: string } = {
      'kong': 'Cloud-native API gateway and platform',
      'nginx': 'High-performance HTTP server and reverse proxy',
      'traefik': 'Modern HTTP reverse proxy and load balancer',
      'envoy': 'High-performance C++ distributed proxy',
      'haproxy': 'Reliable, high-performance TCP/HTTP load balancer'
    };
    return descriptions[type] || 'API Gateway';
  }

  getGatewayTypeIcon(type: string): string {
    const icons: { [key: string]: string } = {
      'kong': 'cloud',
      'nginx': 'dns',
      'traefik': 'router',
      'envoy': 'hub',
      'haproxy': 'balance'
    };
    return icons[type] || 'settings';
  }

  // Status Helpers
  getStatusDisplayName(status: string): string {
    const displayNames: { [key: string]: string } = {
      'active': 'Active',
      'inactive': 'Inactive',
      'error': 'Error',
      'pending': 'Pending'
    };
    return displayNames[status] || status;
  }

  getStatusColor(status: string): string {
    const colors: { [key: string]: string } = {
      'active': 'success',
      'inactive': 'warning',
      'error': 'danger',
      'pending': 'info'
    };
    return colors[status] || 'secondary';
  }

  // Health Status Helpers
  getHealthStatusColor(status: string): string {
    const colors: { [key: string]: string } = {
      'healthy': 'success',
      'degraded': 'warning',
      'unhealthy': 'danger'
    };
    return colors[status] || 'secondary';
  }

  // Credential Type Helpers
  getCredentialTypes(): CredentialType[] {
    return [
      { basic: 'basic' },
      { token: 'token' },
      { apiKey: 'api_key' },
      { tls: 'tls' }
    ];
  }

  getCredentialTypeDisplayName(type: string): string {
    const displayNames: { [key: string]: string } = {
      'basic': 'Basic Authentication',
      'token': 'Bearer Token',
      'api_key': 'API Key',
      'tls': 'TLS Certificate'
    };
    return displayNames[type] || type;
  }

  // Validation Helpers
  validateIntegration(integration: Partial<Integration>): string[] {
    const errors: string[] = [];

    if (!integration.name || integration.name.trim() === '') {
      errors.push('Integration name is required');
    }

    if (!integration.type) {
      errors.push('Gateway type is required');
    }

    if (!integration.endpoints || integration.endpoints.length === 0) {
      errors.push('At least one endpoint is required');
    }

    return errors;
  }

  validateEndpoint(endpoint: Endpoint): string[] {
    const errors: string[] = [];

    if (!endpoint.name || endpoint.name.trim() === '') {
      errors.push('Endpoint name is required');
    }

    if (!endpoint.url || endpoint.url.trim() === '') {
      errors.push('Endpoint URL is required');
    }

    if (!endpoint.protocol || endpoint.protocol.trim() === '') {
      errors.push('Endpoint protocol is required');
    }

    if (!endpoint.port || endpoint.port <= 0) {
      errors.push('Valid endpoint port is required');
    }

    return errors;
  }

  // Mock Data for Development
  getMockIntegrations(): Integration[] {
    return [
      {
        id: '1',
        name: 'Production Kong Gateway',
        type: { kong: 'kong' },
        status: { active: 'active' },
        config: {
          admin_url: 'http://kong-admin:8001',
          proxy_url: 'http://kong-proxy:8000'
        },
        endpoints: [
          {
            id: '1',
            name: 'Admin API',
            url: 'http://kong-admin:8001',
            protocol: 'http',
            port: 8001,
            timeout: 30000
          }
        ],
        health: {
          status: 'healthy',
          message: 'Kong is running',
          lastCheck: new Date().toISOString(),
          latency: 45
        },
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
        lastSync: new Date().toISOString()
      },
      {
        id: '2',
        name: 'Load Balancer NGINX',
        type: { nginx: 'nginx' },
        status: { active: 'active' },
        config: {
          config_path: '/etc/nginx/nginx.conf',
          reload_command: 'nginx -s reload'
        },
        endpoints: [
          {
            id: '2',
            name: 'HTTP',
            url: 'http://nginx:80',
            protocol: 'http',
            port: 80,
            timeout: 30000
          }
        ],
        health: {
          status: 'healthy',
          message: 'NGINX is running',
          lastCheck: new Date().toISOString(),
          latency: 12
        },
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString()
      }
    ];
  }
} 