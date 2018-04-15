package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

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

/**
*con esta rutina listamos el documento movies y lo cargamos en un objeto json para enviarlo a la vista
*
 */
func MovieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results) // hacemos una consulta a la BD del documento movies y en el Find(nil) le ponemos nil para que se traiga todo con All y el & le pasamos todo a la variable results
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados : ", results)
	}
	responseMovies(w, 200, results)
	//json.NewEncoder(w).Encode(movies)
	//fmt.Fprintf(w, "Esta es Listado de peliculas")
}

/**
*en esta rutina es para obtener un documento con el id de la BD
*
 */
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)    // aqui capturamos todas los parametros enviados por la url
	movie_Id := params["id"] //en esta parte obtenemos el id con el nombre que se envio -> /pelicula/{id}

	if !bson.IsObjectIdHex(movie_Id) {
		w.WriteHeader(404)
		return
	}

	objId := bson.ObjectIdHex(movie_Id)

	results := Movie{}

	err := collection.FindId(objId).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	} else {
		responseMovie(w, 200, results)
	}

	fmt.Fprintf(w, "Este es el numero envioado por la url %s", movie_Id)
}

/**
*En esta rutina agregamos un nuevo registro al documento movies de la BD curso-go
*
 */
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
	//** verificar los metodos resposeMovie y responseMovies
	//json.NewEncoder(w).Encode(movieData)               // convertimos a json movieData con la variable w que es el response
	//w.Header().Set("Content-Type", "application/json") // le indicamos que devolevera los datos como json
	//w.WriteHeader(200)                                 // como todo ha ido bien le indicamos que el estatus es 200 de ok

	responseMovie(w, 200, movieData)

}

/**
*en esta rutina es para actualizar un documento con el id de la BD
*
 */
func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)    // aqui capturamos todas los parametros enviados por la url
	movie_Id := params["id"] //en esta parte obtenemos el id con el nombre que se envio -> /pelicula/{id}

	if !bson.IsObjectIdHex(movie_Id) { // comprobamos que sea un id hexadecimal es correcto
		w.WriteHeader(404)
		return
	}

	objId := bson.ObjectIdHex(movie_Id) // lo guardamos en una variable
	decoder := json.NewDecoder(r.Body)  // codifiacos lo que nos llega por el body a json
	var movieData Movie
	err := decoder.Decode(&movieData)
	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}
	defer r.Body.Close()                      // cerramos el body
	document := bson.M{"_id": objId}          // guardamos en una variable donde el document _id sea igual al que le estamos pasando en este caso objId
	change := bson.M{"$set": movieData}       // guardamos en una varible el objeto json que vamos a cambiar en el documento
	err = collection.Update(document, change) // con esta instruccion le pasamos el id del documento y el cambio del json que estamos pasando
	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, movieData)

	fmt.Fprintf(w, "Este es el numero envioado por la url %s", movie_Id)
}

// objeto de tipo Message
type Message struct { // objeto de tipo mensaje para devoleverlo cuado hacemos un DELETE
	Status  string `json:"status"`
	Message string `json:"message"`
}

//*****SETTERS Message

func (this *Message) setStatus(data string) {
	this.Status = data
}
func (this *Message) setMessage(data string) {
	this.Message = data
}

//*****SETTERS Message

/**
*en esta rutina es para obtener un documento con el id de la BD
*
 */
func MovieRemove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)    // aqui capturamos todas los parametros enviados por la url
	movie_Id := params["id"] //en esta parte obtenemos el id con el nombre que se envio -> /pelicula/{id}

	if !bson.IsObjectIdHex(movie_Id) {
		w.WriteHeader(404)
		return
	}

	objId := bson.ObjectIdHex(movie_Id)
	err := collection.RemoveId(objId)

	if err != nil {
		w.WriteHeader(404)
		return
	}
	//results := Message{"Sussecs", "La pelicula con ID " + movie_Id + " ha sido borrada correctamente"}
	message := new(Message) // mensajes con setters
	message.setStatus("Sussecs")
	message.setMessage("La pelicula con ID " + movie_Id + " ha sido borrada correctamente")
	results := message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)

	fmt.Fprintf(w, "Este es el numero envioado por la url %s", movie_Id)
}

func responseMovie(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}
