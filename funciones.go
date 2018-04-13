package main

import "fmt"

func main() {

	holaMundo()
}

func holaMundo() {
	fmt.Println("Hola Mundo")

	var result = operacion(5, 5, "+")
	fmt.Println("Esta es la suma : ", result)

	result = operacion(10, 5, "-")
	fmt.Println("Esta es la Resta : ", result)

	result = operacion(10, 5, "/")
	fmt.Println("Esta es la Division : ", result)

	result = operacion(10, 5, "*")
	fmt.Println("Esta es la Multiplicacion : ", result)
}

func operacion(n1 float32, n2 float32, op string) float32 {

	//var resultado float32 = 0.0

	if op == "+" {
		return n1 + n2
	} else if op == "-" {
		return n1 - n2
	} else if op == "*" {
		return n1 * n2
	} else if op == "/" {
		return n1 / n2
	} else {
		return 0.0
	}

}
