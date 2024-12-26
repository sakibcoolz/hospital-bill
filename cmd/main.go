package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"hospital-bill/constant"
	"hospital-bill/db"
	"hospital-bill/model"
	"os"

	"github.com/kataras/tablewriter"
	"github.com/lensesio/tableprinter"
)

func MakeSpace(str string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("\n")
	}
	fmt.Println("**************************", str, "************************************")

}

func main() {
	db := db.DBConnect()

	printer := tableprinter.New(os.Stdout)

	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor

	db.AutoMigrate(&model.BillMaster{})

	// Define the flag for the JSON file path
	jsonFilePath := flag.String("file", "data/bill.json", "Path to the JSON file")
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

	db.Create(bills)

	err = func() error {
		var results []model.PatientShareSummary
		db.Begin()
		err = db.Model(&model.BillMaster{}). // Point to the model
							Select("patient_name, SUM(patient_share_amount::float) as total").
							Group("patient_name").
							Scan(&results).Error
		if err != nil {
			return err
		}
		fmt.Println("All Amount")
		MakeSpace("PatientShareSummary")
		printer.Print(results)
		fmt.Println()

		return err
	}()

	func() {
		var results []model.PerticularWise
		db.Begin()
		db.Raw(constant.ServiceDateWise).Find(&results)
		MakeSpace("PatientShareSummary")
		fmt.Println("Date wise summary")
		printer.Print(results)
		fmt.Println("Date wise summary End")
	}()

	func() {
		var results []model.DoctorPerticularWise
		db.Begin()
		db.Raw(constant.DoctorServiceSpecific).Find(&results)
		MakeSpace("DoctorPerticularWise")
		fmt.Println("DoctorPerticularWise wise summary")
		printer.Print(results)
		fmt.Println("DoctorPerticularWise wise summary End")
	}()

	func() {
		var results []model.Summary
		db.Begin()
		db.Raw(constant.Summary).Find(&results)
		MakeSpace("Summary")
		fmt.Println("Summary wise summary")
		printer.Print(results)
		fmt.Println("Summary wise summary End")
	}()

	func() {
		var results []model.ItemWiseSpecification
		db.Begin()
		db.Raw(constant.ItemWiseSpecification).Find(&results)
		MakeSpace("ItemWiseSpecification")
		fmt.Println("ItemWiseSpecification wise summary")
		printer.Print(results)
		fmt.Println("ItemWiseSpecification wise summary End")
	}()

}
