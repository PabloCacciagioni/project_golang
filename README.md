# project_golang

This is an example project in Go using GORM and Fiber.

## Requirements

  - Docker
  - Docker Compose 
  - Go 1.16 o superior 

## Configuration and Execution 

### Step 1: Clone the repository

```bash
git clone https://github.com/PabloCacciagioni/project_golang.git
cd project_golang
```
### Step 2: Configure the environment 

Make sure you have Docker and Docker Compose installed on your system.

### Step 3: Lift the containers 

Lift database and application containers using Docker Compose:

```bash
docker-compose up -d
```

This will raise two services: 
- "todoapp": The application in Go. 
- "mysql": The mysql database.

### Step 4: Database migrations 

Runs the database migrations to create the necessary tables. This is done automatically on application initialization when connecting to the database.

### Step 5: Run the application

The application should be running at http://localhost:8000.

Available routes

- ```GET /todos/:id:``` Obtain a todo by ID.
- ```POST /todos:``` Create a new todo.
- ```PUT /todos/:id:``` Update a todo by ID. 
- ```DELETE /todos/:id:``` Delete a todo by ID.


## Test

To run the tests, make sure the Docker containers are running, and then run:

```bash
go test -v routes_test.go
```

## Project structure 


- ```main.go:``` Application entry point.
- ```database/:``` Database connection and configuration.
- ```config/:``` Configuration of the URL used for the database connection.
- ```models/:``` Definition of data models.
- ```routes/:``` Definition of routes and controllers.
- ```Dockerfile:``` Docker configuration file for the application.
- ```docker-compose.yml:``` Docker Compose configuration file.


