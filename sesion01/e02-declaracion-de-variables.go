package main

import "fmt"

// global

func main(){
	//gloval
	const igv float64 = 18.0

	// ambito local
	var x string = "Hola a todos";// forma 1
	fmt.Println(x);
	
	// Se asigna el valor de la variable como tipo de dato
	dato := 20 // forma 2: corta para declarar y inicializar
	fmt.Println("El valor de la Variable es:", dato)

	var(
		nombre string
		edad int
	)

	nombre = "Juan"
	edad = 25

	fmt.Println("El nombre es: ", nombre," Y la edad es:",edad)
}