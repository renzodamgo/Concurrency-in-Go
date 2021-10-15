package main

import (
	"fmt"
	"time"
)

func saludar(){
	fmt.Println("Holasss")
}

func main(){
	
	// No muestra nada pq en una linea de tiempo del proceso principal porque el proceso principal gano a todos
	go saludar()
	time.Sleep(time.Second) // Espere 1 segundo para que el proceso principal espere y consiga que se salude
}