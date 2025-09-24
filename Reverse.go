package testexample

import "fmt"

// function to reverse string
func reverseString(str string) (result string) {
	// iterate over str and prepend to result
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func reverse() {
	str := "hola"

	fmt.Println(str)
	// invoke reverseString
	fmt.Println(reverseString(str))
}
