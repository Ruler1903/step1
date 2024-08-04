package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainHandler(c *gin.Context) {
	c.String(200, "Main page")
}

func main() {
	r := gin.New()

	r.Use()
	adminGroup := r.Group("/", gin.BasicAuth(gin.Accounts{
		"hangeldi": "190318082000",
	}))

	adminGroup.GET("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"secret": "OK",
		})

	})

	adminGroup.GET("/main", mainHandler)

	r.Run(":8080")
}