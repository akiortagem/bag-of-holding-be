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
  - Or:
    ```json
    {
      "error": "user already exists"
    }
    ```
- `500 Internal Server Error`
  - Body:
    ```json
    {
      "error": "failed to create user"
    }
    ```

## Auth

### POST `/api/login`

Logs in a user and returns an access token plus a refresh token.

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
- `200 OK`
  - Body:
    ```json
    {
      "email": "user@example.com",
      "token": "jwt-access-token",
      "refresh_token": "refresh-token"
    }
    ```
- `400 Bad Request`
  - Body:
    ```json
    {
      "error": "invalid request body"
    }
    ```
- `401 Unauthorized` with empty body.
- `500 Internal Server Error` with empty body.

### GET `/protected-health`

Protected version of the health check endpoint.

**Request**
- Header: `Authorization: Bearer <token>`

**Responses**
- `200 OK` with empty body.
- `401 Unauthorized`
  - Body:
    ```json
    {
      "error": "missing or invalid authorization header"
    }
    ```
  - Or:
    ```json
    {
      "error": "missing bearer token"
    }
    ```
  - Or: empty body when the token is invalid.
- `500 Internal Server Error`
  - Body:
    ```json
    {
      "error": "Health check is not OK"
    }
    ```
