package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)

	edad, err := strconv.Atoi(os.Args[2]) // convertir string a entero ejeplo : go run PasarArgumentos.go Rene 40
	if err != nil {
		fmt.Println("error", err)
	}
	if (edad >= 18 || edad == 17) && edad != 25 && edad != 99 {
		fmt.Println("Eres mayor de edad o tienes 17")
	} else if edad == 25 {
		fmt.Println("Tienes 25 años")
	} else if edad == 99 {
		fmt.Println("Pronto moriras")
	} else {
		fmt.Println("Eres mayor de edad")
	}
	fmt.Println("******* SABER SI EL NUMERO ES PAR ********")
	numero, err := strconv.Atoi(os.Args[2]) // convertir string a entero ejeplo : go run PasarArgumentos.go Rene 40
	if err != nil {
		fmt.Println("Error ", err)
	}

	if numero%2 == 0 {
		fmt.Println("El numero es Par ")
	} else {
		fmt.Println("El numero no es Par ")
	}
}
