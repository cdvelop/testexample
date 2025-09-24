package testexample

import "fmt"

// Reverse retorna el string invertido.
// Ejemplo: Reverse("hola") == "aloh"
func Reverse(s string) string {
	// Convertimos a runas para soportar caracteres Unicode correctamente
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	// Pruebas
	fmt.Println(Reverse("hola"))  // aloh
	fmt.Println(Reverse("mundo")) // odnum
	fmt.Println(Reverse("¡Go!"))  // !oG¡
	fmt.Println(Reverse("ábaco")) // ocabá
}
