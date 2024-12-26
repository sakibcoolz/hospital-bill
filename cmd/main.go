package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"hospital-bill/model"
	"hospital-bill/utils"
	"os"
)

func main() {
	// Define the flag for the JSON file path
	jsonFilePath := flag.String("file", "", "Path to the JSON file")
	flag.Parse()

	// Check if the flag is provided
	if *jsonFilePath == "" {
		fmt.Println("Usage: go run main.go -file=<path-to-json-file>")
		return
	}

	// Read the JSON file
	data, err := os.ReadFile(*jsonFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the JSON file into a slice of BillMaster structs
	var bills []model.BillMaster
	err = json.Unmarshal(data, &bills)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Create a nested map[OrderDate]map[Service][]BillMaster
	orderDateMap := make(map[string]map[string][]model.BillMaster)

	// Populate the nested map
	for _, bill := range bills {
		formattedDate := utils.DateFormatter(bill.OrderDate)
		// Check if OrderDate exists in the outer map
		if _, exists := orderDateMap[formattedDate]; !exists {
			orderDateMap[formattedDate] = make(map[string][]model.BillMaster)
		}
		// Append the BillMaster object to the slice under the inner map's Service key
		orderDateMap[formattedDate][bill.Service] = append(orderDateMap[formattedDate][bill.Service], bill)
	}

	// Print the nested map
	fmt.Println("Nested Map (OrderDate -> Service -> []BillMaster):")
	for orderDate, serviceMap := range orderDateMap {
		fmt.Printf("OrderDate: %s\n", orderDate)
		for service, bills := range serviceMap {
			fmt.Printf("  Service: %s\n", service)
			for i, bill := range bills {
				fmt.Printf("    Bill %d:\n", i+1)
				fmt.Printf("      Patient Name: %s\n", bill.PatientName)
				fmt.Printf("      Doctor Name: %s\n", bill.DoctorName)
				fmt.Printf("      Service Amount: %s\n", bill.ServiceAmount)
				fmt.Printf("      Payor Name: %s\n", bill.PayorName)
			}
		}
		fmt.Println()
	}

}
