package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"strconv"
	"net/http"
)


func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/delay/:duration", func(c *gin.Context) {
		duration := c.Param("duration")
		d, _ := strconv.Atoi(duration)
		time.Sleep(time.Duration(d)*time.Second)
		c.String(http.StatusOK, "Delayed %s", duration)
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
