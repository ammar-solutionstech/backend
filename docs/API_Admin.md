# Admin API

## GET /admin

Admin-only endpoint for administrative operations.

**Authentication**: Required (JWT token)

**Permission**: Requires `admin` role (checked via `RequireRoles` middleware)

### Request

**Headers:**
```
Authorization: Bearer <token>
```

### Response

**Success (200 OK):**
Response format depends on implementation in `controllers.AdminOnly`.

**Error (401 Unauthorized):**
Missing or invalid token.

**Error (403 Forbidden):**
User does not have `admin` role.

### Example cURL

```bash
curl -X GET http://localhost:8080/admin \
  -H "Authorization: Bearer <token>"
```

### Notes

- This endpoint uses role-based access control (RBAC) rather than permission-based
- Only users with the `admin` role can access this endpoint
- The endpoint is registered separately from other API routes at `/admin` (not `/api/admin`)

