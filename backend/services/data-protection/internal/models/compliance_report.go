package models

import (
	"time"
)

type ComplianceFramework string

const (
	ComplianceFrameworkGDPR    ComplianceFramework = "gdpr"
	ComplianceFrameworkCCPA    ComplianceFramework = "ccpa"
	ComplianceFrameworkHIPAA   ComplianceFramework = "hipaa"
	ComplianceFrameworkSOX     ComplianceFramework = "sox"
	ComplianceFrameworkPCIDSS  ComplianceFramework = "pci_dss"
	ComplianceFrameworkISO27001 ComplianceFramework = "iso_27001"
	ComplianceFrameworkNIST    ComplianceFramework = "nist"
	ComplianceFrameworkCustom  ComplianceFramework = "custom"
)

type ComplianceStatus string

const (
	ComplianceStatusCompliant           ComplianceStatus = "compliant"
	ComplianceStatusPartiallyCompliant  ComplianceStatus = "partially_compliant"
	ComplianceStatusNonCompliant        ComplianceStatus = "non_compliant"
	ComplianceStatusUnknown             ComplianceStatus = "unknown"
	ComplianceStatusInProgress          ComplianceStatus = "in_progress"
)

type ComplianceSeverity string

const (
	ComplianceSeverityLow      ComplianceSeverity = "low"
	ComplianceSeverityMedium   ComplianceSeverity = "medium"
	ComplianceSeverityHigh     ComplianceSeverity = "high"
	ComplianceSeverityCritical ComplianceSeverity = "critical"
)

type ViolationStatus string

const (
	ViolationStatusOpen       ViolationStatus = "open"
	ViolationStatusInProgress ViolationStatus = "in_progress"
	ViolationStatusResolved   ViolationStatus = "resolved"
	ViolationStatusIgnored    ViolationStatus = "ignored"
	ViolationStatusFalsePositive ViolationStatus = "false_positive"
)

