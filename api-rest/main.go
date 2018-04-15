package main

import (
	"log"
	"net/http"
)

/** main
* AQUI SOLO CREAREMOS UN OBJETO router haciendo referncia a la clase routes.go y se lo pasaremos al server.
*
 */
func main() {
	/*	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hola mundo desde mi servidor web con go")
		})
		server := http.ListenAndServe(":8080", nil)
		log.Fatal(server)*/

	//router := mux.NewRouter().StrictSlash(true)    // creamos la el router donde estaran nuestras rutas
	//router.HandleFunc("/", Index)                  //* creamos la ruta index como segundo parametro pasamos el nombre de una funcion en este caso la funcion Index func Index la cual declaramos abajo
	//router.HandleFunc("/peliculas", MovieList)     //** creamos la ruta contacto como segundo parametro pasamos el nombre de una funcion en este caso la funcion Contact func Index la cual declaramos abajo
	//router.HandleFunc("/pelicula/{id}", MovieShow) //*** crearmos una ruta en la cual le vamos a pasar un parametro con /{parametro}
	router := NewRouter()
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
