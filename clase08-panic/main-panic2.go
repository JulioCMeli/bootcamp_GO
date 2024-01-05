/*
Ejercicio 2 - Imprimir datos


A continuación, vamos a crear un archivo “customers.txt” con información
 de los clientes del estudio.

Ahora que el archivo sí existe, el panic no debe ser lanzado.

Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
Extendemos el código del punto uno para que podamos leer este archivo e
imprimir los datos que contenga. En el caso de no poder leerlo, se debe lanzar un “panic”.
Recordemos que siempre que termina la ejecución, independientemente del
resultado, debemos tener un “defer” que nos indique que la ejecución finalizó.
También recordemos cerrar los archivos al finalizar su uso.

*/

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	//Esto se va a ejecutar a pesar de un panic
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("ejecución finalizada")
	}()

	// open a file
	var fileName = "customers.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("The indicated file was not found or is damaged")
		return
	}
	// defer close file
	defer file.Close()

	//reading file
	fmt.Println(file.Name())
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}

}
