# Sport-Center-Management-System

## Tables:

### Users

- ID (PK, AI)
- FirstName
- LastName
- Email
- PhoneNumber
- Password
- Role

### Cities

- ID (PK, AI)
- CityName

### FieldTypes

- ID (PK, AI)
- TypeName

### Fields

- ID (PK, AI)
- FieldName
- FieldTypeID (FK to FieldTypes)
- CityId (FK to Cities)
- Address
- HourlyRate
- IsActive

### Bookings

- ID (PK, AI)
- UserId (FK to Users)
- FieldId (FK to Fields)
- BookingDate
- StartTime
- EndTime
- TotalHours
- TotalPrice
- BookingStatus
- CreatedAt

### Payments

- ID (PK, AI)
- BookingId (FK to Bookings)
- PaymentMethod
- PaymentStatus
- Amount
- PaidAt

## Relations

### One To One

- One Booking has One Payment

### One To Many

- One City could has Many Fields
- One FieldType could has Many Fields
- One User could has Many Bookings
- One Field could has Many Bookings

### Many To Many

- Many Users could book Many Fields and Many Fields could booked by Many Users

## Table Notes

- Users.Email should be unique
- Users.Password should be hashes
- Users.Role should be
  'ADMIN',
  'CUSTOMER'
- Bookings.BookingStatus should be
  'PENDING' (Default),
  'CONFIRMED',
  'CANCELLED',
  'COMPLETED'
- Payments.PaymentMethod should be
  'CASH',
  'BANK_TRANSFER',
  'E_WALLET',
  'CREDIT_CARD'
- Payments.PaymentStatus should be
  'PENDING' (Default),
  'PAID',
  'FAILED',
  'REFUNDED'
- Bookings.CreateAt should be Timestamp with default current
- Payments.PaidAt should be Timestamp
