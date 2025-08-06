package repository

import (
	"context"
	"time"
	"scopeapi.local/backend/services/threat-detection/internal/models"
)

type AnomalyRepositoryInterface interface {
	// GetRecentAnomalies returns recent anomalies for an entity within a time window
	GetRecentAnomalies(ctx context.Context, entityID string, entityType string, since time.Time) ([]models.Anomaly, error)
	// SaveAnomaly saves a detected anomaly
	SaveAnomaly(ctx context.Context, anomaly *models.Anomaly) error
	// CreateAnomaly creates a new anomaly record
	CreateAnomaly(ctx context.Context, anomaly *models.Anomaly) error
	// GetBaselineStatistics returns baseline statistics for an entity
	GetBaselineStatistics(ctx context.Context, entityID string, entityType string) (map[string]interface{}, error)
	// GetRecentRequestCount returns the recent request count for an entity
	GetRecentRequestCount(ctx context.Context, entityID string, entityType string, duration time.Duration) (int, error)
	// GetBaselineRequestCount returns the baseline request count for an entity
	GetBaselineRequestCount(ctx context.Context, entityID string, entityType string) (int, error)
	// GetBaselineResponseTime returns the baseline response time for an entity
	GetBaselineResponseTime(ctx context.Context, entityID string, entityType string) (float64, error)
	// GetHistoricalCountries returns a list of historical countries for an entity
	GetHistoricalCountries(ctx context.Context, entityID string, entityType string) ([]string, error)
	// GetAnomalies returns a list of anomalies matching a filter
	GetAnomalies(ctx context.Context, filter *models.AnomalyFilter) ([]models.Anomaly, error)
	// GetAnomaly returns a single anomaly by ID
	GetAnomaly(ctx context.Context, anomalyID string) (*models.Anomaly, error)
	// UpdateAnomalyFeedback updates feedback for an anomaly
	UpdateAnomalyFeedback(ctx context.Context, feedback *models.AnomalyFeedback) error
	// GetAnomalyStatistics returns statistics for anomalies
	GetAnomalyStatistics(ctx context.Context, filter *models.AnomalyFilter) (*models.AnomalyStatistics, error)
	// StoreBaselineStatistics stores baseline statistics for an entity
	StoreBaselineStatistics(ctx context.Context, entityID string, baseline map[string]interface{}) error
	// GetModelPerformance returns model performance metrics
	GetModelPerformance(ctx context.Context, modelVersion string) (*models.ModelPerformanceMetric, error)
}

type ThreatRepositoryInterface interface {
	// GetThreatByID fetches a threat by its ID
	GetThreatByID(ctx context.Context, threatID string) (*models.Threat, error)
	// SaveThreat saves a threat event
	SaveThreat(ctx context.Context, threat *models.Threat) error
	// ListThreats returns a list of threats matching a filter
	ListThreats(ctx context.Context, filter *models.ThreatFilter) ([]models.Threat, error)

	// Signature management methods
	GetThreatSignatures(ctx context.Context, filter *models.SignatureFilter) ([]models.ThreatSignature, error)
	UpdateThreatSignature(ctx context.Context, id string, signature *models.ThreatSignature) error
	CreateThreatSignature(ctx context.Context, signature *models.ThreatSignature) error
	DeleteThreatSignature(ctx context.Context, id string) error
}

// In-memory implementation of ThreatRepositoryInterface

type MemoryThreatRepository struct {
	threats     map[string]*models.Threat
	signatures  map[string]*models.ThreatSignature
}

func NewMemoryThreatRepository() *MemoryThreatRepository {
	return &MemoryThreatRepository{
		threats:    make(map[string]*models.Threat),
		signatures: make(map[string]*models.ThreatSignature),
	}
}

func (r *MemoryThreatRepository) GetThreatByID(ctx context.Context, threatID string) (*models.Threat, error) {
	if threat, ok := r.threats[threatID]; ok {
		return threat, nil
	}
	return nil, nil
}

func (r *MemoryThreatRepository) SaveThreat(ctx context.Context, threat *models.Threat) error {
	r.threats[threat.ID] = threat
	return nil
}

func (r *MemoryThreatRepository) ListThreats(ctx context.Context, filter *models.ThreatFilter) ([]models.Threat, error) {
	var result []models.Threat
	for _, threat := range r.threats {
		result = append(result, *threat)
	}
	return result, nil
}

// Signature management methods
func (r *MemoryThreatRepository) GetThreatSignatures(ctx context.Context, filter *models.SignatureFilter) ([]models.ThreatSignature, error) {
	var result []models.ThreatSignature
	for _, sig := range r.signatures {
		// Basic filtering by severity, pattern, signature set, enabled
		if filter != nil {
			if filter.Severity != "" && sig.Severity != filter.Severity {
				continue
			}
			if filter.Pattern != "" && sig.Pattern != filter.Pattern {
				continue
			}
			if filter.SignatureSet != "" && sig.SignatureSet != filter.SignatureSet {
				continue
			}
			if filter.Enabled && !sig.Enabled {
				continue
			}
		}
		result = append(result, *sig)
	}
	return result, nil
}

func (r *MemoryThreatRepository) UpdateThreatSignature(ctx context.Context, id string, signature *models.ThreatSignature) error {
	if _, ok := r.signatures[id]; !ok {
		return nil // Not found
	}
	r.signatures[id] = signature
	return nil
}

func (r *MemoryThreatRepository) CreateThreatSignature(ctx context.Context, signature *models.ThreatSignature) error {
	r.signatures[signature.ID] = signature
	return nil
}

func (r *MemoryThreatRepository) DeleteThreatSignature(ctx context.Context, id string) error {
	delete(r.signatures, id)
	return nil
}
