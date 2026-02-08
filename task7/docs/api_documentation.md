# Task 7 — Task Manager API Documentation

## Overview
Task 7 extends the Task Manager API with user registration, JWT-based authentication, role-based (admin) routes, and MongoDB-backed data models. This document describes the public and protected endpoints, authentication scheme, data models, and common responses.

## Base URL

```
http://localhost:8080
```

## Authentication
- Auth uses JWT signed with the `JWT_SECRET_KEY` environment variable.
- Tokens are returned by `POST /login` and must be sent in the `Authorization` header as `Bearer <token>` for protected routes.
- Token claims include `username` and `role`. Tokens expire after 24 hours.

## Public Endpoints

### Register
**POST** `/register`

Registers a new user. The first registered user is automatically granted the `Admin` role; subsequent users receive the `user` role.

Request (application/json):

```json
{
	"username": "alice",
	"password": "secret"
}
```

Responses:
- `201 Created` — user registered
- `400 Bad Request` — invalid payload
- `500 Internal Server Error` — DB error

### Login
**POST** `/login`

Authenticates a user and returns a JWT.

Request (application/json):

```json
{
	"username": "alice",
	"password": "secret"
}
```

Response (application/json):

```json
{
	"token": "<jwt-token>"
}
```

Errors:
- `400 Bad Request` — invalid payload
- `401 Unauthorized` — invalid credentials

## Protected Endpoints (authenticated)
All endpoints below require the `Authorization: Bearer <token>` header.

### Get All Tasks
**GET** `/tasks`

Returns a list of all tasks.

Response: `200 OK` application/json — array of tasks.

### Get Task by ID
**GET** `/tasks/:id`

Path parameter: `id` (MongoDB ObjectID)

Response:
- `200 OK` — task object
- `404 Not Found` — task not found

## Admin-only Endpoints (role = Admin)
The following routes are restricted to users with the `Admin` role.

### Create Task
**POST** `/tasks`

Request (application/json):

```json
{
	"title": "New Task",
	"description": "Do something",
	"due_date": "2024-01-20T15:00:00Z",
	"status": "Pending"
}
```

Responses:
- `201 Created` — created task
- `400 Bad Request` — validation error

### Update Task
**PUT** `/tasks/:id`

Request (application/json): same fields as create (partial fields may be supported depending on controller implementation).

Responses:
- `200 OK` — updated
- `400 Bad Request` — invalid payload
- `404 Not Found` — task not found

### Delete Task
**DELETE** `/tasks/:id`

Responses:
- `200 OK` — deleted
- `404 Not Found` — task not found

### Promote User to Admin
**PATCH** `/promote/:username`

Promotes the given `username` to the `Admin` role.

Responses:
- `200 OK` — promoted
- `404 Not Found` / `500 Internal Server Error` — on failure

## Data Models

### Task

Fields (Go model in `models/task.go`):

```json
{
	"id": "string (MongoDB ObjectID)",
	"title": "string",
	"description": "string",
	"due_date": "string (ISO 8601)",
	"status": "string"
}
```

### User

Fields (Go model in `models/user.go`):

```json
{
	"id": "string (MongoDB ObjectID)",
	"username": "string",
	"password": "string (hashed)",
	"role": "string (user|Admin)"
}
```

## Errors
API uses standard HTTP status codes and JSON error bodies:

```json
{ "error": "error message" }
```

Common codes:
- `400 Bad Request` — invalid JSON or validation failure
- `401 Unauthorized` — missing/invalid/expired token
- `403 Forbidden` — insufficient role for endpoint
- `404 Not Found` — resource not found
- `500 Internal Server Error` — DB or server error

## Auth Notes & Implementation Caveats
- Tokens are expected in `Authorization: Bearer <token>`.
- Token claims include `username` and `role` (check `models/claim.go`).
- Ensure `JWT_SECRET_KEY` is set in the environment before running the server, otherwise signature verification will fail.

## Running
Set env and run the server from the `task7` folder:

```bash
export JWT_SECRET_KEY=your_secret_here
go run .
```

---