type ComplianceReport struct {
	ID               string                        `json:"id" db:"id"`
	Name             string                        `json:"name" db:"name"`
	Description      string                        `json:"description" db:"description"`
	Type             string                        `json:"type" db:"type"`
	Framework        ComplianceFramework           `json:"framework" db:"framework"`
	Status           ComplianceStatus              `json:"status" db:"status"`
	Summary          ComplianceReportSummary       `json:"summary" db:"summary"`
	FrameworkReports map[string]FrameworkReport    `json:"framework_reports" db:"framework_reports"`
	Violations       []ComplianceViolation         `json:"violations" db:"violations"`
	Recommendations  []ComplianceRecommendation    `json:"recommendations" db:"recommendations"`
	Trends           []ComplianceTrend             `json:"trends" db:"trends"`
	Filter           *ComplianceReportFilter       `json:"filter,omitempty" db:"filter"`
	Metadata         map[string]interface{}        `json:"metadata" db:"metadata"`
	GeneratedAt      time.Time                     `json:"generated_at" db:"generated_at"`
	GeneratedBy      string                        `json:"generated_by" db:"generated_by"`
	ValidFrom        time.Time                     `json:"valid_from" db:"valid_from"`
	ValidTo          time.Time                     `json:"valid_to" db:"valid_to"`
	CreatedAt        time.Time                     `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time                     `json:"updated_at" db:"updated_at"`
}

type ComplianceReportSummary struct {
	TotalValidations      int                            `json:"total_validations"`
	ComplianceRate        float64                        `json:"compliance_rate"`
	OverallScore          float64                        `json:"overall_score"`
	TotalViolations       int                            `json:"total_violations"`
	OpenViolations        int                            `json:"open_violations"`
	ResolvedViolations    int                            `json:"resolved_violations"`
	ViolationsByFramework map[string]int                 `json:"violations_by_framework"`
	ViolationsBySeverity  map[ComplianceSeverity]int     `json:"violations_by_severity"`
	ViolationsByStatus    map[ViolationStatus]int        `json:"violations_by_status"`
	TopViolations         []TopViolation                 `json:"top_violations"`
	ComplianceByAPI       map[string]ComplianceMetric    `json:"compliance_by_api"`
	ComplianceByEndpoint  map[string]ComplianceMetric    `json:"compliance_by_endpoint"`
	RiskScore             float64                        `json:"risk_score"`
	TrendDirection        string                         `json:"trend_direction"`
}

type FrameworkReport struct {
	FrameworkID       string                     `json:"framework_id"`
	FrameworkName     string                     `json:"framework_name"`
	Status            ComplianceStatus           `json:"status"`
	Score             float64                    `json:"score"`
	TotalValidations  int                        `json:"total_validations"`
	CompliantCount    int                        `json:"compliant_count"`
	ViolationCount    int                        `json:"violation_count"`
	ComplianceRate    float64                    `json:"compliance_rate"`
	AverageScore      float64                    `json:"average_score"`
	RequirementStatus map[string]RequirementStatus `json:"requirement_status"`
	TopViolations     []TopViolation             `json:"top_violations"`
	Recommendations   []string                   `json:"recommendations"`
	LastAssessment    time.Time                  `json:"last_assessment"`
}

type RequirementStatus struct {
	RequirementID   string           `json:"requirement_id"`
	RequirementName string           `json:"requirement_name"`
	Status          ComplianceStatus `json:"status"`
	Score           float64          `json:"score"`
	ViolationCount  int              `json:"violation_count"`
	LastChecked     time.Time        `json:"last_checked"`
}

type ComplianceViolation struct {
	ID              string                 `json:"id" db:"id"`
	RuleID          string                 `json:"rule_id" db:"rule_id"`
	RuleName        string                 `json:"rule_name" db:"rule_name"`
	Framework       string                 `json:"framework" db:"framework"`
	Category        string                 `json:"category" db:"category"`
	Severity        ComplianceSeverity     `json:"severity" db:"severity"`
	Status          ViolationStatus        `json:"status" db:"status"`
	Message         string                 `json:"message" db:"message"`
	Description     string                 `json:"description" db:"description"`
	Remediation     string                 `json:"remediation" db:"remediation"`
	APIID           string                 `json:"api_id" db:"api_id"`
	EndpointID      string                 `json:"endpoint_id" db:"endpoint_id"`
	RequestID       string                 `json:"request_id" db:"request_id"`
	Evidence        ViolationEvidence      `json:"evidence" db:"evidence"`
	Impact          ViolationImpact        `json:"impact" db:"impact"`
	Resolution      *ViolationResolution   `json:"resolution,omitempty" db:"resolution"`
	Metadata        map[string]interface{} `json:"metadata" db:"metadata"`
	DetectedAt      time.Time              `json:"detected_at" db:"detected_at"`
	ResolvedAt      *time.Time             `json:"resolved_at,omitempty" db:"resolved_at"`
	CreatedAt       time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at" db:"updated_at"`
	AssignedTo      string                 `json:"assigned_to" db:"assigned_to"`
	DueDate         *time.Time             `json:"due_date,omitempty" db:"due_date"`
}

type ViolationEvidence struct {
	RequestData     string                 `json:"request_data"`
	ResponseData    string                 `json:"response_data"`
	Headers         map[string]string      `json:"headers"`
	QueryParams     map[string]string      `json:"query_params"`
	BodyContent     string                 `json:"body_content"`
	Metadata        map[string]interface{} `json:"metadata"`
	Screenshots     []string               `json:"screenshots"`
	LogEntries      []string               `json:"log_entries"`
	NetworkTraces   []string               `json:"network_traces"`
}

type ViolationImpact struct {
	RiskScore       float64  `json:"risk_score"`
	BusinessImpact  string   `json:"business_impact"`
	TechnicalImpact string   `json:"technical_impact"`
	LegalImpact     string   `json:"legal_impact"`
	FinancialImpact string   `json:"financial_impact"`
	AffectedSystems []string `json:"affected_systems"`
	AffectedUsers   int      `json:"affected_users"`
	DataVolume      int64    `json:"data_volume"`
	Urgency         string   `json:"urgency"`
}

