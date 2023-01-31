package models

import "time"

type Users struct {
	Id        int64     `gorm:"primarykey" json:"id"`
	UserID    string    `gorm:"type:text" json:"user_id"`
	UserName  string    `gorm:"type:varchar(300)" json:"user_name"`
	FirstName string    `gorm:"type:varchar(300)" json:"user_name"`
	LastName  string    `gorm:"type:varchar(300)" json:"user_name"`
	StartDate time.Time `gorm:"type:datetime;default:current_timestamp" json:"start_date"`
}
