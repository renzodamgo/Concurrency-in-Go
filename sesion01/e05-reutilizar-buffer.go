package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Ingrese su nombre: ")
	BufferIn := bufio.NewReader(os.Stdin)
	nombre, _ := BufferIn.ReadString('\n')
	nombre = strings.TrimRight(nombre, "\r\n")

	menu :=
		`
		*****Menu****
		[1] Pizza
		[2] Empanada
		Cual es tu pedido?
	`
	println("Bienvenido", nombre)
	println(menu)

	opc, _ := BufferIn.ReadString('\n')
	opc = strings.TrimRight(opc, "\r\n")

	switch opc{
	case "1":
		fmt.Println("Usted eligió Pizza")
	case "2":
		fmt.Println("Usted eligió Empanada")
	}

	fmt.Println("Gracias por su preferencia")
}
