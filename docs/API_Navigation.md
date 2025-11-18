# Navigation API

All endpoints require JWT authentication and specific permissions from the `Navigation` module.

## Menus

### GET /api/menus

List all menus.

**Permission**: `View Menus` (Module: `Navigation`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "dashboard",
    "title": "Dashboard",
    "uri": "/dashboard"
  },
  {
    "id": 2,
    "name": "users",
    "title": "User Management",
    "uri": "/users"
  }
]
```

### POST /api/menus

Create a menu.

**Permission**: `Create Menu` (Module: `Navigation`)

**Request:**
```json
{
  "name": "settings",
  "title": "Settings",
  "uri": "/settings"
}
```

**Note**: `title` and `uri` are optional fields.

### Standard CRUD

All menus endpoints support:
- `GET /api/menus/{id}` - Get by ID
- `PUT /api/menus/{id}` - Update
- `DELETE /api/menus/{id}` - Delete

---

## Menu Roles

### GET /api/menus/{menuId}/roles

List roles that have access to a menu.

**Permission**: `View Menu Roles` (Module: `Navigation`)

**Response:**
```json
[
  {
    "menu_id": 1,
    "role_id": 1
  },
  {
    "menu_id": 1,
    "role_id": 2
  }
]
```

### POST /api/menus/{menuId}/roles

Grant menu access to a role.

**Permission**: `Add Menu Role` (Module: `Navigation`)

**Request:**
```json
{
  "role_id": 3
}
```

**Response:**
```json
{
  "menu_id": 1,
  "role_id": 3
}
```

### DELETE /api/menus/{menuId}/roles/{roleId}

Revoke menu access from a role.

**Permission**: `Remove Menu Role` (Module: `Navigation`)

**Response:** 204 No Content

---

## Notes

- Menus control navigation visibility based on user roles
- A menu can be accessible to multiple roles
- A role can have access to multiple menus
- The `uri` field defines the route/path for the menu item
- The `title` field is the display name shown in the UI

