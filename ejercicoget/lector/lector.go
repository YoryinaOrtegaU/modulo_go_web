package lector

import (
	"encoding/json"
	"fmt"
	m "goweb/ejercicoget/modelo"
	"os"
)

func LeerJson(rutaArchivo string) []m.Product {

	jsonFile, err := os.Open(rutaArchivo)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	//byteValue, _ := ioutil.ReadAll(jsonFile)

	productsList := m.ProductsList{}
	products := productsList.ProductsList

	//json.Unmarshal(byteValue, &products)

	decoder := json.NewDecoder(jsonFile)

	decoder.Decode(&products)

	return products
}
