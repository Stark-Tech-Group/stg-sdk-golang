# CLAUDE.md - stg-sdk-golang

## Overview

Go SDK for the Stark Platform. Provides a typed client for interacting with the Stark API, including authentication, asset management, telemetry, and search capabilities.

**Go Version**: 1.25.0

## Project Structure

```
stg-sdk-golang/
├── cmd/stg-sdk-golang/main.go    # CLI/example entry point
├── pkg/                           # Public SDK packages
│   ├── api/response/              # API response types
│   ├── azure/                     # Azure Event Hub integration
│   ├── context/                   # Context helpers
│   ├── domain/                    # Domain models (30+ types)
│   ├── env/                       # Environment helpers
│   ├── http/                      # HTTP utilities
│   └── starkapi/                  # Stark API client (core)
├── docker/                        # Docker dev support
└── go.mod
```

## Build & Test Commands

```bash
# Build
go build ./...

# Test all packages
go test ./...

# Test with verbose output
go test -v ./...

# Test specific package
go test -v ./pkg/starkapi/...
go test -v ./pkg/domain/...

# Test with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Verify dependencies
go mod tidy
go mod verify
```

## Key Dependencies

- **Azure**: azure-event-hubs-go/v3 for Event Hub integration
- **Validation**: go-playground/validator/v10
- **Database**: lib/pq for PostgreSQL
- **Logging**: sirupsen/logrus
- **Testing**: stretchr/testify

## Package Overview

### pkg/starkapi (Core Client)

Main SDK client providing access to all API endpoints:

```go
api := starkapi.Client{}
api.Init(host)
api.Login(username, password)

// Available APIs
api.AssetTreeApi   // Asset tree navigation
api.AssetsApi      // Asset management
api.PointApi       // Point/sensor data
api.EquipApi       // Equipment management
api.SiteApi        // Site management
api.SearchApi      // Search queries
api.StatusApi      // API health status
api.ProfileApi     // User profiles
api.ConnApi        // Connections
api.GeoApi         // Geolocation
api.UridApi        // Unique resource IDs
api.TagApi         // Tag management
api.FormsApi       // Form controls
```

### pkg/domain

Domain models for the Stark Platform:
- `Asset`, `AssetTree` - Asset hierarchy
- `Point`, `PointType` - Sensor/data points
- `Equip`, `EquipType` - Equipment
- `Site` - Physical locations
- `Tag`, `TagRef` - Metadata tagging
- `TelemetryMessage` - Telemetry data format
- And more (30+ models)

### pkg/api/response

API response types:
- `AuthResponse` - Authentication result
- `SearchResponse` - Search results
- `StatusResponse` - API status
- `DeleteResponse` - Deletion confirmation

### pkg/azure

Azure Event Hub integration for sending telemetry messages.

### pkg/env

Environment variable helpers for SDK configuration.

### pkg/context

Context helpers for request handling.

### pkg/http

HTTP error handling utilities.

## Environment Variables

| Variable | Description |
|----------|-------------|
| `STG_SDK_API_HOST` | The host API address (e.g., https://stgcapi.staging.starktechgroup.com) |
| `STG_SDK_API_UN` | Your username |
| `STG_SDK_API_PW` | Your password |

## Usage Examples

### Pre-authenticated Client

```go
api := starkapi.Client{}
api.Init(host)
api.Auth(accessToken, username) // Use existing token
```

## Testing

Tests use testify assertions and mock clients:

```bash
# Run all tests
go test ./...

# Run starkapi tests (includes mock client)
go test -v ./pkg/starkapi/...

# Run domain model tests
go test -v ./pkg/domain/...
```

Mock client available in `pkg/starkapi/mock-client_test.go` for testing.

## CI/CD

Custom GitHub Actions workflows (SDK-specific):
- `dev.yml` - Build and release on dev branch
- `feature_bug.yml` - Build on feature/bug branches
- `sonar.yml` - SonarCloud analysis

SonarCloud configured via `sonar-project.properties`.

## SDK Consumer Usage

To use this SDK in another Go project:

```go
import (
    "github.com/Stark-Tech-Group/stg-sdk-golang/pkg/starkapi"
    "github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)
```

Install:
```bash
go get github.com/Stark-Tech-Group/stg-sdk-golang@latest
```
