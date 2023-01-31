package models

import "time"

type Messages struct {
	Id        int64     `gorm:"primary_key;auto_increment" json:"id"`
	MessageID int64     `gorm:"type:bigint" json:"message_id"`
	UserID    string    `gorm:"type:text" json:"user_id"`
	Message   string    `gorm:"type:text" json:"message"`
	Reply     string    `gorm:"type:text" json:"reply"`
	CreatedAt time.Time `gorm:"type:datetime;default:current_timestamp" json:"created_at"`
}
