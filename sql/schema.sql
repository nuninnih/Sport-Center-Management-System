-- create database
CREATE DATABASE IF NOT EXISTS sport_center;

USE sport_center;

-- users table
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    role ENUM('admin', 'customer') NOT NULL
);

-- fields table
CREATE TABLE fields (
    id INT PRIMARY KEY AUTO_INCREMENT,
    field_name VARCHAR(100),
    sport_type VARCHAR(50),
    price INT,
    status VARCHAR(20)
);

-- bookings table
CREATE TABLE bookings (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    field_id INT,
    booking_date DATE,

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (field_id) REFERENCES fields(id)
);