package handler_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nuninnih/Sport-Center-Management-System/handler"
	"github.com/stretchr/testify/assert"
)

func TestRegisterSuccess(t *testing.T) {

	db, mock, err := sqlmock.New()

	assert.NoError(
		t,
		err,
		"Expected no error creating sqlmock, received: %v",
		err,
	)

	defer db.Close()

	h := &handler.Handler{
		DB: db,
	}

	firstName := "Nunin"
	lastName := "Farid"
	email := "nunin@mail.com"
	phoneNumber := "08123456789"
	password := "rahasia123"

	mock.ExpectExec(
		"INSERT INTO Users",
	).
		WithArgs(
			firstName,
			lastName,
			email,
			phoneNumber,
			password,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = h.Register(
		firstName,
		lastName,
		email,
		phoneNumber,
		password,
	)

	assert.NoError(
		t,
		err,
		"Expected no register error, received: %v",
		err,
	)

	assert.NoError(
		t,
		mock.ExpectationsWereMet(),
		"There are unmet SQL mock expectations",
	)
}

func TestRegisterError(t *testing.T) {

	db, mock, err := sqlmock.New()

	assert.NoError(
		t,
		err,
		"Expected no error creating sqlmock, received: %v",
		err,
	)

	defer db.Close()

	h := &handler.Handler{
		DB: db,
	}

	mock.ExpectExec(
		"INSERT INTO Users",
	).
		WithArgs(
			"Nunin",
			"Farid",
			"nunin@mail.com",
			"08123456789",
			"rahasia123",
		).
		WillReturnError(sql.ErrConnDone)

	err = h.Register(
		"Nunin",
		"Farid",
		"nunin@mail.com",
		"08123456789",
		"rahasia123",
	)

	assert.Error(
		t,
		err,
		"Expected register error but received nil",
	)

	assert.NoError(
		t,
		mock.ExpectationsWereMet(),
		"There are unmet SQL mock expectations",
	)
}

func TestLoginSuccess(t *testing.T) {

	db, mock, err := sqlmock.New()

	assert.NoError(
		t,
		err,
		"Expected no error creating sqlmock, received: %v",
		err,
	)

	defer db.Close()

	h := &handler.Handler{
		DB: db,
	}

	rows := sqlmock.NewRows([]string{
		"ID",
		"FirstName",
		"LastName",
		"Email",
		"PhoneNumber",
		"UserRole",
	}).AddRow(
		1,
		"Nunin",
		"Farid",
		"nunin@mail.com",
		"08123456789",
		"ADMIN",
	)

	mock.ExpectQuery(
		"SELECT (.+) FROM Users",
	).
		WithArgs(
			"nunin@mail.com",
			"nunin@mail.com",
			"rahasia123",
		).
		WillReturnRows(rows)

	user, err := h.Login(
		"nunin@mail.com",
		"rahasia123",
	)

	assert.NoError(
		t,
		err,
		"Expected no login error, received: %v",
		err,
	)

	assert.Equal(
		t,
		1,
		user.UserID,
		"Expected UserID = %d, received = %d",
		1,
		user.UserID,
	)

	assert.Equal(
		t,
		"Nunin",
		user.FirstName,
		"Expected FirstName = %s, received = %s",
		"Nunin",
		user.FirstName,
	)

	assert.Equal(
		t,
		"Farid",
		user.LastName,
		"Expected LastName = %s, received = %s",
		"Farid",
		user.LastName,
	)

	assert.Equal(
		t,
		"nunin@mail.com",
		user.Email,
		"Expected Email = %s, received = %s",
		"nunin@mail.com",
		user.Email,
	)

	assert.Equal(
		t,
		"08123456789",
		user.PhoneNumber,
		"Expected PhoneNumber = %s, received = %s",
		"08123456789",
		user.PhoneNumber,
	)

	assert.Equal(
		t,
		"ADMIN",
		user.UserRole,
		"Expected UserRole = %s, received = %s",
		"ADMIN",
		user.UserRole,
	)

	assert.NoError(
		t,
		mock.ExpectationsWereMet(),
		"There are unmet SQL mock expectations",
	)
}

func TestLoginUserNotFound(t *testing.T) {

	db, mock, err := sqlmock.New()

	assert.NoError(
		t,
		err,
		"Expected no error creating sqlmock, received: %v",
		err,
	)

	defer db.Close()

	h := &handler.Handler{
		DB: db,
	}

	rows := sqlmock.NewRows([]string{
		"ID",
		"FirstName",
		"LastName",
		"Email",
		"PhoneNumber",
		"UserRole",
	})

	mock.ExpectQuery(
		"SELECT (.+) FROM Users",
	).
		WithArgs(
			"test@mail.com",
			"test@mail.com",
			"wrongpassword",
		).
		WillReturnRows(rows)

	user, err := h.Login(
		"test@mail.com",
		"wrongpassword",
	)

	assert.NoError(
		t,
		err,
		"Expected no error when user not found, received: %v",
		err,
	)

	assert.Equal(
		t,
		handler.User{},
		user,
		"Expected empty user, received: %+v",
		user,
	)

	assert.NoError(
		t,
		mock.ExpectationsWereMet(),
		"There are unmet SQL mock expectations",
	)
}
