# Geography API

All endpoints require JWT authentication and specific permissions from the `Directory` module.

## Countries

### GET /api/countries

List all countries.

**Permission**: `View Countries` (Module: `Directory`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "United States",
    "code": "US"
  },
  {
    "id": 2,
    "name": "United Kingdom",
    "code": "GB"
  }
]
```

### POST /api/countries

Create a country.

**Permission**: `Create Country` (Module: `Directory`)

**Request:**
```json
{
  "name": "Canada",
  "code": "CA"
}
```

### Standard CRUD

All countries endpoints support:
- `GET /api/countries/{id}` - Get by ID
- `PUT /api/countries/{id}` - Update
- `DELETE /api/countries/{id}` - Delete

---

## Cities

### GET /api/cities

List all cities.

**Permission**: `View Cities` (Module: `Directory`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "New York",
    "country_id": 1
  },
  {
    "id": 2,
    "name": "London",
    "country_id": 2
  }
]
```

### POST /api/cities

Create a city.

**Permission**: `Create City` (Module: `Directory`)

**Request:**
```json
{
  "name": "Toronto",
  "country_id": 3
}
```

### Standard CRUD

All cities endpoints support:
- `GET /api/cities/{id}` - Get by ID
- `PUT /api/cities/{id}` - Update
- `DELETE /api/cities/{id}` - Delete

---

## Locations

### GET /api/locations

List all locations.

**Permission**: `View Locations` (Module: `Directory`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Headquarters Building A",
    "zip_code": "10001",
    "state": "NY",
    "building_number": 1,
    "room_number": 101,
    "latitude": 40.7128,
    "longitude": -74.0060,
    "country_id": 1,
    "city_id": 1,
    "location_map": "Building A, Floor 1, Room 101"
  }
]
```

### POST /api/locations

Create a location.

**Permission**: `Create Location` (Module: `Directory`)

**Request:**
```json
{
  "name": "Branch Office B",
  "zip_code": "90210",
  "state": "CA",
  "building_number": 2,
  "room_number": 205,
  "latitude": 34.0522,
  "longitude": -118.2437,
  "country_id": 1,
  "city_id": 5,
  "location_map": "Building B, Floor 2, Room 205"
}
```

### Standard CRUD

All locations endpoints support:
- `GET /api/locations/{id}` - Get by ID
- `PUT /api/locations/{id}` - Update
- `DELETE /api/locations/{id}` - Delete

---

## Contacts

### GET /api/contacts

List all contacts.

**Permission**: `View Contacts` (Module: `Directory`)

**Response:**
```json
[
  {
    "id": 1,
    "name": "Tech Supplies Inc.",
    "country_id": 1,
    "mobile_number": "+1234567890",
    "phone_number": "+1234567891",
    "website": "https://techsupplies.com",
    "email": "contact@techsupplies.com"
  }
]
```

### POST /api/contacts

Create a contact.

**Permission**: `Create Contact` (Module: `Directory`)

**Request:**
```json
{
  "name": "Hardware Solutions Ltd.",
  "country_id": 2,
  "mobile_number": "+441234567890",
  "phone_number": "+441234567891",
  "website": "https://hardwaresolutions.co.uk",
  "email": "info@hardwaresolutions.co.uk"
}
```

### Standard CRUD

All contacts endpoints support:
- `GET /api/contacts/{id}` - Get by ID
- `PUT /api/contacts/{id}` - Update
- `DELETE /api/contacts/{id}` - Delete

---

## Notes

- All geography resources are reference data used by other entities
- Countries are linked to users (nationality), cities, locations, and contacts
- Locations include GPS coordinates (latitude/longitude) for mapping
- Contacts represent suppliers or partners

