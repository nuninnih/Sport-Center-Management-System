-- CUSTOMER

-- Register User
INSERT INTO users(username, password, role)
VALUES ('komara', '123', 'customer')

-- Login User
SELECT * FROM users
WHERE username = 'komara'
AND password = '123';

-- Check Available Field
SELECT * FROM fields
WHERE status = 'Available'

-- Create Booking Field
INSERT INTO bookings(user_id, field_id, booking_date)
VALUES (1, 2, '2026-05-20');

-- ADMIN

-- Create New Field
INSERT INTO fields(field_name, sport_type, price, status)
VALUES ('Futsal A', 'Futsal', '100000', 'Available')

-- Update Field
UPDATE fields
SET price = 120000
WHERE id = 1;

-- Create New FieldType
INSERT INTO fields(field_name, sport_type, price, status)
VALUES ('Badminton A', 'Badminton', 50000, 'Available')

-- Update FieldType
UPDATE fields
SET sport_type = 'Mini Soccer'
WHERE id = 1;

-- Pay For Booking
SELECT bookings.id, users.username, fields.field_name, fields.price
FROM bookings
JOIN users ON bookings.user_id = users.id
JOIN fields ON bookings.field_id = fields.id    

-- Report

-- Monthly Revenue
SELECT SUM(amount) AS monthly_revenue
FROM payments
WHERE MONTH(payment_date) = MONTH(CURRENT_DATE())
AND YEAR(payment_date) = YEAR(CURRENT_DATE());

-- Most Popular SportType
SELECT fields.sport_type, COUNT(*) AS booking_count
FROM bookings
JOIN fields ON bookings.field_id = fields.id
GROUP BY fields.sport_type
ORDER BY booking_count DESC
LIMIT 1;

-- Most Popular City
SELECT fields.city, COUNT(*) AS booking_count
FROM bookings
JOIN fields ON bookings.field_id = fields.id
GROUP BY fields.city
ORDER BY booking_count DESC
LIMIT 1;

-- Most Popular FieldType
SELECT fields.field_name, COUNT(*) AS booking_count
FROM bookings
JOIN fields ON bookings.field_id = fields.id
GROUP BY fields.field_name
ORDER BY booking_count DESC
LIMIT 1;

-- Most Popular Field
SELECT fields.field_name, COUNT(*) AS booking_count
FROM bookings
JOIN fields ON bookings.field_id = fields.id
GROUP BY fields.field_name
ORDER BY booking_count DESC
LIMIT 1;

-- Most Booking Report (Daily, Upcoming, By Status)
SELECT fields.field_name, COUNT(*) AS booking_count
FROM bookings
JOIN fields ON bookings.field_id = fields.id
WHERE booking_date >= CURRENT_DATE()
GROUP BY fields.field_name
ORDER BY booking_count DESC
LIMIT 1;

-- Most Payment Report (Daily, By Status)
SELECT fields.field_name, SUM(fields.price) AS total_payment
FROM bookings
JOIN fields ON bookings.field_id = fields.id
WHERE booking_date >= CURRENT_DATE()
GROUP BY fields.field_name
ORDER BY total_payment DESC
LIMIT 1;

-- Most Active Customer
SELECT users.username, COUNT(*) AS booking_count
FROM bookings
JOIN users ON bookings.user_id = users.id
GROUP BY users.username
ORDER BY booking_count DESC
LIMIT 1;

-- Payment Status Report
SELECT fields.field_name, payments.status, COUNT(*) AS payment_count
FROM payments
JOIN fields ON bookings.field_id = fields.id
JOIN payments ON bookings.id = payments.booking_id
GROUP BY fields.field_name, payments.status
ORDER BY payment_count DESC;

-- Booking Status Report
SELECT fields.field_name, bookings.status, COUNT(*) AS booking_count
FROM bookings
JOIN fields ON bookings.field_id = fields.id
GROUP BY fields.field_name, bookings.status
ORDER BY booking_count DESC;

-- Peak Hour
SELECT HOUR(booking_date) AS booking_hour, COUNT(*) AS booking_count
FROM bookings
GROUP BY booking_hour
ORDER BY booking_count DESC
LIMIT 1;

-- Customer With Most Total Spend
SELECT users.username, SUM(fields.price) AS total_spend
FROM payments
JOIN users ON payments.user_id = users.id
JOIN fields ON payments.field_id = fields.id
GROUP BY users.username
ORDER BY total_spend DESC
LIMIT 1;