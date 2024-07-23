CREATE SCHEMA IF NOT EXISTS users_schema;

CREATE TABLE IF NOT EXISTS users_schema.user (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    balance FLOAT,
    reserved FLOAT
);