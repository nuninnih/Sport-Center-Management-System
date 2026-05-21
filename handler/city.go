package handler

import (
	"context"
	"fmt"
	"time"
)

type City struct {
	CityID   int
	CityName string
}

func (h *Handler) GetAllCities() ([]City, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := h.DB.QueryContext(ctx, `
	SELECT 
		*
	FROM Cities
	`)
	if err != nil {
		fmt.Println("Error querying data:", err)
		return nil, err
	}
	defer rows.Close()

	var cities []City

	for rows.Next() {
		var CityID int
		var CityName string

		err := rows.Scan(&CityID, &CityName)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		cities = append(cities, City{CityID: CityID, CityName: CityName})
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}
	return cities, nil
}
