# API Quick Reference

## Base URL
```
http://localhost:8080
```

## Authentication
All endpoints except `/login` and `/health` require:
```
Authorization: Bearer <token>
```

## Endpoint Matrix

### Public Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/login` | Authenticate and get token |
| GET | `/health` | Health check |

### User Management (`/api/users`)
| Method | Endpoint | Permission |
|--------|----------|------------|
| GET | `/api/users` | View All Users |
| GET | `/api/users/{id}` | View User |
| POST | `/api/users` | Add User |
| PUT | `/api/users/{id}` | Edit User |
| DELETE | `/api/users/{id}` | Delete User |
| PUT | `/api/users/{id}/suspend` | Suspend User |
| PUT | `/api/users/{id}/permissions` | Update Permissions of User |
| POST | `/api/users/{id}/permissions` | Update Permissions of User |
| DELETE | `/api/users/{id}/permissions` | Update Permissions of User |
| PUT | `/api/users/{id}/roles` | Update Roles of User |
| POST | `/api/users/{id}/roles` | Update Roles of User |
| DELETE | `/api/users/{id}/roles` | Update Roles of User |

### Role Management (`/api/roles`)
| Method | Endpoint | Permission |
|--------|----------|------------|
| GET | `/api/roles` | View All Roles |
| GET | `/api/roles/{id}` | View Role |
| POST | `/api/roles` | Add Role |
| PUT | `/api/roles/{id}` | Edit Role |
| DELETE | `/api/roles/{id}` | Delete Role |
| PUT | `/api/roles/{id}/permissions` | Update Permissions of Role |
| POST | `/api/roles/{id}/permission` | Update Permissions of Role |
| DELETE | `/api/roles/{id}/permission` | Update Permissions of Role |

### Permission Management (`/api/permissions`)
| Method | Endpoint | Permission |
|--------|----------|------------|
| GET | `/api/permissions` | View All Permissions |
| GET | `/api/permissions/{id}` | View Permission |
| POST | `/api/permissions` | Add Permission |
| PUT | `/api/permissions/{id}` | Edit Permission |
| DELETE | `/api/permissions/{id}` | Delete Permission |

### Inventory (`/api/*`)
| Resource | List | View | Create | Update | Delete | Module |
|----------|------|------|--------|--------|--------|--------|
| Brands | View Brands | View Brand | Create Brand | Update Brand | Delete Brand | Inventory |
| Models | View Models | View Model | Create Model | Update Model | Delete Model | Inventory |
| Equipment Types | View Equipment Types | View Equipment Type | Create Equipment Type | Update Equipment Type | Delete Equipment Type | Inventory |
| Operating Systems | View Operating Systems | View Operating System | Create Operating System | Update Operating System | Delete Operating System | Inventory |
| Software Categories | View Software Categories | View Software Category | Create Software Category | Update Software Category | Delete Software Category | Inventory |
| Software | View Software | View Software | Create Software | Update Software | Delete Software | Inventory |
| Equipment | View Equipment | View Equipment | Create Equipment | Update Equipment | Delete Equipment | Inventory |
| Documents | View Documents | View Document | Create Document | Update Document | Delete Document | Inventory |
| Maintenance | View Maintenance Records | View Maintenance Record | Create Maintenance Record | Update Maintenance Record | Delete Maintenance Record | Maintenance |

### Equipment Relationships
| Method | Endpoint | Permission |
|--------|----------|------------|
| GET | `/api/equipment/{id}/software` | View Equipment Software |
| POST | `/api/equipment/{id}/software` | Link Software To Equipment |
| PUT | `/api/equipment/{id}/software/{softwareId}` | Update Equipment Software |
| DELETE | `/api/equipment/{id}/software/{softwareId}` | Unlink Software From Equipment |
| GET | `/api/equipment/{id}/help-desk` | View Equipment Help Desk Links |
| POST | `/api/equipment/{id}/help-desk` | Link Equipment To Help Desk |
| DELETE | `/api/equipment/{id}/help-desk/{helpDeskId}` | Unlink Equipment From Help Desk |
| GET | `/api/equipment/{id}/user-history` | View Equipment User History |
| POST | `/api/equipment/{id}/user-history` | Create Equipment User History |
| PUT | `/api/equipment/{id}/user-history/{userId}/{startDate}` | Update Equipment User History |
| DELETE | `/api/equipment/{id}/user-history/{userId}/{startDate}` | Delete Equipment User History |

