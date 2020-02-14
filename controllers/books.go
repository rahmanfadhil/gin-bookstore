package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rahmanfadhil/gin-boookstore/models"
)

func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var books []models.Book
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
