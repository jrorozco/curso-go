package main

import (
	"fmt"
	"time"
)

func main() {
	momento := time.Now()
	hoy := momento.Weekday()
	switch hoy {
	case 0:
		fmt.Println("Hoy es domingo")
	case 1:
		fmt.Println("Hoy es Lunes")
	case 2:
		fmt.Println("Hoy es Martes")
	case 3:
		fmt.Println("Hoy es Miercoles")
	case 4:
		fmt.Println("Hoy es Jueves")
	case 5:
		fmt.Println("Hoy es Viernes")
	case 6:
		fmt.Println("Hoy es Sabado")
	default:
		fmt.Println("Hoy es otro dia de la semana")
	}

}
