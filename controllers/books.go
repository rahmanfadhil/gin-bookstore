package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rahmanfadhil/gin-boookstore/models"
)

// GET /books
// Find all books
func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var books []models.Book
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}