### Geography (`/api/*`)
| Resource | List | View | Create | Update | Delete | Module |
|----------|------|------|--------|--------|--------|--------|
| Countries | View Countries | View Country | Create Country | Update Country | Delete Country | Directory |
| Cities | View Cities | View City | Create City | Update City | Delete City | Directory |
| Locations | View Locations | View Location | Create Location | Update Location | Delete Location | Directory |
| Contacts | View Contacts | View Contact | Create Contact | Update Contact | Delete Contact | Directory |

### Help Desk (`/api/help-desk/*`)
| Resource | List | View | Create | Update | Delete | Module |
|----------|------|------|--------|--------|--------|--------|
| Help Desk Types | View Help Desk Types | View Help Desk Types | Create Help Desk Type | Update Help Desk Type | Delete Help Desk Type | Help Desk |
| Tickets | View Help Desk Tickets | View Help Desk Ticket | Create Help Desk Ticket | Update Help Desk Ticket | Delete Help Desk Ticket | Help Desk |
| Teams | View Help Desk Teams | View Help Desk Team | Create Help Desk Team | Update Help Desk Team | Delete Help Desk Team | Help Desk |
| Ratings | View Help Desk Ratings | View Help Desk Rating | Create Help Desk Rating | Update Help Desk Rating | Delete Help Desk Rating | Help Desk |
| Transactions | View Help Desk Transactions | View Help Desk Transaction | Create Help Desk Transaction | Update Help Desk Transaction | Delete Help Desk Transaction | Help Desk |
| Transaction Types | View Help Desk Transaction Types | View Help Desk Transaction Type | Create Help Desk Transaction Type | Update Help Desk Transaction Type | Delete Help Desk Transaction Type | Help Desk |

### Help Desk Relationships
| Method | Endpoint | Permission |
|--------|----------|------------|
| GET | `/api/help-desk/teams/{id}/members` | View Team Members |
| POST | `/api/help-desk/teams/{id}/members` | Add Team Member |
| DELETE | `/api/help-desk/teams/{id}/members/{userId}` | Remove Team Member |
| GET | `/api/help-desk/{id}/participants` | View Help Desk Participants |
| POST | `/api/help-desk/{id}/participants` | Add Help Desk Participant |
| DELETE | `/api/help-desk/{id}/participants/{userId}` | Remove Help Desk Participant |
| GET | `/api/help-desk/transactions/{id}/users` | View Transaction Users |
| POST | `/api/help-desk/transactions/{id}/users` | Add Transaction User |
| DELETE | `/api/help-desk/transactions/{id}/users/{userId}` | Remove Transaction User |

### Navigation (`/api/menus`)
| Method | Endpoint | Permission |
|--------|----------|------------|
| GET | `/api/menus` | View Menus |
| GET | `/api/menus/{id}` | View Menu |
| POST | `/api/menus` | Create Menu |
| PUT | `/api/menus/{id}` | Update Menu |
| DELETE | `/api/menus/{id}` | Delete Menu |
| GET | `/api/menus/{id}/roles` | View Menu Roles |
| POST | `/api/menus/{id}/roles` | Add Menu Role |
| DELETE | `/api/menus/{id}/roles/{roleId}` | Remove Menu Role |

### Admin
| Method | Endpoint | Requirement |
|--------|----------|-------------|
| GET | `/admin` | admin role |

## HTTP Status Codes

| Code | Meaning |
|------|---------|
| 200 | OK - Request successful |
| 201 | Created - Resource created |
| 204 | No Content - Resource deleted |
| 400 | Bad Request - Invalid input |
| 401 | Unauthorized - Missing/invalid token |
| 403 | Forbidden - Insufficient permissions |
| 404 | Not Found - Resource not found |
| 500 | Internal Server Error - Server error |

## Common Request Patterns

### Create Resource
```bash
curl -X POST http://localhost:8080/api/{resource} \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{...}'
```

### Update Resource
```bash
curl -X PUT http://localhost:8080/api/{resource}/{id} \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{...}'
```

### Delete Resource
```bash
curl -X DELETE http://localhost:8080/api/{resource}/{id} \
  -H "Authorization: Bearer <token>"
```

## Permission Modules

- **Basic**: User, Role, Permission management
- **Inventory**: Equipment, Software, Brands, Models, etc.
- **Directory**: Countries, Cities, Locations, Contacts
- **Help Desk**: Tickets, Teams, Transactions, Ratings
- **Navigation**: Menu management
- **Maintenance**: Maintenance records

## Notes

- All list endpoints return arrays
- All single resource endpoints return objects
- Update endpoints support partial updates (only send fields to change)
- Relationship endpoints manage many-to-many associations
- Date formats: `YYYY-MM-DD` for dates, ISO 8601 for datetimes
- All IDs are integers

