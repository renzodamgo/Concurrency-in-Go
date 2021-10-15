package main

import "fmt"

func p(c chan int)  {
	c <-5
}

func main() {
	c := make(chan int)

	go p(c)
	go p(c)

	// Si va a recibir una lectura
	n:= <-c // Out 

	q:= <-c // Out Si llega
	// f:= <-c // Out No va a llegar nunca !deadlock

	fmt.Println(n,q)
}