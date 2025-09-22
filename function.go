package testexample

import "fmt"

func GenerarSaludo(nombre string) string {

	fmt.Println("Generando saludo para " + nombre)

	return "Hello, " + nombre + "!"
}
