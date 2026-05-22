package cli

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nuninnih/Sport-Center-Management-System/handler"
	"github.com/olekukonko/tablewriter"
)

func FieldTable(data interface{}) {
	tableData := [][]string{
		{"FieldID", "FieldName", "TypeName", "CityName", "HourlyRate"},
	}

	switch v := data.(type) {

	case handler.Field:
		row := []string{
			strconv.Itoa(v.FieldID),
			v.FieldName,
			v.FieldType,
			v.City,
			strconv.FormatFloat(v.HourlyRate, 'f', 2, 64),
		}
		tableData = append(tableData, row)

	case []handler.Field:
		for _, field := range v {
			row := []string{
				strconv.Itoa(field.FieldID),
				field.FieldName,
				field.FieldType,
				field.City,
				strconv.FormatFloat(field.HourlyRate, 'f', 2, 64),
			}
			tableData = append(tableData, row)
		}

	default:
		fmt.Println("unsupported type")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(tableData[0])
	_ = table.Bulk(tableData[1:])
	_ = table.Render()
}

func FieldStatusTable(data interface{}) {
	tableData := [][]string{
		{"FieldID", "FieldName", "TypeName", "CityName", "HourlyRate", "Is Active"},
	}

	switch v := data.(type) {

	case handler.FieldStatus:
		row := []string{
			strconv.Itoa(v.FieldID),
			v.FieldName,
			v.FieldType,
			v.City,
			strconv.FormatFloat(v.HourlyRate, 'f', 2, 64),
			strconv.FormatBool(v.IsActive),
		}
		tableData = append(tableData, row)

	case []handler.FieldStatus:
		for _, field := range v {
			row := []string{
				strconv.Itoa(field.FieldID),
				field.FieldName,
				field.FieldType,
				field.City,
				strconv.FormatFloat(field.HourlyRate, 'f', 2, 64),
				strconv.FormatBool(field.IsActive),
			}
			tableData = append(tableData, row)
		}

	default:
		fmt.Println("unsupported type")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(tableData[0])
	_ = table.Bulk(tableData[1:])
	_ = table.Render()
}

func CityTable(data interface{}) {
	tableData := [][]string{
		{"ID", "City"},
	}

	switch v := data.(type) {

	case handler.City:
		row := []string{
			strconv.Itoa(v.CityID),
			v.CityName,
		}
		tableData = append(tableData, row)

	case []handler.City:
		for _, city := range v {
			row := []string{
				strconv.Itoa(city.CityID),
				city.CityName,
			}
			tableData = append(tableData, row)
		}

	default:
		fmt.Println("unsupported type")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(tableData[0])
	_ = table.Bulk(tableData[1:])
	_ = table.Render()
}

func TypeTable(data interface{}) {
	tableData := [][]string{
		{"ID", "Type"},
	}

	switch v := data.(type) {

	case handler.Type:
		row := []string{
			strconv.Itoa(v.TypeID),
			v.TypeName,
		}
		tableData = append(tableData, row)

	case []handler.Type:
		for _, tp := range v {
			row := []string{
				strconv.Itoa(tp.TypeID),
				tp.TypeName,
			}
			tableData = append(tableData, row)
		}

	default:
		fmt.Println("unsupported type")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(tableData[0])
	_ = table.Bulk(tableData[1:])
	_ = table.Render()
}

func BookingTable(data interface{}) {
	tableData := [][]string{
		{"BookingID", "Field", "User", "Phone No.", "Booking Date", "Start Time", "End Time"},
	}

	switch v := data.(type) {

	case handler.PendingBooking:
		row := []string{
			strconv.Itoa(v.BookingID),
			v.FieldName,
			v.UserName,
			v.PhoneNumber,
			v.BookingDate,
			v.StartTime,
			v.EndTime,
		}
		tableData = append(tableData, row)

	case []handler.PendingBooking:
		for _, booking := range v {
			row := []string{
				strconv.Itoa(booking.BookingID),
				booking.FieldName,
				booking.UserName,
				booking.PhoneNumber,
				booking.BookingDate,
				booking.StartTime,
				booking.EndTime,
			}
			tableData = append(tableData, row)
		}

	default:
		fmt.Println("unsupported type")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(tableData[0])
	_ = table.Bulk(tableData[1:])
	_ = table.Render()
}

func RenderAvailability(bookings []handler.CheckBooking, Date, City, Type string) {

	slots := []string{
		"08:00",
		"09:00",
		"10:00",
		"11:00",
		"12:00",
		"13:00",
		"14:00",
		"15:00",
		"16:00",
		"17:00",
		"18:00",
		"19:00",
		"20:00",
		"21:00",
		"22:00",
	}

	fieldBookings := map[string][]handler.CheckBooking{}

	fieldNames := []string{}
	fieldExists := map[string]bool{}

	for _, booking := range bookings {

		fieldBookings[booking.FieldName] =
			append(fieldBookings[booking.FieldName], booking)

		if !fieldExists[booking.FieldName] {
			fieldExists[booking.FieldName] = true
			fieldNames = append(fieldNames, booking.FieldName)
		}
	}

	fmt.Println("================================================================================================================")
	fmt.Printf(" %s - %s (%s)\n",
		strings.ToUpper(City),
		strings.ToUpper(Type),
		Date,
	)
	fmt.Println("================================================================================================================")

	fmt.Printf("%-25s", "Field Name")

	for _, slot := range slots {
		fmt.Printf("%-6s", slot)
	}

	fmt.Println()

	fmt.Println("----------------------------------------------------------------------------------------------------------------")

	for _, fieldName := range fieldNames {

		fmt.Printf("%-25s", fieldName)

		fieldBookingList := fieldBookings[fieldName]

		for _, slot := range slots {

			booked := false

			slotHour, err := strconv.Atoi(slot[:2])
			if err != nil {
				continue
			}

			for _, booking := range fieldBookingList {

				if booking.StartTime == "" || booking.EndTime == "" {
					continue
				}

				startHour, err := strconv.Atoi(booking.StartTime[:2])
				if err != nil {
					continue
				}

				endHour, err := strconv.Atoi(booking.EndTime[:2])
				if err != nil {
					continue
				}

				if slotHour >= startHour && slotHour < endHour {
					booked = true
					break
				}
			}

			if booked {
				fmt.Printf("%-5s", "🟥")
			} else {
				fmt.Printf("%-5s", "🟩")
			}
		}

		fmt.Println()
	}

	fmt.Println("================================================================================================================")
	fmt.Println("🟩 Available   🟥 Booked")
	fmt.Println("================================================================================================================")
}

func PaymentMethodTable() {
	fmt.Println("================================")
	fmt.Printf("%-15s", "CODE NUMBER")
	fmt.Printf("%-15s", "PAYMENT METHOD \n")
	fmt.Println("================================")
	fmt.Printf("%-15s", "1")
	fmt.Printf("%-15s", "CASH \n")
	fmt.Println()
	fmt.Printf("%-15s", "2")
	fmt.Printf("%-15s", "BANK_TRANSFER \n")
	fmt.Println()
	fmt.Printf("%-15s", "3")
	fmt.Printf("%-15s", "E_WALLET \n")
	fmt.Println()
	fmt.Printf("%-15s", "4")
	fmt.Printf("%-15s", "CREDIT_CARD \n")
	fmt.Println("================================")
}

func PendingPaymentTable(data interface{}) {
	tableData := [][]string{
		{"BookingID", "Field", "User", "Phone No.", "Booking Date", "Total Payment"},
	}

	switch v := data.(type) {

	case handler.PendingPayment:
		row := []string{
			strconv.Itoa(v.BookingID),
			v.FieldName,
			v.UserName,
			v.PhoneNumber,
			v.BookingDate,
			v.Total,
		}
		tableData = append(tableData, row)

	case []handler.PendingPayment:
		for _, booking := range v {
			row := []string{
				strconv.Itoa(booking.BookingID),
				booking.FieldName,
				booking.UserName,
				booking.PhoneNumber,
				booking.BookingDate,
				booking.Total,
			}
			tableData = append(tableData, row)
		}

	default:
		fmt.Println("unsupported type")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(tableData[0])
	_ = table.Bulk(tableData[1:])
	_ = table.Render()
}

func TableRevenueInAYear(data interface{}) {
	tableData := [][]string{
		{"Total City", "Total Field", "Total Transaction", "Total Revenue"},
	}

	switch v := data.(type) {

	case handler.RevenueInAYear:
		row := []string{
			strconv.Itoa(v.TotalCity),
			strconv.Itoa(v.TotalField),
			strconv.Itoa(v.TotalTransaction),
			strconv.FormatFloat(v.TotalRevenue, 'f', 2, 64),
		}
		tableData = append(tableData, row)

	case []handler.RevenueInAYear:
		for _, booking := range v {
			row := []string{
				strconv.Itoa(booking.TotalCity),
				strconv.Itoa(booking.TotalField),
				strconv.Itoa(booking.TotalTransaction),
				strconv.FormatFloat(booking.TotalRevenue, 'f', 2, 64),
			}
			tableData = append(tableData, row)
		}

	default:
		fmt.Println("unsupported type")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(tableData[0])
	_ = table.Bulk(tableData[1:])
	_ = table.Render()
}

func TableRevenuePerCity(data interface{}) {
	tableData := [][]string{
		{"City Name", "Total Field", "Total Transaction", "Total Revenue"},
	}

	switch v := data.(type) {

	case handler.RevenuePerCity:
		row := []string{
			v.CityName,
			strconv.Itoa(v.TotalField),
			strconv.Itoa(v.TotalTransaction),
			strconv.FormatFloat(v.TotalRevenue, 'f', 2, 64),
		}
		tableData = append(tableData, row)

	case []handler.RevenuePerCity:
		for _, booking := range v {
			row := []string{
				booking.CityName,
				strconv.Itoa(booking.TotalField),
				strconv.Itoa(booking.TotalTransaction),
				strconv.FormatFloat(booking.TotalRevenue, 'f', 2, 64),
			}
			tableData = append(tableData, row)
		}

	default:
		fmt.Println("unsupported type")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header(tableData[0])
	_ = table.Bulk(tableData[1:])
	_ = table.Render()
}

func ReportListTable() {
	fmt.Println("================================")
	fmt.Printf("%-15s", "NO REPORT")
	fmt.Printf("%-25s", "REPORT \n")
	fmt.Println()
	fmt.Printf("%-15s", "1")
	fmt.Printf("%-25s", "REVENUE THIS YEAR \n")
	fmt.Println()
	fmt.Printf("%-15s", "2")
	fmt.Printf("%-25s", "REVENUE PER CITY \n")
	fmt.Println()
	fmt.Printf("%-15s", "3")
	fmt.Printf("%-25s", "-- \n")
	fmt.Println()
	fmt.Printf("%-25s", "4")
	fmt.Printf("%-15s", "-- \n")
	fmt.Println("================================")
}
