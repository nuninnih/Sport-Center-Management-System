package cli

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func MenuCustomer(c *CLI) {
	for {
		promptCustomer := promptui.Select{
			Label: "What do you want to do?",
			Items: []string{"1. Check Available Field", "2. Book", "3. Logout"},
		}

		index, _, err := promptCustomer.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			continue
		}

		switch index + 1 {
		case 1:
			prompt := promptui.Prompt{
				Label:    "Please Enter The City",
				Validate: ValidateStringLength,
			}
			City, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please Enter Field Type",
				Validate: ValidateStringLength,
			}
			Type, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please Enter Date (YYYY-MM-DD)",
				Validate: ValidateStringLength,
			}
			Date, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please Enter Start Time (hh:mm)",
				Validate: ValidateStringLength,
			}
			StartTime, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please Enter End Time (hh:mm)",
				Validate: ValidateStringLength,
			}
			EndTime, _ := prompt.Run()

			fields, err := c.Handler.GetAvailableField(City, Type, Date, StartTime, EndTime)
			if err != nil {
				fmt.Println("Error getting fields:", err)
				return
			}
			if len(fields) != 0 {
				FieldTable(fields)
			} else {
				fmt.Println("Field not found")
			}
		case 2:
			fmt.Println("Not Available Yet")
		case 3:
			fmt.Println("Logout..")
			return
		}
	}
}
