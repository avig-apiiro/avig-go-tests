package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router := r
	router.POST("/somePost", posting)

	group := r.Group("/api/licenses")
	{
		group.POST("/somePost", posting)
		group.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	r.Run()

}

func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Data is valid"})
}
