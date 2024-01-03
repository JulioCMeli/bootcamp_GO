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

// //
type IProduct interface {
	getPrice() float64
}

func getPrice(product IProduct) float64 {
	return product.getPrice()
}

////

// Estructuras
type Small struct {
	Value float64
}
type Medium struct {
	Value float64
}
type Large struct {
	Value float64
}

// Metodos
func (p Small) getPrice() float64 {
	return p.Value
}
func (p Medium) getPrice() float64 {
	return p.Value * 1.03
}
func (p Large) getPrice() float64 {
	return p.Value*1.03 + 2500
}

func main() {

	var p1 Small
	p1.Value = 1
	price := getPrice(p1)
	fmt.Printf("Small: %v -> %v \n", p1.Value, price)

	var p2 Medium
	p2.Value = 1
	price = getPrice(p2)
	fmt.Printf("Medium: %v -> %v \n", p2.Value, price)

	var p3 Large
	p3.Value = 1
	price = getPrice(p3)
	fmt.Printf("Large: %v -> %v \n", p3.Value, price)

}
