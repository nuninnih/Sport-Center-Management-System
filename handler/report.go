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
