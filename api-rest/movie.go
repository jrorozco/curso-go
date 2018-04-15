package main

/** movie.go
*ESTE ES UN OBJETO ENCAPSULADO Y EL MODELO DE LAS PELICULAS LA CUAL OCUPARESMOS PARA REGRESAR UN JSON EN LAS PETICIONES
*
 */
type Movie struct {
	Name     string `json:"name"` // esto es para cuando se cambie a json el nombre del atributo json sera de name y no Name
	Year     int    `json:"year"`
	Director string `json:"director"`
}

type Movies []Movie
