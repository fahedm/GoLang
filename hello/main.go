package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello from main package")
	// this is single line comment
	/*
		this
		is multi line comment
	*/
	var a int = 1029
	var b string = "fahed learning"
	var c float32 = 7.25
	fmt.Println("int", a, "string", b, "float", c)

	x := 129
	y := "type will automatcially be taken"
	z := 6.95
	fmt.Println("int", x, "string", y, "float", z)

	var a1, b1, c1, d1 int = 1, 2, 3, 4
	fmt.Println("variable declared in single line with one type", a1, b1, c1, d1)

	var a2, b2 = 6, "Hello"
	c2, d2 := 7, "World!"
	fmt.Println("variable declared in single line with different types", a2, b2, c2, d2)

}
