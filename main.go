package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string `json:"nombre"`
	LastName string `json:"apellido"`
}

func main() {

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {

		ctx.String(http.StatusOK, "pong")
	})

	router.POST("/saludo", func(ctx *gin.Context) {
		var persona1 Person

		if err := ctx.BindJSON(&persona1); err != nil {
			return
		}
		fmt.Println(persona1)
		message := fmt.Sprintf("Hola, %s %s", persona1.Name, persona1.LastName)
		ctx.String(http.StatusOK, message)
	})
	router.Run(":8080")
}
