-- Drop existing database and user if they exist
DROP DATABASE IF EXISTS breakfast_db;

-- Create user and database
CREATE DATABASE breakfast_db WITH TEMPLATE=template0 OWNER=postgres;

-- Connect to the new database
\connect breakfast_db;

-- Create users table
CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL
);

