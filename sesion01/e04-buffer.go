package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Print("Ingrese un número: ")
	// NewReader: Nueva Instancia de Lectura
	// Stdin: Leer por consola
	// ReadString: Se lee hasta un salto de linea (\n) y retorno de carga (\r)
	bufferIn := bufio.NewReader(os.Stdin)
	str, err :=bufferIn.ReadString('\n')

	if err != nil{
		fmt.Println("Error: ",err.Error())
		// Terminar el programa con un error
		os.Exit(1)
	}
	// Limpiar \r\n de "15\r\n" a "15"
	str = strings.TrimRight(str,"\r\n") 
	// fmt.Print(str)
	num, err := strconv.Atoi(str)

	if err != nil{
		fmt.Println("Error: ",err.Error())
		// Terminar el programa con un error
		os.Exit(1)
	}

	fmt.Println("El doble de: ",num, " es: ", num*2)
}