# User Management API

All endpoints require JWT authentication and specific permissions.

## GET /api/users

List all users with their nationality.

**Permission**: `View All Users` (Module: `Basic`)

### Request

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

### Response

**Success (200 OK):**
```json
[
  {
    "id": 1,
    "first_name": "John",
    "latest_name": "Doe",
    "father_name": "Michael",
    "nationality_id": 1,
    "dob": "1990-05-15T00:00:00Z",
    "work_email": "john.doe@company.com",
    "private_email": "john.doe@personal.com",
    "work_mobile": "+1234567890",
    "private_mobile": "+1234567891",
    "id_number": "123456789",
    "id_type": 1,
    "contact_id": 1,
    "active": 1,
    "nationality": {
      "id": 1,
      "name": "United States",
      "code": "US"
    }
  }
]
```

### Example cURL

```bash
curl -X GET http://localhost:8080/api/users \
  -H "Authorization: Bearer <token>"
```

---

## GET /api/users/{id}

Get a specific user by ID with relationships.

**Permission**: `View User` (Module: `Basic`)

### Request

**Path Parameters:**
- `id` (integer, required): User ID

### Response

**Success (200 OK):**
```json
{
  "id": 1,
  "first_name": "John",
  "latest_name": "Doe",
  "father_name": "Michael",
  "nationality_id": 1,
  "dob": "1990-05-15T00:00:00Z",
  "work_email": "john.doe@company.com",
  "private_email": "john.doe@personal.com",
  "work_mobile": "+1234567890",
  "private_mobile": "+1234567891",
  "id_number": "123456789",
  "id_type": 1,
  "contact_id": 1,
  "active": 1,
  "nationality": {
    "id": 1,
    "name": "United States",
    "code": "US"
  },
  "roles": [
    {
      "id": 1,
      "name": "admin",
      "description": "Administrator role"
    }
  ],
  "permissions": [
    {
      "id": 1,
      "name": "View All Users",
      "module": "Basic"
    }
  ],
  "teams": [
    {
      "id": 1,
      "name": "IT Support Team",
      "description": "Primary IT support team",
      "manager_id": 1,
      "is_active": true
    }
  ]
}
```

**Error (404 Not Found):**
```json
{
  "error": "User not found : record not found"
}
```

### Example cURL

```bash
curl -X GET http://localhost:8080/api/users/1 \
  -H "Authorization: Bearer <token>"
```

---

## POST /api/users

Create a new user.

**Permission**: `Add User` (Module: `Basic`)

### Request

**Body:**
```json
{
  "first_name": "Jane",
  "latest_name": "Smith",
  "father_name": "Robert",
  "nationality_id": 1,
  "dob": "1992-08-20",
  "work_email": "jane.smith@company.com",
  "private_email": "jane.smith@personal.com",
  "password": "securepassword123",
  "work_mobile": "+1234567892",
  "private_mobile": "+1234567893",
  "id_number": "987654321",
  "id_type": 1,
  "contact_id": 2,
  "active": 1
}
```

**Required Fields:**
- `first_name`
- `latest_name`
- `father_name`
- `work_email`
- `password`

### Response

**Success (201 Created):**
```json
{
  "id": 2,
  "first_name": "Jane",
  "latest_name": "Smith",
  "father_name": "Robert",
  "nationality_id": 1,
  "dob": "1992-08-20T00:00:00Z",
  "work_email": "jane.smith@company.com",
  "private_email": "jane.smith@personal.com",
  "work_mobile": "+1234567892",
  "private_mobile": "+1234567893",
  "id_number": "987654321",
  "id_type": 1,
  "contact_id": 2,
  "active": 1
}
```

**Error (400 Bad Request):**
```json
{
  "error": "Invalid input : invalid character"
}
```

### Example cURL

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Jane",
    "latest_name": "Smith",
    "father_name": "Robert",
    "work_email": "jane.smith@company.com",
    "password": "password123"
  }'
```

---

## PUT /api/users/{id}

Update an existing user.

**Permission**: `Edit User` (Module: `Basic`)

### Request

**Path Parameters:**
- `id` (integer, required): User ID

**Body:**
```json
{
  "first_name": "Jane",
  "latest_name": "Smith",
  "father_name": "Robert",
  "nationality_id": 1,
  "dob": "1992-08-20",
  "work_email": "jane.smith@company.com",
  "private_email": "jane.smith@personal.com",
  "password": "securepassword123",
  "work_mobile": "+1234567892",
  "private_mobile": "+1234567893",
  "id_number": "987654321",
  "id_type": 1,
  "contact_id": 2,
  "active": 1
}
```

### Response

**Success (200 OK):**
```json
{
  "id": 1,
  "first_name": "Jane",
  "latest_name": "Smith",
  "father_name": "Robert",
  "nationality_id": 1,
  "dob": "1992-08-20",
  "work_email": "jane.smith@company.com",
  "private_email": "jane.smith@personal.com",
  "password": "securepassword123",
  "work_mobile": "+1234567892",
  "private_mobile": "+1234567893",
  "id_number": "987654321",
  "id_type": 1,
  "contact_id": 2,
  "active": 1
}
```

**Error (404 Not Found):**
```json
{
  "error": "Can't Update user : record not found"
}
```

### Example cURL

```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "work_email": "new.email@company.com"
  }'