type ViolationResolution struct {
	ResolutionType  string                 `json:"resolution_type"`
	Description     string                 `json:"description"`
	Actions         []ResolutionAction     `json:"actions"`
	ResolvedBy      string                 `json:"resolved_by"`
	ResolvedAt      time.Time              `json:"resolved_at"`
	VerifiedBy      string                 `json:"verified_by"`
	VerifiedAt      *time.Time             `json:"verified_at,omitempty"`
	Notes           string                 `json:"notes"`
	Attachments     []string               `json:"attachments"`
	FollowUpRequired bool                  `json:"follow_up_required"`
	FollowUpDate    *time.Time             `json:"follow_up_date,omitempty"`
}

type ResolutionAction struct {
	Action      string                 `json:"action"`
	Description string                 `json:"description"`
	Status      string                 `json:"status"`
	AssignedTo  string                 `json:"assigned_to"`
	DueDate     *time.Time             `json:"due_date,omitempty"`
	CompletedAt *time.Time             `json:"completed_at,omitempty"`
	Notes       string                 `json:"notes"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type TopViolation struct {
	RuleID      string             `json:"rule_id"`
	RuleName    string             `json:"rule_name"`
	Count       int                `json:"count"`
	Severity    ComplianceSeverity `json:"severity"`
	Category    string             `json:"category"`
	Framework   string             `json:"framework"`
	Percentage  float64            `json:"percentage"`
	TrendChange float64            `json:"trend_change"`
}

type ComplianceRecommendation struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Priority    string                 `json:"priority"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Framework   string                 `json:"framework"`
	Category    string                 `json:"category"`
	Actions     []RecommendationAction `json:"actions"`
	Impact      string                 `json:"impact"`
	Effort      string                 `json:"effort"`
	Timeline    string                 `json:"timeline"`
	Resources   []string               `json:"resources"`
	Dependencies []string              `json:"dependencies"`
	Status      string                 `json:"status"`
	CreatedAt   time.Time              `json:"created_at"`
}

type RecommendationAction struct {
	Action      string    `json:"action"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Effort      string    `json:"effort"`
	Timeline    string    `json:"timeline"`
	Owner       string    `json:"owner"`
	DueDate     time.Time `json:"due_date"`
}

type ComplianceTrend struct {
	Date             time.Time        `json:"date"`
	Framework        string           `json:"framework"`
	Status           ComplianceStatus `json:"status"`
	Score            float64          `json:"score"`
	TotalValidations int              `json:"total_validations"`
	CompliantCount   int              `json:"compliant_count"`
	ViolationCount   int              `json:"violation_count"`
	ComplianceRate   float64          `json:"compliance_rate"`
	Change           float64          `json:"change"`
	ChangePercent    float64          `json:"change_percent"`
}

type ComplianceMetric struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	Status           ComplianceStatus `json:"status"`
	Score            float64          `json:"score"`
	ViolationCount   int              `json:"violation_count"`
	ComplianceRate   float64          `json:"compliance_rate"`
	LastAssessment   time.Time        `json:"last_assessment"`
	TrendDirection   string           `json:"trend_direction"`
	RiskLevel        string           `json:"risk_level"`
}

type ComplianceReportFilter struct {
	Frameworks   []string    `json:"frameworks,omitempty"`
	APIIDs       []string    `json:"api_ids,omitempty"`
	EndpointIDs  []string    `json:"endpoint_ids,omitempty"`
	Severities   []string    `json:"severities,omitempty"`
	Statuses     []string    `json:"statuses,omitempty"`
	StartDate    *time.Time  `json:"start_date,omitempty"`
	EndDate      *time.Time  `json:"end_date,omitempty"`
	Categories   []string    `json:"categories,omitempty"`
	RuleIDs      []string    `json:"rule_ids,omitempty"`
	IncludeTrends bool       `json:"include_trends"`
	TrendPeriod  string     `json:"trend_period"`
	Limit        int        `json:"limit,omitempty"`
	Offset       int        `json:"offset,omitempty"`
}

type ComplianceValidation struct {
	ID               string                           `json:"id" db:"id"`
	RequestID        string                           `json:"request_id" db:"request_id"`
	APIID            string                           `json:"api_id" db:"api_id"`
	EndpointID       string                           `json:"endpoint_id" db:"endpoint_id"`
	Frameworks       []string                         `json:"frameworks" db:"frameworks"`
	OverallStatus    ComplianceStatus                 `json:"overall_status" db:"overall_status"`
	OverallScore     float64                          `json:"overall_score" db:"overall_score"`
	ViolationCount   int                              `json:"violation_count" db:"violation_count"`
	WarningCount     int                              `json:"warning_count" db:"warning_count"`
	FrameworkResults map[string]FrameworkValidationResult `json:"framework_results" db:"framework_results"`
	Violations       []ComplianceViolation            `json:"violations" db:"violations"`
	Warnings         []ComplianceWarning              `json:"warnings" db:"warnings"`
	ProcessingTime   time.Duration                    `json:"processing_time" db:"processing_time"`
	Metadata         map[string]interface{}           `json:"metadata" db:"metadata"`
	ValidatedAt      time.Time                        `json:"validated_at" db:"validated_at"`
	CreatedAt        time.Time                        `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time                        `json:"updated_at" db:"updated_at"`
}

