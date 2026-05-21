package handler

import (
	"context"
	"fmt"
	"time"
)

type Field struct {
	FieldID    int
	FieldName  string
	FieldType  string
	Address    string
	City       string
	HourlyRate float64
}

func (h *Handler) GetAvailableField(CityName, TypeName, BookingDate, StartTime, EndTime string) ([]Field, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := h.DB.QueryContext(ctx, `
	SELECT 
		f.ID as FieldID,
		f.FieldName,
		ft.TypeName,
		c.CityName,
		f.HourlyRate
	FROM Fields f
	JOIN FieldTypes ft
	ON f.FieldTypeID = ft.ID
	JOIN Cities c
	ON f.CityID = c.ID
	WHERE
		f.IsActive = TRUE
	AND c.CityName = ?
	AND ft.TypeName = ?
	AND f.ID NOT IN (
    SELECT b.FieldID
    FROM Bookings b
    WHERE b.BookingDate = ?
    AND b.BookingStatus IN ('CONFIRMED', 'COMPLETED')
    AND (
        ? < b.EndTime
        AND ? > b.StartTime
    )
);
	`, CityName, TypeName, BookingDate, EndTime, StartTime)
	if err != nil {
		fmt.Println("Error querying data:", err)
		return nil, err
	}
	defer rows.Close()

	var fields []Field

	for rows.Next() {
		var FieldID int
		var FieldName string
		var TypeName string
		var CityName string
		var HourlyRate float64
		err := rows.Scan(&FieldID, &FieldName, &TypeName, &CityName, &HourlyRate)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		fields = append(fields, Field{FieldID: FieldID, FieldName: FieldName, FieldType: TypeName, City: CityName, HourlyRate: HourlyRate})
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return fields, nil
}
