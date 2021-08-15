package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/subtract", handleSub)
	r.GET("/health", func(c *gin.Context) {
		host, _ := os.Hostname()
		c.JSON(200, host)
	})
	r.Run()
}

func handleSub(c *gin.Context) {
	if subtract, ok := c.GetQueryArray("subtract"); ok {
		c.JSON(200, Subtract(subtract))
	}
}

//Subtract numbers
func Subtract(arr []string) (sum float64) {
	for _, sub := range arr {
		if flt, err := strconv.ParseFloat(sub, 64); err == nil {
			sum -= flt
		}
	}
	return
}
