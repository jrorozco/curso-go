package main

import "fmt"

func main() {

	var limite int = 15
	for i := 0; i <= limite; i++ {
		//fmt.Println("El numero va en : ", i)
		if i%2 == 0 {
			fmt.Println("El numero ", i, "es Par ")
		} else {
			fmt.Println("El numero ", i, " no es Par ")
		}

	}

	peliculas := []string{"Rene", "Raul", "Arnulfo", "Pedro"}

	fmt.Println("********** FOR **********")
	for i := 0; i < len(peliculas); i++ {
		fmt.Println("pelicula es : ", peliculas[i])
	}
	fmt.Println("******* FOR EACH  *******")
	for _, pelicula := range peliculas {
		fmt.Println("Pelicula", pelicula)
	}
}
