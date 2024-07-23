CREATE SCHEMA IF NOT EXISTS reports_schema;

CREATE TABLE IF NOT EXISTS reports_schema.report (
    user_id UUID,
    service_id UUID,
    amount FLOAT,
    info TEXT
);