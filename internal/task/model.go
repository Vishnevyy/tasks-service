package task

type Task struct {
	ID     uint32 `gorm:"primaryKey"`
	UserID uint32 `gorm:"index;not null"`
	Title  string `gorm:"size:255;not null"`
	IsDone bool   `gorm:"not null;default:false"`
}
