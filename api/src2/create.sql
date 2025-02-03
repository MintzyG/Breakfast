-- Create the database
CREATE DATABASE IF NOT EXISTS breakfast_db;

-- Create the user with a password (replace 'your_password' with the desired password)
CREATE USER IF NOT EXISTS 'breakfast_user'@'localhost' IDENTIFIED BY '';

-- Grant all privileges on the database to the user
GRANT ALL PRIVILEGES ON breakfast_db.* TO 'breakfast_user'@'localhost';

-- Apply the privilege changes
FLUSH PRIVILEGES;

