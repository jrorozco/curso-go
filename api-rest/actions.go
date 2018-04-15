package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/** actions.go
* EN ESTA RUTINA PONEMOS LAS FUNCIONES QUE ESTARAN EN UN ARRAY DE OBJETOS router EN LA CLASE routes.go
* EN ESTAS RUTINAS ESTARA TODA LA LOGICA Y QUE ES LO QUE REGRESAREMOS POR POST PUT GET ASI
*
 */
var movies = Movies{
	Movie{"Sin limites", 2013, "Desconocido"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"A todo gas", 2005, "Juan Antonio"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con go")
}

//**
func MovieList(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(movies)
	//fmt.Fprintf(w, "Esta es Listado de peliculas")
}

//***
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)    // aqui capturamos todas los parametros enviados por la url
	movie_Id := params["id"] //en esta parte obtenemos el id con el nombre que se envio -> /pelicula/{id}
	fmt.Fprintf(w, "Este es el numero envioado por la url %s", movie_Id)
}
func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	log.Println(movieData)
	movies = append(movies, movieData)
}
