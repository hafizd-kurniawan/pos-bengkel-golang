-- Create vehicle trading enums first
CREATE TYPE vehicle_ownership_status AS ENUM ('Customer', 'Showroom', 'Workshop');
CREATE TYPE vehicle_condition_status AS ENUM ('Excellent', 'Good', 'Fair', 'Poor');
CREATE TYPE vehicle_sale_status AS ENUM ('Not For Sale', 'For Sale', 'Sold', 'Reserved');
CREATE TYPE reconditioning_job_status AS ENUM ('Pending', 'In Progress', 'Completed', 'Cancelled');
CREATE TYPE reconditioning_detail_type AS ENUM ('Part', 'Service');
CREATE TYPE sales_transaction_type AS ENUM ('Cash', 'Installment');
CREATE TYPE installment_status AS ENUM ('Active', 'Completed', 'Defaulted');
CREATE TYPE installment_payment_status AS ENUM ('Pending', 'Paid', 'Late', 'Skipped');

-- Enhanced vehicles table
CREATE TABLE vehicles (
    vehicle_id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customers(customer_id),
    plate_number VARCHAR(20) UNIQUE NOT NULL,
    brand VARCHAR(100) NOT NULL,
    model VARCHAR(100) NOT NULL,
    type VARCHAR(100) NOT NULL,
    production_year INTEGER NOT NULL,
    chassis_number VARCHAR(100) UNIQUE NOT NULL,
    engine_number VARCHAR(100) UNIQUE NOT NULL,
    color VARCHAR(50) NOT NULL,
    mileage INTEGER,
    fuel_type VARCHAR(20),
    transmission VARCHAR(20),
    ownership_status vehicle_ownership_status NOT NULL DEFAULT 'Customer',
    condition_status vehicle_condition_status NOT NULL DEFAULT 'Good',
    sale_status vehicle_sale_status NOT NULL DEFAULT 'Not For Sale',
    purchase_price DECIMAL(15,2),
    selling_price DECIMAL(15,2),
    estimated_value DECIMAL(15,2),
    condition_notes TEXT,
    internal_notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Create indexes for vehicles table
CREATE INDEX idx_vehicles_customer_id ON vehicles(customer_id);
CREATE INDEX idx_vehicles_ownership_status ON vehicles(ownership_status);
CREATE INDEX idx_vehicles_sale_status ON vehicles(sale_status);
CREATE INDEX idx_vehicles_deleted_at ON vehicles(deleted_at);

-- Vehicle purchase transactions table
CREATE TABLE vehicle_purchase_transactions (
    purchase_transaction_id SERIAL PRIMARY KEY,
    vehicle_id INTEGER NOT NULL REFERENCES vehicles(vehicle_id),
    customer_id INTEGER NOT NULL REFERENCES customers(customer_id),
    purchase_price DECIMAL(15,2) NOT NULL,
    purchase_date TIMESTAMP WITH TIME ZONE NOT NULL,
    payment_method payment_type_enum NOT NULL,
    transaction_status transaction_status NOT NULL DEFAULT 'pending',
    evaluation_notes TEXT,
    payment_reference VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Create indexes for purchase transactions table
CREATE INDEX idx_purchase_transactions_vehicle_id ON vehicle_purchase_transactions(vehicle_id);
CREATE INDEX idx_purchase_transactions_customer_id ON vehicle_purchase_transactions(customer_id);
CREATE INDEX idx_purchase_transactions_date ON vehicle_purchase_transactions(purchase_date);
CREATE INDEX idx_purchase_transactions_status ON vehicle_purchase_transactions(transaction_status);
CREATE INDEX idx_purchase_transactions_deleted_at ON vehicle_purchase_transactions(deleted_at);

-- Vehicle reconditioning jobs table
CREATE TABLE vehicle_reconditioning_jobs (
    reconditioning_job_id SERIAL PRIMARY KEY,
    vehicle_id INTEGER NOT NULL REFERENCES vehicles(vehicle_id),
    job_title VARCHAR(255) NOT NULL,
    job_description TEXT,
    estimated_cost DECIMAL(15,2),
    actual_cost DECIMAL(15,2),
    start_date TIMESTAMP WITH TIME ZONE,
    completion_date TIMESTAMP WITH TIME ZONE,
    status reconditioning_job_status NOT NULL DEFAULT 'Pending',
    assigned_technician_id INTEGER REFERENCES users(user_id),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Create indexes for reconditioning jobs table
CREATE INDEX idx_reconditioning_jobs_vehicle_id ON vehicle_reconditioning_jobs(vehicle_id);
CREATE INDEX idx_reconditioning_jobs_technician_id ON vehicle_reconditioning_jobs(assigned_technician_id);
CREATE INDEX idx_reconditioning_jobs_status ON vehicle_reconditioning_jobs(status);
CREATE INDEX idx_reconditioning_jobs_deleted_at ON vehicle_reconditioning_jobs(deleted_at);

-- Reconditioning details table
CREATE TABLE reconditioning_details (
    detail_id SERIAL PRIMARY KEY,
    reconditioning_job_id INTEGER NOT NULL REFERENCES vehicle_reconditioning_jobs(reconditioning_job_id),
    detail_type reconditioning_detail_type NOT NULL,
    product_id INTEGER REFERENCES products(product_id),
    service_id INTEGER REFERENCES services(service_id),
    description VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    unit_price DECIMAL(15,2) NOT NULL,
    total_price DECIMAL(15,2) NOT NULL,
    usage_date TIMESTAMP WITH TIME ZONE NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Create indexes for reconditioning details table
CREATE INDEX idx_reconditioning_details_job_id ON reconditioning_details(reconditioning_job_id);
CREATE INDEX idx_reconditioning_details_product_id ON reconditioning_details(product_id);
CREATE INDEX idx_reconditioning_details_service_id ON reconditioning_details(service_id);
CREATE INDEX idx_reconditioning_details_type ON reconditioning_details(detail_type);
CREATE INDEX idx_reconditioning_details_deleted_at ON reconditioning_details(deleted_at);

-- Vehicle sales transactions table
CREATE TABLE vehicle_sales_transactions (
    sales_transaction_id SERIAL PRIMARY KEY,
    vehicle_id INTEGER NOT NULL REFERENCES vehicles(vehicle_id),
    customer_id INTEGER NOT NULL REFERENCES customers(customer_id),
    sale_price DECIMAL(15,2) NOT NULL,
    down_payment DECIMAL(15,2),
    sale_date TIMESTAMP WITH TIME ZONE NOT NULL,
    transaction_type sales_transaction_type NOT NULL,
    payment_method payment_type_enum NOT NULL,
    transaction_status transaction_status NOT NULL DEFAULT 'pending',
    payment_reference VARCHAR(255),
    profit_amount DECIMAL(15,2),
    sales_person_id INTEGER REFERENCES users(user_id),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Create indexes for sales transactions table
CREATE INDEX idx_sales_transactions_vehicle_id ON vehicle_sales_transactions(vehicle_id);
CREATE INDEX idx_sales_transactions_customer_id ON vehicle_sales_transactions(customer_id);
CREATE INDEX idx_sales_transactions_sales_person_id ON vehicle_sales_transactions(sales_person_id);
CREATE INDEX idx_sales_transactions_date ON vehicle_sales_transactions(sale_date);
CREATE INDEX idx_sales_transactions_type ON vehicle_sales_transactions(transaction_type);
CREATE INDEX idx_sales_transactions_status ON vehicle_sales_transactions(transaction_status);
CREATE INDEX idx_sales_transactions_deleted_at ON vehicle_sales_transactions(deleted_at);

-- Vehicle installments table
CREATE TABLE vehicle_installments (
    installment_id SERIAL PRIMARY KEY,
    sales_transaction_id INTEGER NOT NULL REFERENCES vehicle_sales_transactions(sales_transaction_id),
    total_amount DECIMAL(15,2) NOT NULL,
    down_payment DECIMAL(15,2) NOT NULL,
    installment_amount DECIMAL(15,2) NOT NULL,
    number_of_installments INTEGER NOT NULL,
    interest_rate DECIMAL(5,2),
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    status installment_status NOT NULL DEFAULT 'Active',
    remaining_balance DECIMAL(15,2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Create indexes for installments table
CREATE INDEX idx_installments_sales_transaction_id ON vehicle_installments(sales_transaction_id);
CREATE INDEX idx_installments_status ON vehicle_installments(status);
CREATE INDEX idx_installments_start_date ON vehicle_installments(start_date);
CREATE INDEX idx_installments_end_date ON vehicle_installments(end_date);
CREATE INDEX idx_installments_deleted_at ON vehicle_installments(deleted_at);

-- Installment payments table
CREATE TABLE installment_payments (
    payment_id SERIAL PRIMARY KEY,
    installment_id INTEGER NOT NULL REFERENCES vehicle_installments(installment_id),
    payment_number INTEGER NOT NULL,
    due_date TIMESTAMP WITH TIME ZONE NOT NULL,
    payment_date TIMESTAMP WITH TIME ZONE,
    due_amount DECIMAL(15,2) NOT NULL,
    paid_amount DECIMAL(15,2),
    late_fee DECIMAL(15,2),
    payment_status installment_payment_status NOT NULL DEFAULT 'Pending',
    payment_method payment_type_enum,
    payment_reference VARCHAR(255),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Create indexes for installment payments table
CREATE INDEX idx_installment_payments_installment_id ON installment_payments(installment_id);
CREATE INDEX idx_installment_payments_payment_number ON installment_payments(installment_id, payment_number);
CREATE INDEX idx_installment_payments_due_date ON installment_payments(due_date);
CREATE INDEX idx_installment_payments_status ON installment_payments(payment_status);
CREATE INDEX idx_installment_payments_deleted_at ON installment_payments(deleted_at);

-- Add constraints for data integrity
ALTER TABLE reconditioning_details ADD CONSTRAINT chk_reconditioning_details_references
    CHECK (
        (detail_type = 'Part' AND product_id IS NOT NULL AND service_id IS NULL) OR
        (detail_type = 'Service' AND service_id IS NOT NULL AND product_id IS NULL)
    );

ALTER TABLE vehicle_sales_transactions ADD CONSTRAINT chk_sales_down_payment
    CHECK (down_payment IS NULL OR down_payment < sale_price);

ALTER TABLE installment_payments ADD CONSTRAINT chk_installment_payment_number
    CHECK (payment_number > 0);

-- Create unique constraint for payment numbers within installments
ALTER TABLE installment_payments ADD CONSTRAINT uk_installment_payment_number
    UNIQUE (installment_id, payment_number);