```

---

## DELETE /api/users/{id}

Delete a user.

**Permission**: `Delete User` (Module: `Basic`)

### Request

**Path Parameters:**
- `id` (integer, required): User ID

### Response

**Success (200 OK):**
```json
{
  "message": "User was deleted successfully",
  "user": {
    "id": 1,
    "first_name": "John",
    "latest_name": "Doe",
    ...
  }
}
```

**Error (404 Not Found):**
```json
{
  "error": "User not found : record not found"
}
```

### Example cURL

```bash
curl -X DELETE http://localhost:8080/api/users/1 \
  -H "Authorization: Bearer <token>"
```

---

## PUT /api/users/{id}/suspend

Activate or deactivate a user.

**Permission**: `Suspend User` (Module: `Basic`)

### Request

**Path Parameters:**
- `id` (integer, required): User ID

**Body:**
```json
{
  "active": "yes"
}
```

or

```json
{
  "active": "no"
}
```

### Response

**Success (200 OK):**
```json
{
  "id": 1,
  "active": 1
}
```

### Example cURL

```bash
curl -X PUT http://localhost:8080/api/users/1/suspend \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"active": "yes"}'
```

---

## PUT /api/users/{id}/permissions

Replace all user permissions with a new set.

**Permission**: `Update Permissions of User` (Module: `Basic`)

### Request

**Path Parameters:**
- `id` (integer, required): User ID

**Body:**
```json
[1, 2, 3, 5]
```

Array of permission IDs.

### Response

**Success (200 OK):**
```json
{
  "message": "Permissions was updated successfully",
  "user": {
    "id": 1,
    "permissions": [
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
        "id": 5,
        "name": "Edit User",
        "module": "Basic"
      }
    ]
  }
}
```

### Example cURL

```bash
curl -X PUT http://localhost:8080/api/users/1/permissions \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '[1, 2, 3, 5]'
```

---

## POST /api/users/{id}/permissions

Add a single permission to a user.

**Permission**: `Update Permissions of User` (Module: `Basic`)

### Request

**Path Parameters:**
- `id` (integer, required): User ID

**Body:**
```json
{
  "id": 4
}
```
JSON include the ID of permission.

### Response

**Success (200 OK):**
```json
{
  "message": "Permission was added successfully to user",
  "user": {
    "id": 1,
    "permissions": [
      {
        "id": 1,
        "name": "View All Users",
        "module": "Basic"
      },
      {
        "id": 4,
        "name": "Delete User",
        "module": "Basic"
      }
    ]
  }
}
```

### Example cURL

```bash
curl -X POST http://localhost:8080/api/users/1/permissions \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"id": 4}'
```

---

## DELETE /api/users/{id}/permissions

Remove a permission from a user.

**Permission**: `Update Permissions of User` (Module: `Basic`)

### Request

**Path Parameters:**
- `id` (integer, required): User ID

**Body:**
```json
{
  "id": 4
}
```

### Response

**Success (200 OK):**
```json
{
  "message": "Permission was deleted successfully from user",
  "user": {
    "id": 1,
    "permissions": [
      {
        "id": 1,
        "name": "View All Users",
        "module": "Basic"
      }
    ]
  }
}
```

### Example cURL

```bash
curl -X DELETE http://localhost:8080/api/users/1/permissions \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"id": 4}'
```

---

## PUT /api/users/{id}/roles

Replace all user roles with a new set.

**Permission**: `Update Roles of User` (Module: `Basic`)

### Request

**Path Parameters:**
- `id` (integer, required): User ID

**Body:**
```json
[1, 2]
```

Array of role IDs.

### Response

**Success (200 OK):**
```json
{
  "message": "Roles was updated successfully",
  "user": {
    "id": 1,
    "roles": [
      {
        "id": 1,
        "name": "admin",
        "description": "Administrator role"
      },
      {
        "id": 2,
        "name": "user",
        "description": "Standard user role"
      }
    ]
  }
}
```

### Example cURL

```bash
curl -X PUT http://localhost:8080/api/users/1/roles \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '[1, 2]'
```

---

## POST /api/users/{id}/roles

Add a role to a user.

**Permission**: `Update Roles of User` (Module: `Basic`)

### Request

**Path Parameters:**
- `id` (integer, required): User ID

**Body:**
```json
{
  "id": 3
}
```
JSON include the ID of role.

### Response

**Success (200 OK):**
```json
{
  "message": "Role was added successfully to user",
  "user": {
    "id": 1,
    "roles": [
      {
        "id": 1,
        "name": "admin",
        "description": "Administrator role"
      },
      {
        "id": 3,
        "name": "manager",
        "description": "Manager role"
      }
    ]
  }
}
```

### Example cURL

```bash
curl -X POST http://localhost:8080/api/users/1/roles \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"id": 3}'
```

---

## DELETE /api/users/{id}/roles

Remove a role from a user.

**Permission**: `Update Roles of User` (Module: `Basic`)

### Request

**Path Parameters:**
- `id` (integer, required): User ID

**Body:**
```json
{
  "id": 3
}
```
JSON include the ID of role.

### Response

**Success (200 OK):**
```json
{
  "message": "Role was deleted successfully from user",
  "user": {
    "id": 1,
    "roles": [
      {
        "id": 1,
        "name": "admin",
        "description": "Administrator role"
      }
    ]
  }
}
```

### Example cURL

```bash
curl -X DELETE http://localhost:8080/api/users/1/roles \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"id": 3}'
```

