package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/subtract", handleSub)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, nil)
	})
	r.Run()
}

func handleSub(c *gin.Context) {
	if subtract, ok := c.GetQueryArray("subtract"); ok {
		fmt.Println(Subtract(subtract))
	}
	c.JSON(200, c.Request.URL.Query())
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
