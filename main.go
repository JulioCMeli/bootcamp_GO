// go run main.go
// go build main.go && ./main
// go mod init -> inicializa el module, archivo .go.mod
// go mod tidy -> gestiona, ordena, actualiza las dependencias

package main

import (
	"fmt"
	//"log"
	//"net/http"
)

var (
	dog string
	cat string
)

func main() {

	var area int = 10

	fullName := "Julio Diaz"

	dog = "dog"
	cat = "cat"

	fmt.Println("Hola, mundo.")
	fmt.Println(dog)
	fmt.Println(area, fullName)
	fmt.Printf("Value: %v \n", fullName)

}

/*
func holaFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo!")
}
func main() {
	http.HandleFunc("/", holaFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}*/
