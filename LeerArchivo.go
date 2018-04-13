package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Lector : ")
	fichero, err := ioutil.ReadFile("fichero.txt")
	showError(err)
	fmt.Println(string(fichero))
}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
