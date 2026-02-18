# ============================================
# Stage 1: Build
# ============================================
FROM golang:1.24-alpine AS builder

WORKDIR /build

# Install git (needed for some Go modules)
RUN apk add --no-cache git ca-certificates tzdata

# Download dependencies first (better layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build static binary (CGO disabled = fully static, no libc dependency)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s -extldflags '-static'" \
    -o /api cmd/api/main.go

# Build migrate tool
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s -extldflags '-static'" \
    -o /migrate cmd/migrate/main.go

# ============================================
# Stage 2: Runtime (distroless — ~2MB base)
# ============================================
FROM gcr.io/distroless/static:nonroot

# Copy timezone data and CA certs from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binaries
COPY --from=builder /api /api
COPY --from=builder /migrate /migrate

# Copy migrations
COPY --from=builder /build/migrations /migrations

# Run as non-root user (distroless/nonroot uses UID 65532)
USER nonroot:nonroot

EXPOSE 3000

ENTRYPOINT ["/api"]
