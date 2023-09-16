package models

import "time"

type Messages struct {
	Id        int64     `gorm:"primary_key;auto_increment" json:"id"`
	MessageID int       `gorm:"type:int" json:"message_id"`
	UserID    string    `gorm:"type:varchar(300)" json:"user_id"`
	Message   string    `gorm:"type:varchar(300)" json:"message"`
	Reply     string    `gorm:"type:varchar(300)" json:"reply"`
	UserName  string    `gorm:"type:varchar(300)" json:"user_name"`
	CreatedAt time.Time `gorm:"type:datetime;default:current_timestamp" json:"created_at"`
}
