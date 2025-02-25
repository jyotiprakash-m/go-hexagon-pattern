# Todo App - Hexagonal Architecture (Go + Gin + PostgreSQL)

## 📌 Overview

This is a **Todo Application** built using **Golang** with **Hexagonal Architecture**, **Gin Web Framework**, and **PostgreSQL** as the database. The project follows **clean architecture principles**, ensuring modularity, maintainability, and scalability.

## 🏗️ Project Structure

```
├── cmd
│   ├── main.go              # Application entry point
│
├── internal
│   ├── adapters
│   │   ├── db
│   │   │   ├── postgres.go  # Database connection
│   │   ├── repository
│   │       ├── todo_repository.go  # Implements repository interface
│   │
│   ├── core
│   │   ├── models
│   │   │   ├── todo.go       # Todo entity
│   │   ├── ports
│   │   │   ├── todo_repository.go  # Repository interface
│   │   ├── service
│   │       ├── todo_service.go  # Business logic layer
│   │
│   ├── handlers
│   │   ├── todo_handler.go  # API Handlers
│
│   ├── migrations
│   │   ├── migrate.go       # Database migrations
│
│   ├── router
│   │   ├── router.go        # API routes
│
├── .env                     # Environment variables
├── go.mod                    # Go module dependencies
├── README.md                 # Documentation
```

---

## 🚀 Getting Started

### 1️⃣ Prerequisites

- Go 1.18+
- PostgreSQL
- Gin Framework

### 2️⃣ Clone the Repository

```sh
git clone https://github.com/yourusername/todo-app.git
cd todo-app
```

### 3️⃣ Configure Environment Variables

Create a `.env` file and add:

```sh
DATABASE_URL=postgres://postgres:password@localhost:5432/todo?sslmode=disable
```

### 4️⃣ Install Dependencies

```sh
go mod tidy
```

### 5️⃣ Run Migrations

```sh
go run cmd/main.go
```

### 6️⃣ Start the Server

```sh
go run cmd/main.go
```

Server runs on **http://localhost:3000**

---

## 🔥 API Endpoints

### Create a Todo

```sh
POST /todos
```

**Request Body:**

```json
{
  "title": "Learn Go"
}
```

**Response:**

```json
{
  "id": 1,
  "title": "Learn Go",
  "completed": false,
  "created_at": "2025-02-25T10:00:00Z"
}
```

### Get All Todos

```sh
GET /todos
```

### Delete Todo

```sh
DELETE /todos/:id
```

### Get TodoById

```sh
GET /todos/:id
```

### Update Todo

```sh
PUT /todos/:id
```

**\* Request Body:**

```json
{
  "id": 1,
  "title": "Learn Go",
  "completed": false
}
```

**Response:**

```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "completed": false
  }
]
```

### Get Todo by ID

```sh
GET /todos/:id
```

**Response:**

```json
{
  "id": 1,
  "title": "Learn Go",
  "completed": false
}
```

### Delete a Todo

```sh
DELETE /todos/:id
```

---

## 🛠️ Tech Stack

- **Golang** - Backend
- **Gin** - Web Framework
- **PostgreSQL** - Database
- **GORM** - ORM
- **Hexagonal Architecture** - Clean & Modular Design

---

## 📌 Author

- **Your Name**
- GitHub: [@yourusername](https://github.com/yourusername)
