package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nuninnih/Sport-Center-Management-System/app/cli"
)

func main() {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		fmt.Println("DSN environment variable not set")
		return
	}

	database, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	if err = database.Ping(); err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer database.Close()

	c := cli.NewCLI()
	c.Run()
}
