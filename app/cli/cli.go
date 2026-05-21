package cli

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/nuninnih/Sport-Center-Management-System/handler"
)

type CLI struct {
	Handler *handler.Handler
}

func NewCLI(handler *handler.Handler) *CLI {
	return &CLI{Handler: handler}
}

func (c *CLI) Run() {
	for {
		prompt := promptui.Select{
			Label: "Select Menu",
			Items: []string{"1. Register", "2. Login", "3. Exit"},
		}

		idx, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			continue
		}

		switch idx + 1 {
		case 1:
			prompt := promptui.Prompt{
				Label:    "Please enter First Name",
				Validate: ValidateStringLength,
			}
			FirstName, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Last Name",
				Validate: ValidateStringLength,
			}
			LastName, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Email",
				Validate: ValidateStringLength,
			}
			Email, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Phone Number",
				Validate: ValidateStringLength,
			}
			PhoneNumber, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Password",
				Validate: ValidateStringLength,
			}
			Password, _ := prompt.Run()

			if FirstName == "" || LastName == "" || Email == "" || PhoneNumber == "" || Password == "" {
				fmt.Println("Invalid input, input cannot be empty")
				continue
			}

			err = c.Handler.Register(FirstName, LastName, Email, PhoneNumber, Password)
			if err != nil {
				fmt.Println("Error Registering New User:", err)
				return
			}
			fmt.Println("Success Registering New User")
		case 2:
			prompt := promptui.Prompt{
				Label:    "Please enter Email or PhoneNumber",
				Validate: ValidateStringLength,
			}
			Account, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Password",
				Validate: ValidateStringLength,
			}
			Password, _ := prompt.Run()

			user, err := c.Handler.Login(Account, Password)
			if err != nil {
				fmt.Println("Login Error:", err)
				return
			}
			if user.UserID != 0 {
				if user.UserRole == "CUSTOMER" {
					// CUSTOMER
					MenuCustomer(c, user)
				} else {
					// ADMIN
				}
			} else {
				fmt.Println("Login Failed!")
			}
		case 3:
			fmt.Println("Exiting the application. Goodbye!")
			return
		}
	}
}

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
