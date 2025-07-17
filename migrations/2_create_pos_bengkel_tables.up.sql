-- Create enums first
CREATE TYPE status_umum AS ENUM ('Aktif', 'Tidak Aktif');
CREATE TYPE product_usage_status AS ENUM ('Jual', 'Pakai Sendiri', 'Rusak');
CREATE TYPE sn_status AS ENUM ('Tersedia', 'Terpakai', 'Rusak');
CREATE TYPE service_status_enum AS ENUM ('Antri', 'Dikerjakan', 'Selesai', 'Diambil', 'Komplain');
CREATE TYPE transaction_status AS ENUM ('pending', 'sukses', 'gagal');
CREATE TYPE purchase_status AS ENUM ('Selesai', 'Pending');
CREATE TYPE payment_type_enum AS ENUM ('tunai', 'transfer', 'cicilan');
CREATE TYPE ap_ar_status AS ENUM ('Belum Lunas', 'Lunas');
CREATE TYPE cash_flow_type AS ENUM ('Pemasukan', 'Pengeluaran');
CREATE TYPE report_type_enum AS ENUM ('Penjualan', 'Keuangan', 'Inventory');
CREATE TYPE report_status AS ENUM ('Pending', 'Selesai', 'Gagal');

-- Foundation & Security Tables
CREATE TABLE outlets (
    outlet_id SERIAL PRIMARY KEY,
    outlet_name VARCHAR(255) NOT NULL,
    branch_type VARCHAR(50) NOT NULL,
    city VARCHAR(100) NOT NULL,
    address TEXT,
    phone_number VARCHAR(20),
    status status_umum NOT NULL DEFAULT 'Aktif',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    outlet_id INTEGER REFERENCES outlets(outlet_id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE role_has_permissions (
    permission_id INTEGER REFERENCES permissions(id) ON DELETE CASCADE,
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (permission_id, role_id)
);

-- Customer & Vehicle Tables
CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    address TEXT,
    status status_umum NOT NULL DEFAULT 'Aktif',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE customer_vehicles (
    vehicle_id SERIAL PRIMARY KEY,
    customer_id INTEGER NOT NULL REFERENCES customers(customer_id),
    plate_number VARCHAR(20) UNIQUE NOT NULL,
    brand VARCHAR(100) NOT NULL,
    model VARCHAR(100) NOT NULL,
    type VARCHAR(100) NOT NULL,
    production_year INTEGER NOT NULL,
    chassis_number VARCHAR(100) UNIQUE NOT NULL,
    engine_number VARCHAR(100) UNIQUE NOT NULL,
    color VARCHAR(50) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Master Data & Inventory Tables
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status status_umum NOT NULL DEFAULT 'Aktif',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE suppliers (
    supplier_id SERIAL PRIMARY KEY,
    supplier_name VARCHAR(255) NOT NULL,
    contact_person_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    address TEXT,
    status status_umum NOT NULL DEFAULT 'Aktif',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE unit_types (
    unit_type_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    status status_umum NOT NULL DEFAULT 'Aktif',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT,
    product_image VARCHAR(255),
    cost_price DECIMAL(15, 2) NOT NULL,
    selling_price DECIMAL(15, 2) NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    sku VARCHAR(100) UNIQUE,
    barcode VARCHAR(100) UNIQUE,
    has_serial_number BOOLEAN NOT NULL DEFAULT FALSE,
    shelf_location VARCHAR(100),
    usage_status product_usage_status NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    category_id INTEGER REFERENCES categories(category_id),
    supplier_id INTEGER REFERENCES suppliers(supplier_id),
    unit_type_id INTEGER REFERENCES unit_types(unit_type_id),
    sourceable_id INTEGER,
    sourceable_type VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE product_serial_numbers (
    serial_number_id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(product_id),
    serial_number VARCHAR(255) UNIQUE NOT NULL,
    status sn_status NOT NULL DEFAULT 'Tersedia',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Service Tables
CREATE TABLE service_categories (
    service_category_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status status_umum NOT NULL DEFAULT 'Aktif',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE services (
    service_id SERIAL PRIMARY KEY,
    service_code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    service_category_id INTEGER NOT NULL REFERENCES service_categories(service_category_id),
    fee DECIMAL(15, 2) NOT NULL,
    status status_umum NOT NULL DEFAULT 'Aktif',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Core Operations Tables
CREATE TABLE service_jobs (
    service_job_id SERIAL PRIMARY KEY,
    service_code VARCHAR(50) UNIQUE NOT NULL,
    queue_number INTEGER NOT NULL,
    customer_id INTEGER NOT NULL REFERENCES customers(customer_id),
    vehicle_id INTEGER NOT NULL REFERENCES customer_vehicles(vehicle_id),
    technician_id INTEGER REFERENCES users(user_id),
    received_by_user_id INTEGER NOT NULL REFERENCES users(user_id),
    outlet_id INTEGER NOT NULL REFERENCES outlets(outlet_id),
    problem_description TEXT NOT NULL,
    technician_notes TEXT,
    status service_status_enum NOT NULL,
    service_in_date TIMESTAMP WITH TIME ZONE NOT NULL,
    picked_up_date TIMESTAMP WITH TIME ZONE,
    complain_date TIMESTAMP WITH TIME ZONE,
    warranty_expires_at DATE,
    next_service_reminder_date DATE,
    down_payment DECIMAL(15, 2) DEFAULT 0,
    grand_total DECIMAL(15, 2) DEFAULT 0,
    technician_commission DECIMAL(15, 2) DEFAULT 0,
    shop_profit DECIMAL(15, 2) DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE service_details (
    detail_id SERIAL PRIMARY KEY,
    service_job_id INTEGER NOT NULL REFERENCES service_jobs(service_job_id),
    item_id INTEGER NOT NULL,
    item_type VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    serial_number_used VARCHAR(255),
    quantity INTEGER NOT NULL,
    price_per_item DECIMAL(15, 2) NOT NULL,
    cost_per_item DECIMAL(15, 2) NOT NULL
);

CREATE TABLE service_job_histories (
    history_id SERIAL PRIMARY KEY,
    service_job_id INTEGER NOT NULL REFERENCES service_jobs(service_job_id),
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    notes TEXT,
    changed_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Transaction Tables
CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    invoice_number VARCHAR(255) UNIQUE NOT NULL,
    transaction_date TIMESTAMP WITH TIME ZONE NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    customer_id INTEGER REFERENCES customers(customer_id),
    outlet_id INTEGER NOT NULL REFERENCES outlets(outlet_id),
    transaction_type VARCHAR(255) NOT NULL,
    status transaction_status NOT NULL DEFAULT 'sukses',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE transaction_details (
    detail_id SERIAL PRIMARY KEY,
    transaction_type VARCHAR(255) NOT NULL,
    transaction_id INTEGER NOT NULL REFERENCES transactions(transaction_id),
    product_id INTEGER REFERENCES products(product_id),
    serial_number_id INTEGER REFERENCES product_serial_numbers(serial_number_id),
    quantity INTEGER NOT NULL,
    unit_price DECIMAL(15, 2) NOT NULL,
    total_price DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE purchase_orders (
    purchase_order_id SERIAL PRIMARY KEY,
    po_code VARCHAR(50) UNIQUE NOT NULL,
    supplier_id INTEGER NOT NULL REFERENCES suppliers(supplier_id),
    outlet_id INTEGER NOT NULL REFERENCES outlets(outlet_id),
    po_date DATE NOT NULL,
    total_amount DECIMAL(15, 2) NOT NULL,
    amount_paid DECIMAL(15, 2) NOT NULL DEFAULT 0,
    change_amount DECIMAL(15, 2) NOT NULL DEFAULT 0,
    payment_type payment_type_enum NOT NULL,
    status purchase_status NOT NULL DEFAULT 'Selesai',
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE purchase_order_details (
    detail_id SERIAL PRIMARY KEY,
    purchase_order_id INTEGER NOT NULL REFERENCES purchase_orders(purchase_order_id),
    product_id INTEGER NOT NULL REFERENCES products(product_id),
    quantity INTEGER NOT NULL,
    cost_price DECIMAL(15, 2) NOT NULL
);

CREATE TABLE vehicle_purchases (
    purchase_id SERIAL PRIMARY KEY,
    purchase_code VARCHAR(50) UNIQUE NOT NULL,
    customer_id INTEGER REFERENCES customers(customer_id),
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    outlet_id INTEGER NOT NULL REFERENCES outlets(outlet_id),
    purchase_date DATE NOT NULL,
    purchase_price DECIMAL(15, 2) NOT NULL,
    vehicle_snapshot TEXT NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Financial Tables
CREATE TABLE payment_methods (
    method_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    status status_umum NOT NULL DEFAULT 'Aktif',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    transaction_id INTEGER NOT NULL REFERENCES transactions(transaction_id),
    method_id INTEGER NOT NULL REFERENCES payment_methods(method_id),
    amount DECIMAL(15, 2) NOT NULL,
    status transaction_status NOT NULL DEFAULT 'sukses',
    payment_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE accounts_payables (
    payable_id SERIAL PRIMARY KEY,
    purchase_order_id INTEGER NOT NULL REFERENCES purchase_orders(purchase_order_id),
    supplier_id INTEGER NOT NULL REFERENCES suppliers(supplier_id),
    total_amount DECIMAL(15, 2) NOT NULL,
    amount_paid DECIMAL(15, 2) NOT NULL DEFAULT 0,
    due_date DATE NOT NULL,
    status ap_ar_status NOT NULL DEFAULT 'Belum Lunas',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE payable_payments (
    payment_id SERIAL PRIMARY KEY,
    payable_id INTEGER NOT NULL REFERENCES accounts_payables(payable_id),
    payment_date DATE NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE accounts_receivables (
    receivable_id SERIAL PRIMARY KEY,
    transaction_id INTEGER NOT NULL REFERENCES transactions(transaction_id),
    customer_id INTEGER NOT NULL REFERENCES customers(customer_id),
    total_amount DECIMAL(15, 2) NOT NULL,
    amount_paid DECIMAL(15, 2) NOT NULL DEFAULT 0,
    due_date DATE NOT NULL,
    status ap_ar_status NOT NULL DEFAULT 'Belum Lunas',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE receivable_payments (
    payment_id SERIAL PRIMARY KEY,
    receivable_id INTEGER NOT NULL REFERENCES accounts_receivables(receivable_id),
    payment_date DATE NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE cash_flows (
    cash_flow_id SERIAL PRIMARY KEY,
    type cash_flow_type NOT NULL,
    source VARCHAR(255) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    date DATE NOT NULL,
    notes TEXT,
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Reporting & Promotion Tables
CREATE TABLE reports (
    report_id SERIAL PRIMARY KEY,
    report_name VARCHAR(255) UNIQUE NOT NULL,
    report_type report_type_enum NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    outlet_id INTEGER REFERENCES outlets(outlet_id),
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    generated_file_path VARCHAR(255),
    status report_status DEFAULT 'Pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

CREATE TABLE promotions (
    promotion_id SERIAL PRIMARY KEY,
    promotion_name VARCHAR(255) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    type VARCHAR(50) CHECK (type IN ('percentage', 'fixed')) NOT NULL,
    value DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by INTEGER
);

-- Create indexes for better performance
CREATE INDEX idx_users_outlet_id ON users(outlet_id);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_customers_phone_number ON customers(phone_number);
CREATE INDEX idx_customer_vehicles_customer_id ON customer_vehicles(customer_id);
CREATE INDEX idx_customer_vehicles_plate_number ON customer_vehicles(plate_number);
CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_supplier_id ON products(supplier_id);
CREATE INDEX idx_products_unit_type_id ON products(unit_type_id);
CREATE INDEX idx_products_sku ON products(sku);
CREATE INDEX idx_products_barcode ON products(barcode);
CREATE INDEX idx_product_serial_numbers_product_id ON product_serial_numbers(product_id);
CREATE INDEX idx_product_serial_numbers_serial_number ON product_serial_numbers(serial_number);
CREATE INDEX idx_services_service_category_id ON services(service_category_id);
CREATE INDEX idx_services_service_code ON services(service_code);
CREATE INDEX idx_service_jobs_customer_id ON service_jobs(customer_id);
CREATE INDEX idx_service_jobs_vehicle_id ON service_jobs(vehicle_id);
CREATE INDEX idx_service_jobs_technician_id ON service_jobs(technician_id);
CREATE INDEX idx_service_jobs_outlet_id ON service_jobs(outlet_id);
CREATE INDEX idx_service_jobs_service_code ON service_jobs(service_code);
CREATE INDEX idx_service_details_service_job_id ON service_details(service_job_id);
CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transactions_customer_id ON transactions(customer_id);
CREATE INDEX idx_transactions_outlet_id ON transactions(outlet_id);
CREATE INDEX idx_transaction_details_transaction_id ON transaction_details(transaction_id);
CREATE INDEX idx_transaction_details_product_id ON transaction_details(product_id);
CREATE INDEX idx_purchase_orders_supplier_id ON purchase_orders(supplier_id);
CREATE INDEX idx_purchase_orders_outlet_id ON purchase_orders(outlet_id);
CREATE INDEX idx_payments_transaction_id ON payments(transaction_id);
CREATE INDEX idx_payments_method_id ON payments(method_id);
CREATE INDEX idx_accounts_payables_purchase_order_id ON accounts_payables(purchase_order_id);
CREATE INDEX idx_accounts_receivables_transaction_id ON accounts_receivables(transaction_id);
CREATE INDEX idx_cash_flows_user_id ON cash_flows(user_id);
CREATE INDEX idx_reports_outlet_id ON reports(outlet_id);
CREATE INDEX idx_reports_user_id ON reports(user_id);