type FrameworkValidationResult struct {
	Framework      string                `json:"framework"`
	Status         ComplianceStatus      `json:"status"`
	Score          float64               `json:"score"`
	ViolationCount int                   `json:"violation_count"`
	WarningCount   int                   `json:"warning_count"`
	Violations     []ComplianceViolation `json:"violations"`
	Warnings       []ComplianceWarning   `json:"warnings"`
	Requirements   []RequirementResult   `json:"requirements"`
	ProcessingTime time.Duration         `json:"processing_time"`
}

type RequirementResult struct {
	RequirementID   string           `json:"requirement_id"`
	RequirementName string           `json:"requirement_name"`
	Status          ComplianceStatus `json:"status"`
	Score           float64          `json:"score"`
	Message         string           `json:"message"`
	Evidence        string           `json:"evidence"`
}

type ComplianceWarning struct {
	ID          string                 `json:"id"`
	RuleID      string                 `json:"rule_id"`
	RuleName    string                 `json:"rule_name"`
	Framework   string                 `json:"framework"`
	Category    string                 `json:"category"`
	Message     string                 `json:"message"`
	Description string                 `json:"description"`
	Suggestion  string                 `json:"suggestion"`
	APIID       string                 `json:"api_id"`
	EndpointID  string                 `json:"endpoint_id"`
	RequestID   string                 `json:"request_id"`
	Metadata    map[string]interface{} `json:"metadata"`
	DetectedAt  time.Time              `json:"detected_at"`
}

type ComplianceValidationRequest struct {
	RequestID   string                 `json:"request_id"`
	APIID       string                 `json:"api_id"`
	EndpointID  string                 `json:"endpoint_id"`
	Frameworks  []string               `json:"frameworks"`
	Content     string                 `json:"content"`
	ContentType string                 `json:"content_type"`
	Source      string                 `json:"source"`
	Rules       []string               `json:"rules"`
	Options     ValidationOptions      `json:"options"`
	Context     map[string]interface{} `json:"context"`
	IPAddress   string                 `json:"ip_address"`
	UserAgent   string                 `json:"user_agent"`
}

type ValidationOptions struct {
	EnableRealTimeValidation bool     `json:"enable_real_time_validation"`
	EnableDeepValidation     bool     `json:"enable_deep_validation"`
	IncludeWarnings          bool     `json:"include_warnings"`
	IncludeRecommendations   bool     `json:"include_recommendations"`
	CustomRules              []string `json:"custom_rules"`
	ValidationDepth          int      `json:"validation_depth"`
	TimeoutSeconds           int      `json:"timeout_seconds"`
}

type ComplianceValidationResult struct {
	RequestID       string                           `json:"request_id"`
	OverallStatus   ComplianceStatus                 `json:"overall_status"`
	OverallScore    float64                          `json:"overall_score"`
	ViolationCount  int                              `json:"violation_count"`
	WarningCount    int                              `json:"warning_count"`
	FrameworkResults map[string]FrameworkValidationResult `json:"framework_results"`
	Violations      []ComplianceViolation            `json:"violations"`
	Warnings        []ComplianceWarning              `json:"warnings"`
	Recommendations []ComplianceRecommendation       `json:"recommendations"`
	ProcessingTime  time.Duration                    `json:"processing_time"`
	Metadata        map[string]interface{}           `json:"metadata"`
	ValidatedAt     time.Time                        `json:"validated_at"`
}

