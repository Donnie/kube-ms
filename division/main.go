package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/calculate", handleDiv)
	r.GET("/health", func(c *gin.Context) {
		host, _ := os.Hostname()
		c.JSON(200, host)
	})
	r.Run()
}

func handleDiv(c *gin.Context) {
	if divide, ok := c.GetQueryArray("val"); ok {
		c.JSON(200, Divide(divide))
	}
}

//Divide numbers
func Divide(arr []string) (sum float64) {
	sum = 1
	for _, div := range arr {
		if flt, err := strconv.ParseFloat(div, 64); err == nil {
			sum /= flt
		}
	}
	return
}
