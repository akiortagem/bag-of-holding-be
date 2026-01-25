# API Endpoints

Base URL: `http://localhost:8080` (default Gin `Run` address)

## Health

### GET `/health`

Returns 200 when the service reports healthy.

**Responses**
- `200 OK` with empty body.
- `500 Internal Server Error`
  - Body:
    ```json
    {
      "error": "Health check is not OK"
    }
    ```

## Users

### POST `/api/users`

Creates a user.

**Request**
- Content-Type: `application/json`
- Body:
  ```json
  {
    "email": "user@example.com",
    "password": "plain-text-password"
  }
  ```

**Responses**
- `201 Created`
  - Body:
    ```json
    {
      "id": 123,
      "email": "user@example.com"
    }
    ```
- `400 Bad Request`
  - Body:
    ```json
    {
      "error": "invalid request body"
    }
    ```
  - Or:
    ```json
    {
      "error": "invalid payload"
    }
    ```
- `500 Internal Server Error`
  - Body:
    ```json
    {
      "error": "failed to create user"
    }
    ```
