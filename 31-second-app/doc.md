# App purpose 
CRUD apis demo 
Modules used:

- [x] [go-chi](https://pkg.go.dev/github.com/go-chi/chi) - API routing 
- [x] [go-chi/render](https://pkg.go.dev/github.com/go-chi/render) - for managing request and response payload
- [x] [lib/pq](https://pkg.go.dev/github.com/lib/pq) - postgres driver, used for interacting with postgres database

Add modules:

    $ go get github.com/go-chi/chi github.com/go-chi/render github.com/lib/pq

# App structure

    ├── db
    │   ├── db.go
    │   └── article.go
    ├── handler
    │   ├── errors.go
    │   ├── handler.go
    │   └── artices.go
    ├── models
    │   └── article.go
    ├── .env
    ├── docker-compose.yml
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── README.md


+ db: database related code
  + db.go: database connection
  + article.go: database related functions

# Build docker image 

  docker build -t tumivn/go-crud .
  docker run -d --rm -p 3100:3000 --name gocrud tumivn/go-crud