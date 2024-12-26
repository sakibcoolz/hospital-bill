package utils

import (
	"fmt"
	"time"
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
