package cli

import (
	"fmt"
	"os"
	"strconv"

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
