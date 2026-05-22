-- DML
-- INSERT DATA TO TABLE
-- Users
INSERT INTO Users
(FirstName, LastName, Email, PhoneNumber, Password, UserRole)
VALUES
('Admin', 'System', 'admin@sporthub.com', '081100000001', 'admin123', 'ADMIN'),
('Andi', 'Saputra', 'andi@gmail.com', '081111111111', 'password', 'CUSTOMER'),
('Budi', 'Santoso', 'budi@gmail.com', '081111111112', 'password', 'CUSTOMER'),
('Citra', 'Lestari', 'citra@gmail.com', '081111111113', 'password', 'CUSTOMER'),
('Dewi', 'Anggraini', 'dewi@gmail.com', '081111111114', 'password', 'CUSTOMER'),
('Eko', 'Pratama', 'eko@gmail.com', '081111111115', 'password', 'CUSTOMER'),
('Farah', 'Putri', 'farah@gmail.com', '081111111116', 'password', 'CUSTOMER'),
('Gilang', 'Ramadhan', 'gilang@gmail.com', '081111111117', 'password', 'CUSTOMER'),
('Hana', 'Permata', 'hana@gmail.com', '081111111118', 'password', 'CUSTOMER'),
('Indra', 'Wijaya', 'indra@gmail.com', '081111111119', 'password', 'CUSTOMER'),
('Joko', 'Susilo', 'joko@gmail.com', '081111111120', 'password', 'CUSTOMER');

-- Cities
INSERT INTO Cities (CityName)
VALUES
('Jakarta'),
('Bandung'),
('Surabaya'),
('Yogyakarta'),
('Bekasi');

-- FieldTypes
INSERT INTO FieldTypes (TypeName)
VALUES
('Futsal'),
('Basketball'),
('Badminton'),
('Tennis'),
('Volleyball'),
('Padel');

-- Fields
INSERT INTO Fields
(CityID, FieldName, FieldTypeID, Address, HourlyRate, IsActive)
VALUES
(1, 'Jakarta Futsal Arena A', 1, 'Jl. Sudirman No. 10', 150000, TRUE),
(1, 'Jakarta Futsal Arena B', 1, 'Jl. Sudirman No. 10', 175000, TRUE),
(1, 'Jakarta Basketball Center', 2, 'Jl. Thamrin No. 5', 200000, TRUE),
(1, 'Jakarta Padel Club', 6, 'Jl. Gatot Subroto No. 20', 250000, TRUE),
(2, 'Bandung Badminton Hall', 3, 'Jl. Asia Afrika No. 7', 80000, TRUE),
(2, 'Bandung Tennis Court', 4, 'Jl. Braga No. 21', 120000, TRUE),
(2, 'Bandung Futsal Elite', 1, 'Jl. Dago No. 88', 140000, TRUE),
(2, 'Bandung Padel Arena', 6, 'Jl. Setiabudi No. 55', 220000, TRUE),
(3, 'Surabaya Sport Arena', 2, 'Jl. Pemuda No. 15', 220000, TRUE),
(3, 'Surabaya Volleyball Court', 5, 'Jl. Darmo No. 12', 100000, TRUE),
(3, 'Surabaya Padel Court', 6, 'Jl. Basuki Rahmat No. 90', 240000, TRUE),
(4, 'Jogja Futsal Center', 1, 'Jl. Malioboro No. 1', 110000, TRUE),
(4, 'Jogja Badminton Court', 3, 'Jl. Kaliurang No. 9', 75000, TRUE),
(4, 'Jogja Padel Space', 6, 'Jl. Solo No. 14', 200000, TRUE),
(5, 'Bekasi Mini Soccer', 1, 'Jl. Ahmad Yani No. 44', 160000, TRUE),
(5, 'Bekasi Tennis Arena', 4, 'Jl. Patriot No. 18', 130000, FALSE),
(5, 'Bekasi Padel Hub', 6, 'Jl. Harapan Indah No. 33', 230000, TRUE);

