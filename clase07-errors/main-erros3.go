/*
Ejercicio 3 - Impuestos de salario #3


Hacé lo mismo que en el ejercicio anterior
pero reformulando el código para que, en reemplazo de “Error()”,
se implemente “errors.New()”. lo contrario, tendrás que imprimir por consola el mensaje “Must pay tax”.

*/

package main

import (
	"errors"
	"fmt"
)

var err = errors.New("Error: salary is LESS than 10000")

func salaryLessThan(salary int) error {
	if salary <= 100000 {
		return err
	}

	return errors.New("Error: salary is MORE than 10000")

}

func main() {

	var salary int = 100000

	var errAux = salaryLessThan(salary)

	var isAnError = errors.Is(errAux, err)

	fmt.Printf("%v\n", isAnError)

}
