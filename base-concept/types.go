package baseconcept

import (
	"fmt"
	"unsafe"
)

func Types() {
	// Following arethe types available in Go
	// bool
	// Numeric Types -
	// int8, int16, int32, int64, int - Signed intergers. int -> int32 or int 64 based on underlying platform (32bit system/64bit system)
	// uint8, uint16, uint32, uint64, uint - Unsigned intergers.
	// float32, float64
	// complex64, complex128
	// byte
	// rune
	// string

	// bool
	// A bool type represents a boolean and is either true or false.
	var a bool = true
	b := false
	c := a && b
	fmt.Println("boolean values - ", a, b, c)
	c = a || b
	fmt.Println("With or cond - ", c)

	// int
	// Always use int. Use signed int's only if it's required
	var i int = 21
	j := 1995
	fmt.Println("int values - ", i, j)

	// The type of a variable can be printed using %T format specifier in Printf function.
	// %d is used to print the size.
	fmt.Printf("type and size of i - %T , %d \n", i, unsafe.Sizeof(i))
	fmt.Printf("type and size of j - %T , %d \n", j, unsafe.Sizeof(j))

	// float
	// float64 is the default type for floating point values
	f1 := 5.264
	f2 := 10.5487
	fmt.Printf("type and size of f1 - %T, %d \n", f1, f1)
	fmt.Printf("type and size of f2 - %T, %d \n", f2, unsafe.Sizeof(f2))

	// complex64/128
	// float32/64with real and imaginary parts
	// func complex(r, i FloatType) ComplexType
	// both imaginary and real parts must be of sametype
	// float32 gives complex64 and float64 fives complex128
	// Shorthand syntax - c := 6 + 7i
	comp1 := complex(2, 3)
	comp2 := 4 + 5i
	fmt.Println("complex values - ", comp1, comp2)
	fmt.Println("Add, multi values - ", comp1+comp2, comp1*comp2)

	// byte is an alias of uint8
	// rune is alias of int32

	// string
	// Strings are a collection of bytes in Go
	fName := "James"
	lName := "Barnes"
	name := fName + " Buchanan " + lName
	fmt.Println("Name => ", name)

	// Type conversion
	//  Go is very strict about explicit typing

	// Below will end up in error
	// i1 := 55      // int
	// j1 := 67.8    // float64
	// sum1 := i1 + j1    // int + float64 not allowed. // This line should be sum1 := i1 + int(j1)
	// fmt.Println(sum1)

	intVal := 20
	var floatVal float64 = float64(intVal)
	fmt.Println("int val and converted int val to float - ", intVal, floatVal)
}
