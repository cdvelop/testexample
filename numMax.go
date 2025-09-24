package testexample

import "fmt"

func NumMax(a int, b int) int {
	if a > b {
		return a
	} else if a < b {
		return b
	} else {
		fmt.Println("Los numeros son iguales")
		return 0
	}
}
