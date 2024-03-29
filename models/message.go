package models

import "time" 

type Messages struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	MessageID int       `gorm:"type:int" json:"message_id"`
	UserID    string    `gorm:"type:varchar(300)" json:"user_id"`
	Message   string    `gorm:"type:varchar(300)" json:"message"`
	Reply     string    `gorm:"type:varchar(300)" json:"reply"`
	UserName  string    `gorm:"type:varchar(300)" json:"user_name"`
	CreatedAt time.Time `gorm:"type:datetime;default:currentTimestamp" json:"created_at"`
        UserPhoto string    `gorm:"type:varchar(500)" json:"user_photo"`
}
