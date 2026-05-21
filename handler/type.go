package handler

import (
	"context"
	"fmt"
	"time"
)

type Type struct {
	TypeID   int
	TypeName string
}

func (h *Handler) GetAllTypes() ([]Type, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := h.DB.QueryContext(ctx, `
	SELECT 
		*
	FROM FieldTypes
	`)
	if err != nil {
		fmt.Println("Error querying data:", err)
		return nil, err
	}
	defer rows.Close()

	var types []Type

	for rows.Next() {
		var TypeID int
		var TypeName string

		err := rows.Scan(&TypeID, &TypeName)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		types = append(types, Type{TypeID: TypeID, TypeName: TypeName})
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return types, nil
}
