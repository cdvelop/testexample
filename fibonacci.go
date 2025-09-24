package testexample

import "fmt"

func Fibonacci(num int) int {

	var numUno int = 0
	var numDos int = 1
	for i := 0; i < num; i++ {
		numPrint := numUno + numDos
		numUno = numDos
		numDos = numPrint
	}

	num = numDos

	fmt.Println(num)

	return num
}
