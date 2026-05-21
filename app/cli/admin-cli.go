package cli

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/nuninnih/Sport-Center-Management-System/helper"
)

func MenuAdmin(c *CLI) {
	for {
		promptCustomer := promptui.Select{
			Label: "What do you want to do?",
			Items: []string{"1. Check Pending Booking", "2. Update Booking Status", "3. Create Payment", "4. Update Status Field", "5. Logout"},
		}

		index, _, err := promptCustomer.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			continue
		}

		switch index + 1 {
		case 1:
			pendingBooking, err := c.Handler.CheckPendingBooking()
			if err != nil {
				fmt.Println("Error check pending booking, ", err)
			}
			BookingTable(pendingBooking)
		case 2:
			prompt := promptui.Prompt{
				Label:    "Please Enter Booking ID",
				Validate: helper.ValidateStringLength,
			}
			bookId, _ := prompt.Run()
			BookingID := helper.GetIntegerInput(bookId)

			prompt = promptui.Prompt{
				Label:    "Please Update Status (CONFIRM / CANCEL / COMPLETE)",
				Validate: helper.ValidateStringLength,
			}
			status, _ := prompt.Run()

			err = c.Handler.UpdateBookingStatus(BookingID, status)
			if err != nil {
				fmt.Println("Error update booking", err)
				return
			}

			fmt.Println("Status Updated successfully")
		case 4:
			data, err := c.Handler.GetAllFieldWithStatus()
			if err != nil {
				fmt.Println("Error check status field, ", err)
			}
			FieldStatusTable(data)

			prompt := promptui.Prompt{
				Label:    "Please Enter Field ID to update active status",
				Validate: helper.ValidateStringLength,
			}
			field, _ := prompt.Run()
			FieldID := helper.GetIntegerInput(field)

			err = c.Handler.UpdateStatusField(FieldID)
			if err != nil {
				fmt.Println("Error update status", err)
				return
			}

			fmt.Println("Field Status Updated successfully")

		case 5:
			fmt.Println("Logout..")
			return
		}
	}
}
