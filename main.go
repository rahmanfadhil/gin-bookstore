package main

import (
	"github.com/rahmanfadhil/gin-boookstore/controllers"
	"github.com/rahmanfadhil/gin-boookstore/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := gin.Default()

	// Connect to database
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrate models
	db.AutoMigrate(&models.Book{})

	// Provide db to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Routes
	r.GET("/books", controllers.FindBooks)

	// Run the server
	r.Run()
}
