import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { GatewayIntegrationService, Integration } from '../../services/gateway-integration.service';

@Component({
  selector: 'app-gateway-integration-overview',
  templateUrl: './gateway-integration-overview.component.html',
  styleUrls: ['./gateway-integration-overview.component.scss']
})
export class GatewayIntegrationOverviewComponent implements OnInit {
  integrations: Integration[] = [];
  loading = false;
  error: string | null = null;

  // Statistics
  totalIntegrations = 0;
  activeIntegrations = 0;
  errorIntegrations = 0;
  pendingIntegrations = 0;

  // Gateway type distribution
  gatewayTypeStats: { [key: string]: number } = {};

  constructor(
    private gatewayIntegrationService: GatewayIntegrationService,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.loadIntegrations();
  }

  loadIntegrations(): void {
    this.loading = true;
    this.error = null;

    // For development, use mock data
    this.integrations = this.gatewayIntegrationService.getMockIntegrations();
    this.calculateStatistics();
    this.loading = false;

    // Uncomment for production
    /*
    this.gatewayIntegrationService.getIntegrations().subscribe({
      next: (integrations) => {
        this.integrations = integrations;
        this.calculateStatistics();
        this.loading = false;
      },
      error: (error) => {
        this.error = 'Failed to load integrations';
        this.loading = false;
        console.error('Error loading integrations:', error);
      }
    });
    */
  }

  calculateStatistics(): void {
    this.totalIntegrations = this.integrations.length;
    this.activeIntegrations = this.integrations.filter(i => i.status.active === 'active').length;
    this.errorIntegrations = this.integrations.filter(i => i.status.error === 'error').length;
    this.pendingIntegrations = this.integrations.filter(i => i.status.pending === 'pending').length;

    // Calculate gateway type distribution
    this.gatewayTypeStats = {};
    this.integrations.forEach(integration => {
      const type = Object.keys(integration.type)[0];
      this.gatewayTypeStats[type] = (this.gatewayTypeStats[type] || 0) + 1;
    });
  }

  onAddIntegration(): void {
    this.router.navigate(['/gateway-integration/integrations/new']);
  }

  onViewIntegration(integration: Integration): void {
    this.router.navigate(['/gateway-integration/integrations', integration.id]);
  }

  onEditIntegration(integration: Integration): void {
    this.router.navigate(['/gateway-integration/integrations', integration.id, 'edit']);
  }

  onDeleteIntegration(integration: Integration): void {
    if (confirm(`Are you sure you want to delete the integration "${integration.name}"?`)) {
      this.gatewayIntegrationService.deleteIntegration(integration.id).subscribe({
        next: () => {
          this.loadIntegrations();
        },
        error: (error) => {
          this.error = 'Failed to delete integration';
          console.error('Error deleting integration:', error);
        }
      });
    }
  }

  onTestIntegration(integration: Integration): void {
    this.gatewayIntegrationService.testIntegration(integration.id).subscribe({
      next: (health) => {
        // Update the integration's health status
        const index = this.integrations.findIndex(i => i.id === integration.id);
        if (index !== -1) {
          this.integrations[index].health = health;
        }
      },
      error: (error) => {
        this.error = 'Failed to test integration';
        console.error('Error testing integration:', error);
      }
    });
  }

  onSyncIntegration(integration: Integration): void {
    this.gatewayIntegrationService.syncIntegration(integration.id).subscribe({
      next: (result) => {
        if (result.success) {
          // Update last sync time
          const index = this.integrations.findIndex(i => i.id === integration.id);
          if (index !== -1) {
            this.integrations[index].lastSync = result.timestamp;
          }
        } else {
          this.error = `Sync failed: ${result.message}`;
        }
      },
      error: (error) => {
        this.error = 'Failed to sync integration';
        console.error('Error syncing integration:', error);
      }
    });
  }

  getGatewayTypeDisplayName(type: any): string {
    const typeKey = Object.keys(type)[0];
    return this.gatewayIntegrationService.getGatewayTypeDisplayName(typeKey);
  }

  getGatewayTypeIcon(type: any): string {
    const typeKey = Object.keys(type)[0];
    return this.gatewayIntegrationService.getGatewayTypeIcon(typeKey);
  }

  getStatusDisplayName(status: any): string {
    const statusKey = Object.keys(status)[0];
    return this.gatewayIntegrationService.getStatusDisplayName(statusKey);
  }

  getStatusColor(status: any): string {
    const statusKey = Object.keys(status)[0];
    return this.gatewayIntegrationService.getStatusColor(statusKey);
  }

  getHealthStatusColor(health?: any): string {
    if (!health) return 'secondary';
    return this.gatewayIntegrationService.getHealthStatusColor(health.status);
  }

  formatDate(dateString: string): string {
    return new Date(dateString).toLocaleString();
  }

  formatLatency(latency: number): string {
    return `${latency}ms`;
  }

  getGatewayTypeRoute(type: any): string {
    const typeKey = Object.keys(type)[0];
    return `/gateway-integration/${typeKey}`;
  }

  onGatewayTypeClick(type: any): void {
    const route = this.getGatewayTypeRoute(type);
    this.router.navigate([route]);
  }

  clearError(): void {
    this.error = null;
  }
} 