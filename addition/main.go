package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/add", handleAdd)
	r.GET("/health", func(c *gin.Context) {
		host, _ := os.Hostname()
		c.JSON(200, host)
	})
	r.Run()
}

func handleAdd(c *gin.Context) {
	if add, ok := c.GetQueryArray("add"); ok {
		c.JSON(200, Add(add))
	}
}

//Add numbers
func Add(arr []string) (sum float64) {
	for _, add := range arr {
		if flt, err := strconv.ParseFloat(add, 64); err == nil {
			sum += flt
		}
	}
	return
}
