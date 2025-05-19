# Blockchain Monitor (BCM)

A Go application that monitors block heights across different blockchain networks and stores the data in ClickHouse.

## Features

- Concurrent monitoring of multiple blockchain endpoints
- Configurable polling intervals per endpoint
- Automatic database migrations
- Structured logging
- YAML-based configuration

## Prerequisites

- Go 1.24.3 or later
- ClickHouse database server
- Access to blockchain RPC endpoints

## Installation

1. Clone the repository
```bash
git clone <git@github.com:ibrahim-dwellir/bcm-practice.git>
cd bcm
```

2. Install dependencies
```bash
go mod download
```

## Configuration

### Environment Variables

Create a `.env` file in the root directory with the following variables:

```properties
DATABASE_HOST=<clickhouse-host>:<port>
DATABASE_NAME=<database-name>
DATABASE_USER=<username>
DATABASE_PASSWORD=<password>
LOG_LEVEL=info  # Optional: debug, info, warn, error (defaults to info)
```

### Endpoint Configuration

Create a (or modify) `config.yaml` file in the root directory:

```yaml
interval: 2  # Default polling interval in seconds
jobs:
  - name: ethereum-sepolia
    endpoint: https://your-ethereum-endpoint
  - name: arbitrum-sepolia
    endpoint: https://your-arbitrum-endpoint
    interval: 5  # Override default interval for specific endpoint
```

## Database Schema

The application creates a `block_heights` table with the following schema:

```
name string
height int64
timestamp datetime
```

## Running the Application

```bash
go run main.go
```

## Building

```bash
go build -o bcm
```

## Project Structure

```
.
├── config.yaml       # Configuration file
├── .env             # Environment variables
├── main.go          # Application entry point
├── database/        # Database connection and migrations
├── job/            # Job execution logic
├── logger/         # Logging configuration
└── timer/          # Timer implementation
```

## Logging

The application uses structured logging with the following levels:
- debug: Detailed information for debugging
- info: General operational information
- warn: Warning messages for potentially harmful situations
- error: Error messages for serious problems
- fatal: Critical errors that stop the program

Set the `LOG_LEVEL` environment variable to control logging verbosity.
