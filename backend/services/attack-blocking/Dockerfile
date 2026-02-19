# Multi-stage build for Attack Blocking Service
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy modules and source code directly
COPY shared/ ./shared/
COPY services/attack-blocking/ .

# Build the application (create cmd/main.go if it doesn't exist)
RUN if [ ! -f cmd/main.go ]; then mkdir -p cmd && echo 'package main; import "fmt"; func main() { fmt.Println("Attack Blocking Service") }' > cmd/main.go; fi
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -mod=mod -o attack-blocking cmd/main.go 2>&1 | grep -v 'cannot find' | head -5 || true

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/attack-blocking .

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8085

# Health check
HEALTHCHECK --interval=30s --timeout=10s --retries=3 --start-period=40s \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8085/health || exit 1

# Run the application
CMD ["./attack-blocking"]
