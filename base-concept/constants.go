package baseconcept

import (
	"fmt"
	"math"
)

func Constants() {
	// const is the keyword used and the value of the variable cannot be channged once assigned
	fmt.Println("chumma")
	// keyword const is used
	const fName = "Steve"
	// Declaring a group of constants
	const (
		name = "Bucky"
		age  = 106
	)
	fmt.Println(" values - ", fName, name, age)

	// The value of a constant should be known at compile time.
	// Hence it cannot be assigned to a value returned by a function call since the function call takes place at run time.
	var b = math.Sqrt(4) // allowed
	// const c := math.Sqrt(9) // not allowed
	fmt.Println("b - ", b)

	// String Constants, Typed and Untyped Constants

	const hello = "Hello!!!" // this is an untyped constant
	// untyped constants have a default type associated with them and they supply it if and only if a line of code demands it
	// like below
	var hello2 = hello
	fmt.Printf("var hello => type - %T, value - %v \n", hello2, hello2)

	const hello3 string = "Hellooowwww" // types constant
	fmt.Println("typed const - ", hello3)

	// Go is a strongly typed language. All variables require an explicit type.
	var defaultName = "Sam"
	type custString string
	var custName custString = "Sam" // The default type of the constant Sam is a string, so after the assignment defaultName is of type string.
	// custName = defaultName // this is not allowed since custString and type of defaultName does not match even though default type of defaultName is string
	// Even though we know that myString is an alias of string, Go's strong typing policy disallows variables of one type to be assigned to another.
	fmt.Println("default and custom strings - ", defaultName, custName)

	// Boolean Constants
	// he same rules for string constants apply to booleans
	const trueConst = true
	type custBool bool
	var defaultBool = trueConst
	var custAns custBool = trueConst // allowed bcz
	// defaultBool = custAns // not allowed
	fmt.Println("lil conf - ", custAns, defaultBool)

	// Integer constants
	const a = 5 //The default type of these kinds of constants can be thought of as being generated on the fly depending on the context
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	// Numeric Expressions
	// Numeric constants are free to be mixed and matched in expressions and a type is needed only when they are assigned to variables or used in any place in code which demands a type.

	var num = 8.8 / 3
	fmt.Printf("num value - %T, %v", num, num, "\n")
}
