package main

import (
	"fmt"
	"time"
)


var n int
func pq(){
	var temp int;
	for i := 0; i<10 ; i++{
		temp=n
		n=temp+1
	}
}

func main(){
	go pq()
	go pq()

	time.Sleep(time.Second)
	fmt.Println(n)
}