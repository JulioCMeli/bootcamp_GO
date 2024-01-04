/*
Ejercicio 2 - Impuestos de salario #2

En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

Creá un error personalizado con un struct que implemente “Error()”
con el mensaje “Error: salary is less than 10000" y lanzalo en caso de que
“salary” sea menor o igual a  10000.
La validación debe ser hecha con la función “Is()” dentro del “main”.

*/

package main

import (
	"errors"
	"fmt"
	"strconv"
)

type CustomError struct {
	Msg  string
	Code int
}

func (e CustomError) Error() string {
	return e.Msg + " " + strconv.Itoa(e.Code)
}

var custErro CustomError = CustomError{Msg: "Error: salary is LESS than 10000", Code: 401}

func salaryLessThan(salary int) error { //CustomError { // podmeos usar cualquiera
	if salary <= 100000 {
		return custErro
	}
	var custErroAux CustomError
	return custErroAux

}

func main() {

	var salary int = 1000001

	var err = salaryLessThan(salary)

	var isAnError = errors.Is(err, custErro)

	fmt.Printf("%v\n", isAnError)

}
