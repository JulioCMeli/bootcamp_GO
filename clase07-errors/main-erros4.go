/*
Ejercicio 4 - Impuestos de salario #4

Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”,
para que el mensaje de error reciba por parámetro el valor de “salary”
indicando que no alcanza el mínimo imponible (el mensaje mostrado por
consola deberá decir: “Error: the minimum taxable amount is 150,000
and the salary entered is: [salary]”, siendo [salary] el valor
de tipo int pasado por parámetro).
*/

package main

import (
	"errors"
	"fmt"
)

var errPattern = "Error: the minimum taxable amount is 150,000 and the salary entered is: [%v]\n"

func salaryLessThan(salary int) error {
	if salary <= 100000 {
		return fmt.Errorf(errPattern, salary)
	}

	return errors.New("Error: salary is MORE than 10000")

}

func main() {

	var salary = 100000

	var errAux = salaryLessThan(salary)

	fmt.Println(errAux)

}
