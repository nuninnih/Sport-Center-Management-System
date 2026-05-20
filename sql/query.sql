-- CUSTOMER

-- Register User
INSERT INTO users(username, password, role)
VALUES ('komara', '123', 'costumer')

-- Login User
SELECT * FROM users
WHERE username = 'komara'
AND password = '123'

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
WHERE id = 1:

-- Pay For Booking

-- Report

-- Monthly Revenue
SUM(price) FROM bookings
JOIN fields ON bookings.field_id = fields.id
WHERE MONTH(booking_date) = MONTH(CURRENT_DATE())
AND YEAR(booking_date) = YEAR(CURRENT_DATE())

-- Most Popular City
SELECT sport_type, COUNT(*) AS booking_count
FROM bookings
JOIN fields ON bookings.field_id = fields.id
GROUP BY sport_type
ORDER BY booking_count DESC
LIMIT 1;

-- Most Popular FieldType
SELECT field_name, COUNT(*) AS booking_count
FROM bookings
JOIN fields ON bookings.field_id = fields.id
GROUP BY field_name
ORDER BY booking_count DESC
LIMIT 1;

-- Most Popular Field
SELECT field_name, COUNT(*) AS booking_count
FROM bookings
JOIN fields ON bookings.field_id = fields.id
GROUP BY field_name
ORDER BY booking_count DESC
LIMIT 1;

-- Most Booking Report (Daily, Upcoming, By Status)
-- Most Payment Report (Daily, By Status)
-- Peak Hour
-- Customer With Most Total Spend
