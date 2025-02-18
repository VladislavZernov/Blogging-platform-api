package models

import "gorm.io/gorm"

type Post struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"size:255;not null"`
	Content  string `gorm:"type:text;not null"`
	UserID   uint   `gorm:"not null" json:"user_id"`
	User     User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments []Comment
}

func MigratePosts(db *gorm.DB) {
	db.AutoMigrate(&Post{})
}
