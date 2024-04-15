package request

type LaboratoryInfo struct {
	Date string `json:"date"`
	Place string `json:"place"`
	Raa []string `json:"raa"`
}

type BookingInfo struct {
	LID uint `json:"lid"`
	ProgramContent string `json:"program_content"`
	Raa []string `json:"raa"`
}


