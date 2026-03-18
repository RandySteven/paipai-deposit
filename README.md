```
  ________                  ____  __.            .__ 
 /  _____/  ____           |    |/ _|____ ______ |__|
/   \  ___ /  _ \   ______ |      < /  _ \\____ \|  |
\    \_\  (  <_> ) /_____/ |    |  (  <_> )  |_> >  |
 \______  /\____/          |____|__ \____/|   __/|__|
        \/                         \/     |__|       
```

A Go backend framework with clean architecture, supporting multiple databases, caching, and message queues.

## Quick Start

### Ask Barista ☕

Create a new project by asking your barista to brew one:

```bash
# Brew a fresh project
curl -fsSL https://raw.githubusercontent.com/RandySteven/go-kopi/v2/barista | bash -s -- brew -n my-project

# Or with your own git remote
curl -fsSL https://raw.githubusercontent.com/RandySteven/go-kopi/v2/barista | bash -s -- brew -n my-project -r https://github.com/youruser/my-project.git
```

### Manual Installation

```bash
# Clone the repository
git clone https://github.com/RandySteven/go-kopi.git my-project
cd my-project

# Run setup
./barista setup

# Change remote to your own repository
./barista remote -r https://github.com/youruser/my-project.git
```

### Barista Commands

| Command | Description |
|---------|-------------|
| `./barista brew -n <name>` | Brew a fresh project (clone + setup + remote) |
| `./barista clone -n <name>` | Clone to a new project directory |
| `./barista setup` | Set up config files and install dependencies |
| `./barista remote -r <url>` | Change git remote to your own repo |
| `./barista refill` | Refill with latest updates from upstream go-kopi |
| `./barista help` | Show help message |

### Keeping Up to Date

After changing your remote, you can still get refills from the original go-kopi:

```bash
./barista refill
```

This will merge the latest changes while preserving your customizations.

## Tech Stack

| Category        | Technology                                      |
|-----------------|-------------------------------------------------|
| Language        | Go 1.24                                         |
| HTTP Router     | [Gorilla Mux](https://github.com/gorilla/mux)   |
| Database        | MySQL, PostgreSQL                               |
| Cache           | Redis (go-redis/v9)                             |
| Message Queue   | NSQ, Kafka                                      |
| Authentication  | JWT (golang-jwt/v5)                             |
| Logging         | Logrus                                          |
| Scheduler       | Cron (robfig/cron/v3)                           |
| Payment         | Midtrans                                        |
| Email           | Gomail                                          |
| Cloud           | AWS SDK (Secrets Manager, S3)                   |
| Search          | Elasticsearch                                   |
| Maps            | Google Maps API                                 |

## Architecture

This project follows a **layered architecture** pattern:

```
┌─────────────────────────────────────────────────────────────┐
│                        HTTP Request                         │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                     handlers/https                          │
│              (Request parsing, validation)                  │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        usecases                             │
│                   (Business logic)                          │
└─────────────────────────────────────────────────────────────┘
                              │
              ┌───────────────┼───────────────┐
              ▼               ▼               ▼
┌─────────────────┐  ┌─────────────┐  ┌───────────────┐
│   repositories  │  │   caches    │  │   NSQ/Kafka   │
│     (MySQL)     │  │   (Redis)   │  │  (Publish)    │
└─────────────────┘  └─────────────┘  └───────────────┘
```

### Async Processing (Consumers)

```
┌─────────────────────────────────────────────────────────────┐
│                    NSQ Message Queue                        │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                   handlers/consumers                        │
│              (Message handlers by topic)                    │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        usecases                             │
│                   (Business logic)                          │
└─────────────────────────────────────────────────────────────┘
                              │
              ┌───────────────┴───────────────┐
              ▼                               ▼
┌─────────────────────┐            ┌─────────────────┐
│     repositories    │            │     caches      │
│       (MySQL)       │            │     (Redis)     │
└─────────────────────┘            └─────────────────┘
```

## Project Structure

| Directory    | Description                                                        |
|--------------|--------------------------------------------------------------------|
| apps         | Application bootstrap - registers handlers, middlewares, usecases  |
| caches       | Redis cache layer implementations                                  |
| cmd          | Entry points: `main`, `migration`, `drop`, `seed`                  |
| entities     | Domain models, request/response payloads                           |
| enums        | Enums and constants                                                |
| files        | Configuration files (YAML, .env)                                   |
| handlers     | HTTP handlers (`https/`) and message consumers (`consumers/`)      |
| interfaces   | Interface definitions for dependency inversion                     |
| middlewares  | HTTP middleware (auth, logging, CORS, rate limiting)               |
| pkg          | Infrastructure clients (MySQL, Redis, NSQ, etc.)                   |
| queries      | Raw SQL queries                                                    |
| repositories | Data access layer                                                  |
| topics       | NSQ topic definitions                                              |
| usecases     | Business logic layer                                               |
| utils        | Helper functions                                                   |

## Getting Started

### Prerequisites

- Go 1.24+
- MySQL
- Redis
- NSQ (optional, for async processing)

### Configuration

1. Copy environment file:

```bash
cp files/env/.env.example files/env/.env
```

2. Copy YAML config:

```bash
cp files/yaml/app.example.yml files/yaml/app.local.yml
```

3. Update `.env` with your environment:

```env
# prod | dev | test | staging
ENV=dev

SECRET_JWT='your-jwt-secret'
```

4. Update `app.local.yml` with your database and server config:

```yaml
server:
  host: "0.0.0.0"
  port: "8080"
  timeout:
    server: 30
    read: 15
    write: 10
    idle: 5

mysql:
  host: "localhost"
  port: "3306"
  dbname: "your_database"
  dbuser: "your_user"
  dbpass: "your_password"

redis:
  host: "localhost"
  port: "6379"

nsq:
  host: "localhost"
  port: "4150"
```

### Make Commands

| Command           | Description                              |
|-------------------|------------------------------------------|
| `make run`        | Run the application                      |
| `make migration`  | Run database migrations                  |
| `make seed`       | Seed the database                        |
| `make drop`       | Drop all database tables                 |
| `make refresh`    | Drop, migrate, and seed (full reset)     |
| `make make_model` | Generate model and repository files      |
| `make run-docker` | Start with Docker Compose                |
| `make stop-docker`| Stop Docker containers                   |

### Running the Application

```bash
# Set up database
make migration
make seed

# Run the server
make run
```

### Docker

```bash
# Start all services
make run-docker

# Stop all services
make stop-docker
```

## Adding New Features

### 1. Add a New HTTP Handler

1. Create handler in `handlers/https/`
2. Define interface in `interfaces/handlers/`
3. Register in `handlers/https/http.go`
4. Add routes in router configuration

### 2. Add a New Consumer

1. Create consumer in `handlers/consumers/`
2. Define topic in `topics/`
3. Register consumer using `Runners.RegisterConsumer(topic, handler)`

### 3. Add a New Model

```bash
make make_model
# Enter model name when prompted
```

This generates:
- `{model}.go` - Model struct
- `{model}_repository.go` - Repository interface

## Environment Support

| Environment | Config File                          |
|-------------|--------------------------------------|
| dev         | `files/yaml/configs/task.local.yml`  |
| staging     | `files/yaml/configs/task.docker.yml` |
| prod        | `files/yaml/configs/task.prod.yml`   |

Set environment in `.env`:

```env
ENV=dev
```

---

<sub>This documentation was written entirely by AI because the developer was too lazy to write it themselves. You're welcome.</sub>
