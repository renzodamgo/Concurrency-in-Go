package main

import "fmt"

func main() {
	fmt.Print("Ingrese un número: ")
	var num float64
	fmt.Scanf("%f", &num) //lectura de consola

	doble := num * 2
	fmt.Println("El doble de ", num, " es ", doble)
}
