-- DML Dummy dml.sql

-- USER
INSERT INTO users (username, email, password, role) VALUES
('komara', 'komara@example.com', '123', 'customer'),
('admin', 'admin@example.com', '123', 'admin');

-- INSERT DATA TO TABLE
INSERT INTO fields (field_name, sport_type, price, status) VALUES
('Futsal A', 'Futsal', 100000, 'Available'),
('Futsal B', 'Futsal', 120000, 'Available'),
('Badminton A', 'Badminton', 50000, 'Available'),
('Badminton B', 'Badminton', 60000, 'Available');