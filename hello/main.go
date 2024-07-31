package main

import (
	"fmt"
	"log"

	"example.com/greetings"
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

	const CONST_NAME_TYPED string = "this is constant variable that will be always same"
	const CONST_VAR_UNTYPED = 2
	fmt.Println("constant variables ", CONST_NAME_TYPED, CONST_VAR_UNTYPED)

	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	var _bool bool = true      // Boolean
	var _int int = 5           // Integer
	var _float float32 = 3.14  // Floating point number
	var _string string = "Hi!" // String

	fmt.Println("Boolean: ", _bool)
	fmt.Println("Integer: ", _int)
	fmt.Println("Float:   ", _float)
	fmt.Println("String:  ", _string)

	fmt.Println("Boolean value will be by default false")
	var bool1 bool = true // typed declaration with initial value
	var bool2 = true      // untyped declaration with initial value
	var bool3 bool        // typed declaration without initial value
	bool4 := true         // untyped declaration with initial value

	fmt.Println(bool1) // Returns true
	fmt.Println(bool2) // Returns true
	fmt.Println(bool3) // Returns false
	fmt.Println(bool4) // Returns true

	var int_x uint = 500
	var uint_y uint = 4500

	fmt.Printf("Type: %T, value: %v \n", int_x, int_x)
	fmt.Printf("Type: %T, value: %v", uint_y, uint_y)

	var float_x float64 = 1.7e+308
	var float_y float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", float_x, float_x)
	fmt.Printf("Type: %T, value: %v\n", float_y, float_y)

	//GoLang Documentaion Tutorial

	//get greeting message and print it
	message := greetings.Hello("Fahed")
	fmt.Println(message)

	//set properties of predefined logger, including 
	//the log entry prefix and a flag to diable printing
	//the time, source file and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//request a greeting message.
	message, err := greetings.Hello_with_Exception("")
	//if error is return print it to console and exit the program

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
