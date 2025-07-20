-- Drop tables in reverse order (to handle foreign key constraints)
DROP TABLE IF EXISTS installment_payments;
DROP TABLE IF EXISTS vehicle_installments;
DROP TABLE IF EXISTS vehicle_sales_transactions;
DROP TABLE IF EXISTS reconditioning_details;
DROP TABLE IF EXISTS vehicle_reconditioning_jobs;
DROP TABLE IF EXISTS vehicle_purchase_transactions;
DROP TABLE IF EXISTS vehicles;

-- Drop enums
DROP TYPE IF EXISTS installment_payment_status;
DROP TYPE IF EXISTS installment_status;
DROP TYPE IF EXISTS sales_transaction_type;
DROP TYPE IF EXISTS reconditioning_detail_type;
DROP TYPE IF EXISTS reconditioning_job_status;
DROP TYPE IF EXISTS vehicle_sale_status;
DROP TYPE IF EXISTS vehicle_condition_status;
DROP TYPE IF EXISTS vehicle_ownership_status;