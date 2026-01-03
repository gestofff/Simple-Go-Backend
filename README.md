Simple Go Backend
=================

Simple Go Backend is an educational REST API project written in Go.
It implements a basic calculation service with persistent storage using PostgreSQL.

The project demonstrates a clean layered architecture (handlers → services → repositories),
basic CRUD operations, expression evaluation, and database interaction via GORM.


Features
--------

- REST API built with Echo
- PostgreSQL database with GORM
- Arithmetic expression evaluation using govaluate
- UUID-based identifiers
- Basic CRUD operations for calculations
- Automatic database migration on startup


Project Structure
-----------------

~
├── cmd/
│   └── main.go                  Application entry point
├── internal/
│   ├── calculationService/      Business logic layer
│   │   ├── orm.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── db/
│   │   └── db.go                Database initialization
│   └── handlers/
│       └── calculationHandlers.go
├── go.mod
└── README.txt


API Endpoints
-------------

Get all calculations
GET /calculations

Create a new calculation
POST /calculations

Request body:
{
  "expression": "2 + 2 * 3"
}

Update a calculation
PATCH /calculations/{id}

Request body:
{
  "expression": "(10 - 4) / 2"
}

Delete a calculation
DELETE /calculations/{id}


Data Model
----------

{
  "id": "uuid",
  "expression": "string",
  "result": "string"
}


Requirements
------------

- Go 1.25.5
- PostgreSQL
- Running PostgreSQL instance on localhost:5432


Running the Project
-------------------

1. Clone the repository
2. Make sure PostgreSQL is running
3. Update database credentials if needed (see note below)
4. Run the application:

go run ./cmd/main.go

The server will start at:
http://localhost:8080


Database Configuration Note
---------------------------

IMPORTANT (Educational Project Notice)

In this educational project, the database connection parameters
(host, user, password, database name) are intentionally hardcoded
in the source code and not moved to environment variables.

This was done for simplicity and learning purposes.
In a real-world application, these values should be stored in
environment variables or configuration files.
This can be implemented easily if needed.


Purpose
-------

This project was created as a learning exercise to practice:

- Backend architecture in Go
- REST API design
- Database interaction
- Separation of concerns
- Service and repository patterns


License
-------

This project is intended for educational use only.
