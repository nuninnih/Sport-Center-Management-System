package handler

import (
	"context"
	"fmt"
	"time"
)

type PendingPayment struct {
	BookingID   int
	FieldName   string
	UserName    string
	PhoneNumber string
	BookingDate string
	Total       string
}

func (h *Handler) CreatePayment(BookingId, PaymentMethod, Amount int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var method string

	switch PaymentMethod {
	case 1:
		method = "CASH"
	case 2:
		method = "BANK_TRANSFER"
	case 3:
		method = "E_WALLET"
	case 4:
		method = "CREDIT_CARD"
	default:
		method = "CASH"
	}

	_, err := h.DB.ExecContext(ctx, "INSERT INTO Payments (BookingId, PaymentMethod, Amount, PaymentStatus, PaidAt) VALUES (?, ?, ?, 'PAID', CURRENT_TIMESTAMP)", BookingId, method, Amount)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return err
	}
	return nil
}

func (h *Handler) CheckPendingPayment() ([]PendingPayment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := h.DB.QueryContext(ctx, `
	SELECT
		b.ID,
		f.FieldName,
		u.FirstName,
		u.PhoneNumber,
		b.BookingDate,
		b.TotalPrice,
		IFNULL(p.ID, '') AS PaymentID
	FROM Fields f
	JOIN Bookings b
		ON f.ID = b.FieldID
	JOIN Users u
		ON b.UserID = u.ID
	LEFT JOIN Payments p
		ON p.BookingID = b.ID
	WHERE p.ID IS NULL
	`)
	if err != nil {
		fmt.Println("Error querying data:", err)
		return nil, err
	}
	defer rows.Close()

	var check []PendingPayment

	for rows.Next() {
		var BookingID int
		var FieldName string
		var UserName string
		var PhoneNumber string
		var BookingDate string
		var TotalPrice string
		var PaymentID string
		err := rows.Scan(&BookingID, &FieldName, &UserName, &PhoneNumber, &BookingDate, &TotalPrice, &PaymentID)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		check = append(check, PendingPayment{BookingID: BookingID, FieldName: FieldName, UserName: UserName, PhoneNumber: PhoneNumber, BookingDate: BookingDate, Total: TotalPrice})
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return check, nil
}
