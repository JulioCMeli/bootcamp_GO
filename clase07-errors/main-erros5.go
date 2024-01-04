/*
Ejercicio 5 -  Impuestos de salario #5

Vamos a hacer que nuestro programa sea un poco más complejo y útil.

Desarrollá las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10 % en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas
sea menor a 80 o un número negativo, la función debe
devolver un error. El mismo tendrá que indicar “Error: the worker cannot have worked less than 80 hours per month”.
*/

package main

import (
	"errors"
	"fmt"
)

var err = errors.New("Error: the worker cannot have worked less than 80 hours per month")
var maxfreeTaxSalary float64 = 150000
var minHours = 80.0
var taxPayment = 0.10

func getPayment(hours float64, payPerHour float64) (float64, error) {
	if hours < float64(minHours) {
		return 0.0, err
	}

	payment := hours * payPerHour

	if payment >= maxfreeTaxSalary {
		payment = payment * (1.0 - taxPayment)
	}

	return payment, nil

}

func main() {

	payment, err := getPayment(80, 1874)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Total Payment: %v\n", payment)

}
