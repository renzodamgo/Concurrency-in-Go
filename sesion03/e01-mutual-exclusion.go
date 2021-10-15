package main

import "fmt"

var turno int = 1

func p()  {
	for{
		fmt.Println("Linea01 SNC P")
		fmt.Println("Linea02 SNC P")
		if(turno != 1){
			//espera
		}
		fmt.Println("Linea01 SC P")
		fmt.Println("Linea01 SC P")
		// Post Proceso
		turno = 2
	}
}

func q()  {
	for{
		fmt.Println("Linea01 SNC Q")
		fmt.Println("Linea02 SNC Q")
		if(turno != 2){
			//espera
		}
		fmt.Println("Linea01 SC Q")
		fmt.Println("Linea01 SC Q")
		// Post Proceso
		turno = 1
	}
}

// Puede haber starvation
func main()  {
	go p()
	q()
}