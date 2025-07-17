# POS Bengkel (Automotive Service Shop POS System)

A comprehensive Point of Sale system for automotive service shops built with Go, Fiber, and PostgreSQL.

## 📋 Table of Contents
- [Architecture](#architecture)
- [🚀 API Documentation](#-api-documentation)
  - [Base URL & Authentication](#base-url--authentication)
  - [Response Format](#response-format)
  - [Foundation APIs](#foundation-apis)
  - [Customer Management APIs](#customer-management-apis)
  - [Inventory Management APIs](#inventory-management-apis)
  - [Service Management APIs](#service-management-apis)
  - [Financial Management APIs](#financial-management-apis)
- [Database Schema](#database-schema)
- [Getting Started](#getting-started)

## Architecture

This project follows Clean Architecture principles with the following layers:

- **Models**: Database entities and business models
- **Repository**: Data access layer with GORM
- **Usecase**: Business logic layer
- **Delivery**: HTTP handlers and API endpoints

### Service Job & Queue Management Quick Reference

#### Service Job Status Workflow
```
Customer Request → Antri → Dikerjakan → Selesai → Diambil
```

#### Key Queue Management Endpoints
```bash
# Get complete queue for outlet
GET /api/v1/queue/{outlet_id}

# Get today's queue  
GET /api/v1/queue/{outlet_id}/today

# Reorder queue for priority handling
PUT /api/v1/queue/{outlet_id}/reorder
```

#### Enhanced Service Job Status Update
```bash
# Update status with technician assignment
PUT /api/v1/service-jobs/{id}/status
{
  "status": "Dikerjakan",
  "user_id": 1,
  "technician_id": 1,
  "notes": "Started maintenance work"
}
```

For complete service job and queue management documentation, see [SERVICE_JOB_QUEUE_DOCS.md](SERVICE_JOB_QUEUE_DOCS.md).

### Testing the Implementation

#### Comprehensive Service Job & Queue Test
```bash
./test_service_job_queue.sh
```

#### Module-Specific Tests
```bash
./test_api.sh              # Foundation APIs
./test_customer_api.sh      # Customer Management
./test_inventory_api.sh     # Inventory Management
./test_service_api.sh       # Service Management
./test_financial_api.sh     # Financial Management
```

---

## 🚀 API Documentation

Complete API reference for POS Bengkel system integration with Flutter applications.

## Base URL & Authentication

**Base URL**: `http://localhost:3000`

**Content-Type**: `application/json`

> **Note**: Authentication endpoints are under development. Currently, all endpoints are accessible without authentication.

## Response Format

All API responses follow a consistent structure:

### Success Response
```json
{
  "status": "success",
  "message": "Operation completed successfully",
  "data": {
    // Response data object
  }
}
```

### Paginated Response
```json
{
  "status": "success",
  "message": "Data retrieved successfully",
  "data": [
    // Array of objects
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "pages": 10
  }
}
```

### Error Response
```json
{
  "status": "error",
  "message": "Error description",
  "error": "Detailed error message"
}
```

## Status Codes

- `200` - OK
- `201` - Created
- `400` - Bad Request
- `404` - Not Found
- `500` - Internal Server Error

---

## Foundation APIs

### Health Check

#### GET /health
Check API health status.

**Response:**
```json
{
  "status": "success",
  "message": "API is running",
  "data": {
    "service": "POS Bengkel API",
    "version": "1.0.0",
    "timestamp": "2024-01-01T10:00:00Z"
  }
}
```

### Users Management

#### POST /api/v1/users
Create a new user.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "outlet_id": 1
}
```

**Validation Rules:**
- `name`: required, min 2 characters
- `email`: required, valid email format
- `password`: required, min 6 characters
- `outlet_id`: optional, must exist in outlets table

**Response:**
```json
{
  "status": "success",
  "message": "User created successfully",
  "data": {
    "user_id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "outlet_id": 1,
    "outlet": {
      "outlet_id": 1,
      "outlet_name": "Main Workshop",
      "branch_type": "Pusat",
      "city": "Jakarta"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/users
List all users with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Users retrieved successfully",
  "data": [
    {
      "user_id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "outlet_id": 1,
      "outlet": {
        "outlet_id": 1,
        "outlet_name": "Main Workshop",
        "branch_type": "Pusat",
        "city": "Jakarta"
      },
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

#### GET /api/v1/users/:id
Get user by ID.

**Path Parameters:**
- `id`: User ID

**Response:**
```json
{
  "status": "success",
  "message": "User retrieved successfully",
  "data": {
    "user_id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "outlet_id": 1,
    "outlet": {
      "outlet_id": 1,
      "outlet_name": "Main Workshop",
      "branch_type": "Pusat",
      "city": "Jakarta"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### PUT /api/v1/users/:id
Update user information.

**Path Parameters:**
- `id`: User ID

**Request Body:**
```json
{
  "name": "John Doe Updated",
  "email": "john.updated@example.com",
  "outlet_id": 2
}
```

**Response:**
```json
{
  "status": "success",
  "message": "User updated successfully",
  "data": {
    "user_id": 1,
    "name": "John Doe Updated",
    "email": "john.updated@example.com",
    "outlet_id": 2,
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T11:00:00Z"
  }
}
```

#### DELETE /api/v1/users/:id
Delete user (soft delete).

**Path Parameters:**
- `id`: User ID

**Response:**
```json
{
  "status": "success",
  "message": "User deleted successfully"
}
```

### Outlets Management

#### POST /api/v1/outlets
Create a new outlet.

**Request Body:**
```json
{
  "outlet_name": "Main Workshop",
  "branch_type": "Pusat",
  "city": "Jakarta",
  "address": "Jl. Merdeka No. 123",
  "phone_number": "021-12345678",
  "status": "Aktif"
}
```

**Validation Rules:**
- `outlet_name`: required
- `branch_type`: required
- `city`: required
- `address`: optional
- `phone_number`: optional
- `status`: optional (default: "Aktif")

**Response:**
```json
{
  "status": "success",
  "message": "Outlet created successfully",
  "data": {
    "outlet_id": 1,
    "outlet_name": "Main Workshop",
    "branch_type": "Pusat",
    "city": "Jakarta",
    "address": "Jl. Merdeka No. 123",
    "phone_number": "021-12345678",
    "status": "Aktif",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/outlets
List all outlets with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Outlets retrieved successfully",
  "data": [
    {
      "outlet_id": 1,
      "outlet_name": "Main Workshop",
      "branch_type": "Pusat",
      "city": "Jakarta",
      "address": "Jl. Merdeka No. 123",
      "phone_number": "021-12345678",
      "status": "Aktif",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

#### GET /api/v1/outlets/:id
Get outlet by ID.

**Path Parameters:**
- `id`: Outlet ID

**Response:**
```json
{
  "status": "success",
  "message": "Outlet retrieved successfully",
  "data": {
    "outlet_id": 1,
    "outlet_name": "Main Workshop",
    "branch_type": "Pusat",
    "city": "Jakarta",
    "address": "Jl. Merdeka No. 123",
    "phone_number": "021-12345678",
    "status": "Aktif",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

---

## Customer Management APIs

### Customers

#### POST /api/v1/customers
Create a new customer.

**Request Body:**
```json
{
  "name": "John Doe",
  "phone_number": "081234567890",
  "address": "Jl. Sudirman No. 456",
  "status": "Aktif"
}
```

**Validation Rules:**
- `name`: required, min 2 characters, max 255 characters
- `phone_number`: required, min 10 characters, max 20 characters, unique
- `address`: optional
- `status`: optional (default: "Aktif")

**Response:**
```json
{
  "status": "success",
  "message": "Customer created successfully",
  "data": {
    "customer_id": 1,
    "name": "John Doe",
    "phone_number": "081234567890",
    "address": "Jl. Sudirman No. 456",
    "status": "Aktif",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/customers
List all customers with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Customers retrieved successfully",
  "data": [
    {
      "customer_id": 1,
      "name": "John Doe",
      "phone_number": "081234567890",
      "address": "Jl. Sudirman No. 456",
      "status": "Aktif",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

#### GET /api/v1/customers/:id
Get customer by ID.

**Path Parameters:**
- `id`: Customer ID

**Response:**
```json
{
  "status": "success",
  "message": "Customer retrieved successfully",
  "data": {
    "customer_id": 1,
    "name": "John Doe",
    "phone_number": "081234567890",
    "address": "Jl. Sudirman No. 456",
    "status": "Aktif",
    "vehicles": [
      {
        "vehicle_id": 1,
        "plate_number": "B1234XYZ",
        "brand": "Toyota",
        "model": "Avanza",
        "production_year": 2020
      }
    ],
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### PUT /api/v1/customers/:id
Update customer information.

**Path Parameters:**
- `id`: Customer ID

**Request Body:**
```json
{
  "name": "John Doe Updated",
  "address": "Jl. Sudirman No. 456 Updated"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Customer updated successfully",
  "data": {
    "customer_id": 1,
    "name": "John Doe Updated",
    "phone_number": "081234567890",
    "address": "Jl. Sudirman No. 456 Updated",
    "status": "Aktif",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T11:00:00Z"
  }
}
```

#### DELETE /api/v1/customers/:id
Delete customer (soft delete).

**Path Parameters:**
- `id`: Customer ID

**Response:**
```json
{
  "status": "success",
  "message": "Customer deleted successfully"
}
```

#### GET /api/v1/customers/search
Search customers by name or phone number.

**Query Parameters:**
- `q`: Search query
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Customers found",
  "data": [
    {
      "customer_id": 1,
      "name": "John Doe",
      "phone_number": "081234567890",
      "address": "Jl. Sudirman No. 456",
      "status": "Aktif",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

### Customer Vehicles

#### POST /api/v1/customer-vehicles
Create a new customer vehicle.

**Request Body:**
```json
{
  "customer_id": 1,
  "plate_number": "B1234XYZ",
  "brand": "Toyota",
  "model": "Avanza",
  "type": "MPV",
  "production_year": 2020,
  "chassis_number": "CH1234567890123456",
  "engine_number": "ENG1234567890",
  "color": "Silver",
  "notes": "Customer vehicle in good condition"
}
```

**Validation Rules:**
- `customer_id`: required, must exist in customers table
- `plate_number`: required, min 3 characters, max 20 characters, unique
- `brand`: required, min 2 characters, max 100 characters
- `model`: required, min 2 characters, max 100 characters
- `type`: required, min 2 characters, max 100 characters
- `production_year`: required, between 1900 and 2030
- `chassis_number`: required, min 10 characters, max 100 characters, unique
- `engine_number`: required, min 5 characters, max 100 characters, unique
- `color`: required, min 2 characters, max 50 characters
- `notes`: optional

**Response:**
```json
{
  "status": "success",
  "message": "Customer vehicle created successfully",
  "data": {
    "vehicle_id": 1,
    "customer_id": 1,
    "plate_number": "B1234XYZ",
    "brand": "Toyota",
    "model": "Avanza",
    "type": "MPV",
    "production_year": 2020,
    "chassis_number": "CH1234567890123456",
    "engine_number": "ENG1234567890",
    "color": "Silver",
    "notes": "Customer vehicle in good condition",
    "customer": {
      "customer_id": 1,
      "name": "John Doe",
      "phone_number": "081234567890"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/customer-vehicles
List all customer vehicles with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Customer vehicles retrieved successfully",
  "data": [
    {
      "vehicle_id": 1,
      "customer_id": 1,
      "plate_number": "B1234XYZ",
      "brand": "Toyota",
      "model": "Avanza",
      "type": "MPV",
      "production_year": 2020,
      "chassis_number": "CH1234567890123456",
      "engine_number": "ENG1234567890",
      "color": "Silver",
      "notes": "Customer vehicle in good condition",
      "customer": {
        "customer_id": 1,
        "name": "John Doe",
        "phone_number": "081234567890"
      },
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

#### GET /api/v1/customers/:id/vehicles
Get all vehicles for a specific customer.

**Path Parameters:**
- `id`: Customer ID

**Response:**
```json
{
  "status": "success",
  "message": "Customer vehicles retrieved successfully",
  "data": [
    {
      "vehicle_id": 1,
      "customer_id": 1,
      "plate_number": "B1234XYZ",
      "brand": "Toyota",
      "model": "Avanza",
      "type": "MPV",
      "production_year": 2020,
      "chassis_number": "CH1234567890123456",
      "engine_number": "ENG1234567890",
      "color": "Silver",
      "notes": "Customer vehicle in good condition",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ]
}
```

---

## Inventory Management APIs

### Categories

#### POST /api/v1/categories
Create a new category.

**Request Body:**
```json
{
  "name": "Spare Parts",
  "status": "Aktif"
}
```

**Validation Rules:**
- `name`: required
- `status`: optional (default: "Aktif")

**Response:**
```json
{
  "status": "success",
  "message": "Category created successfully",
  "data": {
    "category_id": 1,
    "name": "Spare Parts",
    "status": "Aktif",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/categories
List all categories with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Categories retrieved successfully",
  "data": [
    {
      "category_id": 1,
      "name": "Spare Parts",
      "status": "Aktif",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

#### GET /api/v1/categories/:id
Get category by ID.

**Path Parameters:**
- `id`: Category ID

**Response:**
```json
{
  "status": "success",
  "message": "Category retrieved successfully",
  "data": {
    "category_id": 1,
    "name": "Spare Parts",
    "status": "Aktif",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/categories/:id/products
Get all products in a category.

**Path Parameters:**
- `id`: Category ID

**Response:**
```json
{
  "status": "success",
  "message": "Category products retrieved successfully",
  "data": [
    {
      "product_id": 1,
      "product_name": "Brake Pad Toyota Avanza",
      "sku": "BP-TOY-AVZ-001",
      "selling_price": 200000,
      "stock": 25,
      "category_id": 1
    }
  ]
}
```

### Suppliers

#### POST /api/v1/suppliers
Create a new supplier.

**Request Body:**
```json
{
  "supplier_name": "PT Auto Parts Indonesia",
  "contact_person_name": "Budi Santoso",
  "phone_number": "021-87654321",
  "address": "Jl. Industri No. 45, Jakarta",
  "status": "Aktif"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Supplier created successfully",
  "data": {
    "supplier_id": 1,
    "supplier_name": "PT Auto Parts Indonesia",
    "contact_person_name": "Budi Santoso",
    "phone_number": "021-87654321",
    "address": "Jl. Industri No. 45, Jakarta",
    "status": "Aktif",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/suppliers
List all suppliers with pagination.

**Response:**
```json
{
  "status": "success",
  "message": "Suppliers retrieved successfully",
  "data": [
    {
      "supplier_id": 1,
      "supplier_name": "PT Auto Parts Indonesia",
      "contact_person_name": "Budi Santoso",
      "phone_number": "021-87654321",
      "address": "Jl. Industri No. 45, Jakarta",
      "status": "Aktif",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

#### GET /api/v1/suppliers/:id/products
Get all products from a supplier.

**Path Parameters:**
- `id`: Supplier ID

**Response:**
```json
{
  "status": "success",
  "message": "Supplier products retrieved successfully",
  "data": [
    {
      "product_id": 1,
      "product_name": "Brake Pad Toyota Avanza",
      "sku": "BP-TOY-AVZ-001",
      "selling_price": 200000,
      "stock": 25,
      "supplier_id": 1
    }
  ]
}
```

### Unit Types

#### POST /api/v1/unit-types
Create a new unit type.

**Request Body:**
```json
{
  "name": "Pieces",
  "status": "Aktif"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Unit type created successfully",
  "data": {
    "unit_type_id": 1,
    "name": "Pieces",
    "status": "Aktif",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/unit-types
List all unit types.

**Response:**
```json
{
  "status": "success",
  "message": "Unit types retrieved successfully",
  "data": [
    {
      "unit_type_id": 1,
      "name": "Pieces",
      "status": "Aktif",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ]
}
```

### Products

#### POST /api/v1/products
Create a new product.

**Request Body:**
```json
{
  "product_name": "Brake Pad Toyota Avanza",
  "product_description": "High quality brake pad for Toyota Avanza",
  "cost_price": 150000,
  "selling_price": 200000,
  "stock": 25,
  "sku": "BP-TOY-AVZ-001",
  "barcode": "1234567890123",
  "has_serial_number": false,
  "shelf_location": "A1-B2",
  "usage_status": "Jual",
  "is_active": true,
  "category_id": 1,
  "supplier_id": 1,
  "unit_type_id": 1
}
```

**Validation Rules:**
- `product_name`: required
- `cost_price`: required, must be positive number
- `selling_price`: required, must be positive number
- `stock`: required, must be non-negative integer
- `sku`: required, unique
- `category_id`: required, must exist
- `supplier_id`: required, must exist
- `unit_type_id`: required, must exist

**Response:**
```json
{
  "status": "success",
  "message": "Product created successfully",
  "data": {
    "product_id": 1,
    "product_name": "Brake Pad Toyota Avanza",
    "product_description": "High quality brake pad for Toyota Avanza",
    "cost_price": 150000,
    "selling_price": 200000,
    "stock": 25,
    "sku": "BP-TOY-AVZ-001",
    "barcode": "1234567890123",
    "has_serial_number": false,
    "shelf_location": "A1-B2",
    "usage_status": "Jual",
    "is_active": true,
    "category": {
      "category_id": 1,
      "name": "Spare Parts"
    },
    "supplier": {
      "supplier_id": 1,
      "supplier_name": "PT Auto Parts Indonesia"
    },
    "unit_type": {
      "unit_type_id": 1,
      "name": "Pieces"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/products
List all products with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Products retrieved successfully",
  "data": [
    {
      "product_id": 1,
      "product_name": "Brake Pad Toyota Avanza",
      "cost_price": 150000,
      "selling_price": 200000,
      "stock": 25,
      "sku": "BP-TOY-AVZ-001",
      "category": {
        "category_id": 1,
        "name": "Spare Parts"
      },
      "supplier": {
        "supplier_id": 1,
        "supplier_name": "PT Auto Parts Indonesia"
      },
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

#### GET /api/v1/products/:id
Get product by ID.

**Path Parameters:**
- `id`: Product ID

**Response:**
```json
{
  "status": "success",
  "message": "Product retrieved successfully",
  "data": {
    "product_id": 1,
    "product_name": "Brake Pad Toyota Avanza",
    "product_description": "High quality brake pad for Toyota Avanza",
    "cost_price": 150000,
    "selling_price": 200000,
    "stock": 25,
    "sku": "BP-TOY-AVZ-001",
    "barcode": "1234567890123",
    "has_serial_number": false,
    "shelf_location": "A1-B2",
    "usage_status": "Jual",
    "is_active": true,
    "category": {
      "category_id": 1,
      "name": "Spare Parts"
    },
    "supplier": {
      "supplier_id": 1,
      "supplier_name": "PT Auto Parts Indonesia"
    },
    "unit_type": {
      "unit_type_id": 1,
      "name": "Pieces"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### PUT /api/v1/products/:id
Update product information.

**Path Parameters:**
- `id`: Product ID

**Request Body:**
```json
{
  "product_name": "Brake Pad Toyota Avanza - Updated",
  "selling_price": 220000
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Product updated successfully",
  "data": {
    "product_id": 1,
    "product_name": "Brake Pad Toyota Avanza - Updated",
    "selling_price": 220000,
    "updated_at": "2024-01-01T11:00:00Z"
  }
}
```

#### GET /api/v1/products/sku
Get product by SKU.

**Query Parameters:**
- `sku`: Product SKU

**Response:**
```json
{
  "status": "success",
  "message": "Product retrieved successfully",
  "data": {
    "product_id": 1,
    "product_name": "Brake Pad Toyota Avanza",
    "sku": "BP-TOY-AVZ-001",
    "selling_price": 200000,
    "stock": 25
  }
}
```

#### GET /api/v1/products/search
Search products by name or SKU.

**Query Parameters:**
- `q`: Search query
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Products found",
  "data": [
    {
      "product_id": 1,
      "product_name": "Brake Pad Toyota Avanza",
      "sku": "BP-TOY-AVZ-001",
      "selling_price": 200000,
      "stock": 25
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

#### GET /api/v1/products/low-stock
Get products with low stock.

**Query Parameters:**
- `threshold`: Stock threshold (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Low stock products retrieved successfully",
  "data": [
    {
      "product_id": 2,
      "product_name": "Engine Oil Filter",
      "sku": "EOF-001",
      "stock": 5,
      "threshold": 10
    }
  ]
}
```

#### POST /api/v1/products/:id/stock
Update product stock.

**Path Parameters:**
- `id`: Product ID

**Request Body:**
```json
{
  "quantity": -5
}
```

**Validation Rules:**
- `quantity`: required, can be positive (add) or negative (reduce)

**Response:**
```json
{
  "status": "success",
  "message": "Product stock updated successfully",
  "data": {
    "product_id": 1,
    "previous_stock": 25,
    "quantity_changed": -5,
    "current_stock": 20
  }
}
```

---

## Service Management APIs

### Service Categories

#### POST /api/v1/service-categories
Create a new service category.

**Request Body:**
```json
{
  "name": "Engine Services",
  "status": "Aktif"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Service category created successfully",
  "data": {
    "service_category_id": 1,
    "name": "Engine Services",
    "status": "Aktif",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/service-categories
List all service categories with pagination.

**Response:**
```json
{
  "status": "success",
  "message": "Service categories retrieved successfully",
  "data": [
    {
      "service_category_id": 1,
      "name": "Engine Services",
      "status": "Aktif",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

#### GET /api/v1/service-categories/:id
Get service category by ID.

**Path Parameters:**
- `id`: Service Category ID

**Response:**
```json
{
  "status": "success",
  "message": "Service category retrieved successfully",
  "data": {
    "service_category_id": 1,
    "name": "Engine Services",
    "status": "Aktif",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### PUT /api/v1/service-categories/:id
Update service category.

**Path Parameters:**
- `id`: Service Category ID

**Request Body:**
```json
{
  "name": "Engine Services - Updated"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Service category updated successfully",
  "data": {
    "service_category_id": 1,
    "name": "Engine Services - Updated",
    "status": "Aktif",
    "updated_at": "2024-01-01T11:00:00Z"
  }
}
```

#### GET /api/v1/service-categories/:id/services
Get all services in a category.

**Path Parameters:**
- `id`: Service Category ID

**Response:**
```json
{
  "status": "success",
  "message": "Category services retrieved successfully",
  "data": [
    {
      "service_id": 1,
      "service_code": "ENG001",
      "name": "Engine Oil Change",
      "fee": 150000,
      "service_category_id": 1
    }
  ]
}
```

### Services

#### POST /api/v1/services
Create a new service.

**Request Body:**
```json
{
  "service_code": "ENG001",
  "name": "Engine Oil Change",
  "service_category_id": 1,
  "fee": 150000,
  "status": "Aktif"
}
```

**Validation Rules:**
- `service_code`: required, unique
- `name`: required
- `service_category_id`: required, must exist
- `fee`: required, must be positive number
- `status`: optional (default: "Aktif")

**Response:**
```json
{
  "status": "success",
  "message": "Service created successfully",
  "data": {
    "service_id": 1,
    "service_code": "ENG001",
    "name": "Engine Oil Change",
    "service_category_id": 1,
    "fee": 150000,
    "status": "Aktif",
    "service_category": {
      "service_category_id": 1,
      "name": "Engine Services"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/services
List all services with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Services retrieved successfully",
  "data": [
    {
      "service_id": 1,
      "service_code": "ENG001",
      "name": "Engine Oil Change",
      "service_category_id": 1,
      "fee": 150000,
      "status": "Aktif",
      "service_category": {
        "service_category_id": 1,
        "name": "Engine Services"
      },
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

#### GET /api/v1/services/:id
Get service by ID.

**Path Parameters:**
- `id`: Service ID

**Response:**
```json
{
  "status": "success",
  "message": "Service retrieved successfully",
  "data": {
    "service_id": 1,
    "service_code": "ENG001",
    "name": "Engine Oil Change",
    "service_category_id": 1,
    "fee": 150000,
    "status": "Aktif",
    "service_category": {
      "service_category_id": 1,
      "name": "Engine Services"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### PUT /api/v1/services/:id
Update service information.

**Path Parameters:**
- `id`: Service ID

**Request Body:**
```json
{
  "name": "Engine Oil Change - Premium",
  "fee": 180000
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Service updated successfully",
  "data": {
    "service_id": 1,
    "service_code": "ENG001",
    "name": "Engine Oil Change - Premium",
    "fee": 180000,
    "updated_at": "2024-01-01T11:00:00Z"
  }
}
```

#### GET /api/v1/services/code
Get service by code.

**Query Parameters:**
- `service_code`: Service code

**Response:**
```json
{
  "status": "success",
  "message": "Service retrieved successfully",
  "data": {
    "service_id": 1,
    "service_code": "ENG001",
    "name": "Engine Oil Change",
    "fee": 150000
  }
}
```

#### GET /api/v1/services/search
Search services by name or code.

**Query Parameters:**
- `q`: Search query
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Services found",
  "data": [
    {
      "service_id": 1,
      "service_code": "ENG001",
      "name": "Engine Oil Change",
      "fee": 150000,
      "service_category": {
        "service_category_id": 1,
        "name": "Engine Services"
      }
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

---

## Financial Management APIs

### Payment Methods

#### POST /api/v1/payment-methods
Create a new payment method.

**Request Body:**
```json
{
  "name": "Cash",
  "status": "Aktif"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Payment method created successfully",
  "data": {
    "payment_method_id": 1,
    "name": "Cash",
    "status": "Aktif",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/payment-methods
List all payment methods.

**Response:**
```json
{
  "status": "success",
  "message": "Payment methods retrieved successfully",
  "data": [
    {
      "payment_method_id": 1,
      "name": "Cash",
      "status": "Aktif",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    },
    {
      "payment_method_id": 2,
      "name": "Bank Transfer",
      "status": "Aktif",
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ]
}
```

#### GET /api/v1/payment-methods/:id
Get payment method by ID.

**Path Parameters:**
- `id`: Payment Method ID

**Response:**
```json
{
  "status": "success",
  "message": "Payment method retrieved successfully",
  "data": {
    "payment_method_id": 1,
    "name": "Cash",
    "status": "Aktif",
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### PUT /api/v1/payment-methods/:id
Update payment method.

**Path Parameters:**
- `id`: Payment Method ID

**Request Body:**
```json
{
  "name": "Cash Payment - Updated"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Payment method updated successfully",
  "data": {
    "payment_method_id": 1,
    "name": "Cash Payment - Updated",
    "status": "Aktif",
    "updated_at": "2024-01-01T11:00:00Z"
  }
}
```

### Transactions

#### POST /api/v1/transactions
Create a new transaction.

**Request Body:**
```json
{
  "invoice_number": "INV-2024-001",
  "transaction_date": "2024-01-01T10:00:00Z",
  "user_id": 1,
  "customer_id": 1,
  "outlet_id": 1,
  "transaction_type": "Sale",
  "status": "sukses"
}
```

**Validation Rules:**
- `invoice_number`: required, unique
- `transaction_date`: required, ISO 8601 format
- `user_id`: required, must exist
- `customer_id`: optional, must exist if provided
- `outlet_id`: required, must exist
- `transaction_type`: required
- `status`: required

**Response:**
```json
{
  "status": "success",
  "message": "Transaction created successfully",
  "data": {
    "transaction_id": 1,
    "invoice_number": "INV-2024-001",
    "transaction_date": "2024-01-01T10:00:00Z",
    "user_id": 1,
    "customer_id": 1,
    "outlet_id": 1,
    "transaction_type": "Sale",
    "status": "sukses",
    "user": {
      "user_id": 1,
      "name": "John Doe"
    },
    "customer": {
      "customer_id": 1,
      "name": "Jane Smith"
    },
    "outlet": {
      "outlet_id": 1,
      "outlet_name": "Main Workshop"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/transactions
List all transactions with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Transactions retrieved successfully",
  "data": [
    {
      "transaction_id": 1,
      "invoice_number": "INV-2024-001",
      "transaction_date": "2024-01-01T10:00:00Z",
      "user_id": 1,
      "customer_id": 1,
      "outlet_id": 1,
      "transaction_type": "Sale",
      "status": "sukses",
      "user": {
        "user_id": 1,
        "name": "John Doe"
      },
      "customer": {
        "customer_id": 1,
        "name": "Jane Smith"
      },
      "outlet": {
        "outlet_id": 1,
        "outlet_name": "Main Workshop"
      },
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

#### GET /api/v1/transactions/:id
Get transaction by ID.

**Path Parameters:**
- `id`: Transaction ID

**Response:**
```json
{
  "status": "success",
  "message": "Transaction retrieved successfully",
  "data": {
    "transaction_id": 1,
    "invoice_number": "INV-2024-001",
    "transaction_date": "2024-01-01T10:00:00Z",
    "user_id": 1,
    "customer_id": 1,
    "outlet_id": 1,
    "transaction_type": "Sale",
    "status": "sukses",
    "user": {
      "user_id": 1,
      "name": "John Doe"
    },
    "customer": {
      "customer_id": 1,
      "name": "Jane Smith"
    },
    "outlet": {
      "outlet_id": 1,
      "outlet_name": "Main Workshop"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### PUT /api/v1/transactions/:id
Update transaction.

**Path Parameters:**
- `id`: Transaction ID

**Request Body:**
```json
{
  "transaction_type": "Sale - Updated"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Transaction updated successfully",
  "data": {
    "transaction_id": 1,
    "transaction_type": "Sale - Updated",
    "updated_at": "2024-01-01T11:00:00Z"
  }
}
```

#### GET /api/v1/transactions/invoice
Get transaction by invoice number.

**Query Parameters:**
- `invoice_number`: Invoice number

**Response:**
```json
{
  "status": "success",
  "message": "Transaction retrieved successfully",
  "data": {
    "transaction_id": 1,
    "invoice_number": "INV-2024-001",
    "transaction_date": "2024-01-01T10:00:00Z",
    "user_id": 1,
    "customer_id": 1,
    "outlet_id": 1,
    "transaction_type": "Sale",
    "status": "sukses"
  }
}
```

#### GET /api/v1/transactions/status
Get transactions by status.

**Query Parameters:**
- `status`: Transaction status

**Response:**
```json
{
  "status": "success",
  "message": "Transactions retrieved successfully",
  "data": [
    {
      "transaction_id": 1,
      "invoice_number": "INV-2024-001",
      "status": "sukses",
      "transaction_date": "2024-01-01T10:00:00Z"
    }
  ]
}
```

#### GET /api/v1/customers/:id/transactions
Get transactions by customer.

**Path Parameters:**
- `id`: Customer ID

**Response:**
```json
{
  "status": "success",
  "message": "Customer transactions retrieved successfully",
  "data": [
    {
      "transaction_id": 1,
      "invoice_number": "INV-2024-001",
      "transaction_date": "2024-01-01T10:00:00Z",
      "transaction_type": "Sale",
      "status": "sukses"
    }
  ]
}
```

#### GET /api/v1/outlets/:id/transactions
Get transactions by outlet.

**Path Parameters:**
- `id`: Outlet ID

**Response:**
```json
{
  "status": "success",
  "message": "Outlet transactions retrieved successfully",
  "data": [
    {
      "transaction_id": 1,
      "invoice_number": "INV-2024-001",
      "transaction_date": "2024-01-01T10:00:00Z",
      "transaction_type": "Sale",
      "status": "sukses"
    }
  ]
}
```

### Cash Flows

#### POST /api/v1/cash-flows
Create a new cash flow entry.

**Request Body:**
```json
{
  "user_id": 1,
  "outlet_id": 1,
  "flow_type": "Pemasukan",
  "amount": 500000,
  "description": "Sale transaction payment",
  "flow_date": "2024-01-01T10:00:00Z"
}
```

**Validation Rules:**
- `user_id`: required, must exist
- `outlet_id`: required, must exist
- `flow_type`: required, must be "Pemasukan" or "Pengeluaran"
- `amount`: required, must be positive number
- `description`: required
- `flow_date`: required, ISO 8601 format

**Response:**
```json
{
  "status": "success",
  "message": "Cash flow created successfully",
  "data": {
    "cash_flow_id": 1,
    "user_id": 1,
    "outlet_id": 1,
    "flow_type": "Pemasukan",
    "amount": 500000,
    "description": "Sale transaction payment",
    "flow_date": "2024-01-01T10:00:00Z",
    "user": {
      "user_id": 1,
      "name": "John Doe"
    },
    "outlet": {
      "outlet_id": 1,
      "outlet_name": "Main Workshop"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### GET /api/v1/cash-flows
List all cash flows with pagination.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

**Response:**
```json
{
  "status": "success",
  "message": "Cash flows retrieved successfully",
  "data": [
    {
      "cash_flow_id": 1,
      "user_id": 1,
      "outlet_id": 1,
      "flow_type": "Pemasukan",
      "amount": 500000,
      "description": "Sale transaction payment",
      "flow_date": "2024-01-01T10:00:00Z",
      "user": {
        "user_id": 1,
        "name": "John Doe"
      },
      "outlet": {
        "outlet_id": 1,
        "outlet_name": "Main Workshop"
      },
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 1,
    "pages": 1
  }
}
```

#### GET /api/v1/cash-flows/:id
Get cash flow by ID.

**Path Parameters:**
- `id`: Cash Flow ID

**Response:**
```json
{
  "status": "success",
  "message": "Cash flow retrieved successfully",
  "data": {
    "cash_flow_id": 1,
    "user_id": 1,
    "outlet_id": 1,
    "flow_type": "Pemasukan",
    "amount": 500000,
    "description": "Sale transaction payment",
    "flow_date": "2024-01-01T10:00:00Z",
    "user": {
      "user_id": 1,
      "name": "John Doe"
    },
    "outlet": {
      "outlet_id": 1,
      "outlet_name": "Main Workshop"
    },
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

#### PUT /api/v1/cash-flows/:id
Update cash flow.

**Path Parameters:**
- `id`: Cash Flow ID

**Request Body:**
```json
{
  "amount": 600000,
  "description": "Updated sale transaction payment"
}
```

**Response:**
```json
{
  "status": "success",
  "message": "Cash flow updated successfully",
  "data": {
    "cash_flow_id": 1,
    "amount": 600000,
    "description": "Updated sale transaction payment",
    "updated_at": "2024-01-01T11:00:00Z"
  }
}
```

#### GET /api/v1/cash-flows/type
Get cash flows by type.

**Query Parameters:**
- `type`: Flow type ("Pemasukan" or "Pengeluaran")

**Response:**
```json
{
  "status": "success",
  "message": "Cash flows retrieved successfully",
  "data": [
    {
      "cash_flow_id": 1,
      "flow_type": "Pemasukan",
      "amount": 500000,
      "description": "Sale transaction payment",
      "flow_date": "2024-01-01T10:00:00Z"
    }
  ]
}
```

---

## Database Schema

The system implements a complete ERD with 40+ tables covering:

### Foundation & Security
- `users` - User management with outlet assignment
- `outlets` - Branch/shop locations
- `roles` - Role-based access control
- `permissions` - Permission management
- `role_has_permissions` - Role-permission relationships

### Customer & Vehicle Management  
- `customers` - Customer information
- `customer_vehicles` - Vehicle records with detailed specifications

### Master Data & Inventory
- `products` - Product catalog with pricing and stock
- `product_serial_numbers` - Serial number tracking
- `categories` - Product categories
- `suppliers` - Supplier management
- `unit_types` - Units of measurement

### Service Operations
- `services` - Service offerings
- `service_categories` - Service categorization
- `service_jobs` - Core service job management
- `service_details` - Service job line items
- `service_job_histories` - Status change tracking

### Transaction Management
- `transactions` - Transaction records
- `transaction_details` - Transaction line items
- `purchase_orders` - Purchase order management
- `purchase_order_details` - Purchase order line items
- `vehicle_purchases` - Vehicle purchase tracking

### Financial Management
- `payment_methods` - Payment method configuration
- `payments` - Payment records
- `accounts_payables` - Debt management
- `payable_payments` - Debt payment installments
- `accounts_receivables` - Receivable management
- `receivable_payments` - Receivable payment installments
- `cash_flows` - Cash flow tracking

### Reporting & Promotions
- `reports` - Report generation tracking
- `promotions` - Promotional campaigns

---

## Current Implementation Status

### ✅ Fully Implemented & Tested
- **Foundation Module**: Users, Outlets management with complete CRUD operations
- **Customer Module**: Customers, Customer Vehicles with search and relationship management  
- **Inventory Module**: Products, Categories, Suppliers, Unit Types with comprehensive inventory management
- **Service Module**: Services, Service Categories with full service catalog management
- **Financial Module**: Transactions, Payment Methods, Cash Flows with financial tracking
- **Service Job Management**: Complete service job lifecycle with workflow states
- **Queue Management System**: Comprehensive queue system (antrian) for workshop operations

### 🎯 Key Features Successfully Implemented

#### Service Job Management & Queue System
- ✅ **Complete Service Job Workflow**: Antri → Dikerjakan → Selesai → Diambil
- ✅ **Enhanced Status Update API**: Handles technician assignment and validation
- ✅ **Business Logic Validation**: Prevents invalid status transitions
- ✅ **Automatic History Tracking**: All changes logged with timestamps and users
- ✅ **Queue Management APIs**: Complete queue operations for workshop management
- ✅ **Queue Reordering**: Manual priority adjustment for urgent services
- ✅ **Today's Queue Filtering**: Daily operations view
- ✅ **Multi-Outlet Support**: Separate queues per outlet

#### Comprehensive API Coverage
- ✅ **168+ API Endpoints**: Complete coverage for all modules
- ✅ **RESTful Design**: Consistent API patterns across all modules  
- ✅ **Proper Error Handling**: Comprehensive validation and error responses
- ✅ **Pagination Support**: Efficient data handling for large datasets
- ✅ **Search & Filtering**: Advanced search capabilities across modules
- ✅ **Relationship Loading**: Efficient data loading with proper associations

### 📚 Documentation
- ✅ **Comprehensive API Documentation**: Complete endpoint documentation with examples
- ✅ **Service Job & Queue Documentation**: Detailed workflow and usage guide (see [SERVICE_JOB_QUEUE_DOCS.md](SERVICE_JOB_QUEUE_DOCS.md))
- ✅ **Test Scripts**: Multiple test scripts for different modules
- ✅ **Flutter Integration Guide**: Complete guide for mobile app integration

### 🚧 In Progress
- Authentication and authorization middleware
- Role-based access control implementation  
- Advanced reporting APIs
- Real-time notifications

### Features Implemented

#### Security
- Password hashing with bcrypt
- User management with outlet assignment
- Role-based access control foundation

#### Data Management
- Complete CRUD operations for all main entities
- Soft deletes with GORM
- Proper relationships and foreign keys
- Database indexes for performance
- Search and filtering capabilities
- Pagination support

#### API Design
- RESTful API design
- JSON request/response format
- Standardized error handling
- Comprehensive validation
- Detailed response examples

## Flutter Integration Guide

### Installation in Flutter
Add HTTP package to your `pubspec.yaml`:
```yaml
dependencies:
  http: ^1.1.0
  flutter:
    sdk: flutter
```

### Basic API Client Implementation
```dart
import 'dart:convert';
import 'package:http/http.dart' as http;

class PosApiClient {
  static const String baseUrl = 'http://localhost:3000/api/v1';
  
  // Headers for all requests
  Map<String, String> get headers => {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  };

  // Create Customer
  Future<Map<String, dynamic>> createCustomer({
    required String name,
    required String phoneNumber,
    String? address,
  }) async {
    final response = await http.post(
      Uri.parse('$baseUrl/customers'),
      headers: headers,
      body: jsonEncode({
        'name': name,
        'phone_number': phoneNumber,
        'address': address,
        'status': 'Aktif',
      }),
    );
    
    if (response.statusCode == 201) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to create customer: ${response.body}');
    }
  }

  // Get Customers with Pagination
  Future<Map<String, dynamic>> getCustomers({
    int page = 1,
    int limit = 10,
  }) async {
    final response = await http.get(
      Uri.parse('$baseUrl/customers?page=$page&limit=$limit'),
      headers: headers,
    );
    
    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to fetch customers: ${response.body}');
    }
  }

  // Search Products
  Future<Map<String, dynamic>> searchProducts({
    required String query,
    int page = 1,
    int limit = 10,
  }) async {
    final response = await http.get(
      Uri.parse('$baseUrl/products/search?q=$query&page=$page&limit=$limit'),
      headers: headers,
    );
    
    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to search products: ${response.body}');
    }
  }
}
```

### Error Handling in Flutter
```dart
class ApiException implements Exception {
  final String message;
  final int? statusCode;
  
  ApiException(this.message, {this.statusCode});
  
  @override
  String toString() => 'ApiException: $message';
}

// Usage in your Flutter app
try {
  final result = await apiClient.createCustomer(
    name: 'John Doe',
    phoneNumber: '081234567890',
    address: 'Jakarta',
  );
  // Handle success
  print('Customer created: ${result['data']['customer_id']}');
} catch (e) {
  // Handle error
  print('Error: $e');
}
```

### Data Models for Flutter
```dart
class Customer {
  final int customerId;
  final String name;
  final String phoneNumber;
  final String? address;
  final String status;
  final DateTime createdAt;
  final DateTime updatedAt;

  Customer({
    required this.customerId,
    required this.name,
    required this.phoneNumber,
    this.address,
    required this.status,
    required this.createdAt,
    required this.updatedAt,
  });

  factory Customer.fromJson(Map<String, dynamic> json) {
    return Customer(
      customerId: json['customer_id'],
      name: json['name'],
      phoneNumber: json['phone_number'],
      address: json['address'],
      status: json['status'],
      createdAt: DateTime.parse(json['created_at']),
      updatedAt: DateTime.parse(json['updated_at']),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'customer_id': customerId,
      'name': name,
      'phone_number': phoneNumber,
      'address': address,
      'status': status,
      'created_at': createdAt.toIso8601String(),
      'updated_at': updatedAt.toIso8601String(),
    };
  }
}
```

## Getting Started

### Prerequisites
- Go 1.18+
- PostgreSQL
- Git

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd pos-bengkel-golang
```

2. Install dependencies:
```bash
go mod tidy
```

3. Configure database:
Edit `config/config-local.yaml` and update the database connection string:
```yaml
Connection:
    DatabaseApp:
        DriverSource: user=postgres password=postgres sslmode=disable dbname=pos_bengkel host=localhost port=5432
```

4. Create database:
```bash
createdb pos_bengkel
```

5. Run the application:
```bash
go run cmd/api/main.go
```

The application will automatically run database migrations on startup.

### Testing

Test the API endpoints using the provided script:
```bash
./test_api.sh
```

Or manually test individual endpoints:
```bash
# Health check
curl http://localhost:3000/health

# Create outlet
curl -X POST http://localhost:3000/api/v1/outlets \
  -H "Content-Type: application/json" \
  -d '{
    "outlet_name": "Bengkel Utama",
    "branch_type": "Pusat", 
    "city": "Jakarta",
    "address": "Jl. Raya No. 123",
    "phone_number": "021-12345678"
  }'

# Create user
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Admin User",
    "email": "admin@posbengkel.com",
    "password": "password123",
    "outlet_id": 1
  }'
```

## Configuration

The application uses YAML configuration files in the `config/` directory:
- `config-local.yaml` - Local development
- `config-dev.yaml` - Development environment  
- `config-prod.yaml` - Production environment

## Database Migrations

Migrations are automatically run on application startup. The system uses GORM AutoMigrate to create tables based on the model definitions.

## Features Implemented

### Security
- Password hashing with bcrypt
- User management with outlet assignment
- Role-based access control foundation

### Data Management
- Complete CRUD operations for foundation entities
- Soft deletes with GORM
- Proper relationships and foreign keys
- Database indexes for performance

### API Design
- RESTful API design
- JSON request/response format
- Standardized error handling
- Pagination support
- Input validation

## Next Steps

1. Implement remaining modules (Customer, Product, Service, Transaction, Financial)
2. Add authentication and authorization middleware
3. Implement role-based access control
4. Add comprehensive testing
5. Add API documentation with Swagger
6. Implement business-specific features (service scheduling, inventory management, financial reporting)

## Technology Stack

- **Backend**: Go with Fiber framework
- **Database**: PostgreSQL with GORM ORM
- **Architecture**: Clean Architecture
- **Configuration**: Viper with YAML files
- **Logging**: Logrus with structured logging
- **Validation**: Go Playground Validator

## Next Steps

1. ✅ Complete API documentation (DONE)
2. Add authentication and authorization middleware
3. Implement JWT token-based authentication
4. Add role-based access control endpoints
5. Implement service job management workflows
6. Add comprehensive API testing
7. Add Swagger/OpenAPI documentation generation
8. Implement advanced reporting and analytics APIs
9. Add file upload endpoints for documents/images
10. Implement real-time notifications with WebSocket

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support and questions:
- Create an issue in the GitHub repository
- Email: support@posbengkel.com
- Documentation: [API Documentation](README.md)

---

**POS Bengkel** - Complete automotive service shop management system with comprehensive API for Flutter integration.