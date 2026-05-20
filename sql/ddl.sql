-- CREATE DB
CREATE DATABASE IF NOT EXISTS SportCenter

USE SportCenter

-- CREATE TABLES
-- Users
CREATE TABLE IF NOT EXISTS Users(
	ID INT AUTO_INCREMENT PRIMARY KEY,
	FirstName VARCHAR(25) NOT NULL,
	LastName VARCHAR(25),
	Email VARCHAR(100) NOT NULL UNIQUE,
	PhoneNumber VARCHAR(25) NOT NULL,
	Password VARCHAR(255) NOT NULL,
	UserRole ENUM ('CUSTOMER', 'ADMIN') DEFAULT 'CUSTOMER' 
);

-- Cities
CREATE TABLE Cities (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    CityName VARCHAR(100) NOT NULL
);

-- FieldTypes
CREATE TABLE FieldTypes (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    TypeName VARCHAR(100) NOT NULL
);

-- Fields
CREATE TABLE Fields (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    CityID INT NOT NULL,
    FieldName VARCHAR(100) NOT NULL,
    FieldTypeID INT NOT NULL,
    Address TEXT,
    HourlyRate DECIMAL(12,2) NOT NULL,
    IsActive BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (CityID) REFERENCES Cities(ID),
    FOREIGN KEY (FieldTypeID) REFERENCES FieldTypes(ID)
);

-- Bookings
CREATE TABLE Bookings (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT NOT NULL,
    FieldID INT NOT NULL,
    BookingDate DATE NOT NULL,
    StartTime TIME NOT NULL,
    EndTime TIME NOT NULL,
    TotalHours DECIMAL(5,2) NOT NULL,
    TotalPrice DECIMAL(12,2) NOT NULL,
    BookingStatus ENUM(
        'PENDING',
        'CONFIRMED',
        'CANCELLED',
        'COMPLETED'
    ) DEFAULT 'PENDING',
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    FOREIGN KEY (FieldID) REFERENCES Fields(ID)
);

-- Payments
CREATE TABLE Payments (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    BookingID INT NOT NULL,
    PaymentMethod ENUM(
        'CASH',
        'BANK_TRANSFER',
        'E_WALLET',
        'CREDIT_CARD'
    ) NOT NULL,
    PaymentStatus ENUM(
        'PENDING',
        'PAID',
        'FAILED',
        'REFUNDED'
    ) DEFAULT 'PENDING',
    Amount DECIMAL(12,2) NOT NULL CHECK (Amount > 0.0),
    PaidAt TIMESTAMP NULL,
    FOREIGN KEY (BookingID) REFERENCES Bookings(ID)
);

-- DROP ALL TABLES
DROP TABLE IF EXISTS Payments, 
Bookings, Fields, FieldTypes, Cities, Users;