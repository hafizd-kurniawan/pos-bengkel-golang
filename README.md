# POS Bengkel (Automotive Service Shop POS System)

A comprehensive Point of Sale system for automotive service shops built with Go, Fiber, and PostgreSQL.

## Architecture

This project follows Clean Architecture principles with the following layers:

- **Models**: Database entities and business models
- **Repository**: Data access layer with GORM
- **Usecase**: Business logic layer
- **Delivery**: HTTP handlers and API endpoints

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

## Current Implementation Status

### ✅ Completed
- Database models for all 40+ tables
- Database migrations with proper relationships and indexes
- Repository layer with interfaces and GORM implementations
- Usecase layer with business logic and validation
- Delivery layer with HTTP handlers and JSON responses
- Foundation module APIs (Users, Outlets)

### 🚧 In Progress
- Customer module APIs
- Product/Inventory module APIs
- Service module APIs
- Transaction module APIs
- Financial module APIs
- Reporting module APIs

## API Endpoints

### Foundation APIs

#### Users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users` - List users (with pagination)
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

#### Outlets
- `POST /api/v1/outlets` - Create outlet
- `GET /api/v1/outlets` - List outlets (with pagination)
- `GET /api/v1/outlets/:id` - Get outlet by ID

#### Health Check
- `GET /health` - API health status

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

## License

This project is licensed under the MIT License.