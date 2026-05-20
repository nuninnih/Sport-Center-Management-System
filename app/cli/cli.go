package cli

import (
	"errors"
	"fmt"

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
	validateStringLength := func(input string) error {
		if len(input) < 1 || len(input) > 100 {
			return errors.New("input must be at least 1 until 100 characters long\n")
		}
		return nil
	}
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
				Validate: validateStringLength,
			}
			FirstName, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Last Name",
				Validate: validateStringLength,
			}
			LastName, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Email",
				Validate: validateStringLength,
			}
			Email, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Phone Number",
				Validate: validateStringLength,
			}
			PhoneNumber, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please enter Password",
				Validate: validateStringLength,
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
		case 3:
			fmt.Println("Exiting the application. Goodbye!")
			return
		}
	}
}
