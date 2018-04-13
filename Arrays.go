package main

import "fmt"

func main() {
	//Arrays forma 1 unidimencionales
	var peliculas [3]string
	peliculas[0] = "La verdad duele"
	peliculas[1] = "Ciudadano Ejemplar"
	peliculas[2] = "Gran Torino"

	fmt.Println(peliculas)
	fmt.Println(peliculas[0])
	fmt.Println(peliculas[1])
	fmt.Println(peliculas[2])
	//Arrays forma 2 unidimencionales
	peliculas2 := [3]string{"La Vedad duele 2", "Ciudadano Ejemplar 2", "Gran Torino 2"}
	fmt.Println(peliculas2)
	fmt.Println(peliculas2[0])
	fmt.Println(peliculas2[1])
	fmt.Println(peliculas2[2])

	//Array de multiple dimenciones
	var peliculas3 [3][2]string

	peliculas3[0][0] = "La verdad duele"
	peliculas3[0][1] = "Mientras duermes"
	peliculas3[1][0] = "Ciudadano Ejemplar"
	peliculas3[1][1] = "Batman"
	peliculas3[2][0] = "El Señor de los anillos"
	peliculas3[2][1] = "El Hobbit"
	fmt.Println(peliculas3)
	fmt.Println(peliculas3[0][0])
	fmt.Println(peliculas3[0][1])
	fmt.Println(peliculas3[1][0])
	fmt.Println(peliculas3[1][1])
	fmt.Println(peliculas3[2][0])
	fmt.Println(peliculas3[2][1])

}
