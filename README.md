# Go Gin Gorm Swagger (FGA & Hacktiv8 Assigment 2)

## Requirements

- Go 1.x (refer to [installation instructions](https://go.dev/doc/install))
- PostgreSQL database server (refer to [installation instructions](https://www.postgresql.org/))

You can also use docker for this project (refer to [installation instructions](https://www.docker.com/))

## Getting Started

### Clone the repository

```bash
git clone https://github.com/ridhoafwani/gingormpostgres.git
```

### Set up the database
Configure your local PostgreSQL database or a dedicated database instance.

### Database Schema Migrations

The project utilizes migrations for managing schema changes in the database.

Navigate to the project directory:

```bash
cd gingormpostgres
```

Run migrate.go:

```bash
go run migration/migrate.go
```

### Run (Development)
Navigate to the project directory:

```bash
cd gingormpostgres
```

Then run using:

```bash
go run main.go
```

### Build and run (optional)
Navigate to the project directory:

```bash
cd gingormpostgres
```

Build the go binnary:

```bash
go build -o main .
```

Run the application

```bash
./main
```

## VS Code Dev Containers

1. Open the project in VS Code.
2. Install the "Remote-Containers" extension if not already installed.
3. Press F1 and select "Remote-Containers: Open Folder in Container".
4. The development container will be built and started, providing a ready-to-use environment for development.

## GitHub Codespaces

1. Go to your project repository on GitHub.
2. Click on the "Code" button and choose "Open with Codespaces".
3. Codespaces will automatically set up a development environment in the cloud.

## API Documentation

The project includes Swagger documentation for exploring API endpoints and request/response structures. Access the Swagger UI at `http://localhost:<PORT>/swagger/index.html` while the application is running (replace `<PORT>` with the actual port number).
