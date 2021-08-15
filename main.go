package main

import (
	"fmt"

	"github.com/Donnie/kube-ms/addition"
	"github.com/Donnie/kube-ms/division"
	"github.com/Donnie/kube-ms/multiplication"
	"github.com/Donnie/kube-ms/subtraction"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/calc", handleCalc)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, nil)
	})
	r.Run()
}

func handleCalc(c *gin.Context) {
	amount := 0.0
	if add, ok := c.GetQueryArray("add"); ok {
		amount += addition.Add(add)
	}
	if subtract, ok := c.GetQueryArray("subtract"); ok {
		amount += subtraction.Subtract(subtract)
	}
	if divide, ok := c.GetQueryArray("divide"); ok {
		amount *= division.Divide(divide)
	}
	if multiply, ok := c.GetQueryArray("multiply"); ok {
		amount *= multiplication.Multiply(multiply)
	}
	fmt.Println(amount)
	c.JSON(200, amount)
}
