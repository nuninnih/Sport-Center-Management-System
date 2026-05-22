package handler

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type CheckBooking struct {
	FieldID     int
	FieldName   string
	FieldType   string
	CityName    string
	BookingDate string
	StartTime   string
	EndTime     string
}

type PendingBooking struct {
	BookingID   int
	FieldName   string
	UserName    string
	PhoneNumber string
	BookingDate string
	StartTime   string
	EndTime     string
}

func (h *Handler) CheckAvailableSlot(BookingDate, CityName, TypeName string) ([]CheckBooking, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := h.DB.QueryContext(ctx, `
	SELECT
		f.id,
		f.FieldName,
		ft.TypeName,
		c.CityName,
		IFNULL(b.BookingDate, '') AS BookingDate,
		IFNULL(b.StartTime, '') AS StartTime,
		IFNULL(b.EndTime, '') AS EndTime
	FROM Fields f
	LEFT JOIN Bookings b
		ON f.ID = b.FieldID
		AND b.BookingDate = ?
		AND b.BookingStatus IN ('CONFIRMED', 'COMPLETED')
	JOIN FieldTypes ft
		ON f.FieldTypeID = ft.ID
	JOIN Cities c
		ON f.CityID = c.ID
	WHERE
		f.IsActive = TRUE
		AND LOWER(c.CityName) = LOWER(?)
		AND LOWER(ft.TypeName) = LOWER(?);
		;
	`, BookingDate, CityName, TypeName)
	if err != nil {
		fmt.Println("Error querying data:", err)
		return nil, err
	}
	defer rows.Close()

	var check []CheckBooking

	for rows.Next() {
		var FieldID int
		var FieldName string
		var TypeName string
		var CityName string
		var BookingDate string
		var StartTime string
		var EndTime string
		err := rows.Scan(&FieldID, &FieldName, &TypeName, &CityName, &BookingDate, &StartTime, &EndTime)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		check = append(check, CheckBooking{FieldID: FieldID, FieldName: FieldName, FieldType: TypeName, CityName: CityName, BookingDate: BookingDate, StartTime: StartTime, EndTime: EndTime})
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return check, nil
}

func (h *Handler) CreateBooking(UserID, FieldID int, BookingDate, StartTime, EndTime string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	start, err := time.Parse("15:04", StartTime)
	if err != nil {
		fmt.Println("Error parsing start time:", err)
		return err
	}
	end, err := time.Parse("15:04", EndTime)
	if err != nil {
		fmt.Println("Error parsing start time:", err)
		return err
	}
	duration := end.Sub(start)
	totalHours := duration.Hours()

	fieldById, err := h.GetFieldByID(FieldID)
	if err != nil {
		return err
	}
	totalPrice := fieldById.HourlyRate * totalHours

	_, err = h.DB.ExecContext(ctx, "INSERT INTO Bookings (UserID, FieldID, BookingDate, StartTime, EndTime, TotalHours, TotalPrice) VALUES (?, ?, ?, ?, ?, ?, ?)", UserID, FieldID, BookingDate, StartTime, EndTime, totalHours, totalPrice)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return err
	}
	return nil
}

func (h *Handler) UpdateBookingStatus(BookingID int, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var finalStatus string

	switch strings.ToLower(status) {
	case "cancel":
		finalStatus = "CANCELLED"
	case "confirm":
		finalStatus = "CONFIRMED"
	case "complete":
		finalStatus = "COMPLETED"
	}

	_, err := h.DB.ExecContext(ctx, "UPDATE Bookings SET BookingStatus = ? WHERE ID = ?", finalStatus, BookingID)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return err
	}
	return nil
}

func (h *Handler) CheckPendingBooking() ([]PendingBooking, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := h.DB.QueryContext(ctx, `
	SELECT
		b.ID,
		f.FieldName,
		u.FirstName,
		u.PhoneNumber,
		b.BookingDate,
		b.StartTime,
		b.EndTime
	FROM Fields f
	JOIN Bookings b
		ON f.ID = b.FieldID
	JOIN Users u
		ON b.UserID = u.ID
	WHERE
		f.IsActive = TRUE
		AND
		b.BookingStatus = 'PENDING'
		;
	`)
	if err != nil {
		fmt.Println("Error querying data:", err)
		return nil, err
	}
	defer rows.Close()

	var check []PendingBooking

	for rows.Next() {
		var BookingID int
		var FieldName string
		var UserName string
		var PhoneNumber string
		var BookingDate string
		var StartTime string
		var EndTime string
		err := rows.Scan(&BookingID, &FieldName, &UserName, &PhoneNumber, &BookingDate, &StartTime, &EndTime)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		check = append(check, PendingBooking{BookingID: BookingID, FieldName: FieldName, UserName: UserName, PhoneNumber: PhoneNumber, BookingDate: BookingDate, StartTime: StartTime, EndTime: EndTime})
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return check, nil
}
