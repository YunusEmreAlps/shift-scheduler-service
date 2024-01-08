package models

import (
	"time"

	"gorm.io/gorm"
)

type ShiftSchedule struct {
	gorm.Model             // ID, CreatedAt, UpdatedAt, DeletedAt fields will be added to the model automatically
	Alias        string    `json:"alias" gorm:"not null;"`
	Organization JSONB     `json:"organization" gorm:"type:jsonb;not null"`
	Manager      JSONB     `json:"manager" gorm:"type:jsonb;not null"`
	Description  string    `json:"description" gorm:"default:null"`
	Start_Date   time.Time `json:"start_date" gorm:"not null;"`
	End_Date     time.Time `json:"end_date" gorm:"not null;"`
	Shifts       JSONB     `json:"shifts" gorm:"type:jsonb;"`
	Year         int       `json:"year" gorm:"not null;"`
	Status       int       `json:"status" gorm:"not null; default:0"` // 0: pending, 1: approved, 2: rejected
}

// TableName overrides the table name used by User to `users`
func (u ShiftSchedule) TableName() string {
	return "shift_schedule"
}
