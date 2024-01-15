package main

import "fmt"

func main() {
	//test literal
	a := 1_234
	fmt.Println(a)
	b := 12e23
	fmt.Println(b)
	octet := '\141'
	fmt.Println(octet)

	s := "String literal"
	fmt.Println(s)

	rawStr := `abc
\def
""""hgg`

	fmt.Println(rawStr)
}
