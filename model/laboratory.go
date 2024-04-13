package model

type Laboratory struct {
	Teacher
	Place string `json:"place"` // 地点
	Raa []string `json:"raa"` // 可预约时段(Reservations are available) 
}

