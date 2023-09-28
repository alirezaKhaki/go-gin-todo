<p align="center">
  <img src="/gopher.png" height="400">
  <h1 align="center">
    Go Patterns
    <br>
    <a href="http://travis-ci.org/tmrts/go-patterns"><img alt="build-status" src="https://img.shields.io/badge/build-passing-brightgreen.svg?style=flat-square" /></a>
    <a href="https://github.com/sindresorhus/awesome" ><img alt="awesome" src="https://img.shields.io/badge/awesome-%E2%9C%93-ff69b4.svg?style=flat-square" /></a>
    <a href="https://github.com/tmrts/go-patterns/blob/master/LICENSE" ><img alt="license" src="https://img.shields.io/badge/license-Apache%20License%202.0-E91E63.svg?style=flat-square" /></a>
  </h1>
</p>

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

