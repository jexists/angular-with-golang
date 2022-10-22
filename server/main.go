package main

import (
	"fmt"
	"net/http"
	_ "project_1/db"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("???")
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})
	r.Run()
}
