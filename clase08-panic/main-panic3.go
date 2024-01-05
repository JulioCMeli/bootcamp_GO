/*
Ejercicio 3 - Registrando clientes
El mismo estudio del ejercicio anterior,
solicita una funcionalidad para poder registrar
datos de nuevos clientes. Los datos requeridos son:

File
Name
ID
Phone number
Home

Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe.
Para ello, necesitás leer los datos de un array. En caso de que esté repetido,
debes manipular adecuadamente el error como hemos visto hasta aquí. Ese error deberá:
1.- generar un panic;

2.- lanzar por consola el mensaje: “Error: client already exists”, y continuar
con la ejecución del programa normalmente.

Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe,
desarrollá una función para validar que todos los datos a registrar de un cliente
contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores.
Uno de ellos tendrá que ser del tipo error para el caso de que se ingrese por parámetro
algún valor cero (recordá los valores cero de cada tipo de dato, ej: 0, “”, nil).
Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics,
se deberán imprimir por consola los siguientes mensajes: “End of execution” y “Several errors were detected at runtime”. Utilizá defer para cumplir con este requerimiento.

Requerimientos generales:

Utilizá recover para recuperar el valor de los panics que puedan surgir
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
Generá algún error, personalizandolo a tu gusto utilizando alguna de las
funciones de Go (realiza también la validación pertinente para el caso de error retornado).

*/

package main

import (
	"errors"
	"fmt"
)

type Customer struct {
	ID          int
	File        string
	Name        string
	PhoneNumber string
	Home        string
}

var errValue = errors.New("Invalid value for Customer")

func addCustomer(c Customer, lstCustomers *[]Customer) {

	fmt.Printf("Adding customer:  %v \n", c)

	// defer function to recover from panic in same fuinciton scope
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	var _, err = validCustomerValues(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	var r = customerExist(c.ID, *lstCustomers)
	if r {
		var msg = fmt.Sprintf("Error: client  with ID [%v] already exists", c.ID)
		panic(msg)
	} else {
		*lstCustomers = append(*lstCustomers, c)
	}
}

func customerExist(id int, lstCustomers []Customer) bool {
	for _, c := range lstCustomers {
		if c.ID == id {
			return true
		}
	}
	return false
}

func validCustomerValues(c Customer) (bool, error) {
	if c.ID == 0 {
		return false, errValue
	}
	if c.File == "" {
		return false, errValue
	}
	if c.Name == "" {
		return false, errValue
	}
	if c.PhoneNumber == "" {
		return false, errValue
	}
	if c.Home == "" {
		return false, errValue
	}
	return true, nil
}

func main() {

	//Esto se va a ejecutar a pesar de un panic
	defer func() {
		fmt.Println("End of execution")
		fmt.Println("Several errors were detected at runtime")
	}()

	var lstCustomers []Customer

	var c1 Customer
	c1.ID = 1
	c1.File = "tbd"
	c1.Name = "Julio"
	c1.PhoneNumber = "tbd"
	c1.Home = "tbd"
	var c2 = Customer{
		ID:          1,
		File:        "tbd",
		Name:        "tbd",
		PhoneNumber: "tbd",
		Home:        "tbd"}
	var c3 = Customer{
		ID:          3,
		File:        "tbd",
		Name:        "Cesar",
		PhoneNumber: "tbd",
		Home:        "tbd"}

	addCustomer(c1, &lstCustomers)
	fmt.Printf("Customers: %v \n", lstCustomers)

	addCustomer(c2, &lstCustomers)
	fmt.Printf("Customers: %v \n", lstCustomers)

	addCustomer(c3, &lstCustomers)
	fmt.Printf("Customers: %v \n", lstCustomers)

}
