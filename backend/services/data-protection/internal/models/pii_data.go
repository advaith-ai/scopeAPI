package models

import (
	"time"
)

type PIIType string

const (
	PIITypeEmail           PIIType = "email"
	PIITypePhone           PIIType = "phone"
	PIITypeSSN             PIIType = "ssn"
	PIITypeCreditCard      PIIType = "credit_card"
	PIITypeIPAddress       PIIType = "ip_address"
	PIITypeName            PIIType = "name"
	PIITypeAddress         PIIType = "address"
	PIITypeDateOfBirth     PIIType = "date_of_birth"
	PIITypePassport        PIIType = "passport"
	PIITypeDriversLicense  PIIType = "drivers_license"
	PIITypeBankAccount     PIIType = "bank_account"
	PIITypeHealthRecord    PIIType = "health_record"
	PIITypeBiometric       PIIType = "biometric"
	PIITypeCustom          PIIType = "custom"
)

type PIISensitivityLevel string

const (
	PIISensitivityLow      PIISensitivityLevel = "low"
	PIISensitivityMedium   PIISensitivityLevel = "medium"
	PIISensitivityHigh     PIISensitivityLevel = "high"
	PIISensitivityCritical PIISensitivityLevel = "critical"
)

type PIIDetectionMethod string

const (
	PIIDetectionRegex     PIIDetectionMethod = "regex"
	PIIDetectionML        PIIDetectionMethod = "ml"
	PIIDetectionDictionary PIIDetectionMethod = "dictionary"
	PIIDetectionContext   PIIDetectionMethod = "context"
	PIIDetectionHybrid    PIIDetectionMethod = "hybrid"
)

type PIIData struct {
	ID                string                 `json:"id" db:"id"`
	RequestID         string                 `json:"request_id" db:"request_id"`
	APIID             string                 `json:"api_id" db:"api_id"`
	EndpointID        string                 `json:"endpoint_id" db:"endpoint_id"`
	Type              PIIType                `json:"type" db:"type"`
	Value             string                 `json:"value" db:"value"`
	MaskedValue       string                 `json:"masked_value" db:"masked_value"`
	FieldName         string                 `json:"field_name" db:"field_name"`
	FieldPath         string                 `json:"field_path" db:"field_path"`
	Location          PIILocation            `json:"location" db:"location"`
	SensitivityLevel  PIISensitivityLevel    `json:"sensitivity_level" db:"sensitivity_level"`
	DetectionMethod   PIIDetectionMethod     `json:"detection_method" db:"detection_method"`
	DetectionScore    float64                `json:"detection_score" db:"detection_score"`
	Context           PIIContext             `json:"context" db:"context"`
	Classification    PIIClassification      `json:"classification" db:"classification"`
	ComplianceFlags   []string               `json:"compliance_flags" db:"compliance_flags"`
	ProcessingActions []PIIProcessingAction  `json:"processing_actions" db:"processing_actions"`
	Metadata          map[string]interface{} `json:"metadata" db:"metadata"`
	DetectedAt        time.Time              `json:"detected_at" db:"detected_at"`
	UpdatedAt         time.Time              `json:"updated_at" db:"updated_at"`
}

type PIILocation struct {
	Source     string `json:"source"`      // request, response, header, query, body
	Section    string `json:"section"`     // specific section within source
	StartIndex int    `json:"start_index"` // character position start
	EndIndex   int    `json:"end_index"`   // character position end
	LineNumber int    `json:"line_number"` // line number if applicable
}

type PIIContext struct {
	SurroundingText   string            `json:"surrounding_text"`
	FieldDescription  string            `json:"field_description"`
	DataFormat        string            `json:"data_format"`
	ValidationRules   []string          `json:"validation_rules"`
	BusinessContext   string            `json:"business_context"`
	UserConsent       bool              `json:"user_consent"`
	LegalBasis        string            `json:"legal_basis"`
	RetentionPeriod   string            `json:"retention_period"`
	ProcessingPurpose string            `json:"processing_purpose"`
	ThirdPartySharing bool              `json:"third_party_sharing"`
	Attributes        map[string]string `json:"attributes"`
}

