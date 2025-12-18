FROM golang:1.21-alpine AS build

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 go build -o bin/challenge

# Production stage - minimal image
FROM alpine:3.19

# Install runtime dependencies and create non-root user
RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -g 1000 appgroup \
    && adduser -u 1000 -G appgroup -s /bin/sh -D appuser

WORKDIR /app

# Copy binary from build stage
COPY --from=build /go/src/app/bin/challenge /app/challenge

# Set ownership
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose application port
EXPOSE 10000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:10000/v1/users || exit 1

ENTRYPOINT ["/app/challenge"]
