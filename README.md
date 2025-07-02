
# 🎬 Go Movies REST API

A simple REST API written in Golang using the Gorilla Mux router to manage a collection of movies with basic CRUD operations.

## 📦 Features

- List all movies
- Retrieve a movie by ID
- Create a new movie
- Update an existing movie
- Delete a movie

## 🛠️ Tech Stack

- Golang
- Gorilla Mux (`github.com/gorilla/mux`)
- net/http
- JSON encoding/decoding

## 📁 Project Structure

```
├── CRUDAPI
├── go.mod
├── go.sum
├── main.go
└── README.md
``` 

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install)

---

### 🔧 Run Locally

```bash
# Download dependencies
go mod tidy

# Run the server
go run main.go
```
#### The server will start on: http://localhost:8000

## 📬 API Endpoints

| Method | Endpoint       | Description           |
| ------ | -------------- | --------------------- |
| GET    | `/movies`      | Get all movies        |
| GET    | `/movies/{id}` | Get a single movie    |
| POST   | `/movies`      | Add a new movie       |
| PUT    | `/movies/{id}` | Update existing movie |
| DELETE | `/movies/{id}` | Delete a movie        |


## 📝 Notes

#### This is an in-memory API. All data is lost when the server restarts.
