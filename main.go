package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.POST("/create", func(c *gin.Context) {
		var json User
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Name:", json.name)
		fmt.Println("Age:", json.age)

		c.JSON(http.StatusOK, gin.H{"status": "You are in"})
	})

	router.Run(":8080")
}

type User struct {
	name string `form:"name" json:"user"`
	age  int    `form:"age" json:"age"`
}
