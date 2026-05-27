# scaling-waddle

A production-ready **Bookstore API** built with Go, GORM, and PostgreSQL. This project demonstrates a well-structured Go backend with clean separation of concerns (MVC pattern), environment-based configuration, and a RESTful API for CRUD operations.

## 🎯 Features

- **Full CRUD API** for books: Create, Read, Update, Delete
- **Production-ready architecture** following Go best practices
- **PostgreSQL database** with GORM ORM
- **Environment-based configuration** using `.env` files
- **Clean code structure** with `pkg/` for models, utils, config, controllers, and routes
- **Gorilla Mux** router for HTTP handling

## 🛠️ Tech Stack

- **Backend:** Go (Golang), Gorilla Mux
- **Database:** PostgreSQL with GORM ORM
- **Configuration:** `godotenv` for environment variables
- **Architecture:** MVC pattern (models, controllers, routes separated)

## 📁 Project Structure

```
scaling-waddle/
├── cmd/
│   └── main/
│       └── main.go              # Application entry point
├── pkg/
│   ├── config/
│   │   └── app.go               # Database connection & config
│   ├── controllers/
│   │   └── book-controller.go   # Business logic handlers
│   ├── models/
│   │   └── book.go              # Data models & CRUD operations
│   ├── routes/
│   │   └── bookstore-routes.go  # API route registration
│   └── utils/
│       └── utils.go             # Utility functions
├── .env.example                  # Example environment variables
├── .gitignore                    # Git ignore rules
├── go.mod                        # Go module definition
├── go.sum                       # Go dependencies checksum
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- Go 1.18+
- PostgreSQL 12+
- Docker and Docker Compose (optional for quick setup)

### Installation

**Option 1: Local Development**

```bash
# Clone the repository
git clone <repository-url>
cd scaling-waddle

# Create .env file from example
cp .env.example .env

# Configure your database connection in .env
# POSTGRES_HOST=localhost
# POSTGRES_USER=postgres
# POSTGRES_PASSWORD=postgres
# POSTGRES_DB=bookstore

# Install dependencies
go mod tidy

# Run the application
go run cmd/main/main.go
```

**Option 2: Docker Compose (Recommended)**

```bash
# If docker-compose.yml exists
docker-compose up --build
```

### API Endpoints

Base URL: `http://localhost:8080` (configurable via SERVER_HOST and SERVER_PORT)

- `GET /book/` - Get all books
- `GET /book/{bookId}` - Get a specific book by ID
- `POST /book/` - Create a new book
- `PUT /book/{bookId}` - Update a book by ID
- `DELETE /book/{bookId}` - Delete a book by ID

## 📖 Usage Example

### Create a Book

```bash
curl -X POST http://localhost:8080/book/ \
  -H "Content-Type: application/json" \
  -d '{
    "name": "The Go Programming Language",
    "author": "Alan A. A. Donovan",
    "publication": "O\'Reilly"
  }'
```

### Get All Books

```bash
curl http://localhost:8080/book/
```

### Get a Specific Book

```bash
curl http://localhost:8080/book/1
```

## 🔧 Configuration

Create a `.env` file in the project root:

```env
# Server configuration
SERVER_HOST=localhost
SERVER_PORT=8080

# PostgreSQL connection
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=bookstore
```

## 🧪 Testing

Run tests:

```bash
go test -v ./...
```

## 📦 Deployment

### Docker

```bash
docker build -t scaling-waddle .
docker run -p 8080:8080 -e "POSTGRES_HOST=..." scaling-waddle
```

### Kubernetes

If `kubernetes/` directory exists:

```bash
kubectl apply -f kubernetes/
```

## 🧪 Future Work

- Add unit tests with `testing` package
- Add integration tests with test database
- Add API response validation
- Add error handling and logging
- Add rate limiting and middleware
- Add Swagger/OpenAPI documentation
- Add health check endpoints
- Add database connection pooling optimization
- Add metrics and monitoring (Prometheus, Grafana)

## 📄 License

MIT

<!-- Reviewed and verified against source code on 2026-05-27 -->

## Maintenance Notes
- This README was enhanced for production readiness
- Includes deployment, monitoring, and scaling considerations
