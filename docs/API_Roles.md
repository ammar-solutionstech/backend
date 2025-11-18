# Role Management API

All endpoints require JWT authentication and specific permissions.

## GET /api/roles

List all roles.

**Permission**: `View All Roles` (Module: `Basic`)

### Response

**Success (200 OK):**
```json
[
  {
    "id": 1,
    "name": "admin",
    "description": "Administrator role with full access"
  },
  {
    "id": 2,
    "name": "user",
    "description": "Standard user role"
  }
]
```

### Example cURL

```bash
curl -X GET http://localhost:8080/api/roles \
  -H "Authorization: Bearer <token>"
```

---

## GET /api/roles/{id}

Get a specific role by ID with permissions.

**Permission**: `View Role` (Module: `Basic`)

### Response

**Success (200 OK):**
```json
{
  "id": 1,
  "name": "admin",
  "description": "Administrator role with full access",
  "Permissions": [
    {
      "id": 1,
      "name": "View All Users",
      "module": "Basic"
    }
  ]
}
```

### Example cURL

```bash
curl -X GET http://localhost:8080/api/roles/1 \
  -H "Authorization: Bearer <token>"
```

---

## POST /api/roles

Create a new role.

**Permission**: `Add Role` (Module: `Basic`)

### Request

**Body:**
```json
{
  "name": "manager",
  "description": "Manager role with elevated permissions"
}
```

### Response

**Success (201 Created):**
```json
{
  "id": 3,
  "name": "manager",
  "description": "Manager role with elevated permissions"
}
```

### Example cURL

```bash
curl -X POST http://localhost:8080/api/roles \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "manager",
    "description": "Manager role with elevated permissions"
  }'
```

---

## PUT /api/roles/{id}

Update an existing role.

**Permission**: `Edit Role` (Module: `Basic`)

### Request

**Body:**
```json
{
  "name": "senior_manager",
  "description": "Senior manager role"
}
```

### Response

**Success (200 OK):**
```json
{
  "id": 3,
  "name": "senior_manager",
  "description": "Senior manager role"
}
```

### Example cURL

```bash
curl -X PUT http://localhost:8080/api/roles/3 \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "senior_manager",
    "description": "Senior manager role"
  }'
```

---

## DELETE /api/roles/{id}

Delete a role.

**Permission**: `Delete Role` (Module: `Basic`)

### Response

**Success (204 No Content)**

### Example cURL

```bash
curl -X DELETE http://localhost:8080/api/roles/3 \
  -H "Authorization: Bearer <token>"
```

---

## PUT /api/roles/{id}/permissions

Replace all role permissions with a new set.

**Permission**: `Update Permissions of Role` (Module: `Basic`)

### Request

**Body:**
```json
[1, 2, 3, 4, 5]
```

Array of permission IDs.

### Response

**Success (200 OK):**
```json
{
  "id": 1,
  "name": "admin",
  "description": "Administrator role",
  "Permissions": [
    {
      "id": 1,
      "name": "View All Users",
      "module": "Basic"
    },
    {
      "id": 2,
      "name": "View User",
      "module": "Basic"
    },
    {
      "id": 3,
      "name": "Add User",
      "module": "Basic"
    },
    {
      "id": 4,
      "name": "Edit User",
      "module": "Basic"
    },
    {
      "id": 5,
      "name": "Delete User",
      "module": "Basic"
    }
  ]
}
```

### Example cURL

```bash
curl -X PUT http://localhost:8080/api/roles/1/permissions \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '[1, 2, 3, 4, 5]'
```

---

## POST /api/roles/{id}/permission

Add a permission to a role.

**Permission**: `Update Permissions of Role` (Module: `Basic`)

### Request

**Body:**
```json
{
  "id": 6
}
```

### Response

**Success (200 OK):**
```json
{
  "id": 1,
  "name": "admin",
  "description": "Administrator role",
  "Permissions": [
    {
      "id": 1,
      "name": "View All Users",
      "module": "Basic"
    },
    {
      "id": 6,
      "name": "Suspend User",
      "module": "Basic"
    }
  ]
}
```

### Example cURL

```bash
curl -X POST http://localhost:8080/api/roles/1/permission \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"id": 6}'
```

---

## DELETE /api/roles/{id}/permission

Remove a permission from a role.

**Permission**: `Update Permissions of Role` (Module: `Basic`)

### Request

**Body:**
```json
{
  "id": 6
}
```

### Response

**Success (200 OK):**
```json
{
  "id": 1,
  "name": "admin",
  "description": "Administrator role",
  "Permissions": [
    {
      "id": 1,
      "name": "View All Users",
      "module": "Basic"
    }
  ]
}
```

### Example cURL

```bash
curl -X DELETE http://localhost:8080/api/roles/1/permission \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"id": 6}'
```

