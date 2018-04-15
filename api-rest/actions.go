package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
)

/** actions.go
* EN ESTA RUTINA PONEMOS LAS FUNCIONES QUE ESTARAN EN UN ARRAY DE OBJETOS router EN LA CLASE routes.go
* EN ESTAS RUTINAS ESTARA TODA LA LOGICA Y QUE ES LO QUE REGRESAREMOS POR POST PUT GET ASI
*
 */
/*var movies = Movies{
	Movie{"Sin limites", 2013, "Desconocido"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"A todo gas", 2005, "Juan Antonio"},
}*/

//CONEXION A MONGO
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

// obtenemos la coleccion de documentos de la base de datos curso_go del documento movies
var collection = getSession().DB("curso_go").C("movies")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con go")
}

//**
func MovieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results) // hacemos una consulta a la BD del documento movies y en el Find(nil) le ponemos nil para que se traiga todo con All y el & le pasamos todo a la variable results
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados : ", results)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
	//json.NewEncoder(w).Encode(movies)
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
	//log.Println(movieData)
	//session := getSession()
	//session.DB("curso_go").C("movies").Insert(movieData)
	err = collection.Insert(movieData)
	if err != nil { // si algo salio mal devolvemos un error 500
		w.WriteHeader(500)
		return
	}
	//movies = append(movies, movieData)
	json.NewEncoder(w).Encode(movieData)               // convertimos a json movieData con la variable w que es el response
	w.Header().Set("Content-Type", "application/json") // le indicamos que devolevera los datos como json
	w.WriteHeader(200)                                 // como todo ha ido bien le indicamos que el estatus es 200 de ok
}
