# Authentication API

## POST /login

Authenticate a user and receive a JWT token.

**Authentication**: Not required (public endpoint)

### Request

**Headers:**
```
Content-Type: application/json
```

**Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

### Response

**Success (200 OK):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJyb2xlcyI6WyJhZG1pbiJdLCJwZXJtaXNzaW9ucyI6WzEsMiwzXSwiZXhwIjoxNzA1MjM0NTYwfQ.example_signature"
}
```

**Error (400 Bad Request):**
```json
{
  "error": "invalid request"
}
```

**Error (401 Unauthorized):**
```json
{
  "error": "invalid email or password"
}
```

### Example cURL

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

### Notes

- The email field is case-insensitive and trimmed of whitespace
- The token contains user ID, roles, and permissions
- Store the token securely and include it in subsequent requests

---

## GET /health

Check if the API server is running.

**Authentication**: Not required (public endpoint)

### Request

No parameters required.

### Response

**Success (200 OK):**
```
ok
```

### Example cURL

```bash
curl http://localhost:8080/health
```

### Notes

- This endpoint does not check database connectivity
- Useful for load balancer health checks
- Returns plain text "ok" response

