# Inventory Management API

All endpoints require JWT authentication and specific permissions from the `Inventory` or `Maintenance` modules.

## Brands

### GET /api/brands

List all brands.

**Permission**: `View Brands` (Module: `Inventory`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Dell"
  },
  {
    "id": 2,
    "name": "HP"
  }
]
```

### POST /api/brands

Create a brand.

**Permission**: `Create Brand` (Module: `Inventory`)

**Request:**
```json
{
  "name": "Lenovo"
}
```

**Response:**
```json
{
  "id": 3,
  "name": "Lenovo"
}
```

---

## Models

### GET /api/models

List all models.

**Permission**: `View Models` (Module: `Inventory`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Latitude 5520",
    "brand_id": 1
  }
]
```

### POST /api/models

Create a model.

**Permission**: `Create Model` (Module: `Inventory`)

**Request:**
```json
{
  "name": "ThinkPad X1 Carbon",
  "brand_id": 3
}
```

---

## Equipment Types

### GET /api/equipment-types

List all equipment types.

**Permission**: `View Equipment Types` (Module: `Inventory`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Laptop",
    "description": "Portable computer"
  },
  {
    "id": 2,
    "name": "Desktop",
    "description": "Desktop computer"
  }
]
```

### POST /api/equipment-types

Create an equipment type.

**Permission**: `Create Equipment Type` (Module: `Inventory`)

**Request:**
```json
{
  "name": "Server",
  "description": "Server hardware"
}
```

---

## Operating Systems

### GET /api/operating-systems

List all operating systems.

**Permission**: `View Operating Systems` (Module: `Inventory`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Windows",
    "description": "Microsoft Windows",
    "version": "11",
    "architectures": "x64, ARM64"
  },
  {
    "id": 2,
    "name": "Linux",
    "description": "Linux distribution",
    "version": "Ubuntu 22.04",
    "architectures": "x64, ARM64"
  }
]
```

### POST /api/operating-systems

Create an operating system.

**Permission**: `Create Operating System` (Module: `Inventory`)

**Request:**
```json
{
  "name": "macOS",
  "description": "Apple macOS",
  "version": "Ventura",
  "architectures": "ARM64, x64"
}
```

---

## Software Categories

### GET /api/software-categories

List all software categories.

**Permission**: `View Software Categories` (Module: `Inventory`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Productivity"
  },
  {
    "id": 2,
    "name": "Development"
  }
]
```

### POST /api/software-categories

Create a software category.

**Permission**: `Create Software Category` (Module: `Inventory`)

**Request:**
```json
{
  "name": "Security"
}
```

---

## Software

### GET /api/software

List all software.

**Permission**: `View Software` (Module: `Inventory`)

**Response:**
```json
[
  {
    "id": 1,
    "software_name": "Microsoft Office",
    "category_id": 1,
    "license_exp_date": 1735689600
  }
]
```

**Note**: `license_exp_date` is a Unix timestamp (seconds since epoch).

### POST /api/software

Create software.

**Permission**: `Create Software` (Module: `Inventory`)

**Request:**
```json
{
  "software_name": "Visual Studio Code",
  "category_id": 2,
  "license_exp_date": 1735689600
}
```

---

## Equipment

### GET /api/equipment

List all equipment.

**Permission**: `View Equipment` (Module: `Inventory`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Laptop-001",
    "description": "Dell Latitude 5520",
    "state": true,
    "serial_number": "SN123456",
    "equipment_type_id": 1,
    "model_id": 1,
    "production_date": "2023-01-15",
    "ip_v4": "192.168.1.100",
    "ip_v6": "2001:0db8::1",
    "mac_address": "00:1B:44:11:3A:B7",
    "operating_system_id": 1,
    "location_id": 1,
    "processors": "Intel Core i7",
    "country_of_region": 1,
    "user_id": 1,
    "warranty_start_date": "2023-01-15",
    "warranty_end_date": "2026-01-15",
    "supplier_id": 1
  }
]
```

### POST /api/equipment

Create equipment.

**Permission**: `Create Equipment` (Module: `Inventory`)

**Request:**
```json
{
  "name": "Laptop-002",
  "description": "HP EliteBook 850",
  "state": true,
  "serial_number": "SN789012",
  "equipment_type_id": 1,
  "model_id": 2,
  "production_date": "2024-03-20",
  "ip_v4": "192.168.1.101",
  "ip_v6": "2001:0db8::2",
  "mac_address": "00:1B:44:11:3A:B8",
  "operating_system_id": 1,
  "location_id": 1,
  "processors": "Intel Core i5",
  "country_of_region": 1,
  "user_id": 2,
  "warranty_start_date": "2024-03-20",
  "warranty_end_date": "2027-03-20",
  "supplier_id": 1
}
```

---

## Equipment Software Relationships

### GET /api/equipment/{equipmentId}/software

List software installed on equipment.

**Permission**: `View Equipment Software` (Module: `Inventory`)

**Response:**
```json
[
  {
    "software_id": 1,
    "equipment_id": 1,
    "software_version": "2021",
    "software_size": 2048,
    "license": "Enterprise",
    "software": {
      "id": 1,
      "software_name": "Microsoft Office",
      "category_id": 1,
      "license_exp_date": 1735689600
    }
  }
]
```

