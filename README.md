 The architecture of this Go web application project, designed as a REST API for resource management, leverages the gorilla/mux framework for efficient routing and gorm for database interactions. Here's a detailed explanation of its architecture and execution instructions, including setting up a PostgreSQL database using Docker.


### Prerequisites

- Go installed (verify with `go version`).
- Docker installed for setting up PostgreSQL.


## Project Architecture

### Components

- **Web Server/API:** Utilizes `gorilla/mux` for handling HTTP requests and routing.
- **Data Access:** Employs `gorm` for Object-Relational Mapping (ORM), facilitating CRUD operations with the database.

### Directory Structure

- `/controllers`: Logic for handling requests and responses. Controllers parse requests, interact with models, and return responses.
- `/db`: Database connection configurations and the database instance setup.
- `/migrations`: Scripts for database migrations to create and update the database schema.
- `/models`: Data models and associated business logic. Defines the structure of the database tables and their relationships.
- `/routes`: Route definitions and their association with controllers. Maps endpoints to controller functions.
