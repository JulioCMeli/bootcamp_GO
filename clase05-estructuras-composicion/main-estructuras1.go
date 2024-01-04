/*
Ejercicio 1
Crear un programa que cumpla los siguiente puntos:

X Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definido desde main().
*/

package main

import "fmt"

var Products = make([]Product, 0, 3)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

// Metodo
func (p Product) Save() {
	Products = append(Products, p)
}

// Metodo
func (p Product) GetAll() {
	fmt.Printf("GetAll: %v\n", Products)
}

// Funcion
func getById(i int) (Product, string) {
	for _, p := range Products {
		if p.ID == i {
			return p, "OK"
		}
	}
	var pAux Product
	return pAux, "NOT_FOUND"
}

func main() {

	var p1 Product
	p1.ID = 1
	p1.Name = "Galletas"
	p1.Price = 10.1
	p1.Description = "Chocolate"
	p1.Category = "Dulceria"

	p1.Save()
	p1.GetAll()

	var p2 Product
	p2.ID = 2
	p2.Name = "Cocacola"
	p2.Price = 13.5
	p2.Description = "Bebida"
	p2.Category = "Bebibles"

	p2.Save()
	p2.GetAll()

	p3, err := getById(10)
	fmt.Printf("Prduct: %v, error: %v \n", p3, err)

}
