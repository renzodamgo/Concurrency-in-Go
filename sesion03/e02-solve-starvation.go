package main

import "fmt"

var wantp bool = false
var wantq bool = true

func p()  {
	for{
		fmt.Println("Linea01 SNC P")
		fmt.Println("Linea02 SNC P")
		for wantq{
			//espera
		}
		wantp = true
		fmt.Println("Linea 01 SC P")
		fmt.Println("Linea 01 SC P")
		wantp = false
	}
}

func q()  {
	for{
		fmt.Println("Linea01 SNC Q")
		fmt.Println("Linea02 SNC Q")
		for wantp{
			//espera
		}
		wantq = true
		fmt.Println("Linea 01 SC Q")
		fmt.Println("Linea 02 SC Q")
		wantq = false
	}
}

// Ocurriria un deadlock
func main()  {
	go p()
	q()
}