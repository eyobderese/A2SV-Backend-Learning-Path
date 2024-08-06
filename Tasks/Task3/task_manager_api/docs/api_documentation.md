# API Documentation

## Endpoints

### GET /tasks

Fetches all tasks.

Full URL: `http://localhost:8080/tasks`

#### Parameters

None

#### Response

- `200 OK` on success

```json
{
  "data": [
    {
      "id": "1",
      "title": "Task 1",
      "description": "First task",
      "due_date": "2024-08-06T15:31:14.5417452+03:00",
      "status": "Pending"
    },
    {
      "id": "2",
      "title": "Task 2",
      "description": "Second task",
      "due_date": "2024-08-07T15:31:14.5417452+03:00",
      "status": "In Progress"
    },
    {
      "id": "3",
      "title": "Task 3",
      "description": "Third task",
      "due_date": "2024-08-08T15:31:14.5427942+03:00",
      "status": "Completed"
    }
  ]
}
```

### GET /tasks/{id}

Fetches a task by its ID.

Full URL: `http://localhost:8080/tasks/{id}`

#### Parameters

- `id` (integer): The ID of the task to fetch

#### Response

- `200 OK` on success

#### Error Code

- `404 NOT FOUND` on fail

```json
{
  "data": {
    "id": "1",
    "title": "Task 1",
    "description": "First task",
    "due_date": "2024-08-06T15:31:14.5417452+03:00",
    "status": "Pending"
  }
}
```

### POST /tasks

Creates a new task.

Full URL: `http://localhost:8080/tasks`

#### Parameters

- `title` (string): The title of the task
- `description` (string): The description of the task
- `due_date` (string): The due date of the task
- `status` (string): The status of the task

Request example:

```json
{
  "title": "New Task",
  "description": "New task description",
  "due_date": "2024-09-06T15:31:14.5417452+03:00",
  "status": "Pending"
}
```

#### Response

`200 OK` on sucess

```json
{
  "data": {
    "id": "4",
    "title": "New Task",
    "description": "New task description",
    "due_date": "2024-09-06T15:31:14.5417452+03:00",
    "status": "Pending"
  }
}
```

`404 Not Found` on error

```json
{
  "error": "Invalid data provided"
}
```

### PUT /tasks/{id}

Updates a task by its ID.

Full URL: `http://localhost:8080/tasks/{id}`

#### Parameters

- `id` (integer): The ID of the task to update
- `title` (string, optional): The new title of the task
- `description` (string, optional): The new description of the task
- `due_date` (string, optional): The new due date of the task
- `status` (string, optional): The new status of the task

Request example:

```json
{
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2024-09-06T15:31:14.5417452+03:00",
  "status": "In Progress"
}
```

#### Response

- `200 OK` on success

```json
{
  "data": {
    "id": "1",
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-09-06T15:31:14.5417452+03:00",
    "status": "In Progress"
  }
}
```

`404 Not Found` on error

```json
{
  "error": "Task not found"
}
```

### DELETE /tasks/{id}

Deletes a task by its ID.

Full URL: `http://localhost:8080/tasks/{id}`

#### Parameters

- `id` (integer): The ID of the task to delete

#### Response

- `204 No Content` on success

- `404 Not Found` if the task with the specified ID does not exist

```json
{
  "error": "Task not found"
}
```
