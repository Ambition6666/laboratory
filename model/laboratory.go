package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Laboratory struct {
	gorm.Model 
	TID        uint      `json:"-"`
	TName      string    `json:"name" gorm:"-"`
	TPhone     string    `json:"phone" gorm:"-"`
	Date       string    `json:"date" gorm:"index"`
	Place      string    `json:"place"` // 地点
	Raa        StringArr `json:"raa"`   // 可预约时段(Reservations are available)
}

func NewLaboratory(date string, place string, raa []string, tid uint) *Laboratory {
	return &Laboratory{
		Date:  date,
		Place: place,
		Raa:   raa,
		TID:   tid,
	}
}

type StringArr []string

// 实现 driver.Valuer 接口
func (s StringArr) Value() (driver.Value, error) {
	if s == nil {
		return "[]", nil
	}
	return json.Marshal(s)
}

// 实现 sql.Scanner 接口
func (s *StringArr) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan Array value:", value))
	}
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, s)
	}
	*s = make([]string, 0)
	return nil
}

// 钩子函数查询教师信息
func (l *Laboratory) AfterFind(tx *gorm.DB) (err error) {
	t := new(Teacher)
	tx.Where("id = ?", l.TID).Find(&t)
	l.TName = t.UINFO.Name
	l.TPhone = t.UINFO.Phone
	return
}
