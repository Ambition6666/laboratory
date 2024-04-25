package model

import "gorm.io/gorm"

type TAppointment struct {
	gorm.Model
	TID                      uint   `json:"tid" gorm:"index"`
	TName                    string `json:"t_name"`
	TPhone                   string `json:"t_phone"`
	CourseName               string `json:"course_name"`
	CourseNo                 string `json:"course_no"`
	CourseTime               int    `json:"course_time"`
	ExperimentNum            int    `json:"experiment_num"`
	ExperimentKind           string `json:"experiment_kind"`
	ExperimentKind2          string `json:"experiment_kind2"`
	OutlineRequirements      string `json:"outline_requirements"`
	Campus                   string `json:"campus"`
	ExperimentHU             string `json:"experiment_hu"`
	Class                    string `json:"class"`
	ExperimentHUNum          int    `json:"experiment_hu_num"`
	Raa                      string `json:"raa"`
	WeekTime                 int    `json:"week_time"`
	WeekCourseTime           int    `json:"week_course_time"`
	EnvironmentalRequirement string `json:"environmental_requirement"`
	ProgramContent           string `json:"program_content"`
	Reason                   string `json:"reason"`
}

func (t *TAppointment) BeforeCreate(tx *gorm.DB) (err error) {
	te := new(Teacher)
	err = tx.Where("id = ?", t.TID).Find(te).Error
	if err != nil {
		return
	}

	t.TName = te.UINFO.Name
	t.TPhone = te.UINFO.Phone
	return
}
