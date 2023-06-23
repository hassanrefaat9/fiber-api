# REST API with fiber
This API allows you to perform CRUD operations on products using the Go programming language, Fiber framework, PostgreSQL database, and GORM ORM.

## Prerequisites

Before running the API, ensure that you have the following dependencies installed:

- Go programming language (version X.X.X)
- Fiber framework (version X.X.X)
- PostgreSQL database
- GORM ORM (version X.X.X)

## Installation

1. Clone the repository:
```sh
git clone https://github.com/your/repository.git
```
2. Install the required dependencies:
```sh
go get -u github.com/gofiber/fiber
go get -u gorm.io/gorm 
go get -u gorm.io/driver/postgres
```
1. Set up the PostgreSQL database and configure the connection in your application.
```go
// Example database configuration
dsn := "host=127.0.0.1 user=youruser password=yourpassword dbname=yourdatabase port=5432 sslmode=disable"

```

## Endpoints
The following endpoints are available for interacting with the Product API:

### Get all products
- **URL:** `/product/`
- **Method:** `GET`
- **Description:** Retrieves all products from the database.
- **Response:** Returns a JSON array containing all products.
```json
{
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2023-06-23T11:34:04.036596+03:00",
      "UpdatedAt": "2023-06-23T11:34:04.036596+03:00",
      "DeletedAt": null,
      "Title": "Sample Product",
      "description": "This is a sample product description.",
      "amount": 10
    },
    {
      "ID": 2,
      "CreatedAt": "2023-06-23T11:48:16.36595+03:00",
      "UpdatedAt": "2023-06-23T11:48:16.36595+03:00",
      "DeletedAt": null,
      "Title": "Sample Product2",
      "description": "This is a sample product description.2",
      "amount": 10
    }
  ],
  "message": "All products",
  "status": "success"
}

```

### Get a specific product

- **URL:** `/product/:id`
- **Method:** `GET`
- **Description:** Retrieves a specific product based on its ID.
- **Parameters:**
    - `id` (path parameter) - ID of the product to retrieve.
- **Response:** Returns a JSON object representing the product.
```json
{
  "data": {
    "ID": 1,
    "CreatedAt": "2023-06-23T11:34:04.036596+03:00",
    "UpdatedAt": "2023-06-23T11:34:04.036596+03:00",
    "DeletedAt": null,
    "Title": "Sample Product",
    "description": "This is a sample product description.",
    "amount": 10
  },
  "message": "Product found",
  "status": "success"
}

```

### Create a new product

- **URL:** `/product/`
- **Method:** `POST`
- **Description:** Creates a new product.
- **Request Body:** Requires a JSON object containing the following fields:
    - `title` (string) - Title of the product.
    - `description` (string) - Description of the product.
    - `amount` (number) - Amount of the product.
- **Response:** Returns a JSON object representing the created product.
```json
{
    "data": {
        "ID": 2,
        "CreatedAt": "2023-06-23T11:48:16.365950748+03:00",
        "UpdatedAt": "2023-06-23T11:48:16.365950748+03:00",
        "DeletedAt": null,
        "Title": "Sample Product2",
        "description": "This is a sample product description.2",
        "amount": 10
    },
    "message": "Created product",
    "status": "success"
}
```

### Delete a product

- **URL:** `/product/:id`
- **Method:** `DELETE`
- **Description:** Deletes a specific product based on its ID.
- **Parameters:**
    - `id` (path parameter) - ID of the product to delete.
- **Response:** Returns a success message indicating the deletion of the product.
```json
{
    "data": null,
    "message": "Product successfully deleted",
    "status": "success"
}
```

## Error Handling

The API provides appropriate error handling and returns standard HTTP status codes in case of any failures. Possible error responses include:
- `400 Bad Request`: Invalid request parameters or missing fields.
- `404 Not Found`: The requested resource (product) does not exist.
- `500 Internal Server Error`: Internal server error occurred.