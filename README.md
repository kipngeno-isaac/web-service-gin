# Simple Go API for Album Management (Gin Web Framework)

This project demonstrates a basic API built with Go and the Gin web framework, enabling CRUD operations on albums. Explore essential Go concepts and RESTful API design while creating and managing album data.

## Features:

- Lists all albums in the database (GET /albums).
- Retrieves details of a specific album by ID (GET /albums/:id).
- Saves a new album to the database (POST /albums).
- Error handling and meaningful HTTP status codes.
- Clear instructions for setup and usage.

### Getting Started:

**Prerequisites:**
    Go installation (https://golang.org/)
    A code editor or IDE (e.g., VS Code, GoLand)

**Clone the repository:**
Bash

`git clone https://github.com/kipngeno-isaac/web-service-gin.git`

Use code with caution. Learn more

**Install dependencies:**
Bash
```
cd album-api
go mod download
```
Use code with caution. Learn more

**Run the server:**
Bash

`go run main.go`

Use code with caution. Learn more

The server will typically start on port 8080, but check the code for exact configuration.

### Access the API:

Use tools like Postman, curl, or your preferred REST client to interact with the API:

| Method       | Endpoint          | Description                                    |
|--------------|------------------|------------------------------------------------|
| GET           | /albums          | Retrieve a list of all albums                    |
| GET           | /albums/:id       | Get details of a specific album by ID (404 for not found) |
| POST          | /albums          | Create a new album (JSON payload with name and artist) |
| PUT          | /albums/:id       | Update an existing album (JSON payload with optional name and artist) |
| DELETE        | /albums/:id       | Delete an album (404 for not found)                 |

**Example Request (GET):**

GET /albums HTTP/1.1
Content-Type: application/json
```
[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    },
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
    }
]
```

Use code with caution. Learn more

Customization and Enhancement:

Integrate with a database for persistent storage (e.g., PostgreSQL, SQLite).
Implement data validation and error handling for enhanced robustness.
Add search and filtering functionalities for album queries.
Consider performance optimization and scalability for larger datasets or user loads.
Implement authentication and authorization for robust security.