package main

import (
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

	router := mux.NewRouter().StrictSlash(true) // creamos la el router donde estaran nuestras rutas
	router.HandleFunc("/", Index)               //* creamos la ruta index como segundo parametro pasamos el nombre de una funcion en este caso la funcion Index func Index la cual declaramos abajo
	router.HandleFunc("/contacto", Contact)     //** creamos la ruta contacto como segundo parametro pasamos el nombre de una funcion en este caso la funcion Contact func Index la cual declaramos abajo
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}

//*
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con go")
}

//**
func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es la pagina de contacto")
}
