package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Vishnevyy/tasks-service/internal/task"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to init DB: %v", err)
	}

	// создаём таблицу tasks, если её нет
	if err := DB.AutoMigrate(&task.Task{}); err != nil {
		log.Fatalf("auto-migrate failed: %v", err)
	}
}
