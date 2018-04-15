package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	/*	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hola mundo desde mi servidor web con go")
		})
		server := http.ListenAndServe(":8080", nil)
		log.Fatal(server)*/

	router := mux.NewRouter().StrictSlash(true)    // creamos la el router donde estaran nuestras rutas
	router.HandleFunc("/", Index)                  //* creamos la ruta index como segundo parametro pasamos el nombre de una funcion en este caso la funcion Index func Index la cual declaramos abajo
	router.HandleFunc("/peliculas", MovieList)     //** creamos la ruta contacto como segundo parametro pasamos el nombre de una funcion en este caso la funcion Contact func Index la cual declaramos abajo
	router.HandleFunc("/pelicula/{id}", MovieShow) //*** crearmos una ruta en la cual le vamos a pasar un parametro con /{parametro}
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}

//*
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con go")
}

//**
func MovieList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"Sin limites", 2013, "Desconocido"},
		Movie{"Batman Begins", 1999, "Scorsese"},
		Movie{"A todo gas", 2005, "Juan Antonio"},
	}

	json.NewEncoder(w).Encode(movies)
	//fmt.Fprintf(w, "Esta es Listado de peliculas")
}

//***
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)    // aqui capturamos todas los parametros enviados por la url
	movie_Id := params["id"] //en esta parte obtenemos el id con el nombre que se envio -> /pelicula/{id}
	fmt.Fprintf(w, "Este es el numero envioado por la url %s", movie_Id)
}
