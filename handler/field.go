package handler

import (
	"context"
	"database/sql"
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

func (h *Handler) GetAllField() ([]Field, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := h.DB.QueryContext(ctx, `
	SELECT 
		f.ID as FieldID,
		f.FieldName,
		ft.TypeName,
		f.address,
		c.CityName,
		f.HourlyRate
	FROM Fields f
	JOIN FieldTypes ft
	ON f.FieldTypeID = ft.ID
	JOIN Cities c
	ON f.CityID = c.ID
	WHERE
		f.IsActive = TRUE;
	`)
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
		var Address string
		var CityName string
		var HourlyRate float64
		err := rows.Scan(&FieldID, &FieldName, &TypeName, &Address, &CityName, &HourlyRate)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		fields = append(fields, Field{FieldID: FieldID, FieldName: FieldName, FieldType: TypeName, Address: Address, City: CityName, HourlyRate: HourlyRate})
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return fields, nil
}

func (h *Handler) GetFieldByID(id int) (Field, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := h.DB.QueryRowContext(ctx, `
	SELECT 
		f.ID as FieldID,
		f.FieldName,
		ft.TypeName,
		f.address,
		c.CityName,
		f.HourlyRate
	FROM Fields f
	JOIN FieldTypes ft
	ON f.FieldTypeID = ft.ID
	JOIN Cities c
	ON f.CityID = c.ID
	WHERE f.ID = ?
	`, id)

	var field Field
	err := row.Scan(&field.FieldID, &field.FieldName, &field.FieldType, &field.Address, &field.City, &field.HourlyRate)
	if err != nil {
		if err == sql.ErrNoRows {
			return Field{}, nil
		}
		fmt.Println("Error scanning row:", err)
		return Field{}, err
	}
	return field, nil
}
