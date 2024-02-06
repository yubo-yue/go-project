package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	showBasicType()
	explictTypeConv()
	initArrayVar()
}

func showBasicType() {
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

	fmt.Println("Show complex type")
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x / y)
	fmt.Println(x * y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))

	var first rune = 'J'
	fmt.Println(first)
}

func explictTypeConv() {
	var x int = 10
	var y float64 = 30.2

	var sum1 float64 = float64(x) + y
	var sum2 int = x + int(y)
	fmt.Println(sum1, ",", sum2)
}

func initArrayVar() {
	var x [3]int
	var y = [3]int{10, 20, 30}
	//sparse array
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	var k = [...]int{4, 5, 6}
	var j = [3]int{4, 5, 6}

	fmt.Println(x, y, z, k)
	fmt.Println("k == j is ", k == j)

}
