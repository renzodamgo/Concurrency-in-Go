package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//estructura Si no tiene letra capital no te devuelve el JSON
type Alumno struct {
	Codigo string `json:"cod"`
	Nombre string `json:"nom"`
	Dni    int    `json:"dni"`
}

var listaAlumnos []Alumno

func cargarAlumnos() {
	listaAlumnos = []Alumno{
		{"20211862", "Juan Merino", 40127558},
		{"20212342", "Miguel Mejía", 40127558},
		{"20216532", "Maria Gonzales", 40565454},
	}
}

func mostrarHome(response http.ResponseWriter, request *http.Request) {
	log.Println("Ingresa a endpoint home")
	response.Header().Set("Content-Type", "text/html") //cabecera del mensaje
	//cuerpo del mensaje
	io.WriteString(response, `
		<html>
			<head></head>
			</body><h2>Mi API de Alumnos</h2></body>
		</html>
	`)
}

func listarAlumnos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json") // contenido del mensaje json
	//codificar
	jsonBytes, _ := json.MarshalIndent(listaAlumnos, "", " ")
	io.WriteString(response, string(jsonBytes))

}

func buscarAlumno(response http.ResponseWriter, request *http.Request) {
	//recuperación de parámetros
	cod := request.FormValue("codigo") // http://localhost:9000/buscar?codigo=20212342
	log.Println(cod)

	response.Header().Set("Content-Type", "application/json")
	//logica de busqueda
	for _, oAlumno := range listaAlumnos {
		if oAlumno.Codigo == cod {
			jsonBytes, _ := json.MarshalIndent(oAlumno, "", "")
			io.WriteString(response, string(jsonBytes))
		}
	}
}

func agregarAlumno(response http.ResponseWriter, request *http.Request) {
	// Validar que la posicion se realiza por el método POST
	if request.Method == "POST" {
		if request.Header.Get("Content-Type") == "application/json" {
			ioutil.ReadAll(request.Body)
		} else {
			http.Error(response, "Contenido Invállido", http.StatusBadRequest)
		}
	} else {
		//lanzar error
		http.Error(response, "Método Inválido", http.StatusMethodNotAllowed)
	}
}

func manejadorSolicitudes() {
	mux := http.NewServeMux()
	//endpoint
	mux.HandleFunc("/home", mostrarHome)
	mux.HandleFunc("/listar", listarAlumnos)
	mux.HandleFunc("/buscar", buscarAlumno)
	mux.HandleFunc("/agregar", agregarAlumno)

	//puerto
	log.Fatal(http.ListenAndServe(":9000", mux))
}
func main() {
	cargarAlumnos()
	manejadorSolicitudes()
}
