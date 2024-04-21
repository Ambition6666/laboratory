package model

import (
	"gorm.io/gorm"
)

// 学生预约实验室
type Appointment struct {
	LID            uint        `json:"lid" gorm:"primaryKey;autoIncrement:false"`
	LaboratoryInfo *Laboratory `json:"laboratory_info" gorm:"-"`
	SID            uint        `json:"sid" gorm:"primaryKey;autoIncrement:false"`
	StudentInfo    *Student    `json:"student_info" gorm:"-"`
	ProgramContent string      `json:"program_content"`
	Raa            StringArr   `json:"araa"` // 可预约时段(Reservations are available)
}

func NewAppointment(sid uint, lid uint, pc string, raa []string) *Appointment {
	return &Appointment{
		LID:            lid,
		SID:            sid,
		ProgramContent: pc,
		Raa:            raa,
	}
}

// 钩子函数查询信息
func (a *Appointment) AfterFind(tx *gorm.DB) (err error) {
	if a.LaboratoryInfo == nil {
		tx.Where("id = ?", a.LID).Find(&a.LaboratoryInfo)
	}
	if a.StudentInfo == nil {
		tx.Where("id = ?", a.SID).Find(&a.StudentInfo)
	}
	return
}
