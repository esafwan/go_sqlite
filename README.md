# Go SQLite REST API Example

This repository contains an example of a RESTful API implemented in Go, using SQLite as the database. It demonstrates basic CRUD (Create, Read, Update, Delete) operations on a user entity and showcases quick and efficient setup for REST APIs in Go.

## Features

- Basic CRUD operations for a user entity.
- SQLite database integration.
- Pagination support for list queries.
- RESTful endpoints accessible via tools like Postman.

## Getting Started

### Prerequisites

- Go (version 1.x)
- SQLite3

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/esafwan/go_sqlite

2. Navigate to the project directory:
   ```bash
   cd your-repo-name
   ```
3. Install dependencies (if any):
   ```bash
   go mod tidy
   ```

### Running the Application

1. Start the server:
   ```bash
   go run main.go
   ```
2. The server will start on `localhost:8080`. You can interact with the API using tools like Postman or cURL.

## API Endpoints

### User CRUD Operations

- **List Users (GET /users)**
  - Supports pagination with query parameters `page` and `perPage`.
- **Add User (POST /users)**
  - Accepts JSON body with `name`, `age`, and `class`.
- **Edit User (PUT /users/{id})**
  - Accepts JSON body with `name`, `age`, and `class`.
- **Delete User (DELETE /users/{id})**
  - Deletes the user with the specified ID.

## Project Structure

- `main.go`: The entry point of the application, setting up the server and routes.
- `user/user.go`: Contains the user model and handlers for CRUD operations.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the [MIT License](LICENSE).
