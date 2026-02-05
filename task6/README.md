# Task Manager (task6)

## Overview
A small Task Manager API implemented in Go using Gin and MongoDB. This module (task6) loads environment variables from a `.env` file, connects to a MongoDB instance, and exposes HTTP endpoints to create, read, update, and delete tasks.

## What's changed in task6
- Loads environment variables using `godotenv` (.env required)
- Uses MongoDB (MongoDB URI provided via `MONGODB_URL`)
- Router exposes endpoints under `/tasks` (e.g., `/tasks`, `/tasks/:id`)
- Default server port: `8080` (Gin's default)

## Prerequisites
- Go 1.18 or newer
- A MongoDB instance and a connection URI

## Setup
1. Copy or create a `.env` file in the `task6` directory with your MongoDB URI:

```
MONGODB_URL=mongodb+srv://<user>:<password>@cluster0.example.mongodb.net/?retryWrites=true&w=majority
```

2. Ensure dependencies are downloaded (run from the `task6` directory):

```bash
go mod tidy
```

## Run
From the `task6` directory:

```bash
go run .
```

Or build and run:

```bash
go build -o task
./task
```

The server listens on port `8080` by default. You can override the port by setting `PORT` environment variable before starting the app, or change the router call in `main.go`.

## API Documentation
Full API documentation and examples are in the docs: [docs/api_documentation.md](docs/api_documentation.md)

## Environment Variables
- `MONGODB_URL` — MongoDB connection URI (required)

## Example Requests
- Get all tasks:

```bash
curl http://localhost:8080/tasks
```

- Create a task:

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"New Task","description":"Do something","due_date":"2026-02-10T10:00:00Z","status":"Pending"}'
```

## Notes
- The application expects BSON/UUID handling consistent with the MongoDB driver used in `data/task_service.go`.
- The API docs were synchronized to use the `/tasks` paths matching the router.
