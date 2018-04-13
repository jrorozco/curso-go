package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Escribir en archivo")

	nuevo_texto := "\n" + os.Args[1]

	//escribir := ioutil.WriteFile("ficheroEscribir.txt", nuevo_texto, 0777)
	//fmt.Println(escribir)

	fichero, err := os.OpenFile("ficheroEscribir.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	showError(err)

	escribir, err := fichero.WriteString(nuevo_texto)
	fmt.Println(escribir)
	showError(err)

	fichero.Close()

	texto, err := ioutil.ReadFile("ficheroEscribir.txt")
	showError(err)

	fmt.Println(string(texto))
}
func showError(e error) {
	if e != nil {
		panic(e)
	}
}
