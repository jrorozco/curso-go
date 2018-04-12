package main

import "fmt"

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	var gorra_negra = Gorra{
		marca:  "nike",
		color:  "negro",
		precio: 25.50,
		plana:  false}
	fmt.Println("Este es un tipo personalizado", gorra_negra)

	// esta es otra forma de expresar un estrutura personalizada
	var gorra_roja = Gorra{"adidas", "roja", 20.50, true} // si sabemos el orden no es necesario poner antes los datos asociados
	fmt.Println("Esta es otra forma de expresar un objeto personalizado ", gorra_roja)
	// para hacer referencia a cada atributo se ejecuta de la siquiente manera
	fmt.Println("De esta forma hacemos referencia a los atributos :")
	fmt.Println("Marca ", gorra_negra.marca)
	fmt.Println("Color", gorra_negra.color)
	fmt.Println("Precio", gorra_negra.precio)
	fmt.Println("Es plana", gorra_negra.plana)
}
