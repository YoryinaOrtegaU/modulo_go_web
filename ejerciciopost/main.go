package main

import (
	"errors"
	"fmt"
	l "goweb/ejerciciopost/lector"
	"goweb/ejerciciopost/modelo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AñadirProducto(productoNuevo modelo.Product, sliceProd *modelo.ProductsList) error {

	sliceProd1 := sliceProd.ProductsList
	tamanoSlice := len(sliceProd1)

	ultimoProd := sliceProd1[tamanoSlice-1]
	ultimoId := ultimoProd.Id

	nuevoId := ultimoId + 1
	productoNuevo.Id = nuevoId

	for _, prod := range sliceProd1 {
		if prod.CodeValue == productoNuevo.CodeValue {
			return errors.New("CodeValue no valido, ya existe en el slice")
		}
	}

	//sliceProd1 = append(sliceProd1, productoNuevo)

	sliceProd.ProductsList = append(sliceProd.ProductsList, productoNuevo)

	fmt.Println("1)", sliceProd.ProductsList)
	return nil

}

func findById(id int, sliceProd modelo.ProductsList) modelo.Product {
	sliceProd1 := sliceProd.ProductsList
	for _, prod := range sliceProd1 {
		if prod.Id == id {
			return prod
		}
	}
	panic("no encontró el product")
}

func main() {

	// slice de productos
	sliceProd := l.LeerJson("products.json")

	//Servidor
	router := gin.Default()

	//añadir un producto al slice
	router.POST("/products", func(ctx *gin.Context) {
		var prod modelo.Product

		if err := ctx.BindJSON(&prod); err != nil {
			fmt.Println(err)
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}
		err := AñadirProducto(prod, &sliceProd)
		if err != nil {
			fmt.Println(err)
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println("2)", sliceProd)

		ctx.String(http.StatusOK, "producto añadido")

	})

	//Buscar por Id el producto agregado
	router.GET("/products/:id", func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}

		producto := findById(id, sliceProd)

		ctx.IndentedJSON(http.StatusOK, gin.H{"productoBuscado": producto})
	})

	router.Run(":8080")
}
