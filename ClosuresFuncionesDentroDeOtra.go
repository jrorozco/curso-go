package main

import (
	"fmt"
)

func main() {
	fmt.Println(gorras(3, "Dlls"))
	patanlon("rojo", "largo", "sin bolsillos", "nike")

}

func gorras(pedido float32, moneda string) (string, float32, string) {
	//Funcion dentro de otra
	precio := func() float32 {
		return pedido * 7
	}
	return "El precio  de las gorras pedidas es : ", precio(), moneda
}

//funciones con argumentos
func patanlon(caracteristicas ...string) {
	for _, caracteristica := range caracteristicas {
		fmt.Println(caracteristica)
	}
}
