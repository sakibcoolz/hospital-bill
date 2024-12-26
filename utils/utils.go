package utils

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func DateFormatter(input string) string {
	layout := "01/02/2006 03:04 PM"
	t, err := time.Parse(layout, input)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	// Format the date to "MM/DD/YYYY"
	return t.Format("01/02/2006")
}

func GetColumnNames(db *gorm.DB, model interface{}) ([]string, error) {
	// Parse the model's schema
	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(model); err != nil {
		return nil, err
	}

	// Collect column names
	var columnNames []string
	for _, field := range stmt.Schema.Fields {
		columnNames = append(columnNames, field.DBName) // `DBName` contains the column name
	}

	return columnNames, nil
}
