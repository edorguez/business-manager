# Agent Guidelines for Business Manager

This document provides guidelines for AI agents working on the Business Manager codebase. It covers build commands, testing, linting, and code style conventions for both Go backend services and Next.js frontend.

## Overview

The Business Manager is a monorepo containing:
- **Go backend services** (auth-svc, company-svc, customer-svc, etc.) using gRPC, PostgreSQL, and SQLc
- **Next.js frontend** (web) with TypeScript, Chakra UI, Tailwind CSS
- **Shared protobuf definitions** in `proto/`
- **Shared Go packages** in `shared/`

## Build Commands

### Go Services
```bash
# Generate protobuf Go code
make proto-all

# Run a specific service (e.g., auth-svc)
cd services/auth-svc && go run cmd/main.go

# Build all services for production
make docker-image-build-all-prod

# Build all services for development
make docker-image-build-all-dev

# Run development Docker compose
make docker-compose-up-dev

# Stop development Docker compose
make docker-compose-down-dev

# Run production Docker compose
make docker-compose-up-prod

# Stop production Docker compose
make docker-compose-down-prod
```

### Frontend (Next.js)
```bash
cd web

# Install dependencies
npm install

# Run development server
npm run dev

# Build for production
npm run build

# Start production server
npm start

# Lint code
npm run lint
```

## Linting and Formatting

### Go
- Use `gofmt` to format Go code. No custom linting configuration is present.
- Ensure imports are grouped: standard library, third-party, internal.
- Run `go vet ./...` to check for suspicious constructs.

### Frontend
- ESLint configuration extends `next/core-web-vitals`.
- Prettier configuration: 2-space tabs, no tabs.
- Run linting with `npm run lint` in the `web` directory.

## Testing

### Go
- Test files are named `*_test.go` and reside alongside the code they test.
- Use `go test ./...` to run all tests.
- Use `go test ./services/auth-svc/pkg/...` to test a specific package.
- To run a single test: `go test -v ./services/auth-svc/pkg/services -run TestSignUp`
- The project uses `testify` for assertions (already in go.mod).

### Frontend
- No dedicated test framework is configured. Use `jest` or `vitest` if added later.
- For now, focus on manual verification.

## Code Style Guidelines

### Go

#### Imports
Group imports in the following order:
1. Standard library
2. Third-party packages
3. Internal packages (from shared or other services)

Example:
```go
import (
    "context"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/spf13/viper"

    pbauth "github.com/edorguez/business-manager/shared/pb/auth"
    "github.com/edorguez/business-manager/shared/util/jwt_manager"
)
```

#### Error Handling
- For gRPC services, errors are returned as part of the response message (using `status` and `error` fields). Do not return Go errors directly; instead, populate the `Error` field and return `nil` error.
- Use `fmt.Println` for logging (avoid more sophisticated logging libraries unless added).
- For database operations, check `sql.ErrNoRows` and handle appropriately.

#### Naming Conventions
- Use `PascalCase` for exported functions and types.
- Use `camelCase` for unexported functions and variables.
- Use `UPPER_SNAKE_CASE` for constants.
- Struct field names should be `PascalCase` (matching protobuf generated code).
- Interface names: `PascalCase` with `er` suffix if applicable (e.g., `Validator`).

#### Function Signatures
- Keep receiver names short (e.g., `s *AuthService`).
- Use `ctx context.Context` as first parameter for functions that need context.
- Return errors as the last return value (except in gRPC service methods).
- Service structs should embed the generated `Unimplemented...Server` (e.g., `pb.UnimplementedAuthServiceServer`).

#### Logging
- Use `fmt.Println` for informational logs.
- Include service name and operation in log messages (e.g., `"Auth Service : Sign Up - SUCCESS"`).

### TypeScript / React

#### Imports
Order imports:
1. React and Next.js core
2. Third-party libraries
3. Internal components and utilities
4. Styles and assets

Use absolute imports with `@/*` alias (configured in `tsconfig.json`).

#### Components
- Use functional components with TypeScript interfaces for props.
- Prefer `"use client"` directive for client components.
- Use Chakra UI components (`Button`, `Box`, etc.) with custom variants defined in `chakraui.config.ts`.
- Tailwind CSS classes should be used for layout and spacing; Chakra UI for interactive components.

#### Styling
- Combine Tailwind CSS with Chakra UI styling props.
- Use `className` for Tailwind classes; use `sx` or `style` prop for dynamic styles.
- Custom colors are defined in `tailwind.config.ts` (e.g., `maincolor`, `thirdcolor`).

#### State Management
- Use `zustand` for global state (already installed).
- For local state, use `useState` and `useEffect`.

#### Error Handling
- Use try/catch for async operations.
- Display user-friendly error messages via Chakra UI toast or alert components.

#### Naming Conventions
- Components: `PascalCase` (e.g., `FloatingItem`).
- Functions and variables: `camelCase`.
- Interfaces and types: `PascalCase`.
- Constants: `UPPER_SNAKE_CASE` or `camelCase` depending on usage.

## Protobuf Definitions

- Proto files are in `proto/` directory.
- Use `proto3` syntax.
- Follow snake_case for field names (e.g., `name_format_url`).
- Use PascalCase for message and service names.
- Generate Go code with `make proto-all`; output goes to `shared/pb/`.

## Database

- Each service has its own database schema and migrations.
- SQLc is used to generate type-safe Go code from SQL queries. Each service has a `sqlc.yml` config.
- To regenerate SQLc code after modifying SQL queries, run `sqlc generate` in the service directory.
- Migration files are located in each service's `pkg/db/migration/` directory.
- Run migrations with `make migrateup-all` (requires environment variables).

## Environment Variables

- Copy `example.env` to `.env` and fill in required values.
- Frontend environment variables are in `web/.env`.
- Environment variables are loaded via `config.LoadConfig()` in Go services.

## Commit Conventions

- Use descriptive commit messages that explain the "why" not just the "what".
- Follow existing patterns in the git log.

## Agent-Specific Notes

- When editing Go code, ensure the service compiles (`go build ./...`).
- When editing frontend code, run `npm run lint` to avoid style issues.
- When adding new dependencies, update `go.mod` or `web/package.json` accordingly.
- Never commit `.env` files or secrets.
- If you add a new feature, consider adding appropriate tests.

## Quick Reference

| Task | Command |
|------|---------|
| Generate protobuf | `make proto-all` |
| Run all tests | `go test ./...` |
| Run frontend lint | `cd web && npm run lint` |
| Start dev environment | `make docker-compose-up-dev` |
| Stop dev environment | `make docker-compose-down-dev` |
| Build frontend | `cd web && npm run build` |

---

*This file is intended to help AI agents understand the project structure and conventions. Update it as the codebase evolves.*