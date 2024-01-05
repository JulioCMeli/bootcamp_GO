/*
Crear un programa que cumpla los siguiente puntos:

Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definido desde main().
*/

package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func main() {
	var a int = 20    /* actual variable declaration */
	var p_a *int = &a /* pointer variable declaration and asigment to "a" address */

	p_a = &a /* store address of a in pointer variable*/

	fmt.Printf("Address of a variable: %x\n", &a)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in p_a variable: %x\n", p_a)

	/* access the value using the pointer */
	fmt.Printf("Value of *p_a variable: %d\n", *p_a)
}
