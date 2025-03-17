CREATE SCHEMA IF NOT EXISTS schema_name;
CREATE TABLE if NOT EXISTS schema_name.orders (
    id bigserial NOT NULL constraint orders_pk primary key,
    name text NOT NULL,
    price integer default 0
);