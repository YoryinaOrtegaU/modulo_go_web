package imple

import (
	p "goweb/ejercicoget/modelo"
)

type ProductService struct {
}

func (p ProductService) ObtenerById(slice []p.Product, id int) p.Product {

	for _, product := range slice {
		if product.Id == id {
			return product
		}
	}
	panic("no encontró el product")
}

func (ser ProductService) ObtenerByPrice(slice []p.Product, price float64) []p.Product {

	var productosEncontrados []p.Product

	for _, product := range slice {
		if product.Price > price {
			productosEncontrados = append(productosEncontrados, product)
		}
	}
	return productosEncontrados
}

func (p ProductService) Guardar() p.Product {
	panic("no implementado")
}

func (p ProductService) Actualizar() p.Product {
	panic("no implementado")
}

func (p ProductService) Borrar() p.Product {
	panic("no implementado")
}
