# Service Job Management & Queue System Documentation

## Overview

The POS Bengkel system includes a comprehensive service job management and queue system (sistem antrian) designed for automotive service workshops. This system handles the complete lifecycle of service jobs from initial customer request to completion and pickup.

## Service Job Workflow

### Service Job States

1. **Antri** - Service job is in queue waiting to be processed
2. **Dikerjakan** - Service job is currently being worked on by a technician  
3. **Selesai** - Service job has been completed
4. **Diambil** - Vehicle has been picked up by customer

### Workflow Transitions

```
Customer Request → Antri → Dikerjakan → Selesai → Diambil
```

## Queue Management System (Sistem Antrian)

### Features

- **Automatic Queue Positioning**: New service jobs are automatically assigned queue numbers based on service date and outlet
- **Queue Reordering**: Operators can manually reorder service jobs in the queue for priority handling
- **Today's Queue**: Filter queue to show only today's service jobs for daily operations
- **Multi-Outlet Support**: Each outlet maintains its own separate queue

### Queue API Endpoints

#### Get Complete Queue
```bash
GET /api/v1/queue/{outlet_id}
```
Returns all service jobs in queue for the specified outlet.

#### Get Today's Queue
```bash
GET /api/v1/queue/{outlet_id}/today
```
Returns only today's service jobs in queue for the specified outlet.

#### Reorder Queue
```bash
PUT /api/v1/queue/{outlet_id}/reorder
```
**Request Body:**
```json
{
  "service_job_ids": [3, 1, 2]
}
```
Reorders service jobs according to the provided array order.

## Service Job Management APIs

### Create Service Job
```bash
POST /api/v1/service-jobs
```
**Request Body:**
```json
{
  "customer_id": 1,
  "vehicle_id": 1,
  "received_by_user_id": 1,
  "outlet_id": 1,
  "problem_description": "Engine making strange noise during acceleration",
  "service_in_date": "2025-07-17T08:00:00Z",
  "down_payment": 100000
}
```

### Update Service Job Status
```bash
PUT /api/v1/service-jobs/{id}/status
```

#### Basic Status Update
```json
{
  "status": "Selesai",
  "user_id": 1,
  "notes": "Service completed successfully"
}
```

#### Status Update with Technician Assignment
```json
{
  "status": "Dikerjakan",
  "user_id": 1,
  "technician_id": 1,
  "notes": "Started brake system inspection"
}
```

### Business Logic Validation

- **Technician Assignment**: A technician must be assigned before a service job can be started (status "Dikerjakan")
- **Status Transitions**: Service jobs follow a logical progression through states
- **History Tracking**: All status changes are automatically tracked with timestamps and user information
- **Automatic Calculations**: Totals and commissions are calculated when service is completed

## Service Job History

Every status change and important event is tracked in the service job history:

```bash
GET /api/v1/service-jobs/{id}/histories
```

**Response:**
```json
{
  "status": "success",
  "message": "Service job histories retrieved successfully",
  "data": [
    {
      "history_id": 1,
      "service_job_id": 1,
      "user_id": 1,
      "notes": "Service job created",
      "changed_at": "2025-07-17T08:00:00Z",
      "user": {
        "user_id": 1,
        "name": "Admin User"
      }
    }
  ]
}
```

## Integration with Other Modules

### Customer & Vehicle Management
- Service jobs are linked to specific customers and their vehicles
- Customer information is automatically included in service job responses
- Vehicle details are tracked for service history

### User Management
- Service jobs track who received the job initially
- Technician assignment for actual work
- All status changes are logged with user information

### Outlet Management
- Each outlet maintains its own service job queue
- Queue numbers are outlet-specific
- Service jobs can only be managed within their assigned outlet

## Usage Examples

### Complete Service Job Workflow

1. **Create Service Job**
```bash
curl -X POST http://localhost:3000/api/v1/service-jobs \
  -H "Content-Type: application/json" \
  -d '{
    "customer_id": 1,
    "vehicle_id": 1,
    "received_by_user_id": 1,
    "outlet_id": 1,
    "problem_description": "Regular maintenance service",
    "service_in_date": "2025-07-17T10:00:00Z",
    "down_payment": 50000
  }'
```

2. **Check Queue Position**
```bash
curl http://localhost:3000/api/v1/queue/1/today
```

3. **Assign Technician and Start Work**
```bash
curl -X PUT http://localhost:3000/api/v1/service-jobs/1/status \
  -H "Content-Type: application/json" \
  -d '{
    "status": "Dikerjakan",
    "user_id": 1,
    "technician_id": 1,
    "notes": "Started maintenance work"
  }'
```

4. **Complete Service**
```bash
curl -X PUT http://localhost:3000/api/v1/service-jobs/1/status \
  -H "Content-Type: application/json" \
  -d '{
    "status": "Selesai",
    "user_id": 1,
    "notes": "Maintenance completed successfully"
  }'
```

5. **Mark as Picked Up**
```bash
curl -X PUT http://localhost:3000/api/v1/service-jobs/1/status \
  -H "Content-Type: application/json" \
  -d '{
    "status": "Diambil",
    "user_id": 1,
    "notes": "Vehicle picked up by customer"
  }'
```

### Queue Management Examples

**Reorder Queue for Priority Service:**
```bash
curl -X PUT http://localhost:3000/api/v1/queue/1/reorder \
  -H "Content-Type: application/json" \
  -d '{
    "service_job_ids": [3, 1, 2]
  }'
```

**Check Today's Workload:**
```bash
curl http://localhost:3000/api/v1/queue/1/today
```

## Error Handling

The system includes comprehensive error handling and validation:

- **Invalid Status Transitions**: Prevents illogical status changes
- **Missing Technician**: Prevents starting work without technician assignment
- **Not Found**: Proper handling when service jobs, users, or outlets don't exist
- **Validation Errors**: Clear error messages for invalid request data

## Best Practices

1. **Daily Queue Management**: Use the today's queue endpoint for daily operations
2. **Priority Handling**: Use queue reordering for urgent or priority services
3. **Technician Assignment**: Always assign technicians before starting work
4. **Status Updates**: Include meaningful notes for all status changes
5. **History Tracking**: Regularly check service job history for audit trails

## Testing

Use the provided test script to validate all functionality:

```bash
./test_service_job_queue.sh
```

This script demonstrates the complete service job lifecycle and queue management features.