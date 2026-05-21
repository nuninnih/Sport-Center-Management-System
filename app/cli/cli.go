package cli

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/nuninnih/Sport-Center-Management-System/handler"
	"github.com/nuninnih/Sport-Center-Management-System/helper"
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
				Validate: helper.ValidateStringLength,
			}
			FirstName, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Last Name",
				Validate: helper.ValidateStringLength,
			}
			LastName, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Email",
				Validate: helper.ValidateStringLength,
			}
			Email, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Phone Number",
				Validate: helper.ValidateStringLength,
			}
			PhoneNumber, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Password",
				Validate: helper.ValidateStringLength,
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
				Validate: helper.ValidateStringLength,
			}
			Account, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Password",
				Validate: helper.ValidateStringLength,
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
					MenuAdmin(c)
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
