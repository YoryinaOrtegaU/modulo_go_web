package repositorio

import (
	p "goweb/ejercicoget/modelo"
)

type Repositorio interface {
	ObtenerById(slice []p.Product, id int) p.Product
	ObtenerByPrice(slice []p.Product, price float64) []p.Product
	Guardar() p.Product
	Actualizar() p.Product
	Borrar()
}
