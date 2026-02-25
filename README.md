# gin-project

A simple REST API built with Go and Gin. This is a learning/sandbox project to explore Go development.

## Stack

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)

## Project Structure

```
.
├── main.go
├── main_test.go
├── go.mod
├── go.sum
└── internal/
    ├── router/
    │   └── router.go     # Route definitions
    └── handler/
        └── hello.go      # Request handlers
```

## Endpoints

| Method | Path     | Description        |
|--------|----------|--------------------|
| GET    | `/`      | Health check       |
| GET    | `/hello` | Returns hello message |

## Running

```bash
go run main.go
```

The server starts on `http://localhost:8080`.

## Testing

```bash
go test ./...
```
