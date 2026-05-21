package cli

import (
	"bufio"
	"fmt"
	"os"
)

type CLI struct{}

// create new cli
func NewCLI() *CLI {
	return &CLI{}
}

// run application
func (c *CLI) Run() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== SPORT CENTER ===")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Exit")

	fmt.Print("Choose menu: ")

	input, _ := reader.ReadString('\n')

	fmt.Println("Your choice:", input)

	switch input {

	case "1\n":
		// login process
		fmt.Println("Login Success")

		// menu after login
		fmt.Println("=== MAIN MENU ===")
		fmt.Println("1. Booking Field")
		fmt.Println("2. View Available Fields")
		fmt.Println("3. Report")
		fmt.Println("4. Payment")
		fmt.Println("5. Logout")

		fmt.Print("Choose menu: ")

		input, _ = reader.ReadString('\n')

		fmt.Println("Your choice:", input)

		switch input {
		case "1\n":
			fmt.Println("Booking Field")
		case "2\n":
			fmt.Println("View Available Fields")
		case "3\n":
			fmt.Println("Report")
		case "4\n":
			fmt.Println("Payment")
		case "5\n":
			fmt.Println("Logout Success")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	case "2\n":
		fmt.Println("Register Success")
	case "3\n":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}