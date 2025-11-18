# Help Desk API

All endpoints require JWT authentication and specific permissions from the `Help Desk` module.

## Help Desk Types

### GET /api/help-desk-types

List all help desk types.

**Permission**: `View Help Desk Types` (Module: `Help Desk`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Hardware Issue",
    "description": "Hardware-related problems",
    "team_id": 1,
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

### POST /api/help-desk-types

Create a help desk type.

**Permission**: `Create Help Desk Type` (Module: `Help Desk`)

**Request:**
```json
{
  "name": "Software Issue",
  "description": "Software-related problems",
  "team_id": 1,
  "is_active": true
}
```

### Standard CRUD

All help desk types endpoints support:
- `GET /api/help-desk-types/{id}` - Get by ID
- `PUT /api/help-desk-types/{id}` - Update
- `DELETE /api/help-desk-types/{id}` - Delete

---

## Help Desk Tickets

### GET /api/help-desk

List all help desk tickets.

**Permission**: `View Help Desk Tickets` (Module: `Help Desk`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Laptop won't start",
    "create_date": "2024-01-15",
    "help_desk_type_id": 1,
    "portal_user_id": 1,
    "description": "Laptop shows black screen on boot",
    "state": "open",
    "resolve_date": "2024-01-20",
    "parent_id": null,
    "project_id": null
  }
]
```

### POST /api/help-desk

Create a help desk ticket.

**Permission**: `Create Help Desk Ticket` (Module: `Help Desk`)

**Request:**
```json
{
  "name": "Printer not working",
  "create_date": "2024-12-01",
  "help_desk_type_id": 1,
  "portal_user_id": 2,
  "description": "Printer shows error code E-01",
  "state": "open",
  "resolve_date": "2024-12-05",
  "parent_id": null,
  "project_id": 1
}
```

**Note**: `parent_id` allows linking to parent tickets for ticket hierarchies.

### Standard CRUD

All help desk endpoints support:
- `GET /api/help-desk/{id}` - Get by ID
- `PUT /api/help-desk/{id}` - Update
- `DELETE /api/help-desk/{id}` - Delete

---

## Help Desk Teams

### GET /api/help-desk/teams

List all help desk teams.

**Permission**: `View Help Desk Teams` (Module: `Help Desk`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "IT Support Team",
    "description": "Primary IT support team",
    "user_manager_id": 1,
    "is_active": true,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

### POST /api/help-desk/teams

Create a help desk team.

**Permission**: `Create Help Desk Team` (Module: `Help Desk`)

**Request:**
```json
{
  "name": "Network Support Team",
  "description": "Network infrastructure support",
  "user_manager_id": 2,
  "is_active": true
}
```

### Standard CRUD

All teams endpoints support:
- `GET /api/help-desk/teams/{id}` - Get by ID
- `PUT /api/help-desk/teams/{id}` - Update
- `DELETE /api/help-desk/teams/{id}` - Delete

---

## Team Members

### GET /api/help-desk/teams/{teamId}/members

List team members.

**Permission**: `View Team Members` (Module: `Help Desk`)

**Response:**
```json
[
  {
    "user_id": 1,
    "help_desk_team_id": 1
  },
  {
    "user_id": 2,
    "help_desk_team_id": 1
  }
]
```

### POST /api/help-desk/teams/{teamId}/members

Add member to team.

**Permission**: `Add Team Member` (Module: `Help Desk`)

**Request:**
```json
{
  "user_id": 3
}
```

**Response:**
```json
{
  "help_desk_team_id": 1,
  "user_id": 3
}
```

### DELETE /api/help-desk/teams/{teamId}/members/{userId}

Remove member from team.

**Permission**: `Remove Team Member` (Module: `Help Desk`)

**Response:** 204 No Content

---

## Help Desk Participants

### GET /api/help-desk/{helpDeskId}/participants

List participants in a help desk ticket.

**Permission**: `View Help Desk Participants` (Module: `Help Desk`)

**Response:**
```json
[
  {
    "user_id": 1,
    "help_desk_id": 1
  },
  {
    "user_id": 2,
    "help_desk_id": 1
  }
]
```

### POST /api/help-desk/{helpDeskId}/participants

Add participant to help desk ticket.

**Permission**: `Add Help Desk Participant` (Module: `Help Desk`)

**Request:**
```json
{
  "user_id": 3
}
```

**Response:**
```json
{
  "help_desk_id": 1,
  "user_id": 3
}
```

### DELETE /api/help-desk/{helpDeskId}/participants/{userId}

Remove participant from help desk ticket.

**Permission**: `Remove Help Desk Participant` (Module: `Help Desk`)

**Response:** 204 No Content

---

## Help Desk Ratings

### GET /api/help-desk/ratings

List all ratings.

**Permission**: `View Help Desk Ratings` (Module: `Help Desk`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Excellent service",
    "rate": 5,
    "help_desk_id": 1
  }
]
```

