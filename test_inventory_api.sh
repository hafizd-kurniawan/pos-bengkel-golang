#!/bin/bash

echo "Testing Inventory Management API Endpoints"
echo "==========================================="

BASE_URL="http://localhost:3000"

# Test 1: Create category
echo "1. Testing create category..."
curl -s -X POST "$BASE_URL/api/v1/categories" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Spare Parts",
    "status": "Aktif"
  }' | jq .

echo ""

# Test 2: Create supplier
echo "2. Testing create supplier..."
curl -s -X POST "$BASE_URL/api/v1/suppliers" \
  -H "Content-Type: application/json" \
  -d '{
    "supplier_name": "PT Auto Parts Indonesia",
    "contact_person_name": "Budi Santoso",
    "phone_number": "021-87654321",
    "address": "Jl. Industri No. 45, Jakarta",
    "status": "Aktif"
  }' | jq .

echo ""

# Test 3: Create unit type
echo "3. Testing create unit type..."
curl -s -X POST "$BASE_URL/api/v1/unit-types" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Pieces",
    "status": "Aktif"
  }' | jq .

echo ""

# Test 4: Create product
echo "4. Testing create product..."
curl -s -X POST "$BASE_URL/api/v1/products" \
  -H "Content-Type: application/json" \
  -d '{
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
  }' | jq .

echo ""

# Test 5: List products
echo "5. Testing list products..."
curl -s "$BASE_URL/api/v1/products" | jq .

echo ""

# Test 6: Get product by SKU
echo "6. Testing get product by SKU..."
curl -s "$BASE_URL/api/v1/products/sku?sku=BP-TOY-AVZ-001" | jq .

echo ""

# Test 7: Search products
echo "7. Testing search products..."
curl -s "$BASE_URL/api/v1/products/search?q=brake" | jq .

echo ""

# Test 8: Update product stock
echo "8. Testing update product stock..."
curl -s -X POST "$BASE_URL/api/v1/products/1/stock" \
  -H "Content-Type: application/json" \
  -d '{
    "quantity": -5
  }' | jq .

echo ""

# Test 9: Get low stock products
echo "9. Testing get low stock products..."
curl -s "$BASE_URL/api/v1/products/low-stock?threshold=30" | jq .

echo ""

# Test 10: List categories
echo "10. Testing list categories..."
curl -s "$BASE_URL/api/v1/categories" | jq .

echo ""

# Test 11: List suppliers
echo "11. Testing list suppliers..."
curl -s "$BASE_URL/api/v1/suppliers" | jq .

echo ""

# Test 12: List unit types
echo "12. Testing list unit types..."
curl -s "$BASE_URL/api/v1/unit-types" | jq .

echo ""

# Test 13: Get products by category
echo "13. Testing get products by category..."
curl -s "$BASE_URL/api/v1/categories/1/products" | jq .

echo ""

# Test 14: Get products by supplier
echo "14. Testing get products by supplier..."
curl -s "$BASE_URL/api/v1/suppliers/1/products" | jq .

echo ""

# Test 15: Update product
echo "15. Testing update product..."
curl -s -X PUT "$BASE_URL/api/v1/products/1" \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Brake Pad Toyota Avanza - Updated",
    "selling_price": 220000
  }' | jq .

echo ""

echo "Inventory API testing completed!"