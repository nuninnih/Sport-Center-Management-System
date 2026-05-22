package handler

import (
	"context"
	"fmt"
	"time"
)

type RevenueInAYear struct {
	TotalCity        int
	TotalField       int
	TotalTransaction int
	TotalRevenue     float64
}

type RevenuePerCity struct {
	CityName         string
	TotalField       int
	TotalTransaction int
	TotalRevenue     float64
}

type FieldMostRevenue struct {
	FieldName     string
	FieldType     string
	CityName      string
	TotalBookings int
	TotalRevenue  float64
}

type RevenuePerType struct {
	FieldType    string
	TotalRevenue float64
}

type MostSpender struct {
	UserID        int
	UserName      string
	PhoneNumber   string
	TotalBookings int
	TotalHours    float64
	TotalSpend    float64
}

func (h *Handler) ReportRevenueInAYear() ([]RevenueInAYear, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := h.DB.QueryContext(ctx, `
	SELECT
		(SELECT COUNT(*) FROM Cities) AS TotalCity,
		(SELECT COUNT(*) FROM Fields WHERE IsActive = TRUE) AS TotalField,
		COUNT(p.ID) AS TotalTransaction,
		IFNULL(SUM(p.Amount), 0) AS TotalRevenue
	FROM Payments p
	WHERE
		p.PaymentStatus = 'PAID'
		AND YEAR(p.PaidAt) = YEAR(CURRENT_DATE);`)

	if err != nil {
		fmt.Println("Error querying data:", err)
		return nil, err
	}
	defer rows.Close()

	var report []RevenueInAYear

	for rows.Next() {
		var TotalCity int
		var TotalField int
		var TotalTransaction int
		var TotalRevenue float64
		err := rows.Scan(&TotalCity, &TotalField, &TotalTransaction, &TotalRevenue)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		report = append(report, RevenueInAYear{TotalCity: TotalCity, TotalField: TotalField, TotalTransaction: TotalTransaction, TotalRevenue: TotalRevenue})
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return report, nil
}

func (h *Handler) ReportRevenuePerCity() ([]RevenuePerCity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := h.DB.QueryContext(ctx, `
	SELECT
		c.CityName,
		COUNT(f.ID) AS TotalField,
		COUNT(p.ID) AS TotalTransactions,
		SUM(p.Amount) AS TotalRevenue
	FROM Payments p
	JOIN Bookings b
		ON p.BookingID = b.ID
	JOIN Fields f
		ON b.FieldID = f.ID
	JOIN Cities c
		ON f.CityID = c.ID
	WHERE p.PaymentStatus = 'PAID'
	AND YEAR(p.PaidAt) = YEAR(CURRENT_DATE)
	GROUP BY c.ID, c.CityName
	ORDER BY TotalRevenue DESC;`)

	if err != nil {
		fmt.Println("Error querying data:", err)
		return nil, err
	}
	defer rows.Close()

	var report []RevenuePerCity

	for rows.Next() {
		var CityName string
		var TotalField int
		var TotalTransaction int
		var TotalRevenue float64
		err := rows.Scan(&CityName, &TotalField, &TotalTransaction, &TotalRevenue)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		report = append(report, RevenuePerCity{CityName: CityName, TotalField: TotalField, TotalTransaction: TotalTransaction, TotalRevenue: TotalRevenue})
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return report, nil
}

func (h *Handler) ReportFieldWithMostRevenue() ([]FieldMostRevenue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := h.DB.QueryContext(ctx, `
	SELECT
		f.FieldName,
		ft.TypeName AS FieldType,
		c.CityName,
		COUNT(b.ID) AS TotalBookings,
		SUM(p.Amount) AS TotalRevenue
	FROM Payments p
	JOIN Bookings b
		ON p.BookingID = b.ID
	JOIN Fields f
		ON b.FieldID = f.ID
	JOIN FieldTypes ft
		ON f.FieldTypeID = ft.ID
	JOIN Cities c
		ON f.CityID = c.ID
	WHERE p.PaymentStatus = 'PAID'
	AND YEAR(p.PaidAt) = YEAR(CURRENT_DATE)
	GROUP BY f.ID
	ORDER BY TotalRevenue DESC;`)

	if err != nil {
		fmt.Println("Error querying data:", err)
		return nil, err
	}
	defer rows.Close()

	var report []FieldMostRevenue

	for rows.Next() {
		var FieldName string
		var FieldType string
		var CityName string
		var TotalBookings int
		var TotalRevenue float64
		err := rows.Scan(&FieldName, &FieldType, &CityName, &TotalBookings, &TotalRevenue)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		report = append(report, FieldMostRevenue{FieldName: FieldName, FieldType: FieldType, CityName: CityName, TotalBookings: TotalBookings, TotalRevenue: TotalRevenue})
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return report, nil
}

func (h *Handler) ReportRevenuePerType() ([]RevenuePerType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := h.DB.QueryContext(ctx, `
	SELECT
		ft.TypeName,
		IFNULL(SUM(p.Amount), 0) AS TotalRevenue
	FROM Payments p
	JOIN Bookings b
		ON p.BookingID = b.ID
	JOIN Fields f
		ON b.FieldID = f.ID
	JOIN FieldTypes ft
		ON f.FieldTypeID = ft.ID
	WHERE
		p.PaymentStatus = 'PAID'
		AND YEAR(p.PaidAt) = YEAR(CURRENT_DATE)
	GROUP BY ft.ID, ft.TypeName
	ORDER BY TotalRevenue DESC;`)

	if err != nil {
		fmt.Println("Error querying data:", err)
		return nil, err
	}
	defer rows.Close()

	var report []RevenuePerType

	for rows.Next() {
		var FieldType string
		var TotalRevenue float64
		err := rows.Scan(&FieldType, &TotalRevenue)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		report = append(report, RevenuePerType{FieldType: FieldType, TotalRevenue: TotalRevenue})
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return report, nil
}

func (h *Handler) ReportMostSpender() ([]MostSpender, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := h.DB.QueryContext(ctx, `
	SELECT
		u.ID,
		CONCAT(u.FirstName, ' ', u.LastName) AS CustomerName,
		u.PhoneNumber,
		COUNT(b.ID) AS TotalBookings,
		SUM(b.TotalHours) AS TotalHours,
		SUM(b.TotalPrice) AS TotalSpent
	FROM Users u
	JOIN Bookings b
		ON u.ID = b.UserID
	WHERE u.UserRole = 'CUSTOMER'
	AND b.BookingStatus IN ('CONFIRMED', 'COMPLETED')
	GROUP BY u.ID
	HAVING SUM(b.TotalPrice) > 1000000
	ORDER BY TotalSpent DESC;`)

	if err != nil {
		fmt.Println("Error querying data:", err)
		return nil, err
	}
	defer rows.Close()

	var report []MostSpender

	for rows.Next() {
		var UserID int
		var UserName string
		var PhoneNumber string
		var TotalBookings int
		var TotalHours float64
		var TotalSpend float64
		err := rows.Scan(&UserID, &UserName, &PhoneNumber, &TotalBookings, &TotalHours, &TotalSpend)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		report = append(report, MostSpender{UserID: UserID, UserName: UserName, PhoneNumber: PhoneNumber, TotalBookings: TotalBookings, TotalHours: TotalHours, TotalSpend: TotalSpend})
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return report, nil
}
