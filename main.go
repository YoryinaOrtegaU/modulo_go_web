package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string
	LastName string
}

func main() {

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {

		ctx.String(http.StatusOK, "pong")
	})

	var body Person
	router.POST("/saludo", func(ctx *gin.Context) {
		ctx.BindJSON(&body)

		message := fmt.Sprintf("Hola, %s %s", body.Name, body.LastName)
		ctx.String(http.StatusOK, message)
	})
	router.Run(":8080")
}
