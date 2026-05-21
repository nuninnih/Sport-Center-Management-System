package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nuninnih/Sport-Center-Management-System/app/cli"
	"github.com/nuninnih/Sport-Center-Management-System/db"
	"github.com/nuninnih/Sport-Center-Management-System/handler"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.Close()

	handler := handler.NewHandler(db)
	cli := cli.NewCLI(handler)
	cli.Run()
}
