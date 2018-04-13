package main

import "fmt"

/**
* Slice es como un array pero sin un limite en particular osea un array dinamico
 */
func main() {
	peliculas := []string{"La Vedad duele 2", "Ciudadano Ejemplar 2", "Gran Torino 2"}
	fmt.Println(peliculas)
	fmt.Println(peliculas[0])
	fmt.Println(peliculas[1])
	fmt.Println(peliculas[2])

	//con esta instruccion agregamos un elemento al array o al slice

	peliculas = append(peliculas, "Sin Limites")

	fmt.Println(peliculas)

	// mas funciones utilies
	// para saber el numero de elementos de un slice
	fmt.Println(len(peliculas))
	// para que solo me muestre un cirto rango de elementos
	fmt.Println(peliculas[0:2])

}
