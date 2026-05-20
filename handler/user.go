package handler

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	UserID      int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}

func (h *Handler) Register(FirstName, LastName, Email, PhoneNumber, Password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := h.DB.ExecContext(ctx, "INSERT INTO Users (FirstName, LastName, Email, PhoneNumber, Password) VALUES (?, ?, ?, ?, ?)", FirstName, LastName, Email, PhoneNumber, Password)
	if err != nil {
		fmt.Println("Error registering new user:", err)
		return err
	}
	return nil
}

func (h *Handler) Login(Email, PhoneNumber, Password string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := h.DB.QueryRowContext(ctx, `
	SELECT 
		*
	FROM User
	WHERE 
	 (Email = ? OR PhoneNumber = ? )
	 AND
	 Password = ?
	`, Email, PhoneNumber, Password)

	var user User
	err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, nil
		}
		fmt.Println("Error scanning row:", err)
		return User{}, err
	}
	return user, nil
}
