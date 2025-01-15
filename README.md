# JWT Authentication in Go

A simple and secure JWT (JSON Web Token) authentication service implemented in Go. This project demonstrates how to implement token-based authentication with features like login, protected routes, and token management.

## Features

- JWT-based authentication system
- Secure token generation and validation
- Cookie-based token storage
- Configurable token expiration
- Structured logging using slog

## Prerequisites

- Go 1.23.3 or higher
- github.com/golang-jwt/jwt/v5 package

## Installation

1. Clone the repository:
```bash
git clone https://github.com/minhaz11/jwt-auth.git
cd jwt-auth
```
2. Install dependencies:
```bash
go mod tidy
```

## Usage
 1. Build and run the project using Make:
```bash
make run
```
Or manually:
```bash
go build -o bin/auth
./bin/auth
```
or
```bash
go run main.go
```
### API Endpoints
#### POST /login
Authenticates a user and issues a JWT token.

#### Request Body:
```json
{
    "username": "user1",
    "password": "password1"
}
```
#### Response
    Success: sets a cookie with the JWT token "_token"
    Failure: returns 401 Unauthorized

### GET /home
Protected route that requires a valid JWT token.

#### Request Headers:
    Requires valid JWT token "_token" in the cookie

#### Response
    Success: returns a welcome message
    Failure: returns 401 Unauthorized