type PIIClassification struct {
	Category          string   `json:"category"`
	Subcategory       string   `json:"subcategory"`
	DataSubject       string   `json:"data_subject"`
	ProcessingBasis   string   `json:"processing_basis"`
	RetentionClass    string   `json:"retention_class"`
	AccessLevel       string   `json:"access_level"`
	EncryptionRequired bool    `json:"encryption_required"`
	MaskingRequired   bool     `json:"masking_required"`
	AuditRequired     bool     `json:"audit_required"`
	ComplianceFrameworks []string `json:"compliance_frameworks"`
}

type PIIProcessingAction struct {
	Action      string                 `json:"action"`
	Status      string                 `json:"status"`
	Timestamp   time.Time              `json:"timestamp"`
	Details     map[string]interface{} `json:"details"`
	Result      string                 `json:"result"`
	Error       string                 `json:"error,omitempty"`
}

type PIIDetectionRequest struct {
	RequestID     string                 `json:"request_id"`
	APIID         string                 `json:"api_id"`
	EndpointID    string                 `json:"endpoint_id"`
	Content       string                 `json:"content"`
	ContentType   string                 `json:"content_type"`
	Source        string                 `json:"source"`
	DetectionRules []PIIDetectionRule    `json:"detection_rules"`
	Options       PIIDetectionOptions    `json:"options"`
	Context       map[string]interface{} `json:"context"`
	IPAddress     string                 `json:"ip_address"`
	UserAgent     string                 `json:"user_agent"`
}

type PIIDetectionRule struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Type        PIIType            `json:"type"`
	Pattern     string             `json:"pattern"`
	Method      PIIDetectionMethod `json:"method"`
	Sensitivity PIISensitivityLevel `json:"sensitivity"`
	Enabled     bool               `json:"enabled"`
	Priority    int                `json:"priority"`
	Conditions  []RuleCondition    `json:"conditions"`
	Actions     []RuleAction       `json:"actions"`
}

type RuleCondition struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type RuleAction struct {
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config"`
}

type PIIDetectionOptions struct {
	EnableMLDetection     bool    `json:"enable_ml_detection"`
	EnableRegexDetection  bool    `json:"enable_regex_detection"`
	EnableContextAnalysis bool    `json:"enable_context_analysis"`
	MinConfidenceScore    float64 `json:"min_confidence_score"`
	MaxScanDepth          int     `json:"max_scan_depth"`
	IncludeMaskedData     bool    `json:"include_masked_data"`
	EnableRealTimeAlerts  bool    `json:"enable_real_time_alerts"`
	CustomRules           []string `json:"custom_rules"`
}

type PIIDetectionResult struct {
	RequestID       string                 `json:"request_id"`
	DetectedPII     []PIIData              `json:"detected_pii"`
	Summary         PIIDetectionSummary    `json:"summary"`
	Recommendations []string               `json:"recommendations"`
	ProcessingTime  time.Duration          `json:"processing_time"`
	Metadata        map[string]interface{} `json:"metadata"`
	DetectedAt      time.Time              `json:"detected_at"`
}

type PIIDetectionSummary struct {
	TotalPIIFound       int                            `json:"total_pii_found"`
	PIIByType           map[PIIType]int                `json:"pii_by_type"`
	PIIBySensitivity    map[PIISensitivityLevel]int    `json:"pii_by_sensitivity"`
	PIIByMethod         map[PIIDetectionMethod]int     `json:"pii_by_method"`
	HighRiskPII         int                            `json:"high_risk_pii"`
	ComplianceImpact    []string                       `json:"compliance_impact"`
	RecommendedActions  []string                       `json:"recommended_actions"`
	OverallRiskScore    float64                        `json:"overall_risk_score"`
}

type PIIMaskingRequest struct {
	RequestID    string                 `json:"request_id"`
	PIIDataItems []PIIData              `json:"pii_data_items"`
	MaskingRules []PIIMaskingRule       `json:"masking_rules"`
	Options      PIIMaskingOptions      `json:"options"`
	Context      map[string]interface{} `json:"context"`
}

type PIIMaskingRule struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	PIIType     PIIType         `json:"pii_type"`
	Method      PIIMaskingMethod `json:"method"`
	Pattern     string          `json:"pattern"`
	Replacement string          `json:"replacement"`
	Enabled     bool            `json:"enabled"`
	Priority    int             `json:"priority"`
	Conditions  []RuleCondition `json:"conditions"`
}

type PIIMaskingMethod string

