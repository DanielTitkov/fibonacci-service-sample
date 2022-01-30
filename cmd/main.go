package main

import (
	"log"

	"github.com/DanielTitkov/fibonacci-service-sample/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("service starting...")

	r := gin.Default()
	//nolint
	r.SetTrustedProxies(nil)
	r.Use(cors.Default())
	r.LoadHTMLGlob("templates/*.html")
	// TODO: add middleware to forcefully stop requests which are taking too long

	v1 := r.Group("/api/v1")
	v1.GET("/getFibNumber", handler.GetFibonacciNumberHandlerFunc)

	log.Fatalln("failed to run server", r.Run("0.0.0.0:1123"))
}
