package main

import (
	"log"

	"github.com/DanielTitkov/fibonacci-service-sample/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("service starting...")

	r := setupServer("templates/*.html")

	log.Fatalln("failed to run server", r.Run("0.0.0.0:1123"))
}

func setupServer(templatesPath string) *gin.Engine {
	r := gin.Default()
	//nolint
	r.SetTrustedProxies(nil)
	r.Use(cors.Default())
	r.LoadHTMLGlob(templatesPath)
	// TODO: add middleware to forcefully stop requests which are taking too long

	r.GET("/", handler.GetFibonacciNumberHandlerFunc)

	return r
}
