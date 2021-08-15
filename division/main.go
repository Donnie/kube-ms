package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/divide", handleDiv)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, nil)
	})
	r.Run()
}

func handleDiv(c *gin.Context) {
	if divide, ok := c.GetQueryArray("divide"); ok {
		fmt.Println(Divide(divide))
	}
	c.JSON(200, c.Request.URL.Query())
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
