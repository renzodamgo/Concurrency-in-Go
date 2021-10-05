// Funciones
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Print("Ingres un número: ")
	bufferIn := bufio.NewReader(os.Stdin)
	dato , _ := bufferIn.ReadString('\n')
	dato = strings.TrimRight(dato, "\r\n")
	num, _ := strconv.Atoi(dato)
	evaluar(num)

}

func evaluar(num int){
	if num%2 == 0{
		fmt.Println("El numero ", num,"es dvisible entre 2")
	}
	if num%3 == 0{
		fmt.Println("El numero ", num,"es dvisible entre 3")
	}
}