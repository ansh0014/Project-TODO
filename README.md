# Project-TODO

A simple TODO application built with Go, featuring user registration, login with JWT authentication, and CRUD operations for TODO items.

## Features

- User registration and login with password hashing using bcrypt.
- JWT-based authentication middleware to protect API routes.
- CRUD operations for managing TODO items.
- Environment variable configuration using a .env file.

## Project Structure

- `main.go`: Application entry point, loads environment variables, sets up routes, and starts the server.
- `config/config.go`: Loads environment variables from the .env file.
- `controllers/auth.go`: Handles user registration and login.
- `controllers/controllers.go`: Handles TODO item operations.
- `middleware/jwt.go`: JWT authentication middleware.
- `model/model.go`: Defines User and Todo data models.
- `routes/routes.go`: Defines API routes and applies middleware.

## Setup

1. Create a `.env` file in the project root with the following content:

```
PORT=8080
JWT_SECRET=your_jwt_secret_key
```

2. Run `go mod tidy` to install dependencies.
3. Start the server with `go run main.go`.

## API Endpoints

- `POST /Register`: Register a new user.
- `POST /login`: Login and receive a JWT token.
- `GET /api/todos`: Get all TODO items (requires JWT).
- `POST /api/todos`: Add a new TODO item (requires JWT).
- `PUT /api/todos/{id}`: Mark a TODO item as done (requires JWT).
- `DELETE /api/todos/{id}`: Delete a TODO item (requires JWT).

This project is a good  for Revision of Go web development with authentication and RESTful APIs.
