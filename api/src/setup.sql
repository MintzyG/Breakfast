DROP DATABASE IF EXISTS breakfast_db;

CREATE DATABASE breakfast_db WITH TEMPLATE=template0 OWNER=postgres;
\connect breakfast_db;

CREATE TABLE users (
  id UUID PRIMARY KEY UNIQUE NOT NULL,
  first_name VARCHAR(31) NOT NULL,
  last_name VARCHAR(31) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password TEXT NOT NULL
);

CREATE TABLE categories (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id UUID NOT NULL,
  title VARCHAR(127) NOT NULL UNIQUE,
  description VARCHAR(255),
  emoji VARCHAR(15) NOT NULL,
  color CHAR(7) NOT NULL,
  text_color CHAR(7) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE yogurt (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id UUID NOT NULL,
  emoji VARCHAR(15) NOT NULL,
  title VARCHAR(127) NOT NULL,
  description VARCHAR(255),
  completed BOOLEAN NOT NULL,
  difficulty INT NOT NULL,
  task_size INT NOT NULL,
  priority INT NOT NULL,
  category_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);
