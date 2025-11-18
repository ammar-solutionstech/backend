# Permission Management API

All endpoints require JWT authentication and specific permissions.

## GET /api/permissions

List all permissions.

**Permission**: `View All Permissions` (Module: `Basic`)

### Response

**Success (200 OK):**
```json
[
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
  },
  {
    "id": 6,
    "name": "View Brands",
    "module": "Inventory"
  }
]
```

### Example cURL

```bash
curl -X GET http://localhost:8080/api/permissions \
  -H "Authorization: Bearer <token>"
```

---

## GET /api/permissions/{id}

Get a specific permission by ID.

**Permission**: `View Permission` (Module: `Basic`)

### Response

**Success (200 OK):**
```json
{
  "id": 1,
  "name": "View All Users",
  "module": "Basic"
}
```

**Error (404 Not Found):**
```json
{
  "error": "Permission not found"
}
```

### Example cURL

```bash
curl -X GET http://localhost:8080/api/permissions/1 \
  -H "Authorization: Bearer <token>"
```

---

## POST /api/permissions

Create a new permission.

**Permission**: `Add Permission` (Module: `Basic`)

### Request

**Body:**
```json
{
  "name": "View Equipment",
  "module": "Inventory"
}
```

**Required Fields:**
- `name`
- `module`

### Response

**Success (201 Created):**
```json
{
  "id": 7,
  "name": "View Equipment",
  "module": "Inventory"
}
```

**Error (400 Bad Request):**
```json
{
  "error": "Invalid input"
}
```

### Example cURL

```bash
curl -X POST http://localhost:8080/api/permissions \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "View Equipment",
    "module": "Inventory"
  }'
```

---

## PUT /api/permissions/{id}

Update an existing permission.

**Permission**: `Edit Permission` (Module: `Basic`)

### Request

**Body:**
```json
{
  "name": "View All Equipment",
  "module": "Inventory"
}
```

### Response

**Success (200 OK):**
```json
{
  "id": 7,
  "name": "View All Equipment",
  "module": "Inventory"
}
```

### Example cURL

```bash
curl -X PUT http://localhost:8080/api/permissions/7 \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "View All Equipment",
    "module": "Inventory"
  }'
```

---

## DELETE /api/permissions/{id}

Delete a permission.

**Permission**: `Delete Permission` (Module: `Basic`)

### Response

**Success (204 No Content)**

**Error (404 Not Found):**
```json
{
  "error": "Permission not found"
}
```

### Example cURL

```bash
curl -X DELETE http://localhost:8080/api/permissions/7 \
  -H "Authorization: Bearer <token>"
```

