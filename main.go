package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
}

func main() {
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	// Configurations from environment variables
	Config.SetEnvPrefix("SIMPLE_SERVICE")
	Config.AutomaticEnv()

	service_tag := Config.GetString("tag")
	if service_tag == "" {
		service_tag = "service"
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
		c.String(http.StatusOK, "%s delayed %s seconds on %s", service_tag, duration, hostName)
	})

	router.GET("/header", func(c *gin.Context) {
		c.JSON(http.StatusOK, c.Request.Header)
	})

	router.POST("/header", func(c *gin.Context) {
		c.JSON(http.StatusOK, c.Request.Header)
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