const (
	PIIMaskingMethodRedaction    PIIMaskingMethod = "redaction"
	PIIMaskingMethodPartialMask  PIIMaskingMethod = "partial_mask"
	PIIMaskingMethodTokenization PIIMaskingMethod = "tokenization"
	PIIMaskingMethodEncryption   PIIMaskingMethod = "encryption"
	PIIMaskingMethodHashing      PIIMaskingMethod = "hashing"
	PIIMaskingMethodSynthetic    PIIMaskingMethod = "synthetic"
	PIIMaskingMethodFormat       PIIMaskingMethod = "format_preserving"
)

type PIIMaskingOptions struct {
	PreserveFormat     bool   `json:"preserve_format"`
	MaskingCharacter   string `json:"masking_character"`
	PartialMaskRatio   float64 `json:"partial_mask_ratio"`
	TokenizationKey    string `json:"tokenization_key"`
	EncryptionKey      string `json:"encryption_key"`
	HashingSalt        string `json:"hashing_salt"`
	RetainStructure    bool   `json:"retain_structure"`
	AuditMasking       bool   `json:"audit_masking"`
}

type PIIMaskingResult struct {
	RequestID      string                 `json:"request_id"`
	MaskedData     []PIIMaskedData        `json:"masked_data"`
	Summary        PIIMaskingSummary      `json:"summary"`
	ProcessingTime time.Duration          `json:"processing_time"`
	Metadata       map[string]interface{} `json:"metadata"`
	MaskedAt       time.Time              `json:"masked_at"`
}

type PIIMaskedData struct {
	OriginalPII   PIIData `json:"original_pii"`
	MaskedValue   string  `json:"masked_value"`
	MaskingMethod PIIMaskingMethod `json:"masking_method"`
	Token         string  `json:"token,omitempty"`
	Success       bool    `json:"success"`
	Error         string  `json:"error,omitempty"`
}

type PIIMaskingSummary struct {
	TotalItemsMasked    int                         `json:"total_items_masked"`
	SuccessfulMasking   int                         `json:"successful_masking"`
	FailedMasking       int                         `json:"failed_masking"`
	MaskingByMethod     map[PIIMaskingMethod]int    `json:"masking_by_method"`
	MaskingByType       map[PIIType]int             `json:"masking_by_type"`
	ProcessingErrors    []string                    `json:"processing_errors"`
}

type PIIFilter struct {
	APIID            string                  `json:"api_id,omitempty"`
	EndpointID       string                  `json:"endpoint_id,omitempty"`
	PIIType          PIIType                 `json:"pii_type,omitempty"`
	SensitivityLevel PIISensitivityLevel     `json:"sensitivity_level,omitempty"`
	DetectionMethod  PIIDetectionMethod      `json:"detection_method,omitempty"`
	StartDate        *time.Time              `json:"start_date,omitempty"`
	EndDate          *time.Time              `json:"end_date,omitempty"`
	ComplianceFlags  []string                `json:"compliance_flags,omitempty"`
	MinScore         *float64                `json:"min_score,omitempty"`
	MaxScore         *float64                `json:"max_score,omitempty"`
	Limit            int                     `json:"limit,omitempty"`
	Offset           int                     `json:"offset,omitempty"`
}

type PIIStatistics struct {
	TotalPIIDetected     int                            `json:"total_pii_detected"`
	PIIByType            map[PIIType]int                `json:"pii_by_type"`
	PIIBySensitivity     map[PIISensitivityLevel]int    `json:"pii_by_sensitivity"`
	PIIByAPI             map[string]int                 `json:"pii_by_api"`
	PIIByEndpoint        map[string]int                 `json:"pii_by_endpoint"`
	DetectionTrends      []PIITrend                     `json:"detection_trends"`
	ComplianceImpact     map[string]int                 `json:"compliance_impact"`
	RiskDistribution     map[string]int                 `json:"risk_distribution"`
	ProcessingActions    map[string]int                 `json:"processing_actions"`
	GeneratedAt          time.Time                      `json:"generated_at"`
}

type PIITrend struct {
	Date             time.Time `json:"date"`
	PIIDetected      int       `json:"pii_detected"`
	HighRiskPII      int       `json:"high_risk_pii"`
	ComplianceIssues int       `json:"compliance_issues"`
}
