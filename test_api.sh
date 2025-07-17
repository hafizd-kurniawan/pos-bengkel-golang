#!/bin/bash

# Test script for POS Bengkel API endpoints
# This script tests the basic functionality of our API

echo "Testing POS Bengkel API Endpoints"
echo "=================================="

BASE_URL="http://localhost:3000"

# Test 1: Health check
echo "1. Testing health check endpoint..."
curl -s "$BASE_URL/health" | jq .

echo ""

# Test 2: Root endpoint
echo "2. Testing root endpoint..."
curl -s "$BASE_URL/" | jq .

echo ""

# Test 3: Create outlet
echo "3. Testing create outlet..."
curl -s -X POST "$BASE_URL/api/v1/outlets" \
  -H "Content-Type: application/json" \
  -d '{
    "outlet_name": "Bengkel Utama",
    "branch_type": "Pusat",
    "city": "Jakarta",
    "address": "Jl. Raya No. 123",
    "phone_number": "021-12345678",
    "status": "Aktif"
  }' | jq .

echo ""

# Test 4: List outlets
echo "4. Testing list outlets..."
curl -s "$BASE_URL/api/v1/outlets" | jq .

echo ""

# Test 5: Create user
echo "5. Testing create user..."
curl -s -X POST "$BASE_URL/api/v1/users" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Admin User",
    "email": "admin@posbengkel.com",
    "password": "password123",
    "outlet_id": 1
  }' | jq .

echo ""

# Test 6: List users
echo "6. Testing list users..."
curl -s "$BASE_URL/api/v1/users" | jq .

echo ""

echo "API testing completed!"