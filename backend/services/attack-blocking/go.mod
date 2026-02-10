module scopeapi.local/backend/services/attack-blocking

go 1.21

require (
	github.com/gin-gonic/gin v1.10.1
	github.com/google/uuid v1.6.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/lib/pq v1.10.9
	shared v0.0.0
)

replace shared => ../../shared
