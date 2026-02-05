# Task Manager API Documentation

## Overview
This API provides endpoints for managing tasks in a task management system. All endpoints return JSON responses.

## Base URL
```
http://localhost:8080
```

## Endpoints

### 1. Get All Tasks
**GET** `/tasks`

Retrieves a list of all tasks.

#### Response
- **Status Code:** 200 OK
- **Content-Type:** application/json

```json
{
  "task": [
    {
      "id": "1",
      "title": "Task 1",
      "description": "First task",
      "due_date": "2024-01-15T10:30:00Z",
      "status": "Pending"
    },
    {
      "id": "2",
      "title": "Task 2",
      "description": "Second task",
      "due_date": "2024-01-16T10:30:00Z",
      "status": "In Progress"
    }
  ]
}
```

---

### 2. Get Task by ID
**GET** `/tasks/{id}`

Retrieves a specific task by its ID.

#### Parameters
- `id` (path parameter): The unique identifier of the task

#### Response
- **Status Code:** 200 OK
- **Content-Type:** application/json

```json
{
  "id": "1",
  "title": "Task 1",
  "description": "First task",
  "due_date": "2024-01-15T10:30:00Z",
  "status": "Pending"
}
```

#### Error Response
- **Status Code:** 404 Not Found

```json
{
  "error": "Task not found"
}
```

---

### 3. Create Task
**POST** `/tasks`

Creates a new task.

#### Request
- **Content-Type:** application/json

```json
{
  "id": "4",
  "title": "New Task",
  "description": "Description of the new task",
  "due_date": "2024-01-20T15:00:00Z",
  "status": "Pending"
}
```

#### Response
- **Status Code:** 201 Created
- **Content-Type:** application/json

Returns the created task object.

```json
{
  "id": "4",
  "title": "New Task",
  "description": "Description of the new task",
  "due_date": "2024-01-20T15:00:00Z",
  "status": "Pending"
}
```

#### Error Response
- **Status Code:** 400 Bad Request

```json
{
  "error": "invalid JSON format"
}
```

---

### 4. Update Task
**PUT** `/tasks/{id}`

Updates an existing task by its ID.

#### Parameters
- `id` (path parameter): The unique identifier of the task to update

#### Request
- **Content-Type:** application/json

```json
{
  "title": "Updated Task Title",
  "description": "Updated description",
  "due_date": "2024-01-25T12:00:00Z",
  "status": "Completed"
}
```

#### Response
- **Status Code:** 200 OK
- **Content-Type:** application/json

```json
{
  "message": "Task update successfully"
}
```

#### Error Responses
- **Status Code:** 400 Bad Request

```json
{
  "error": "invalid JSON format"
}
```

- **Status Code:** 404 Not Found

```json
{
  "error": "task not found"
}
```

---

### 5. Delete Task
**DELETE** `/tasks/{id}`

Deletes a task by its ID.

#### Parameters
- `id` (path parameter): The unique identifier of the task to delete

#### Response
- **Status Code:** 200 OK
- **Content-Type:** application/json

```json
{
  "message": "Task deleted successfully"
}
```

#### Error Response
- **Status Code:** 404 Not Found

```json
{
  "error": "task not found"
}
```

## Data Models

### Task
```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "due_date": "string (ISO 8601 date-time format)",
  "status": "string"
}
```

#### Field Descriptions
- `id`: Unique identifier for the task
- `title`: Title of the task
- `description`: Detailed description of the task
- `due_date`: Due date and time in ISO 8601 format (e.g., "2024-01-15T10:30:00Z")
- `status`: Current status of the task (e.g., "Pending", "In Progress", "Completed")

## Error Handling
The API uses standard HTTP status codes and returns error messages in JSON format:

```json
{
  "error": "error message description"
}
```

Common error codes:
- `400 Bad Request`: Invalid request data or malformed JSON
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server-side error


