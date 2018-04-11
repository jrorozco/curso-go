package main

import "fmt"

func main() {
	// VARIABLES
	var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "Rene"
	var apellidos string = "Orozco"
	var prueba bool = true
	var floatVariable float32 = 1.024
	// CONSTANTES
	const year int = 2018
	fmt.Println(year)
	fmt.Println("esta es un float", floatVariable)
	fmt.Println("Este es un boolean ", prueba)
	pais := "Espana"
	fmt.Println("mi nombre es " + nombre + " " + apellidos + " Pais " + pais)
	fmt.Println(suma)
	fmt.Println(resta)
}
