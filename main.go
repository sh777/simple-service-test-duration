package main

import (
	"fmt"
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
		fmt.Println(c.Request.Header)
		c.JSON(http.StatusOK, c.Request.Header)
	})

	router.POST("/header", func(c *gin.Context) {
		fmt.Println(c.Request.Header)
		c.JSON(http.StatusOK, c.Request.Header)
	})

	router.GET("/header/:key/:value", func(c *gin.Context) {
		//fmt.Println(c.Request.Header)
		key := c.Param("key")
		value := c.Param("value")
		if c.Request.Header.Get(key) == value {
			c.JSON(http.StatusOK, key+": "+value)
		} else {
			c.JSON(http.StatusBadRequest, "Excepting "+value+" but getting "+key+": "+c.Request.Header.Get(key))
		}

	})

	router.GET("/response/:code", func(c *gin.Context) {
		code := c.Param("code")
		intCode, err := strconv.Atoi(code)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Excepting Integer but getting "+code)
		} else {
			c.String(intCode, "%s responsed %s on %s", service_tag, code, hostName)
		}

	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
