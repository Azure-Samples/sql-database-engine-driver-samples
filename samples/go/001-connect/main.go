package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/TwiN/go-color"
	"github.com/joho/godotenv"
	"github.com/microsoft/go-mssqldb"
)

func main() {
	godotenv.Load()

	credential := os.Getenv("SQL_CREDENTIAL")

	if credential == "" {
		panic("Please set the environment variable using an \".env\" file.")
	}

	fmt.Println(color.InBold(color.InBlue(fmt.Sprintf("%c\tConnecting to SQL Server...", '\u2699'))))

	connector, err := mssql.NewConnector(credential)
	if err != nil {
		panic(err)
	}
	db := sql.OpenDB(connector)
	defer db.Close()

	query := "SELECT @@VERSION"

	fmt.Println(color.InBlue(fmt.Sprintf("Query:\t%s", color.InUnderline(query))))

	var version string
	err = db.QueryRow(query).Scan(&version)
	if err != nil {
		panic(err)
	}

	if version != "" {
		fmt.Println(color.InGreen(fmt.Sprintf("Result:\t%s", version)))
	}
}