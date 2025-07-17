#!/bin/bash

echo "Testing Service Management API Endpoints"
echo "========================================"

BASE_URL="http://localhost:3000"

# Test 1: Create service category
echo "1. Testing create service category..."
curl -s -X POST "$BASE_URL/api/v1/service-categories" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Engine Services",
    "status": "Aktif"
  }' | jq .

echo ""

# Test 2: Create another service category
echo "2. Testing create another service category..."
curl -s -X POST "$BASE_URL/api/v1/service-categories" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Electrical Services",
    "status": "Aktif"
  }' | jq .

echo ""

# Test 3: List service categories
echo "3. Testing list service categories..."
curl -s "$BASE_URL/api/v1/service-categories" | jq .

echo ""

# Test 4: Get service category by ID
echo "4. Testing get service category by ID..."
curl -s "$BASE_URL/api/v1/service-categories/1" | jq .

echo ""

# Test 5: Create service
echo "5. Testing create service..."
curl -s -X POST "$BASE_URL/api/v1/services" \
  -H "Content-Type: application/json" \
  -d '{
    "service_code": "ENG001",
    "name": "Engine Oil Change",
    "service_category_id": 1,
    "fee": 150000,
    "status": "Aktif"
  }' | jq .

echo ""

# Test 6: Create another service
echo "6. Testing create another service..."
curl -s -X POST "$BASE_URL/api/v1/services" \
  -H "Content-Type: application/json" \
  -d '{
    "service_code": "ENG002",
    "name": "Engine Tune Up",
    "service_category_id": 1,
    "fee": 350000,
    "status": "Aktif"
  }' | jq .

echo ""

# Test 7: List services
echo "7. Testing list services..."
curl -s "$BASE_URL/api/v1/services" | jq .

echo ""

# Test 8: Get service by ID
echo "8. Testing get service by ID..."
curl -s "$BASE_URL/api/v1/services/1" | jq .

echo ""

# Test 9: Get service by code
echo "9. Testing get service by code..."
curl -s "$BASE_URL/api/v1/services/code?service_code=ENG001" | jq .

echo ""

# Test 10: Search services
echo "10. Testing search services..."
curl -s "$BASE_URL/api/v1/services/search?q=engine" | jq .

echo ""

# Test 11: Get services by category
echo "11. Testing get services by category..."
curl -s "$BASE_URL/api/v1/service-categories/1/services" | jq .

echo ""

# Test 12: Update service
echo "12. Testing update service..."
curl -s -X PUT "$BASE_URL/api/v1/services/1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Engine Oil Change - Premium",
    "fee": 180000
  }' | jq .

echo ""

# Test 13: Update service category
echo "13. Testing update service category..."
curl -s -X PUT "$BASE_URL/api/v1/service-categories/1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Engine Services - Updated"
  }' | jq .

echo ""

echo "Service API testing completed!"