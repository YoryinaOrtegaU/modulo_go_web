package lector

import (
	"encoding/json"
	"fmt"
	m "goweb/ejerciciopost/modelo"
	"os"
)

func LeerJson(rutaArchivo string) m.ProductsList {

	jsonFile, err := os.Open(rutaArchivo)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	productsList := m.ProductsList{}
	sliceProd := &productsList.ProductsList

	decoder := json.NewDecoder(jsonFile)
	decoder.Decode(sliceProd)

	return productsList
}
