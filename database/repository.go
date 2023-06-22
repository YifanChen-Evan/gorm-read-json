package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type EmployeeRepository interface {
	LoadDataFromJSON(filepath string) error

	// haven't implement database operations part
	AddEmployee(employee *Employee) error
	GetEmployeeByName(name string) (*Employee, error)
	UpdateEmployeeByName(employee *Employee) error
	DeleteEmployeeByName(name string) error
}

// create a function load data from local JSON file and store data in database
func LoadDataFromJSON(filePath string) error {
	// -------------read local JSON file-------------
	jsonData, err := ioutil.ReadFile(filePath) // read local JSON file
	if err != nil {                            // fail
		return fmt.Errorf("failed to read JSON file: %w", err)
	}

	// -------------parse local JSON file-------------
	var employees []Employee                   // define an empty slice to store parsed JSON data
	err = json.Unmarshal(jsonData, &employees) // parse "jsonData" and store to "employees" slice
	if err != nil {                            // fail
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// -------------store initial data in database-------------
	// using "Create()" method of Gorm to populate into database
	err = DB.Create(&employees).Error // pass a slice to insert
	if err != nil {                   // fail
		return fmt.Errorf("failed to insert data in database: %w", err)
	}

	return nil // success
}
