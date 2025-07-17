#!/bin/bash

echo "Testing Financial Management API Endpoints"
echo "=========================================="

BASE_URL="http://localhost:3000"

# Test 1: Create payment method
echo "1. Testing create payment method..."
curl -s -X POST "$BASE_URL/api/v1/payment-methods" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Cash",
    "status": "Aktif"
  }' | jq .

echo ""

# Test 2: Create another payment method
echo "2. Testing create another payment method..."
curl -s -X POST "$BASE_URL/api/v1/payment-methods" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Bank Transfer",
    "status": "Aktif"
  }' | jq .

echo ""

# Test 3: List payment methods
echo "3. Testing list payment methods..."
curl -s "$BASE_URL/api/v1/payment-methods" | jq .

echo ""

# Test 4: Get payment method by ID
echo "4. Testing get payment method by ID..."
curl -s "$BASE_URL/api/v1/payment-methods/1" | jq .

echo ""

# Test 5: Create transaction
echo "5. Testing create transaction..."
curl -s -X POST "$BASE_URL/api/v1/transactions" \
  -H "Content-Type: application/json" \
  -d '{
    "invoice_number": "INV-2024-001",
    "transaction_date": "2024-07-17T10:00:00Z",
    "user_id": 1,
    "customer_id": 1,
    "outlet_id": 1,
    "transaction_type": "Sale",
    "status": "sukses"
  }' | jq .

echo ""

# Test 6: Create another transaction
echo "6. Testing create another transaction..."
curl -s -X POST "$BASE_URL/api/v1/transactions" \
  -H "Content-Type: application/json" \
  -d '{
    "invoice_number": "INV-2024-002",
    "transaction_date": "2024-07-17T11:00:00Z",
    "user_id": 1,
    "customer_id": 2,
    "outlet_id": 1,
    "transaction_type": "Sale",
    "status": "sukses"
  }' | jq .

echo ""

# Test 7: List transactions
echo "7. Testing list transactions..."
curl -s "$BASE_URL/api/v1/transactions" | jq .

echo ""

# Test 8: Get transaction by ID
echo "8. Testing get transaction by ID..."
curl -s "$BASE_URL/api/v1/transactions/1" | jq .

echo ""

# Test 9: Get transaction by invoice number
echo "9. Testing get transaction by invoice number..."
curl -s "$BASE_URL/api/v1/transactions/invoice?invoice_number=INV-2024-001" | jq .

echo ""

# Test 10: Create cash flow
echo "10. Testing create cash flow..."
curl -s -X POST "$BASE_URL/api/v1/cash-flows" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "outlet_id": 1,
    "flow_type": "Pemasukan",
    "amount": 500000,
    "description": "Sale transaction payment",
    "flow_date": "2024-07-17T10:00:00Z"
  }' | jq .

echo ""

# Test 11: Create another cash flow
echo "11. Testing create another cash flow..."
curl -s -X POST "$BASE_URL/api/v1/cash-flows" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "outlet_id": 1,
    "flow_type": "Pengeluaran",
    "amount": 100000,
    "description": "Office supplies purchase",
    "flow_date": "2024-07-17T09:00:00Z"
  }' | jq .

echo ""

# Test 12: List cash flows
echo "12. Testing list cash flows..."
curl -s "$BASE_URL/api/v1/cash-flows" | jq .

echo ""

# Test 13: Get cash flow by ID
echo "13. Testing get cash flow by ID..."
curl -s "$BASE_URL/api/v1/cash-flows/1" | jq .

echo ""

# Test 14: Get cash flows by type
echo "14. Testing get cash flows by type..."
curl -s "$BASE_URL/api/v1/cash-flows/type?type=Pemasukan" | jq .

echo ""

# Test 15: Get transactions by status
echo "15. Testing get transactions by status..."
curl -s "$BASE_URL/api/v1/transactions/status?status=sukses" | jq .

echo ""

# Test 16: Get transactions by customer
echo "16. Testing get transactions by customer..."
curl -s "$BASE_URL/api/v1/customers/1/transactions" | jq .

echo ""

# Test 17: Get transactions by outlet
echo "17. Testing get transactions by outlet..."
curl -s "$BASE_URL/api/v1/outlets/1/transactions" | jq .

echo ""

# Test 18: Update payment method
echo "18. Testing update payment method..."
curl -s -X PUT "$BASE_URL/api/v1/payment-methods/1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Cash Payment - Updated"
  }' | jq .

echo ""

# Test 19: Update transaction
echo "19. Testing update transaction..."
curl -s -X PUT "$BASE_URL/api/v1/transactions/1" \
  -H "Content-Type: application/json" \
  -d '{
    "transaction_type": "Sale - Updated"
  }' | jq .

echo ""

# Test 20: Update cash flow
echo "20. Testing update cash flow..."
curl -s -X PUT "$BASE_URL/api/v1/cash-flows/1" \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 600000,
    "description": "Updated sale transaction payment"
  }' | jq .

echo ""

echo "Financial API testing completed!"