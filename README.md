# User Addresses Project

## Project Description

### Part 1 - Reading Massive Amount of Data and Inserting into a Relational Database

This project reads information for 1 million users from a JSON file and concurrently
inserts this data into a relational database. Each user's information is represented
as a JSON object, and the data is stored in two separate tables: `users` and `addresses`.
The relationship between these tables is one-to-many, meaning each user can have zero or
more addresses. 10 concurrent processes are used to insert the data into the database.

### Part 2 - Building an API to Access User Information

This project builds an API to access user information stored in the database.
The API must include endpoints to create, read, update, and delete user information.
The user information contains 4 fields: `first name`, `last name`, `address`, and `ID`.

## Features

- User information include at least 4 fields: first name, last name, address, and ID.
- Uses PostgreSQL as the database.
- Uses channels to send data from the reading phase to the writing phase.
- Uses an ORM like Gorm.
- Follow three-layer architecture such as Clean Architecture.
- Uses the Echo library to build the API.
- Uses worker pool design pattern to manage concurrent processing.

## Project FrameWorks

- **Database**: PostgreSQL
- **ORM**: Gorm
- **API Framework**: Echo
- **Architecture**: Clean Architecture

## API Endpoints

- **Create User**: `POST /user`
- **Get User by ID**: `GET /user/:id`
- **Update User**: `PUT /user`
- **Delete User**: `DELETE /user/:id`

## Setup and Run

1. **Clone the repository**:
    ```sh
    git clone
    ```
2. **Setup the postgresql database and config the connection**
3. **Run the desired service from ./cmd package**:
    ```sh
    go run ./cmd/<service-name>
    ``` 

## Example `curl` Commands

- **Create User**:
  ```sh
  curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1995a603-816f-431f-a424-942baab6f6a7",
    "name": "John Doe",
    "email": "john.doe@example.com",
    "addresses": [
      {
        "street": "123 Main St",
        "city": "Anytown",
        "state": "CA",
        "zip": "12345"
      }
    ]
  }'
  ```

- **Get User by ID**:
  ```sh
  curl -X GET http://localhost:8080/user/1995a603-816f-431f-a424-942baab6f6a7
  ```

- **Update User**:
  ```sh
  curl -X PUT http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1995a603-816f-431f-a424-942baab6f6a7",
    "name": "John Doe Updated",
    "email": "john.doe.updated@example.com",
    "addresses": [
      {
        "street": "123 Main St",
        "city": "Anytown",
        "state": "CA",
        "zip": "12345"
      }
    ]
  }'
  ```

- **Delete User**:
  ```sh
  curl -X DELETE http://localhost:8080/user/1995a603-816f-431f-a424-942baab6f6a7
  ```

## License

This project is licensed under the MIT License.