-- Bookings
INSERT INTO Bookings
(
    UserID,
    FieldID,
    BookingDate,
    StartTime,
    EndTime,
    TotalHours,
    TotalPrice,
    BookingStatus
)
VALUES
(2, 1, '2026-05-01', '18:00:00', '20:00:00', 2, 300000, 'COMPLETED'),
(2, 6, '2026-05-03', '19:00:00', '21:00:00', 2, 280000, 'COMPLETED'),
(3, 3, '2026-05-02', '15:00:00', '17:00:00', 2, 400000, 'COMPLETED'),
(3, 7, '2026-05-06', '18:00:00', '20:00:00', 2, 440000, 'CONFIRMED'),
(4, 4, '2026-05-04', '08:00:00', '10:00:00', 2, 160000, 'COMPLETED'),
(4, 10, '2026-05-05', '13:00:00', '15:00:00', 2, 150000, 'COMPLETED'),
(5, 5, '2026-05-07', '09:00:00', '11:00:00', 2, 240000, 'CANCELLED'),
(6, 8, '2026-05-08', '16:00:00', '18:00:00', 2, 200000, 'COMPLETED'),
(7, 9, '2026-05-09', '19:00:00', '21:00:00', 2, 220000, 'COMPLETED'),
(8, 11, '2026-05-10', '20:00:00', '22:00:00', 2, 320000, 'CONFIRMED'),
(9, 1, '2026-05-11', '10:00:00', '12:00:00', 2, 300000, 'COMPLETED'),
(10, 6, '2026-05-12', '18:00:00', '20:00:00', 2, 280000, 'COMPLETED'),
(11, 3, '2026-05-13', '17:00:00', '19:00:00', 2, 400000, 'PENDING'),
(2, 1, '2026-05-14', '19:00:00', '21:00:00', 2, 300000, 'COMPLETED'),
(3, 1, '2026-05-15', '20:00:00', '22:00:00', 2, 300000, 'COMPLETED'),
(4, 1, '2026-05-16', '18:00:00', '20:00:00', 2, 300000, 'COMPLETED'),
(5, 4, '2026-05-17', '09:00:00', '11:00:00', 2, 160000, 'COMPLETED'),
(6, 7, '2026-05-18', '18:00:00', '20:00:00', 2, 440000, 'COMPLETED'),
(7, 7, '2026-05-19', '19:00:00', '21:00:00', 2, 440000, 'COMPLETED'),
(2, 13, '2026-05-20', '19:00:00', '21:00:00', 2, 500000, 'COMPLETED'),
(3, 14, '2026-05-21', '18:00:00', '20:00:00', 2, 440000, 'COMPLETED'),
(4, 15, '2026-05-22', '20:00:00', '22:00:00', 2, 480000, 'CONFIRMED'),
(5, 16, '2026-05-23', '16:00:00', '18:00:00', 2, 400000, 'COMPLETED'),
(6, 17, '2026-05-24', '09:00:00', '11:00:00', 2, 460000, 'PENDING'),
(7, 13, '2026-05-25', '18:00:00', '20:00:00', 2, 500000, 'COMPLETED'),
(8, 14, '2026-05-26', '19:00:00', '21:00:00', 2, 440000, 'COMPLETED'),
(9, 15, '2026-05-27', '17:00:00', '19:00:00', 2, 480000, 'CANCELLED');

-- Payments
INSERT INTO Payments
(
    BookingID,
    PaymentMethod,
    PaymentStatus,
    Amount,
    PaidAt
)
VALUES
(1, 'E_WALLET', 'PAID', 300000, '2026-05-01 17:30:00'),
(2, 'BANK_TRANSFER', 'PAID', 280000, '2026-05-03 18:30:00'),
(3, 'CREDIT_CARD', 'PAID', 400000, '2026-05-02 14:00:00'),
(4, 'BANK_TRANSFER', 'PENDING', 440000, NULL),
(5, 'CASH', 'PAID', 160000, '2026-05-04 07:30:00'),
(6, 'E_WALLET', 'PAID', 150000, '2026-05-05 12:30:00'),
(7, 'BANK_TRANSFER', 'REFUNDED', 240000, '2026-05-07 08:00:00'),
(8, 'CASH', 'PAID', 200000, '2026-05-08 15:00:00'),
(9, 'E_WALLET', 'PAID', 220000, '2026-05-09 18:00:00'),
(10, 'CREDIT_CARD', 'PENDING', 320000, NULL),
(11, 'BANK_TRANSFER', 'PAID', 300000, '2026-05-11 09:00:00'),
(12, 'E_WALLET', 'PAID', 280000, '2026-05-12 17:00:00'),
(13, 'BANK_TRANSFER', 'FAILED', 400000, NULL),
(14, 'E_WALLET', 'PAID', 300000, '2026-05-14 18:00:00'),
(15, 'E_WALLET', 'PAID', 300000, '2026-05-15 19:00:00'),
(16, 'BANK_TRANSFER', 'PAID', 300000, '2026-05-16 17:00:00'),
(17, 'CASH', 'PAID', 160000, '2026-05-17 08:00:00'),
(18, 'CREDIT_CARD', 'PAID', 440000, '2026-05-18 17:00:00'),
(19, 'E_WALLET', 'PAID', 440000, '2026-05-19 18:00:00'),
(20, 'CREDIT_CARD', 'PAID', 500000, '2026-05-20 18:00:00'),
(21, 'E_WALLET', 'PAID', 440000, '2026-05-21 17:00:00'),
(22, 'BANK_TRANSFER', 'PENDING', 480000, NULL),
(23, 'CASH', 'PAID', 400000, '2026-05-23 15:00:00'),
(24, 'E_WALLET', 'PENDING', 460000, NULL),
(25, 'CREDIT_CARD', 'PAID', 500000, '2026-05-25 17:00:00'),
(26, 'BANK_TRANSFER', 'PAID', 440000, '2026-05-26 18:00:00'),
(27, 'BANK_TRANSFER', 'REFUNDED', 480000, '2026-05-27 16:00:00');

