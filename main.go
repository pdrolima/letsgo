// Executable program start from the main package
package main

import (
	"fmt"
	m "math"
	"reflect"
)

// The main function is the starting point of the program
func main() {

	fmt.Println("Hello, World!")

	const PI float64 = 3.1415
	var radius = 3.2 // Type inference (float64)

	// := is the short assignment and operator and it is used to declare and assign variables in one line, it also infer the type of the variable
	// = is the assignment operator
	// A declared variable must be used or it will throw an error
	// More examples: https://stackoverflow.com/a/45654233/8273185
	area := m.Pi * m.Pow(radius, 2)
	area = 1 // you can reassign the value of a variable
	// area := 1 // you can't redeclare a variable with the same name

	// Print the area of the circle
	fmt.Printf("Area of the circle: %f\n", area)

	// Other ways to declare constants
	const (
		first  = "First"
		second = 2
		third  = 3.0
	)

	// It is possible to declare multiple variables in one line
	var (
		a = 1
		b = 2
		c = 3
	)

	fmt.Println(a, b, c)

	var d, e bool = true, false
	f, g, h := 4, false, "Hello"

	fmt.Println(d, e)
	fmt.Println(f, g, h)

	// Ways of printing
	x := 3.1415
	// fmt.Println("The value of x is: " + x) // You can concatenate a float64 with a string
	xs := fmt.Sprint(x) // Convert a float64 to a string
	fmt.Println("The value of x is: " + xs)

	// You can use the %v to print a variable
	fmt.Printf("The value of x is: %v\n", x)
	// You can use .2f to print a float64 with 2 decimal places
	fmt.Printf("The value of x is: %.2f\n", x)

	number := 1
	floating := 3.1415
	text := "Hello"
	boolean := true
	// %d is for integer
	// %f is for float64
	// %s is for string
	// %t is for boolean
	// %v is for different types
	fmt.Printf("\n%d %1.f %s %t", number, floating, text, boolean)

	// Basic types in Go
	// https://golang.org/ref/spec#Types
	// https://golang.org/ref/spec#Numeric_types
	// https://golang.org/ref/spec#String_types
	// https://golang.org/ref/spec#Boolean_types
	// reflect package is used to get the type of a variable
	newNumber := 100000
	// uint8 is an unsigned 8-bit integer (0 to 255)
	// uint16 is an unsigned 16-bit integer (0 to 65535)
	// uint32 is an unsigned 32-bit integer (0 to 4294967295)
	// uint64 is an unsigned 64-bit integer (0 to 18446744073709551615)
	fmt.Printf("\nThe type of number is: %v", reflect.TypeOf(newNumber))

	var cls byte = 255 // byte is an alias for uint8
	fmt.Printf("\nThe type of b is: %v", reflect.TypeOf(cls))

	// signed integers
	// int8 is a signed 8-bit integer (-128 to 127)
	// int16 is a signed 16-bit integer (-32768 to 32767)
	// int32 is a signed 32-bit integer (-2147483648 to 2147483647)
	// int64 is a signed 64-bit integer (-9223372036854775808 to 9223372036854775807)
	// int is an alias for int32
	// int is the default integer type in Go
	i1 := m.MaxInt64
	fmt.Printf("\nThe max value of int is: %v", i1)
	fmt.Printf("\nThe type of i1 is: %v", reflect.TypeOf(i1))

	var i2 rune = 'a' // rune is a mapping of int32 to a Unicode code point
	fmt.Printf("\nThe type of i2 is: %v", reflect.TypeOf(i2))

	// Real numbers
	var reais float32 = 1.2345
	fmt.Printf("\nThe type of reais is: %v", reflect.TypeOf(reais))

	// boolean
	bo := true
	fmt.Printf("\nThe type of bo is: %v\n", reflect.TypeOf(bo))
	fmt.Println(!bo)
}
