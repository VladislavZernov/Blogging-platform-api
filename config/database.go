package config

import (
	"fmt"
	"github.com/VladislavZernov/blogging-platform-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=blog_user password=password dbname=blogging_platform port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Ошибка подключения к БД:", err)
	}

	// Выполняем миграции
	fmt.Println("🔄 Выполняем миграции...")
	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatal("❌ Ошибка при миграции:", err)
	}
	fmt.Println("✅ Миграции выполнены!")

	DB = db
	fmt.Println("✅ Успешное подключение к PostgreSQL!")
}