### POST /api/equipment/{equipmentId}/software

Link software to equipment.

**Permission**: `Link Software To Equipment` (Module: `Inventory`)

**Request:**
```json
{
  "software_id": 2,
  "software_version": "1.85.0",
  "software_size": 512,
  "license": "MIT"
}
```

**Response:**
```json
{
  "software_id": 2,
  "equipment_id": 1,
  "software_version": "1.85.0",
  "software_size": 512,
  "license": "MIT"
}
```

### PUT /api/equipment/{equipmentId}/software/{softwareId}

Update software link details.

**Permission**: `Update Equipment Software` (Module: `Inventory`)

**Request:**
```json
{
  "software_version": "1.86.0",
  "license": "MIT v2"
}
```

### DELETE /api/equipment/{equipmentId}/software/{softwareId}

Unlink software from equipment.

**Permission**: `Unlink Software From Equipment` (Module: `Inventory`)

**Response:** 204 No Content

---

## Equipment Help Desk Links

### GET /api/equipment/{equipmentId}/help-desk

List help desk tickets linked to equipment.

**Permission**: `View Equipment Help Desk Links` (Module: `Inventory`)

**Response:**
```json
[
  {
    "equipment_id": 1,
    "help_desk_id": 5
  }
]
```

### POST /api/equipment/{equipmentId}/help-desk

Link equipment to help desk ticket.

**Permission**: `Link Equipment To Help Desk` (Module: `Inventory`)

**Request:**
```json
{
  "help_desk_id": 6
}
```

**Response:**
```json
{
  "equipment_id": 1,
  "help_desk_id": 6
}
```

### DELETE /api/equipment/{equipmentId}/help-desk/{helpDeskId}

Unlink equipment from help desk ticket.

**Permission**: `Unlink Equipment From Help Desk` (Module: `Inventory`)

**Response:** 204 No Content

---

## Equipment User History

### GET /api/equipment/{equipmentId}/user-history

List user assignment history for equipment.

**Permission**: `View Equipment User History` (Module: `Inventory`)

**Response:**
```json
[
  {
    "equipment_id": 1,
    "user_id": 1,
    "start_date": "2024-01-01",
    "end_date": "2024-06-30",
    "comment": "Initial assignment"
  },
  {
    "equipment_id": 1,
    "user_id": 2,
    "start_date": "2024-07-01",
    "end_date": "2024-12-31",
    "comment": "Reassigned"
  }
]
```

### POST /api/equipment/{equipmentId}/user-history

Add user history entry.

**Permission**: `Create Equipment User History` (Module: `Inventory`)

**Request:**
```json
{
  "user_id": 3,
  "start_date": "2025-01-01",
  "end_date": "2025-12-31",
  "comment": "New assignment"
}
```

**Note**: `end_date` must be after `start_date`.

### PUT /api/equipment/{equipmentId}/user-history/{userId}/{startDate}

Update user history entry.

**Permission**: `Update Equipment User History` (Module: `Inventory`)

**Request:**
```json
{
  "end_date": "2025-06-30",
  "comment": "Updated assignment"
}
```

**Note**: Date format in URL: `YYYY-MM-DD`

### DELETE /api/equipment/{equipmentId}/user-history/{userId}/{startDate}

Delete user history entry.

**Permission**: `Delete Equipment User History` (Module: `Inventory`)

**Response:** 204 No Content

---

## Documents

### GET /api/documents

List all documents.

**Permission**: `View Documents` (Module: `Inventory`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Invoice-001.pdf",
    "document_size": "2.5MB",
    "picture": "base64encodeddata...",
    "document_type": "invoice",
    "equipment_id": 1,
    "supplier_id": 1,
    "help_desk_id": 0
  }
]
```

**Note**: `picture` field contains base64-encoded binary data.

### POST /api/documents

Create a document.

**Permission**: `Create Document` (Module: `Inventory`)

**Request:**
```json
{
  "name": "Warranty-001.pdf",
  "document_size": "1.2MB",
  "picture": "base64encodeddata...",
  "document_type": "warranty",
  "equipment_id": 1,
  "supplier_id": 1,
  "help_desk_id": 0
}
```

**Note**: Encode binary file content as base64 string for the `picture` field.

---

## Maintenance

### GET /api/maintenance

List all maintenance records.

**Permission**: `View Maintenance Records` (Module: `Maintenance`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Quarterly Maintenance",
    "start_date": "2024-01-01",
    "end_date": "2024-01-15",
    "contact_id": 1
  }
]
```

### POST /api/maintenance

Create a maintenance record.

**Permission**: `Create Maintenance Record` (Module: `Maintenance`)

**Request:**
```json
{
  "name": "Annual Service",
  "start_date": "2024-12-01",
  "end_date": "2024-12-05",
  "contact_id": 2
}
```

---

## Standard CRUD Operations

All inventory resources support standard REST operations:

- **GET /api/{resource}** - List all
- **GET /api/{resource}/{id}** - Get by ID
- **POST /api/{resource}** - Create
- **PUT /api/{resource}/{id}** - Update (partial updates supported)
- **DELETE /api/{resource}/{id}** - Delete

All require appropriate permissions from the `Inventory` or `Maintenance` modules.

