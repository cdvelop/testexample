package testexample

import "fmt"

func Max() {
	a := 42
	b := 27

	maximum := max(a, b)
	fmt.Printf("El max de %d y %d es %d\n", a, b, maximum)

}
