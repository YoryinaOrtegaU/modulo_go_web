package lector

import (
	"encoding/json"
	"fmt"
	m "goweb/ejerciciopost/modelo"
	"io/ioutil"
	"os"
)

func LeerJson(rutaArchivo string) m.ProductsList {

	jsonFile, err := os.Open(rutaArchivo)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	productsList := m.ProductsList{}
	sliceProd := &productsList.ProductsList

	err = json.Unmarshal(byteValue, &sliceProd)
	if err != nil {
		fmt.Println(err)
	}

	b, err := json.Marshal(sliceProd)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))

	return productsList
}
