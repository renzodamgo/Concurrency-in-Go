package main

import "fmt"

func main(){
	//declaración de un arreglo
	arreglo := [7]int{15,3,2,7,8,9,1}

	//manejo del for
	for i, v := range arreglo{
		fmt.Printf("El valor de v es %d en la posición #%d\n",v,i)
	}

	//for infinito como while(true)
	i := 0
	for {
		fmt.Println(i)
		if i == 4{
			//Salir del Bucle
			break
		}
		i++
		
	}
	
	// for con condición while i<10
	for i<10{
		fmt.Println("Continua")
		fmt.Println(i)
		i++
	}
	// For clásico
	for i:=1; i<=10; i++{
		if i%2 == 0{
			fmt.Println(i)
		}
	}
}