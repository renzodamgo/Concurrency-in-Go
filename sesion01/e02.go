package main

import "fmt"

// global

func main(){
	// ambito local
	var x string = "Hola a todos";// forma 1
	fmt.Println(x);
	// Se asigna el valor de la variable como tipo de dato
	dato := 20 // forma 2: corta para declarar y inicializar
	fmt.Println(dato)
}