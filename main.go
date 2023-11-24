package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juniovitorino/golang-bookstore-rest-api/controllers"
	"github.com/juniovitorino/golang-bookstore-rest-api/models"
)

func main() {
	r :=
		gin.Default()
	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/book/:id", controllers.FindBook)
	r.PATCH("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Server is running..."})
	})

	r.Run()
}
