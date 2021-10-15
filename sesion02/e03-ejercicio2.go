package main

import (
	"fmt"
	"time"
)

var n int

func pq(){
	var temp int
	temp = n
	time.Sleep(time.Second)
	n= temp + 1
}

func main(){
	go pq()
	go pq()

	time.Sleep(time.Second)
	fmt.Println(n)
}