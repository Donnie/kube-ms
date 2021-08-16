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
	r.GET("/calculate", handleCalc)
	r.GET("/health", func(c *gin.Context) {
		host, _ := os.Hostname()
		c.JSON(200, host)
	})
	r.Run()
}

func handleCalc(c *gin.Context) {
	amount := 0.0

	calcHandler := make(map[string]func(prev *float64, add []string))
	calcHandler["add"] = getAdd
	calcHandler["subtract"] = getSubtract
	calcHandler["multiply"] = getMultiply
	calcHandler["divide"] = getDivide

	queries := c.Request.URL.Query()
	for param, values := range queries {
		calcHandler[param](&amount, values)
	}

	c.JSON(200, amount)
}

func getAdd(prev *float64, add []string) {
	url := "http://addition:8080/calculate"
	*prev = *prev + parseFloat(doHTTP(url, add))
}

func getSubtract(prev *float64, add []string) {
	url := "http://subtraction:8080/calculate"
	*prev = *prev + parseFloat(doHTTP(url, add))
}

func getMultiply(prev *float64, add []string) {
	url := "http://multiplication:8080/calculate"
	*prev = *prev * parseFloat(doHTTP(url, add))
}

func getDivide(prev *float64, add []string) {
	url := "http://division:8080/calculate"
	*prev = *prev * parseFloat(doHTTP(url, add))
}

func parseFloat(input string) (output float64) {
	if flt, err := strconv.ParseFloat(input, 64); err == nil {
		output = flt
	}
	return
}

func doHTTP(url string, arr []string) string {
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	for _, ad := range arr {
		q.Add("val", ad)
	}
	req.URL.RawQuery = q.Encode()

	res, err := http.Get(req.URL.String())
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
