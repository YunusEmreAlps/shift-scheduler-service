package models

import (
	"time"

	"gorm.io/gorm"
)

type ShiftSchedule struct {
	gorm.Model             // ID, CreatedAt, UpdatedAt, DeletedAt fields will be added to the model automatically
	Alias        string    `json:"alias" gorm:"not null;"`
	Description  string    `json:"description" gorm:"default:null"`
	Frequency    int       `json:"frequency" gorm:"not null; default:1"` // 1, 2, 3, 4, 5, 6, 7 (days of the week)
	Start_Date   time.Time `json:"start_date" gorm:"not null;"`
	End_Date     time.Time `json:"end_date" gorm:"not null;"`
	Year         int       `json:"year" gorm:"not null;"`
	Status       int       `json:"status" gorm:"not null; default:0"` // 0: pending, 1: approved, 2: rejected
	Organization JSONB     `json:"organization" gorm:"type:jsonb;not null"`
	Manager      JSONB     `json:"manager" gorm:"type:jsonb;not null"`
	Shifts       JSONB     `json:"shifts" gorm:"type:jsonb;"`
}

// TableName overrides the table name used by User to `users`
func (u ShiftSchedule) TableName() string {
	return "shift_schedule"
}
