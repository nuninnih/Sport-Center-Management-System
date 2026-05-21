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
}