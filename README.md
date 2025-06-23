# Wowza API

This is a backend service for the Wowza application, built with Go. It provides a RESTful API for user authentication, password management, and other core functionalities.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Usage](#usage)
  - [Running the Application](#running-the-application)
  - [Makefile Commands](#makefile-commands)
- [API Documentation](#api-documentation)
  - [Swagger](#swagger)
  - [Postman](#postman)
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)

## Features

- User sign-up and sign-in
- Secure password handling with hashing
- Password reset functionality
- JWT-based authentication using Paseto
- RESTful API with detailed logging
- Configuration management for different environments
- Database migrations for PostgreSQL
- Caching layer with Dragonfly (Redis compatible)

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.20 or higher)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [swag](https://github.com/swaggo/swag)

### Installation

1.  **Clone the repository:**

    ```sh
    git clone https://github.com/nordew/wowza/
    cd wowza
    ```

2.  **Install Go dependencies:**

    ```sh
    go mod tidy
    ```

3.  **Install developer tools:**
    ```sh
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

### Configuration

The application uses a `.env` file for environment variables. Create a `.env` file in the root of the project by copying the example:

```sh
cp .env.example .env
```

Then, update the `.env` file with your configuration.

**Environment Variables:**

| Variable               | Description                             | Default Value     |
| ---------------------- | --------------------------------------- | ----------------- |
| `POSTGRES_HOST`        | PostgreSQL host                         | `localhost`       |
| `POSTGRES_PORT`        | PostgreSQL port                         | `5432`            |
| `POSTGRES_USER`        | PostgreSQL username                     | `user`            |
| `POSTGRES_PASSWORD`    | PostgreSQL password                     | `password`        |
| `POSTGRES_DB`          | PostgreSQL database name                | `wowza`           |
| `DRAGONFLY_HOST`       | DragonflyDB/Redis host                  | `localhost`       |
| `DRAGONFLY_PORT`       | DragonflyDB/Redis port                  | `6379`            |
| `PASETO_SYMMETRIC_KEY` | 32-byte key for Paseto token encryption | `your-secret-key` |
| `HTTP_SERVER_PORT`     | Port for the HTTP server to listen on   | `8080`            |

## Usage

### Running the Application

1.  **Start the services (PostgreSQL and Dragonfly):**

    ```sh
    docker-compose up -d
    ```

2.  **Apply database migrations:**

    ```sh
    make migrate-up
    ```

3.  **Run the application:**

    ```sh
    make run
    ```

The API will be available at `http://localhost:8080`.

### Makefile Commands

This project uses a `Makefile` to streamline common tasks.

| Command          | Description                                                                   |
| ---------------- | ----------------------------------------------------------------------------- |
| `run`            | Run the application.                                                          |
| `build`          | Build the application binary.                                                 |
| `swagger`        | Generate Swagger documentation.                                               |
| `migrate-up`     | Apply all 'up' database migrations.                                           |
| `migrate-down`   | Apply all 'down' database migrations.                                         |
| `migrate-create` | Create a new migration file. (e.g., `make migrate-create name=add_new_table`) |
| `help`           | Show all available commands.                                                  |

## API Documentation

### Swagger

This project uses Swagger for API documentation.

1.  **Generate the Swagger docs:**

    ```sh
    make swagger
    ```

2.  **Access the documentation:**

    After running the application, the Swagger UI will be available at `http://localhost:8080/swagger/index.html`.

### Postman

A Postman collection is available in the root of the project: `wowza_postman_collection.json`. You can import this into Postman to easily test the API endpoints.

## Project Structure

The project follows a standard Go project layout:

```
├── cmd/api/            # Main application entrypoint
├── configs/            # Configuration files (e.g., config.yml)
├── docs/               # Swagger documentation files
├── internal/           # Private application logic
│   ├── app/            # Application setup and initialization
│   ├── cache/          # Caching logic (Dragonfly/Redis)
│   ├── config/         # Configuration loading
│   ├── dto/            # Data Transfer Objects
│   ├── entity/         # Core business entities
│   ├── handler/http/   # HTTP handlers and routing
│   ├── service/        # Business logic services
│   └── storage/        # Database interactions
├── migrations/         # Database migration files
├── pkg/                # Reusable packages
│   ├── db/             # Database clients (Postgres, Dragonfly)
│   ├── generator/      # Code/token generation utilities
│   ├── hash/           # Password hashing utilities
│   ├── logger/         # Logging setup
│   └── paseto/         # Paseto token management
├── Makefile            # Makefile for common commands
└── docker-compose.yml  # Docker Compose for development services
```

## Technologies Used

- **Language:** [Go](https://golang.org/)
- **Framework:** [Fiber](https://gofiber.io/)
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **Cache:** [Dragonfly](https://www.dragonflydb.io/) (Redis compatible)
- **Authentication:** [Paseto](https://paseto.io/)
- **Migrations:** [golang-migrate](https://github.com/golang-migrate/migrate)
- **API Documentation:** [Swagger](https://swagger.io/)
- **Containerization:** [Docker](https://www.docker.com/)
