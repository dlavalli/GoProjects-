# REST API with Gin Framework

This is a simple REST API built using the Gin Framework. It provides CRUD operations for a user resource.

## Project Structure

The project follows the standard Go project layout:

```
myapp/
|___ cmd/
|    |___myapp/                 # Main application entry point
|        |____ main.go
|___ internal/                  # Private application and library code
|    |___ auth/                 # Authentication logic
|    |___ db/                   # Database access code
|    |___ handlers/             # HTTP request handlers
|    |___ middleware/           # Custom middleware
|___ pkg/                       # Public library code
|    |___ utils/                # Utility functions
|___ api/                       # API definitions
|___ configs/                   # Configuration files
|___ deployments/               # Deployment configurations (Docker, etc.)
|___ scripts/                   # Scripts for setup, migration, etc.
|___ web/                       # Web frontend
|___ test/                      # Additional test data and utilities
|___ build/                     # Build artifacts
```

## API Endpoints

The API provides the following endpoints:

### Health Check

- `GET /health`: Returns the health status of the API

### Hello World

- `GET /api/v1/hello`: Returns a simple hello message

### Users

- `GET /api/v1/users`: Returns a list of all users
- `GET /api/v1/users/:id`: Returns a specific user by ID
- `POST /api/v1/users`: Creates a new user
- `PUT /api/v1/users/:id`: Updates an existing user
- `DELETE /api/v1/users/:id`: Deletes a user

## Running the Application

### Prerequisites

- Go 1.24 or higher

### Local Development

1. Clone the repository
2. Navigate to the project directory
3. Run the application:

```bash
go run cmd/myapp/main.go
```

The API will be available at http://localhost:8080

### Using Docker

1. Build the Docker image:

```bash
docker build -f deployments/Dockerfile -t myapp .
```

2. Run the Docker container:

```bash
docker run -p 8080:8080 myapp
```

The API will be available at http://localhost:8080

## Testing

Run the tests with:

```bash
go test ./...
```

## API Examples

### Create a User

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

### Get All Users

```bash
curl http://localhost:8080/api/v1/users
```

### Get User by ID

```bash
curl http://localhost:8080/api/v1/users/1
```

### Update a User

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "John Updated", "email": "john.updated@example.com"}'
```

### Delete a User

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```