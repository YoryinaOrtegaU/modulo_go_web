package main

import (
	
	l "goweb/ejercicoget/lector"
	"net/http"
	"strconv"

	r "goweb/ejercicoget/repositorio/imple"

	"github.com/gin-gonic/gin"
)

func main() {

	productos := l.LeerJson("products.json")

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {

		ctx.String(http.StatusOK, "pong")
	})
	// Lista de productos
	router.GET("/products", func(ctx *gin.Context) {

		ctx.IndentedJSON(http.StatusOK, gin.H{"listaProductos": productos})
	})

	// Producto por ID
	router.GET("/products/:id", func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}

		var ProductService = r.ProductService{}

		producto := ProductService.ObtenerById(productos, id)

		ctx.IndentedJSON(http.StatusOK, gin.H{"productoBuscado": producto})
	})

	// Producto por ID con query
	router.GET("/products/search", func(ctx *gin.Context) {

		price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Param",
			})
		}

		var ProductService = r.ProductService{}

		producto := ProductService.ObtenerByPrice(productos, price)

		ctx.IndentedJSON(http.StatusOK, gin.H{"productosPrecioMayor": producto})
	})

	router.Run(":8080")
}
