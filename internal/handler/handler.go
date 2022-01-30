package handler

import (
	"net/http"
	"time"

	"github.com/DanielTitkov/fibonacci-service-sample/internal/fib"
	"github.com/gin-gonic/gin"
)

const sequenceLimit = 200000 // it takes around 100 ms

type fibonacciRequest struct {
	N *int `form:"n"`
}

func GetFibonacciNumberHandlerFunc(c *gin.Context) {
	var request fibonacciRequest
	//nolint
	c.Bind(&request)

	if request.N == nil || *request.N < 0 || *request.N > sequenceLimit {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"requestPresent": false,
			"limit":          sequenceLimit,
		})
		return
	}

	start := time.Now()
	fibN := fib.N(*request.N)
	elapsed := time.Since(start)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"n":              *request.N,
		"requestPresent": true,
		"fibN":           fibN,
		"elapsed":        elapsed,
	})
}
