module scopeapi.local/backend/services/threat-detection

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
	github.com/spf13/viper v1.18.2
	shared v0.0.0
)

replace shared => ../../shared