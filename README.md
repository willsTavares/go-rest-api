This Go project is a RESTful API designed to manage a collection of personalities. It provides endpoints for creating, retrieving, updating, and deleting personality records. 

Each personality is represented by a unique ID and contains details that can be accessed or modified through the API. The project leverages the Gorilla Mux router for handling HTTP requests and GORM for database operations, ensuring efficient data manipulation and retrieval. With a focus on simplicity and functionality, this API serves as a robust backend for applications requiring personality data management.

### Prerequisites

- Go installed (verify with `go version`).
- Docker installed for setting up PostgreSQL.

### Initialize Project locally

- Run `docker-compose up` -d to start the PostgreSQL container.
- Execute `go run main.go`

## Project Architecture

### Components

- **Web Server/API:** Utilizes `gorilla/mux` for handling HTTP requests and routing.
- **Data Access:** Employs `gorm` for Object-Relational Mapping (ORM), facilitating CRUD operations with the database.

### Directory Structure

- `/controllers`: Logic for handling requests and responses. Controllers parse requests, interact with models, and return responses.
- `/database`: Database connection configurations and the database instance setup.
- `/migrations`: Scripts for database migrations to create and update the database schema.
- `/models`: Data models and associated business logic. Defines the structure of the database tables and their relationships.
- `/routes`: Route definitions and their association with controllers. Maps endpoints to controller functions.
