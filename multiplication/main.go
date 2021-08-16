package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/calculate", handleMul)
	r.GET("/health", func(c *gin.Context) {
		host, _ := os.Hostname()
		c.JSON(200, host)
	})
	r.Run()
}

func handleMul(c *gin.Context) {
	if multiply, ok := c.GetQueryArray("val"); ok {
		c.JSON(200, Multiply(multiply))
	}
}

//Multiply numbers
func Multiply(arr []string) (sum float64) {
	sum = 1
	for _, mul := range arr {
		if flt, err := strconv.ParseFloat(mul, 64); err == nil {
			sum *= flt
		}
	}
	return
}