type ComplianceRule struct {
	ID              string                 `json:"id" db:"id"`
	Name            string                 `json:"name" db:"name"`
	Description     string                 `json:"description" db:"description"`
	Framework       string                 `json:"framework" db:"framework"`
	Category        string                 `json:"category" db:"category"`
	Severity        ComplianceSeverity     `json:"severity" db:"severity"`
	Type            string                 `json:"type" db:"type"`
	Enabled         bool                   `json:"enabled" db:"enabled"`
	Priority        int                    `json:"priority" db:"priority"`
	Conditions      []RuleCondition        `json:"conditions" db:"conditions"`
	Actions         []RuleAction           `json:"actions" db:"actions"`
	Requirements    []string               `json:"requirements" db:"requirements"`
	Tags            []string               `json:"tags" db:"tags"`
	References      []string               `json:"references" db:"references"`
	Remediation     string                 `json:"remediation" db:"remediation"`
	Examples        []RuleExample          `json:"examples" db:"examples"`
	Metadata        map[string]interface{} `json:"metadata" db:"metadata"`
	CreatedAt       time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at" db:"updated_at"`
	CreatedBy       string                 `json:"created_by" db:"created_by"`
	UpdatedBy       string                 `json:"updated_by" db:"updated_by"`
	Version         string                 `json:"version" db:"version"`
	LastTested      *time.Time             `json:"last_tested,omitempty" db:"last_tested"`
}

type RuleCondition struct {
	Field       string                 `json:"field"`
	Operator    string                 `json:"operator"`
	Value       string                 `json:"value"`
	CaseSensitive bool                 `json:"case_sensitive"`
	Weight      float64                `json:"weight"`
	Context     map[string]interface{} `json:"context"`
	Negated     bool                   `json:"negated"`
}

type RuleAction struct {
	Type        string                 `json:"type"`
	Config      map[string]interface{} `json:"config"`
	Enabled     bool                   `json:"enabled"`
	Priority    int                    `json:"priority"`
	Conditions  []string               `json:"conditions"`
	Parameters  map[string]string      `json:"parameters"`
}

type RuleExample struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Input       string `json:"input"`
	Expected    string `json:"expected"`
	Violation   bool   `json:"violation"`
}

type ComplianceRuleFilter struct {
	Framework string `json:"framework,omitempty"`
	Category  string `json:"category,omitempty"`
	Severity  string `json:"severity,omitempty"`
	Enabled   *bool  `json:"enabled,omitempty"`
	Name      string `json:"name,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	Limit     int    `json:"limit,omitempty"`
	Offset    int    `json:"offset,omitempty"`
}

type ComplianceStatusFilter struct {
	Frameworks []string `json:"frameworks,omitempty"`
	APIIDs     []string `json:"api_ids,omitempty"`
	Statuses   []string `json:"statuses,omitempty"`
}

type ComplianceStatus struct {
	OverallStatus     ComplianceStatus             `json:"overall_status"`
	FrameworkStatuses map[string]FrameworkStatus   `json:"framework_statuses"`
	Summary           ComplianceStatusSummary      `json:"summary"`
	LastUpdated       time.Time                    `json:"last_updated"`
}

type FrameworkStatus struct {
	FrameworkID      string           `json:"framework_id"`
	FrameworkName    string           `json:"framework_name"`
	Status           ComplianceStatus `json:"status"`
	Score            float64          `json:"score"`
	ActiveViolations int              `json:"active_violations"`
	LastAssessment   time.Time        `json:"last_assessment"`
}

type ComplianceStatusSummary struct {
	TotalFrameworks     int `json:"total_frameworks"`
	CompliantFrameworks int `json:"compliant_frameworks"`
	ActiveViolations    int `json:"active_violations"`
	RecentViolations    int `json:"recent_violations"`
}

type ComplianceValidationFilter struct {
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	Frameworks  []string   `json:"frameworks,omitempty"`
	APIIDs      []string   `json:"api_ids,omitempty"`
	EndpointIDs []string   `json:"endpoint_ids,omitempty"`
	Statuses    []string   `json:"statuses,omitempty"`
	Limit       int        `json:"limit,omitempty"`
	Offset      int        `json:"offset,omitempty"`
}
