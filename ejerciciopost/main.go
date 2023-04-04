package main

import (
	"encoding/json"
	"fmt"
	l "goweb/ejerciciopost/lector"
	"goweb/ejerciciopost/modelo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AñadirProducto(productoNuevo modelo.Product, sliceProd modelo.ProductsList) {

	sliceProd1 := sliceProd.ProductsList
	tamanoSlice := len(sliceProd1)

	ultimoProd := sliceProd1[tamanoSlice-1]
	ultimoId := ultimoProd.Id

	nuevoId := ultimoId + 1
	productoNuevo.Id = nuevoId

	for _, prod := range sliceProd1 {
		if prod.CodeValue == productoNuevo.CodeValue {
			fmt.Println("CodeValue no valido, ya existe en el slice")
		}
	}

	sliceProd1 = append(sliceProd1, productoNuevo)

}

func main() {

	// slice de productos
	sliceProd := l.LeerJson("products.json")
	fmt.Println(sliceProd)

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
		AñadirProducto(prod, sliceProd)

		ctx.String(http.StatusOK, "producto añadido")

		prodc, _ := json.Marshal(prod)

		fmt.Println(string(prodc))
	})

	fmt.Println(sliceProd)
	router.Run(":8080")
}