### POST /api/help-desk/ratings

Create a rating.

**Permission**: `Create Help Desk Rating` (Module: `Help Desk`)

**Request:**
```json
{
  "name": "Good response time",
  "rate": 4,
  "help_desk_id": 1
}
```

**Note**: `rate` is typically 1-5 scale.

### Standard CRUD

All ratings endpoints support:
- `GET /api/help-desk/ratings/{id}` - Get by ID
- `PUT /api/help-desk/ratings/{id}` - Update
- `DELETE /api/help-desk/ratings/{id}` - Delete

---

## Help Desk Transactions

### GET /api/help-desk/transactions

List all transactions.

**Permission**: `View Help Desk Transactions` (Module: `Help Desk`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Initial investigation",
    "transaction_type_id": 1,
    "date_time": "2024-01-15T10:30:00Z",
    "time": 2.5,
    "help_desk_id": 1
  }
]
```

**Note**: `time` is in hours (float).

### POST /api/help-desk/transactions

Create a transaction.

**Permission**: `Create Help Desk Transaction` (Module: `Help Desk`)

**Request:**
```json
{
  "name": "Follow-up call",
  "transaction_type_id": 2,
  "date_time": "2024-01-16T14:00:00Z",
  "time": 0.5,
  "help_desk_id": 1
}
```

### Standard CRUD

All transactions endpoints support:
- `GET /api/help-desk/transactions/{id}` - Get by ID
- `PUT /api/help-desk/transactions/{id}` - Update
- `DELETE /api/help-desk/transactions/{id}` - Delete

---

## Transaction Users

### GET /api/help-desk/transactions/{transactionId}/users

List users assigned to a transaction.

**Permission**: `View Transaction Users` (Module: `Help Desk`)

**Response:**
```json
[
  {
    "user_id": 1,
    "transaction_id": 1
  }
]
```

### POST /api/help-desk/transactions/{transactionId}/users

Add user to transaction.

**Permission**: `Add Transaction User` (Module: `Help Desk`)

**Request:**
```json
{
  "user_id": 2
}
```

**Response:**
```json
{
  "transaction_id": 1,
  "user_id": 2
}
```

### DELETE /api/help-desk/transactions/{transactionId}/users/{userId}

Remove user from transaction.

**Permission**: `Remove Transaction User` (Module: `Help Desk`)

**Response:** 204 No Content

---

## Transaction Types

### GET /api/help-desk/transaction-types

List all transaction types.

**Permission**: `View Help Desk Transaction Types` (Module: `Help Desk`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Investigation",
    "expected_time": 2.0
  },
  {
    "id": 2,
    "name": "Resolution",
    "expected_time": 1.5
  }
]
```

**Note**: `expected_time` is in hours (float).

### POST /api/help-desk/transaction-types

Create a transaction type.

**Permission**: `Create Help Desk Transaction Type` (Module: `Help Desk`)

**Request:**
```json
{
  "name": "Follow-up",
  "expected_time": 0.5
}
```

### Standard CRUD

All transaction types endpoints support:
- `GET /api/help-desk/transaction-types/{id}` - Get by ID
- `PUT /api/help-desk/transaction-types/{id}` - Update
- `DELETE /api/help-desk/transaction-types/{id}` - Delete

