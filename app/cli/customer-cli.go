package cli

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/nuninnih/Sport-Center-Management-System/handler"
	"github.com/nuninnih/Sport-Center-Management-System/helper"
)

func MenuCustomer(c *CLI, user handler.User) {
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
			allCity, err := c.Handler.GetAllCities()
			CityTable(allCity)

			prompt := promptui.Prompt{
				Label:    "Please Enter The City",
				Validate: helper.ValidateStringLength,
			}
			City, _ := prompt.Run()

			allType, err := c.Handler.GetAllTypes()
			TypeTable(allType)

			prompt = promptui.Prompt{
				Label:    "Please Enter Field Type",
				Validate: helper.ValidateStringLength,
			}
			Type, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please Enter Date (YYYY-MM-DD)",
				Validate: helper.ValidateStringLength,
			}
			Date, _ := prompt.Run()

			fields, err := c.Handler.CheckAvailableSlot(Date, City, Type)
			if err != nil {
				fmt.Println("Error getting fields:", err)
				return
			}
			if len(fields) != 0 {
				RenderAvailability(fields, Date, City, Type)
			} else {
				fmt.Printf("There is no available %v field in %v \n", Type, City)
			}
		case 2:
			allField, err := c.Handler.GetAllField()
			FieldTable(allField)

			prompt := promptui.Prompt{
				Label:    "Please Enter Field ID",
				Validate: helper.ValidateStringLength,
			}
			field, _ := prompt.Run()
			FieldID := helper.GetIntegerInput(field)

			prompt = promptui.Prompt{
				Label:    "Please Enter Date (YYYY-MM-DD)",
				Validate: helper.ValidateStringLength,
			}
			Date, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please Enter Start Time (hh:mm)",
				Validate: helper.ValidateStringLength,
			}
			StartTime, _ := prompt.Run()

			prompt = promptui.Prompt{
				Label:    "Please Enter End Time (hh:mm)",
				Validate: helper.ValidateStringLength,
			}
			EndTime, _ := prompt.Run()

			err = c.Handler.CreateBooking(user.UserID, FieldID, Date, StartTime, EndTime)

			if err != nil {
				fmt.Println("Error create booking", err)
				return
			}

			fmt.Println("Booking created successfully")
		case 3:
			fmt.Println("Logout..")
			return
		}
	}
}
