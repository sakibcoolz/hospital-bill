package model

// Define the BillMaster struct
type BillMaster struct {
	SNo                        string `gorm:"column:s_no;primaryKey" json:"S.No" header:"S.No"`
	EpisodeType                string `gorm:"column:episode_type" json:"Episode Type" header:"Episode Type"`
	EpisodeNo                  string `gorm:"column:episode_no" json:"Episode No" header:"Episode No"`
	PatientName                string `gorm:"column:patient_name" json:"Patient Name" header:"Patient Name"`
	DoctorName                 string `gorm:"column:doctor_name" json:"Doctor Name" header:"Doctor Name"`
	AdmitDoctorName            string `gorm:"column:admit_doctor_name" json:"Admit Doctor Name" header:"Admit Doctor Name"`
	OrderNo                    string `gorm:"column:order_no" json:"Order No" header:"Order No"`
	OrderDate                  string `gorm:"column:order_date" json:"Order Date" header:"Order Date"`
	OrderBy                    string `gorm:"column:order_by" json:"Order By" header:"Order By"`
	OrderLocation              string `gorm:"column:order_location" json:"Order Location" header:"Order Location"`
	DischargeDate              string `gorm:"column:discharge_date" json:"Discharge Date" header:"Discharge Date"`
	PayorType                  string `gorm:"column:payor_type" json:"PayorType" header:"PayorType"`
	PayorName                  string `gorm:"column:payor_name" json:"Payor Name" header:"Payor Name"`
	Item                       string `gorm:"column:item" json:"Item" header:"Item"`
	OracleCode                 string `gorm:"column:oracle_code" json:"Oracle Code" header:"Oracle Code"`
	Service                    string `gorm:"column:service" json:"Service" header:"Service"`
	Quantity                   string `gorm:"column:quantity" json:"Quantity" header:"Quantity"`
	ServiceAmount              string `gorm:"column:service_amount" json:"Service Amount" header:"Service Amount"`
	PatientShareAmount         string `gorm:"column:patient_share_amount" json:"Patient Share Amount" header:"Patient Share Amount"`
	PayorShareAmount           string `gorm:"column:payor_share_amount" json:"Payor Share Amount" header:"Payor Share Amount"`
	PatientDiscountAmount      string `gorm:"column:patient_discount_amount" json:"Patient Discount Amount" header:"Patient Discount Amount"`
	PayorDiscountAmount        string `gorm:"column:payor_discount_amount" json:"Payor Discount Amount" header:"Payor Discount Amount"`
	DiscretionaryFixedDiscount string `gorm:"column:discretionary_fixed_discount" json:"Discretionary Fixed Discount" header:"Discretionary Fixed Discount"`
	MarkupAmount               string `gorm:"column:markup_amount" json:"Markup Amount" header:"Markup Amount"`
	Package                    string `gorm:"column:package" json:"PACKAGE" header:"PACKAGE"`
	Charged                    string `gorm:"column:charged" json:"Charged" header:"Charged"`
}

type PatientShareSummary struct {
	PatientName string  `header:"patient_name" json:"patient_name"`
	Total       float64 `header:"total" json:"total"`
}

type Person struct {
	Firstname string  `header:"first name"`
	Lastname  string  `header:"last name"`
	Total     float64 `header:"total"`
}

type PerticularWise struct {
	Datewise   string  `json:"datewise" header:"Datewise"`
	Service    string  `json:"service" header:"Service"`
	Amount     float64 `json:"amount" header:"Amount"`
	DailyTotal float64 `json:"daily_total" header:"DailyTotal"`
}

type DoctorPerticularWise struct {
	Datewise   string  `json:"datewise" header:"Datewise"`
	Service    string  `json:"service" header:"Service"`
	Amount     float64 `json:"amount" header:"Amount"`
	DoctorName string  `json:"doctor_name" header:"Doctor Name"`
	DailyTotal float64 `json:"daily_total" header:"DailyTotal"`
}

type Summary struct {
	Service string  `json:"service" header:"Service"`
	Amount  float64 `json:"amount" header:"Amount"`
}

type ItemWiseSpecification struct {
	Datewise   string  `json:"datewise" header:"Datewise"`
	Service    string  `json:"service" header:"Service"`
	Amount     float64 `json:"amount" header:"Amount"`
	OracleCode string  `json:"oracle_code" header:"Code"`
	Item       string  `json:"item" header:"Items"`
	Quantity   float64 `json:"quantity" header:"Quantity"`
}
