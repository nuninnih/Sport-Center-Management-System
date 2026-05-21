package helper

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func ValidateStringLength(input string) error {
	if len(input) < 1 || len(input) > 100 {
		return errors.New("input must be at least 1 until 100 characters long\n")
	}
	return nil
}

func GetIntegerInput(numb string) int {
	numbInt, _ := strconv.Atoi(numb)
	return numbInt
}

func IsDateAfterToday(input string) error {
	inputDate, err := time.ParseInLocation("2006-01-02", input, time.Local)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return err
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	if inputDate.Before(today) {
		return errors.New("Minimal input is today")
	}
	return nil
}
