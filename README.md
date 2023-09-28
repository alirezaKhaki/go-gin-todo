# Go-gin with modular dependency injection

Clean architecture template with gin framework, go-fx as dependency container, gorm as orm for database related operations.

## Getting Started

### Prerequisites

- go >= 1.20
- node >= 16 & npm

## Running the project

- Copy `.env.example` to `.env`
- To run in develop mode

```sh
    go mod tidy
    npm i nodemon -g
    ./strat-dev.sh
```

- To run

```sh
go run main.go
```

- Go to `localhost:5000` to verify if the server works.

## Implemented Features

- Dependency Injection (go-fx)
- Routing (gin web framework)
- Environment Files
- Middlewares
- Database Setup (postgres)
- Models Setup and Automigrate (gorm)
- Repositories
- Implementing Basic CRUD Operation
- Authentication (JWT)
- Live code refresh

## Todos

- [x] Use of Modules and injection of them.
- [x] Use of Repository pattern.
- [x] Use of Interfaces.
- [ ] Use of Go routines and concurrency patterns.
- [ ] Logic for Todo list
- [ ] Swagger documentation examples
- [ ] Unit testing examples.
- [ ] File upload middleware.
- [ ] Implementing DDD.

