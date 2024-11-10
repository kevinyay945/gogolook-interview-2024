# Task API Documentation

## Overview

This is a simple REST API for managing tasks, built using Go and the Echo framework. The API allows clients to create, retrieve, and list tasks. The project demonstrates how to use Swagger for API documentation and provides a sample implementation of common RESTful operations.

## Features

- List all tasks
- Create a new task
- Update a new task
- Delete a new task

## Technologies Used

- **Golang** (Version 1.22 or higher): Programming language used to build the API.
- **Echo**: Web framework used for routing and handling HTTP requests.
- **validator**: Library used for input validation.
- **GORM**: ORM library used for database interaction.
- **PostgreSQL**: Main database used for data storage.
- **Swagger (oapi-codegen)**: Used for generating API based on OpenAPI specification.
- **Docker & Docker Compose**: Used for containerizing the application and setting up dependencies.

## Installation

1. **Clone the repository**:

   ```sh
   git clone https://github.com/kevinyay945/gogolook.git
   cd task-api
   ```

2. **Install dependencies**:

   ```sh
   go mod tidy
   ```

3. **Run the initialization script to install required tools**:

   ```sh
   make init
   ```

4. **Generate OpenAPI types and server code**:

   ```sh
   make openapi_http
   ```
5. **Generate Mock Interface**:

   ```sh
   go generate ./...
   ```
6. **Run the application using Docker Compose**:

   ```sh
   docker-compose up --build
   ```

## Docker Setup

The project includes a `docker-compose.yaml` file to simplify setting up the application and its dependencies.

### Docker Compose Services

- **Database (PostgreSQL)**: The PostgreSQL container runs version 13.16 and creates a database with the provided credentials.
    - **Service Name**: `db`
    - **Port**: `5432`
    - **Environment Variables**:
        - `POSTGRES_USER=myuser`
        - `POSTGRES_PASSWORD=mypassword`
        - `POSTGRES_DB=mydb`
    - **Volume**: Persists data in a Docker volume named `db-data`.

- **Application (Go Task API)**: The application is built and run in a container, with connections to the PostgreSQL container.
    - **Service Name**: `app`
    - **Port**: `5000`
    - **Environment Variables**:
        - `PORT=5000`
        - `HOST=0.0.0.0`
        - `POSTGRES_HOST=db`
        - `POSTGRES_PORT=5432`
        - `POSTGRES_USER=myuser`
        - `POSTGRES_PASSWORD=mypassword`
        - `POSTGRES_DB=mydb`
        - `POSTGRES_QUERY_PARAMS=?sslmode=disable`

## Makefile

To simplify the setup process, use the provided `Makefile` to install dependencies and generate Swagger files.

### Available Makefile Commands

- **`make init`**: Installs necessary tools such as `mockgen` and `oapi-codegen`.

  ```makefile
  .PHONY: init
  init:
    	go install go.uber.org/mock/mockgen@v0.5.0
    	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0
  ```

- **`make openapi_http`**: Generates the Swagger types and server code.

  ```makefile
  .PHONY: openapi_http
  openapi_http:
    	oapi-codegen -generate types -o "http/openapi_types.gen.go" -package "http" "assets/swagger/swagger.yml"
    	oapi-codegen -generate server -o "http/openapi_api.gen.go" -package "http" "assets/swagger/swagger.yml"
  ```

## Swagger Documentation

The API uses Swagger to generate interactive documentation.

After running the server, you can access the Swagger UI at:

[Swagger UI](http://localhost:5000/api-docs/index.html)

## Example Task Structure

```json
{
  "id": 1,
  "name": "Sample Task"
}
```

## Contact

For support or any questions, feel free to reach out:

- **Email**: [kevinyay945@gmail.com](mailto:kevinyay945@gmail.com)
- **Website**: [http://www.kevinyay945.com](http://www.kevinyay945.com)

