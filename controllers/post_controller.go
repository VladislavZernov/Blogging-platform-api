package controllers

import (
	"github.com/VladislavZernov/blogging-platform-api/config"
	"github.com/VladislavZernov/blogging-platform-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 📌 1. Создание поста
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка на существование пользователя
	var user models.User
	if err := config.DB.First(&user, post.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// Сохранение поста
	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// 📌 2. Получение всех постов
func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := config.DB.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// 📌 3. Получение поста по ID
func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := config.DB.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// 📌 4. Обновление поста
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&post)
	c.JSON(http.StatusOK, post)
}

// 📌 5. Удаление поста
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Post{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
