# CRUD API using Gin and GORM in Go
This is a basic CRUD (Create, Read, Update, Delete) API built in Go using the Gin web framework and the GORM ORM. 
The API allows you to manage "Blog Post" records, storing their Title and Body in a PostgreSQL database.

## Prerequisites
1. Go (1.16 or higher)
2. [PostgreSQL](https://www.elephantsql.com/) (optional, the API can be modified to use other databases)

## Installation 
1. Clone the repository :
```
git clone https://github.com/yourusername/crud-api-gin-gorm.git
cd crud-api-gin-gorm
```

2. Install the required packages :
```
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```
## Usage
1. Start the server
```
go run main.go
```
The server will be running at http://localhost:3001.

2. Available API endpoints:

* GET /posts: Fetches all posts from the database.
* POST /posts: Creates a new posts by providing JSON data in the request body.
* GET /posts/:id: Fetches a posts by its ID.
* PUT /posts/:id: Updates a posts by its ID using JSON data in the request body.
* DELETE /posts/:id: Deletes a posts by its ID.
