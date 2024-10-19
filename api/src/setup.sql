DROP DATABASE IF EXISTS breakfast_db;

CREATE DATABASE breakfast_db WITH TEMPLATE=template0 OWNER=postgres;
\connect breakfast_db;

CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL,
    title VARCHAR(32) NOT NULL,
    description VARCHAR(255) NOT NULL,
    emoji VARCHAR(16) NOT NULL,
    color CHAR(7) NOT NULL,
    text_color CHAR(7) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
