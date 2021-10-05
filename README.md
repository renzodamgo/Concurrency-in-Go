# Go
Correr un programa go
```go
go run sesion01/e01.go
```

# Inicializar:
Inicializar con el tipo de dato:
```go
var x string = "Hola a todos";
```
Inicializar asignando el tipo de dato de la variable:
```go
//Se le asigna el tipo int por que la variable es int
x := 20;
```
Asignar varias variables:
```go
var(
  nombre string
  edad int
)

nombre = "Juan"
edad = 25
```
Inicializar constantes:
```go
const igv float64 = 18.0
```

# Ingresar datos
Forma 1:
```go
fmt.Print("Ingrese un número: ")
var num float64
//%f: Formatear a tipo flotante
//&num: Asigna a la direccion de memoria
fmt.Scanf("%f", &num)
```
