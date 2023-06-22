package main

import (
	"fmt"

	"github.com/YifanChen-Evan/gorm-read-json/database"
)

func main() {
	// connect database
	err := database.ConnectDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to the database")

	// create table automatically by using Grom
	err = database.DB.AutoMigrate(&database.Employee{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Table created successfully")

	// read local JSON file and store data in database
	err = database.LoadDataFromJSON("../data/simple.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data inserted into the database")

	fmt.Println("Hello")
}
