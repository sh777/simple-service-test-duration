package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/delay/:duration", func(c *gin.Context) {
		duration := c.Param("duration")
		d, _ := strconv.Atoi(duration)
		time.Sleep(time.Duration(d) * time.Second)
		c.String(http.StatusOK, "v2 delayed %s seconds on %s", duration, hostName)
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
