/*
Ejercicio 1 - Impuestos de salario #1

En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

Creá un error personalizado con un struct que implemente “Error()”
con el mensaje “Error: the salary entered does not reach the taxable minimum"
 y lanzalo en caso de que “salary” sea menor a 150.000.
 De lo contrario, tendrás que imprimir por consola el mensaje “Must pay tax”.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

type CustomError struct {
	Msg  string
	Code int
}

func (e *CustomError) Error() string {
	return e.Msg + " " + strconv.Itoa(e.Code)
}

func main() {

	var salary int = 200000

	var custErro CustomError
	custErro.Msg = "Error: the salary entered does not reach the taxable minimum"
	custErro.Code = 401

	if salary < 150000 {
		fmt.Println(custErro)
		os.Exit(1)
	}

	//var err error = errors.New("this is an error")

	fmt.Printf("Must pay tax\n")

}
