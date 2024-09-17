# Gin-Tin

A simple REST API built with Gin and GORM.

## Getting Started

### Prerequisites

- Go (1.20 or later)
- Docker
- Docker Compose

### Installation

1. Clone the repository:

```bash
git clone https://github.com/tinwritescode/gin-tin.git
```

2. Navigate to the project directory:

```bash
cd gin-tin
```

```bash
go run cmd/api/main.go
```

### Running the Application

The application will run on `http://localhost:8080`.

### API Endpoints

- `GET /books`: Get all books.
- `POST /books`: Create a new book.
- `DELETE /books/:id`: Delete a book by ID.
