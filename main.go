package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/calc", handleCalc)
	r.GET("/health", func(c *gin.Context) {
		host, _ := os.Hostname()
		c.JSON(200, host)
	})
	r.Run()
}

func handleCalc(c *gin.Context) {
	amount := 0.0
	add := c.QueryArray("add")
	subtract := c.QueryArray("subtract")
	multiply := c.QueryArray("multiply")
	divide := c.QueryArray("divide")

	addAmount := getAdd(add)
	subtractAmount := getSubtract(subtract)
	multiplyAmount := getMultiply(multiply)
	divideAmount := getDivide(divide)

	amount += addAmount
	amount += subtractAmount
	amount *= multiplyAmount
	amount *= divideAmount
	c.JSON(200, amount)
}

func getAdd(add []string) float64 {
	req, _ := http.NewRequest("GET", "http://addition:8080/add", nil)
	return doHTTP(req, add, "add")
}

func getSubtract(add []string) float64 {
	req, _ := http.NewRequest("GET", "http://subtraction:8080/subtract", nil)
	return doHTTP(req, add, "subtract")
}

func getMultiply(add []string) float64 {
	req, _ := http.NewRequest("GET", "http://multiplication:8080/multiply", nil)
	return doHTTP(req, add, "multiply")
}

func getDivide(add []string) float64 {
	req, _ := http.NewRequest("GET", "http://division:8080/divide", nil)
	return doHTTP(req, add, "divide")
}

func doHTTP(req *http.Request, arr []string, qu string) float64 {
	q := req.URL.Query()
	for _, ad := range arr {
		q.Add(qu, ad)
	}
	req.URL.RawQuery = q.Encode()

	res, err := http.Get(req.URL.String())
	if err != nil {
		return 0
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	text := string(body)
	if flt, err := strconv.ParseFloat(text, 64); err == nil {
		return flt
	}
	return 0.0
}
