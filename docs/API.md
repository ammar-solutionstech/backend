# ITaaS API Documentation

## Base URL

```
http://{{host name/IP}}:{{Port}}
```

## Overview

The ITaaS (IT as a Service) API provides comprehensive endpoints for managing users, roles, permissions, inventory, help desk tickets, and more. All endpoints except `/login` and `/health` require JWT authentication.

## Authentication

### JWT Token Flow

1. **Login**: Send credentials to `/login` endpoint
2. **Receive Token**: Response contains a JWT token
3. **Use Token**: Include token in `Authorization` header for all protected endpoints

### Authorization Header Format

```
Authorization: Bearer <token>
```

### Token Expiration

Tokens expire after the duration specified in `JWT_EXP_MINUTES` environment variable (default: 15 minutes).

## Common Headers

All requests should include:

```
Content-Type: application/json
Authorization: Bearer <token>  (required for protected endpoints)
```

## Error Handling

### Standard Error Response Format

```json
{
  "error": "Error message describing what went wrong"
}
```

### HTTP Status Codes

- **200 OK**: Request successful
- **201 Created**: Resource created successfully
- **204 No Content**: Resource deleted successfully
- **400 Bad Request**: Invalid request parameters or body
- **401 Unauthorized**: Missing or invalid authentication token
- **403 Forbidden**: User lacks required permission
- **404 Not Found**: Resource not found
- **500 Internal Server Error**: Server error occurred

### Example Error Response

```json
{
  "error": "user not found"
}
```

## Success Response Format

### Single Resource

```json
{
  "id": 1,
  "name": "Example Resource",
  ...
}
```

### Collection

```json
[
  {
    "id": 1,
    "name": "Resource 1"
  },
  {
    "id": 2,
    "name": "Resource 2"
  }
]
```

## Date/Time Formats

- **Date**: `YYYY-MM-DD` (e.g., `2025-01-15`)
- **DateTime**: ISO 8601 format (e.g., `2025-01-15T10:30:00Z`)
- **Time**: `HH:MM:SS` (e.g., `14:30:00`)

## Pagination

Currently, list endpoints return all records. Pagination may be added in future versions.

## Common Patterns

### CRUD Operations

Most resources follow standard REST patterns:

- `GET /api/{resource}` - List all resources
- `GET /api/{resource}/{id}` - Get specific resource
- `POST /api/{resource}` - Create new resource
- `PUT /api/{resource}/{id}` - Update resource (partial updates supported)
- `DELETE /api/{resource}/{id}` - Delete resource

### Relationship Endpoints

Many resources have relationship endpoints for managing associations:

- `GET /api/{resource}/{id}/{relation}` - List related resources
- `POST /api/{resource}/{id}/{relation}` - Add relationship
- `DELETE /api/{resource}/{id}/{relation}/{relatedId}` - Remove relationship

## Permission System

All protected endpoints require:
1. Valid JWT token
2. Specific permission (name and module)

Permissions are checked via middleware. If a user lacks the required permission, a `403 Forbidden` response is returned.

## Field Validation

- Required fields must be provided in request bodies
- String fields are trimmed of whitespace
- ID fields must be positive integers
- Foreign key references are validated to exist

## Nullable Fields

Fields marked with `omitempty` in JSON tags are optional and can be `null` or omitted from requests.

## Related Documentation

- [Authentication API](API_Auth.md)
- [User Management API](API_Users.md)
- [Role Management API](API_Roles.md)
- [Permission Management API](API_Permissions.md)
- [Inventory API](API_Inventory.md)
- [Geography API](API_Geography.md)
- [Help Desk API](API_HelpDesk.md)
- [Navigation API](API_Navigation.md)
- [Admin API](API_Admin.md)
- [Quick Reference](API_QuickReference.md)

