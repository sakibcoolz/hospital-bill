package model

// Define the BillMaster struct
type BillMaster struct {
	SNo                        string `json:"S.No"`
	EpisodeType                string `json:"Episode Type"`
	EpisodeNo                  string `json:"Episode No"`
	PatientName                string `json:"Patient Name"`
	DoctorName                 string `json:"Doctor Name"`
	AdmitDoctorName            string `json:"Admit Doctor Name"`
	OrderNo                    string `json:"Order No"`
	OrderDate                  string `json:"Order Date"`
	OrderBy                    string `json:"Order By"`
	OrderLocation              string `json:"Order Location"`
	DischargeDate              string `json:"Discharge Date"`
	PayorType                  string `json:"PayorType"`
	PayorName                  string `json:"Payor Name"`
	Item                       string `json:"Item"`
	OracleCode                 string `json:"Oracle Code"`
	Service                    string `json:"Service"`
	Quantity                   string `json:"Quantity"`
	ServiceAmount              string `json:"Service Amount"`
	PatientShareAmount         string `json:"Patient Share Amount"`
	PayorShareAmount           string `json:"Payor Share Amount"`
	PatientDiscountAmount      string `json:"Patient Discount Amount"`
	PayorDiscountAmount        string `json:"Payor Discount Amount"`
	DiscretionaryFixedDiscount string `json:"Discretionary Fixed Discount"`
	MarkupAmount               string `json:"Markup Amount"`
	Package                    string `json:"PACKAGE"`
	Charged                    string `json:"Charged"`
}
