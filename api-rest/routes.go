package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

/** routes.go
* EN ESTA PARTE ESTARA UN ARRAY DE RUTAS DONDE TENDREMOS QUE AGREGAR MAS OBJETOS Route PARA AGREGAR NUEVAS RUTAS
*A PARTE DE ESTO TENDREMOS QUE AGREGAR METODOS A  LA CLASE actios.go DONDE ESTARA TODA LA LOGICA DE LO QUE NOS LLEGA
*POR POST PUT DELETE GET
*
 */
type Route struct {
	Name       string
	Methods    string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Methods).Path(route.Pattern).Name(route.Name).Handler(route.HandleFunc)

	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"Get",
		"/",
		Index,
	},
	Route{
		"MovieList",
		"Get",
		"/peliculas",
		MovieList,
	},
	Route{
		"MovieShow",
		"Get",
		"/pelicula/{id}",
		MovieShow,
	},
	Route{
		"MovieAdd",
		"POST",
		"/pelicula",
		MovieAdd,
	},
	Route{
		"MovieUpdate",
		"PUT",
		"/pelicula/{id}",
		MovieUpdate,
	},
	Route{
		"MovieRemove",
		"DELETE",
		"/pelicula/{id}",
		MovieRemove,
	},
}
