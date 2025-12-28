<p align="center">
  <a href="http://go.dev" target="blank"><img src="go.png" width="200" alt="Nest Logo" /></a>
</p>

# IAM Service - Identity and Access Management

A production-ready Identity and Access Management (IAM) service built with Go Fiber, featuring JWT-based authentication, permission-based authorization, and comprehensive user management capabilities.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Installation & Setup](#installation--setup)
- [Environment Configuration](#environment-configuration)
- [Database Management](#database-management)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [License](#license)

## Features

### Authentication & Authorization
- **JWT-based dual-token authentication** (access + refresh tokens)
- **Permission-based authorization** (roles group permissions for easier management)
- Secure password hashing with bcrypt
- Token renewal mechanism
- Automatic token cleanup (scheduled daily at midnight)
- Bearer token middleware protection

### User Management
- Full CRUD operations
- Role assignment to users (many-to-many relationship)
- Pagination with customizable page size
- Search and filtering by role
- Email validation
- Password strength validation

### Role Management
- CRUD operations for roles
- Permission assignment to roles (many-to-many relationship)
- Auto-generated unique role keys from names
- Collision-safe key generation
- Prevent deletion if role is assigned to users

### Permission Management
- CRUD operations for permissions
- Auto-generated unique permission keys
- Predefined permission constants
- Prevent deletion if permission is assigned to roles

### Background Jobs
- Asynq integration for async task processing
- Scheduled token cleanup job (daily at midnight)
- Job worker with configurable concurrency
- Redis-backed job queue

## Tech Stack

**Core Framework**
- Go 1.25
- [Fiber v2](https://gofiber.io/) - Fast HTTP web framework

**Database & ORM**
- PostgreSQL
- [GORM](https://gorm.io/) - ORM for Go
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate) - Database migrations

**Authentication & Security**
- [golang-jwt/jwt v5](https://github.com/golang-jwt/jwt) - JWT token generation and validation
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Password hashing
- [validator v10](https://github.com/go-playground/validator) - Request validation

**Background Jobs**
- [Asynq](https://github.com/hibiken/asynq) - Async task queue
- Redis - Job queue backend
- [Cron v3](https://github.com/robfig/cron) - Job scheduling

**Development Tools**
- [Air](https://github.com/air-verse/air) - Hot reload for Go
- [godotenv](https://github.com/joho/godotenv) - Environment variable loading
- [Logrus](https://github.com/sirupsen/logrus) - Structured logging

## Prerequisites

- **Go** 1.25 or higher
- **PostgreSQL** 13+ (running instance)
- **Redis** 6+ (for background jobs)
- **Make** (optional, for using Makefile commands)

## Installation & Setup

### Step 1: Clone the repository

```bash
git clone <repository-url>
cd go-fiber-iam
```

### Step 2: Install dependencies

```bash
go mod download
```

### Step 3: Environment Configuration

Copy the example environment file and configure it:

```bash
cp .env.example .env
```

Edit `.env` with your configuration (see [Environment Configuration](#environment-configuration) section for details).

### Step 4: Database Setup

Run migrations to create database tables:

```bash
make migrate-up db=main_db
```

Seed initial data (creates default admin user and permissions):

```bash
make seed db=main_db
```

### Step 5: Run the application

Development mode with hot reload (recommended):

```bash
make air
```

Or standard development mode:

```bash
make dev
```

Run background job worker (in a separate terminal):

```bash
make job-worker
```

The service will be available at `http://localhost:3000`

## Environment Configuration

Create a `.env` file in the project root with the following variables:

```env
# Application
PORT=3000

# PostgreSQL Database Configuration
MAIN_DB_HOST=localhost
MAIN_DB_PORT=5432
MAIN_DB_USER=postgres
MAIN_DB_PASS=postgres
MAIN_DB_NAME=iam_service_db
MAIN_DB_SSL_MODE=disable
MAIN_DB_TIMEZONE=Asia/Jakarta

# Database Connection Pool
MAIN_DB_MAX_IDLE_CONNS=5
MAIN_DB_MAX_OPEN_CONNS=10
MAIN_DB_MAX_IDLE_CONNS_IN_MINUTES=10
MAIN_DB_MAX_LIFETIME_CONNS_IN_MINUTES=60

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT Configuration
JWT_ACCESS_TOKEN_SECRET=your-access-token-secret-change-in-production
JWT_ACCESS_TOKEN_EXPIRED_IN=30m

JWT_REFRESH_TOKEN_SECRET=your-refresh-token-secret-change-in-production
JWT_REFRESH_TOKEN_EXPIRED_IN=24d

# Background Job Scheduler
SCHEDULER_CONCURRENCY=1
```

**Important Security Notes:**
- Change `JWT_ACCESS_TOKEN_SECRET` and `JWT_REFRESH_TOKEN_SECRET` to strong, unique values in production
- Use `MAIN_DB_SSL_MODE=require` in production environments
- Set strong passwords for database and Redis

## Database Management

### Migrations

Run all pending migrations:

```bash
make migrate-up db=main_db
```

Rollback the last migration:

```bash
make migrate-down db=main_db
```

Force migration to a specific version:

```bash
make migrate-force version=1 db=main_db
```

**Available Migrations:**
1. `00000000000001_init_iam_tables` - Creates users, roles, permissions, user_roles, role_permissions tables
2. `00000000000002_init_notification_table` - Creates notifications table
3. `00000000000003_add_user_token_table` - Creates user_tokens table for refresh token storage

### Seeders

Run all seeders:

```bash
make seed db=main_db
```

Run a specific seeder:

```bash
make seed module=user db=main_db
make seed module=role db=main_db
make seed module=permission db=main_db
```

**Available Seeders:**
- `user` - Creates default admin user (admin@admin.com)
- `role` - Seeds default roles (Super Admin)
- `permission` - Seeds system permissions (all permissions: *)
- `role_permission` - Assigns permissions to roles
- `user_role` - Assigns roles to users

## Running the Application

### Development Mode

**With hot reload (recommended):**

```bash
make air
```

Air will automatically rebuild and restart the application when you make changes to `.go` files.

**Standard mode:**

```bash
make dev
```

### Background Job Worker

The job worker processes async tasks like token cleanup. Run it in a separate terminal:

```bash
make job-worker
```

### Production Mode

Build the binary:

```bash
go build -o bin/iam-service main.go
```

Run the service:

```bash
./bin/iam-service
```

Build and run the job worker:

```bash
go build -o bin/job-worker interfaces/job_worker/main.go
./bin/job-worker
```

### Makefile Commands

```bash
# Development
make dev           # Run development server (no hot reload)
make air           # Run with hot reload (recommended)
make air-init      # Initialize Air configuration

# Background Jobs
make job-worker    # Run background job worker

# Database Migrations
make migrate-up db=main_db      # Run all migrations up
make migrate-down db=main_db    # Rollback last migration
make migrate-force version=1 db=main_db  # Force to version

# Database Seeders
make seed db=main_db              # Run all seeders
make seed module=user db=main_db  # Run specific seeder
```

## API Documentation

📚 **Coming Soon:** Comprehensive API documentation with request/response examples will be available via [Apidog](https://apidog.com).

## Project Structure

```
go-fiber-iam/
├── constants/              # Application constants
│   ├── app_env.go         # Environment constants
│   ├── db_list.go         # Database identifiers
│   ├── http_ctx.go        # HTTP context keys
│   ├── pagination.go      # Pagination defaults
│   └── permission.go      # Permission constants
│
├── infrastructure/        # Core infrastructure layer
│   ├── config/           # Configuration management
│   │   ├── global.go     # Global config (PORT, APP_ENV)
│   │   ├── jwt.go        # JWT configuration
│   │   ├── main_db.go    # Database configuration
│   │   ├── redis.go      # Redis configuration
│   │   └── scheduler.go  # Scheduler configuration
│   │
│   ├── databases/        # Database connections
│   │   ├── maindb/       # Main database setup
│   │   └── helpers/      # Database helpers (PostgreSQL)
│   │
│   ├── dtos/             # Data Transfer Objects
│   │   └── params/       # Request parameters (pagination)
│   │
│   ├── errors/           # Custom error types
│   ├── integrations/     # External integrations
│   │   └── asyncq/       # Asynq (job queue) integration
│   │
│   ├── messages/         # Response messages
│   ├── types/            # Custom types
│   ├── utils/            # Utility functions
│   └── validators/       # Custom validators
│
├── interfaces/           # Application entry points
│   ├── cli/             # CLI commands
│   │   ├── migrate/     # Migration CLI
│   │   └── seeder/      # Seeder CLI
│   │
│   ├── http/            # HTTP interface
│   │   ├── middlewares/ # HTTP middlewares
│   │   │   ├── auth_bearer.go    # JWT authentication
│   │   │   ├── dto_validator.go  # DTO validation
│   │   │   └── error_handler.go  # Error handling
│   │   │
│   │   ├── modules/iam/ # IAM module handlers
│   │   │   ├── auth/       # Auth handlers & routes
│   │   │   ├── user/       # User handlers & routes
│   │   │   ├── role/       # Role handlers & routes
│   │   │   └── permission/ # Permission handlers & routes
│   │   │
│   │   └── app_route.go # Main route registration
│   │
│   └── job_worker/      # Background job worker
│       ├── modules/     # Job handlers
│       └── worker/      # Worker initialization
│
├── migrations/          # Database migrations
│   └── main_db/        # Main database migrations
│       ├── 00000000000001_init_iam_tables.up.sql
│       ├── 00000000000001_init_iam_tables.down.sql
│       ├── 00000000000002_init_notification_table.up.sql
│       ├── 00000000000002_init_notification_table.down.sql
│       ├── 00000000000003_add_user_token_table.up.sql
│       └── 00000000000003_add_user_token_table.down.sql
│
├── models/             # Database models
│   └── maindb/
│       ├── base.go     # Base model (UUID, timestamps, soft delete)
│       └── iam/        # IAM models
│           ├── user.go
│           ├── role.go
│           ├── permission.go
│           ├── user_role.go
│           ├── role_permission.go
│           └── user_token.go
│
├── modules/            # Business logic modules
│   └── iam/
│       ├── auth/       # Authentication module
│       │   ├── dtos/          # Request/Response DTOs
│       │   │   ├── params/    # Request parameters
│       │   │   └── results/   # Response DTOs
│       │   ├── jobs/          # Background jobs
│       │   │   └── clear_expired_token.go
│       │   ├── messages/      # Error messages
│       │   ├── repositories/  # Data access layer
│       │   │   ├── auth_query_repository.go
│       │   │   └── auth_store_repository.go
│       │   └── services/      # Business logic
│       │       └── auth_service.go
│       │
│       ├── user/       # User management module
│       │   ├── dtos/
│       │   ├── messages/
│       │   ├── repositoires/  # Data access layer
│       │   └── services/
│       │
│       ├── role/       # Role management module
│       │   ├── dtos/
│       │   ├── messages/
│       │   ├── repositories/
│       │   └── services/
│       │
│       └── permission/ # Permission management module
│           ├── dtos/
│           ├── messages/
│           ├── repositories/
│           └── services/
│
├── seeders/            # Database seeders (JSON files)
│   └── main_db/
│       ├── users.json
│       ├── roles.json
│       ├── permissions.json
│       ├── role_permissions.json
│       └── user_roles.json
│
├── .air.toml           # Air hot reload configuration
├── .env.example        # Environment variables template
├── .gitignore          # Git ignore patterns
├── go.mod              # Go module dependencies
├── go.sum              # Dependency checksums
├── main.go             # Application entry point
└── makefile            # Build and run commands
```

**Architecture Patterns:**

- **Clean Architecture**: Clear separation of concerns with distinct layers (interfaces, infrastructure, modules, models)
- **Repository Pattern**: Data access is abstracted through repository interfaces
- **Service Layer**: Business logic is isolated in service layers, independent of HTTP handlers
- **Module-based**: Each feature (auth, user, role, permission) is self-contained with its own DTOs, repositories, and services

## License

This project is licensed under the BSD 3-Clause License - see the [LICENSE.md](LICENSE.md) file for details.


---

**Made with Go Fiber** | [Report Issues](https://github.com/yusril-adr/go-fiber-iam/issues) | [Contribute](https://github.com/yusril-adr/go-fiber-iam/pulls)
