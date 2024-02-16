# Todo APP

This is a simple rest api for CRUD oprations on todo list items. It uses jwt in cookies
for stateless authentication and authorization, and mongoDB Atlas via mongo-go-driver as the database, and gin-gonic as a framework.
There are structs for all responses to make errors and responses easier to manage on the front end, and also
for flexibility should you want to change something about the responses.

## Getting Started
Clone this repo. Run go mod tidy in your Terminal and the go run main.go. You are now ready to hit all the API's.

## Project Structure
The project is structured into backend components, with contoller, routes, Database in each different folder. The backend is implemented in Golang with a MongoDB database. 

### Overview

The backend serves as the API server for the TODO_APP. It handles API calls initiated from the POSTMAN, managing the application's database. Built using GOLANG, the backend provides a RESTful API, and the database is powered by MOGODB. JWT Authentication is used to authenticate user resgistering in application.


### API Endpoints


#### 1. Create User
- **Endpoint**: `?api/v1/auth/register`
- **Method**: `POST`
- **Description**: Creates a new user for the Todo and enrolls him to specified batch.
- **Request Body**:
  ```json
  {
    "name": "example_username",
    "email": "example_email"
    "password": "example_password",
  }
  ```

#### 2. Login
- **Endpoint**: `/api/v1/auth/login`
- **Method**: `POST`
- **Description**: Validates user credentials for login.
- **Request Body**:
  ```json
  {
    "name":"example_name"
    "email": "example_username",
    "password": "example_password"
  }
  ```

#### 3. GET all task
- **Endpoint**: `/api/v1/todos`
- **Method**: `GET`
  ```
#### 4. Create Task
- **Endpoint**: `/api/v1/todos`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "title":"example_name"
    "description": "example_username",
    "done": "example_password",
    "created_at":"",
    "updated_at":"",
    "user":""
  }
  ```
  #### 5. Update Task
- **Endpoint**: `/api/v1/todos/:id`
- **Method**: `PUT`
- **Request Body**:
  ```json
  {
    "title":"example_name"
    "description": "example_username",
    "done": "example_password",
    "created_at":"",
    "updated_at":"",
    "user":""
  }

#### 6.  Delete Task
- **Endpoint**: `/api/v1/todos/:id`
- **Method**: `DELETE`

## Improvements

I think it would be cool to implement a way to revoke jwt. It is not the most practical thing to do but it would be
possible here with a redis black list for revoked tokens. You would need to make sure to delete the black listed tokens from
the list after they expire.

It needs a request rate limiter.
