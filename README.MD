# Go Project Template

A clean and structured Go project template using Fiber framework.

## 📁 Project Structure

```bash
.
├── app/ # Application bootstrapping
│ ├── application.go # Main application setup
│ ├── database.go # Databases setup  
│ ├── logger.go # Logger setup  
│ ├── redis.go # Redis setup  
│ └── task.go # Task management setup  
│
├── cmd/ # Command line interface
│ ├── root.go # Root cobra command
│ └── serve.go # Server start command
│
├── docs/ # Documentation
│
├── config/ # Configuration
│ ├── config.example.yml # Configuration file example
│ └── config.yml # Configuration file
│
│
├── internal/ # Private application code
│ ├── config/ # Configuration
│ ├── database/ # Databases
│ ├── ent/ # Ent schema files
│ ├── mocks/ # Mock files
│ ├── producers/ # Producers
│ ├── task/ # Tasks
│ ├── utils/ # Utility functions
│ └── helper/ # Helper functions
│
├── pkg/ # Shared packages
│
└── main.go # Application entry point
```

## 🔧 Technology Stack

- [Ants](https://github.com/panjf2000/ants) - High-performance goroutine pool
- [Uber FX](https://uber-go.github.io/fx/) - Dependency injection
- [Viper](https://github.com/spf13/viper) - Configuration management
- [Cobra](https://github.com/spf13/cobra) - CLI commands

## 🏗️ Architecture

This template follows clean architecture principles:

1. **Domain Layer** (`internal/tasks/`)

   - Cronjob entities (e.g., Cronjob, Schedule, Task)
   - Repository interfaces for cronjob-related data
   - Service interfaces for cronjob-related operations

2. **Application Layer** (`internal/tasks/`)

   - Business logic for cronjob management
   - Use case orchestration for scheduling, execution, and monitoring
   - Domain service implementation for cronjob workflows

3. **Infrastructure Layer** (`internal/repositories/`)

   - Database implementations for cronjob-related data
   - External service integrations (e.g., monitoring tools, notification systems)
   - Repository implementations for cronjob data persistence


## 🚀 Getting Started

### installation and run

1. install dependencies

   ```bash
   go mod tidy
   ```

2. configure environment variables in config.yml

   ```bash
   nano config/config.yml.example
   nano config/config.yml
   ```

3. run the server

   ```bash
   go run main.go serve
   ```

4. Also you can set environment variable mappings by writing the following command before executing command `go run main.go serve`:

   ```bash
   APP_SERVER_PORT=8004 APP_SERVER_HOST=0.0.0.0 go run main.go serve
   ```

## 📚 ArangoDB

To see the instructions on how to work with ArangoDB in this application : [ArangoDB Documentation](./docs/arangodb.md)
