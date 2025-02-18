package models

import "gorm.io/gorm"

type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text;not null"`
	UserID  uint   `gorm:"not null"`
	PostID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID"`
	Post    Post   `gorm:"foreignKey:PostID"`
}

func MigrateComments(db *gorm.DB) {
	db.AutoMigrate(&Comment{})
}
