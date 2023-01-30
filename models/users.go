package models

type Users struct {
	Id       int64  `gorm:"primarykey" json:"id"`
	UserID   string `gorm:"type:text" json:"user_id"`
	UserName string `gorm:"type:varchar(300)" json:"user_name"`
}
