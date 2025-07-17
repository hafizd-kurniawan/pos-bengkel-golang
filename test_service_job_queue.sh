#!/bin/bash

# Test script for Service Job and Queue Management Features
# This script demonstrates the complete service job workflow and queue management

echo "Testing Service Job Management & Queue System"
echo "=============================================="

BASE_URL="http://localhost:3000/api/v1"

# Test 1: Create a new service job
echo "1. Creating a new service job..."
curl -s -X POST "$BASE_URL/service-jobs" \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": 1,
    "vehicle_id": 1,
    "received_by_user_id": 1,
    "outlet_id": 1,
    "problem_description": "Brake system inspection and maintenance",
    "service_in_date": "2025-07-17T14:00:00Z",
    "down_payment": 75000
  }' | jq '{ status: .status, message: .message, service_job_id: .data.service_job_id, queue_number: .data.queue_number }'

echo ""

# Test 2: Check current queue for outlet 1
echo "2. Checking current queue for outlet 1..."
curl -s "$BASE_URL/queue/1" | jq '{ 
  status: .status, 
  message: .message, 
  queue_count: (.data | length),
  jobs: [.data[] | { 
    service_job_id: .service_job_id, 
    queue_number: .queue_number, 
    status: .status,
    problem: .problem_description,
    customer: .customer.name
  }]
}'

echo ""

# Test 3: Check today's queue
echo "3. Checking today's queue for outlet 1..."
curl -s "$BASE_URL/queue/1/today" | jq '{ 
  status: .status, 
  message: .message, 
  today_queue_count: (.data | length),
  jobs: [.data[] | { 
    service_job_id: .service_job_id, 
    queue_number: .queue_number, 
    service_date: .service_in_date,
    customer: .customer.name
  }]
}'

echo ""

# Test 4: Update service job status with technician assignment
echo "4. Assigning technician and starting work on service job..."
curl -s -X PUT "$BASE_URL/service-jobs/3/status" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "Dikerjakan",
    "user_id": 1,
    "technician_id": 1,
    "notes": "Started brake system inspection - assigned to technician"
  }' | jq .

echo ""

# Test 5: Check queue again to see status change
echo "5. Checking queue after status update..."
curl -s "$BASE_URL/queue/1" | jq '{ 
  status: .status, 
  jobs: [.data[] | { 
    service_job_id: .service_job_id, 
    queue_number: .queue_number, 
    status: .status,
    technician_assigned: (.technician_id != null),
    customer: .customer.name
  }]
}'

echo ""

# Test 6: Reorder queue (move job 3 to first position)
echo "6. Reordering queue - moving job 3 to first position..."
curl -s -X PUT "$BASE_URL/queue/1/reorder" \
  -H "Content-Type: application/json" \
  -d '{
    "service_job_ids": [3, 2, 1]
  }' | jq .

echo ""

# Test 7: Verify queue reordering
echo "7. Verifying queue reordering..."
curl -s "$BASE_URL/queue/1" | jq '{ 
  status: .status, 
  message: .message,
  reordered_queue: [.data[] | { 
    service_job_id: .service_job_id, 
    queue_number: .queue_number, 
    status: .status,
    customer: .customer.name
  }]
}'

echo ""

# Test 8: Complete a service job
echo "8. Completing service job..."
curl -s -X PUT "$BASE_URL/service-jobs/3/status" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "Selesai",
    "user_id": 1,
    "notes": "Brake system inspection completed - all components in good condition"
  }' | jq .

echo ""

# Test 9: Mark service job as picked up
echo "9. Marking service job as picked up..."
curl -s -X PUT "$BASE_URL/service-jobs/3/status" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "Diambil",
    "user_id": 1,
    "notes": "Vehicle picked up by customer - service completed"
  }' | jq .

echo ""

# Test 10: Check service job history
echo "10. Checking service job history..."
curl -s "$BASE_URL/service-jobs/3/histories" | jq '{ 
  status: .status, 
  message: .message,
  history_count: (.data | length),
  history: [.data[] | { 
    notes: .notes, 
    changed_at: .changed_at,
    user: .user.name
  }]
}'

echo ""

# Test 11: Final queue status
echo "11. Final queue status (should show remaining jobs)..."
curl -s "$BASE_URL/queue/1" | jq '{ 
  status: .status, 
  remaining_jobs: (.data | length),
  jobs: [.data[] | { 
    service_job_id: .service_job_id, 
    queue_number: .queue_number, 
    status: .status,
    customer: .customer.name
  }]
}'

echo ""
echo "Service Job and Queue Management testing completed!"
echo "=============================================="