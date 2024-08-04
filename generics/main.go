package main
import "fmt"

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints), // specify type arguments in brackets
		SumIntsOrFloats[string, float64](floats))
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 { // takes a map of string to int64 values
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 { // takes a map of string to float64 values
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

/* single generic function can receive a map containing either integer or float values, 
this will replace separate functions for int and float 

generic function should declare type parameters in addition to its ordinary function parameters. 
These type parameters make the function generic, enabling it to work with arguments of different types. 
we'll call the function with type arguments and ordinary function arguments.

Type parameter, will also have a meta-type, let's say how many types it can support.

In the calling code as well, we need to specify which type we will be using. 

Also, we need to make sure that type parameter supports all the operations in generic function.
a string parameter will fail in our case
*/

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.

/* K type parameter is type constraint comparable. 
the comparable constraint is predeclared in Go. 
It allows any type whose values may be used as an operand of the comparison operators == and !=. 
Go requires that map keys be comparable. 
declaring K as comparable is necessary so you can use K as the key in the map variable. 
It also ensures that calling code uses an allowable type for map keys.

V type parameter a constraint that is a union of two types: int64 and float64. 
Using | specifies a union of the two types, meaning that this constraint allows either type. 
Either type will be permitted by the compiler as an argument in the calling code.

m argument is of type map[K]V, where K and V are the types already specified for the type parameters. 
Note that we know map[K]V is a valid map type because K is a comparable type. 
If we hadn’t declared K comparable, the compiler would reject the reference to map[K]V.
*/
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	// two type parameters (inside the square brackets), K and V, and 
	// one argument that uses the type parameters, m of type map[K]V. 
	// The function returns a value of type V.

    var s V
    for _, v := range m {
        s += v
    }
    return s
}

/* when you run the program,
in each call the compiler replaced the type parameters with the concrete types specified in that call.

*/