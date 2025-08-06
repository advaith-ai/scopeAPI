package models

import "time"

type SignatureDetectionRequest struct {
	RequestID   string                 `json:"request_id"`
	EntityID    string                 `json:"entity_id"`
	EntityType  string                 `json:"entity_type"`
	Payload     map[string]interface{} `json:"payload"`
	Timestamp   time.Time              `json:"timestamp"`
}

type SignatureDetectionResult struct {
	ResultID    string    `json:"result_id"`
	SignatureID string    `json:"signature_id"`
	Matched     bool      `json:"matched"`
	Details     string    `json:"details"`
	DetectedAt  time.Time `json:"detected_at"`
}

type ThreatSignature struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Pattern     string    `json:"pattern"`
	Severity    string    `json:"severity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Type        string    `json:"type"`
	Category    string    `json:"category"`
	RiskScore   float64   `json:"risk_score"`
	Confidence  float64   `json:"confidence"`
	Tags        []string  `json:"tags"`
	SignatureSet string   `json:"signature_set"`
	Enabled     bool      `json:"enabled"`
	Rules       []SignatureRule `json:"rules"`
}

type SignatureFilter struct {
	Severity     string    `json:"severity,omitempty"`
	Pattern      string    `json:"pattern,omitempty"`
	SignatureSet string    `json:"signature_set,omitempty"`
	Enabled      bool      `json:"enabled,omitempty"`
}

type SignatureTestResult struct {
	TestID      string    `json:"test_id"`
	SignatureID string    `json:"signature_id"`
	Passed      bool      `json:"passed"`
	Details     string    `json:"details"`
	TestedAt    time.Time `json:"tested_at"`
}

type SignatureMatch struct {
	SignatureID   string    `json:"signature_id"`
	SignatureName string    `json:"signature_name"`
	SignatureType string    `json:"signature_type"`
	Category      string    `json:"category"`
	Severity      string    `json:"severity"`
	RiskScore     float64   `json:"risk_score"`
	Confidence    float64   `json:"confidence"`
	Description   string    `json:"description"`
	MatchedField  string    `json:"matched_field"`
	MatchedValue  string    `json:"matched_value"`
	RuleMatched   string    `json:"rule_matched"`
	RuleOperator  string    `json:"rule_operator"`
	RuleValue     string    `json:"rule_value"`
	MatchedAt     time.Time `json:"matched_at"`
	Metadata      map[string]interface{} `json:"metadata"`
}

type SignatureRule struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Pattern     string `json:"pattern"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
	Field       string `json:"field"`
	Operator    string `json:"operator"`
	Value       string `json:"value"`
	IntValue    int    `json:"int_value"`
	Weight      float64 `json:"weight"`
}

type SignatureStatistics struct {
	TotalSignatures int `json:"total_signatures"`
	Matched         int `json:"matched"`
	Unmatched       int `json:"unmatched"`
}

type SignatureOptimizationResult struct {
	OptimizedSignatures int    `json:"optimized_signatures"`
	Details             string `json:"details"`
} 