package handler

import (
	"database/sql"
	"fmt"
)

// handler struct
type Handler struct {
	DB *sql.DB
}

// create handler
func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		DB: db,
	}
}

// register user
func (h *Handler) Register(username, password, role string) {

	query := `
	INSERT INTO users(username, password, role)
	VALUES (?, ?, ?)
	`

	_, err := h.DB.Exec(query, username, password, role)

	if err != nil {
		fmt.Println("register failed")
		return
	}

	fmt.Println("register success")
}
