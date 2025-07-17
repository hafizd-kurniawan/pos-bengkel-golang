#!/bin/bash

echo "Testing Customer Management API Endpoints"
echo "========================================="

BASE_URL="http://localhost:3000"

# Test 1: Create customer
echo "1. Testing create customer..."
curl -s -X POST "$BASE_URL/api/v1/customers" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "phone_number": "081234567890",
    "address": "Jl. Merdeka No. 123",
    "status": "Aktif"
  }' | jq .

echo ""

# Test 2: Create another customer
echo "2. Testing create another customer..."
curl -s -X POST "$BASE_URL/api/v1/customers" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Smith",
    "phone_number": "081234567891",
    "address": "Jl. Sudirman No. 456",
    "status": "Aktif"
  }' | jq .

echo ""

# Test 3: List customers
echo "3. Testing list customers..."
curl -s "$BASE_URL/api/v1/customers" | jq .

echo ""

# Test 4: Get customer by ID
echo "4. Testing get customer by ID..."
curl -s "$BASE_URL/api/v1/customers/1" | jq .

echo ""

# Test 5: Create customer vehicle
echo "5. Testing create customer vehicle..."
curl -s -X POST "$BASE_URL/api/v1/customer-vehicles" \
  -H "Content-Type: application/json" \
  -d '{
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
  }' | jq .

echo ""

# Test 6: List customer vehicles
echo "6. Testing list customer vehicles..."
curl -s "$BASE_URL/api/v1/customer-vehicles" | jq .

echo ""

# Test 7: Get customer vehicles by customer ID
echo "7. Testing get customer vehicles by customer ID..."
curl -s "$BASE_URL/api/v1/customers/1/vehicles" | jq .

echo ""

# Test 8: Search customers
echo "8. Testing search customers..."
curl -s "$BASE_URL/api/v1/customers/search?q=John" | jq .

echo ""

# Test 9: Update customer
echo "9. Testing update customer..."
curl -s -X PUT "$BASE_URL/api/v1/customers/1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe Updated",
    "address": "Jl. Merdeka No. 123 Updated"
  }' | jq .

echo ""

echo "Customer API testing completed!"