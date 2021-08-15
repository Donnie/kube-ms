package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/add", handleAdd)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, nil)
	})
	r.Run()
}

func handleAdd(c *gin.Context) {
	if add, ok := c.GetQueryArray("add"); ok {
		fmt.Println(Add(add))
	}
	c.JSON(200, c.Request.URL.Query())
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
