package baseconcept

import (
	"fmt"
)

// An array is a collection of elements that belong to the same type.

// Mixing values of different types, for example,
// an array that contains both strings and integers is not allowed in Go.
func Arrays() {
	fmt.Println("In arrays")
	var a [3]int //int array with length 3
	fmt.Println(a)
	// The index of an array starts from 0 and ends at length - 1, like JavaScript
	a[0] = 12
	fmt.Println(a)

	b := [3]int{123, 1, 2} // short hand declaration to create array
	fmt.Println(b)

	// It is not necessary that all elements in an array have to be assigned a value during short hand declaration.
	c := [3]int{234}
	fmt.Println(c)

	// You can even ignore the length of the array in the declaration and replace it with ...
	// and let the compiler find the length for you.
	d := [...]int{345} // ... makes the compiler determine the length
	// d[1] = 10 // errors out as invalid array index 1 (out of bounds for 1-element array)go
	fmt.Println(d)
	// The size of the array is a part of the type
	// Because of this, arrays cannot be resized.

	// e := [3]int{5, 78, 8}
	// var f [5]int
	// f = e //not possible since [3]int and [5]int are distinct types

	// Length of an array
	fmt.Println("Length of array d - ", len(d))
}

// Arrays are value types
// Arrays in Go are value types and not reference types.
// This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable.
// If changes are made to the new variable, it will not be reflected in the original array.
func ArrayAreValueTypes() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

	// Similarly when arrays are passed to functions as parameters,
	// they are passed by value and the original array is unchanged.
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
}
func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

// Iterating arrays
func IteratingArrays() {
	a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}

	// Go provides a better and concise way to iterate over an array by using the range form of the for loop.
	// range returns both the index and the value at that index -- like map or forEach in javascript
	sum := 0.0
	for index, value := range a {
		fmt.Printf("%d - %.2f\n", index, value)
		sum += value
	}
	fmt.Printf("sum value %f", sum)

	// In case you want only the value and want to ignore the index,
	// you can do this by replacing the index with the _ blank identifier.
	// for _, v := range a { //ignores index
	// }

}

// Multidimensional arrays
func MultiDimensionalArrays() {
	a := [3][2]string{
		{"tiger", "cat-family"},
		{"lion", "cat-family"},
		{"cheetah", "cat-family"},
		// this tailing comma is necessary. The compiler will complain if you omit this comma
		// This is because of the fact that the lexer automatically inserts semicolons according to simple rules.
		// Please read https://golang.org/doc/effective_go.html#semicolons to know more on this
	}
	printArray(a)

	var b [3][2]string
	b[0][0] = "tiger"
	b[0][1] = "wild-animal"
	b[1][0] = "tiger"
	b[1][1] = "wild-animal"
	printArray(b)
}
func printArray(array [3][2]string) {
	for _, valueSet := range array {
		for _, value := range valueSet {
			fmt.Printf("%s -", value)
		}
		fmt.Printf("\n")
	}
}

// Although arrays seem to be flexible enough, they come with the restriction that they are of fixed length.
// It is not possible to increase the length of an array.
// This is where slices come into the picture
