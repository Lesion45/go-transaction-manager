CREATE SCHEMA IF NOT EXISTS orders_schema;

CREATE TABLE IF NOT EXISTS orders_schema.order (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    service_id UUID,
    amount FLOAT,
    info TEXT
);