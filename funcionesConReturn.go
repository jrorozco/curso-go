package main

import (
	"fmt"
)

func main() {

	fmt.Println(devolverTexto())
	fmt.Println(devolverDosDatos())
}

//devolver un solo dato
func devolverTexto() string {
	var texto string = "Texto por defecto"
	return texto
}

//devolver dos datos
func devolverDosDatos() (string, string) {
	var dato1 string = "Hola"
	var dato2 string = "Mundo"
	return dato1, dato2
}
