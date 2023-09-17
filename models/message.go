package models

import "time"

type Messages struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	MessageID int       `gorm:"type:int" json:"message_id"`
	UserID    string    `gorm:"type:varchar(300)" json:"user_id"`
	Message   string    `gorm:"type:varchar(300)" json:"message"`
	Reply     string    `gorm:"type:varchar(300)" json:"reply"`
	UserName  string    `gorm:"type:varchar(300)" json:"user_name"`
	CreatedAt string    `gorm:"column:created_at" json:"created_at"` // Perubahan di sini
}
func (m *Messages) AfterFind(tx *gorm.DB) error {
	// Konversi string ke time.Time
	parseTime, err := time.Parse("2006-01-02 15:04:05", m.CreatedAt)
	if err != nil {
		return err
	}
	m.CreatedAt = parseTime.Format(time.RFC3339)
	return nil
} 
