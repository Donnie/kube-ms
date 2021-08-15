package multiplication

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func serve() (err error) {
	r := gin.Default()
	r.GET("/multiply", handleMul)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, nil)
	})
	r.Run()
	return
}

func handleMul(c *gin.Context) {
	if multiply, ok := c.GetQueryArray("multiply"); ok {
		fmt.Println(Multiply(multiply))
	}
	c.JSON(200, c.Request.URL.Query())
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
