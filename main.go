package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set middleware
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Heath Check
	r.GET("/heathCheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	api := r.Group("/api")

	{
		api.GET("todo")
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
