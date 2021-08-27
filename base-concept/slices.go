package baseconcept

import (
	"fmt"
)

// A slice is a convenient, flexible and powerful wrapper on top of an array.
// Slices do not own any data on their own. They are just references to existing arrays.
func Slices() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array - ", array)
	var slice []int = array[1:4] // creates a slice from a[1] to a[3], excluding given last index, like [1, 4) --> start to end-1
	fmt.Println("Slice - ", slice)

	// Another way to create
	c := []int{1, 2, 3} // //creates and array and returns a slice reference, which is c
	fmt.Println(c)
}

// A slice does not own any data of its own.
// It is just a representation of the underlying array.
// Any modifications done to the slice will be reflected in the underlying array.
func ModifyingSlice() {
	arr1 := [...]int{11, 12, 13, 14, 15, 16, 17}
	sl1 := arr1[1:]
	fmt.Println("array before", arr1)
	for i, _ := range sl1 {
		sl1[i]++ // v++ did not work
	}
	fmt.Println("change 1 - array after update to sl1", arr1)

	// When a number of slices share the same underlying array,
	// the changes that each one makes will be reflected in the array.
	sl2 := arr1[:] //creates a slice which contains all elements of the array
	sl2[0] = 22
	fmt.Println("change 2 - array after update to sl2", arr1)
}

// length and capacity of a slice
// The length of the slice is the number of elements in the slice.
// The capacity of the slice is the number of elements in the underlying array,
// starting from the index from which the slice is created.
func LengthAndCapacityOfSlices() {
	animalArr := [...]string{"cat", "dog", "wolf", "lion", "goat", "cow", "pig"} // this array length is 7
	slice := animalArr[1:4]
	fmt.Println("slice - ", slice)
	fmt.Printf("length of slice %d capacity %d \n", len(slice), cap(slice)) //length of slice is 3 and capacity is 6

	// A slice can be re-sliced upto its capacity.
	// Anything beyond that will cause the program to throw a run time error.
	slice = slice[:cap(slice)] //re-slicing slice till its capacity
	fmt.Println("slice after re-slicing - ", slice)
	fmt.Println("After re-slicing length is", len(slice), "and capacity is", cap(slice))

}

// creating a slice using make
// func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity.
// The capacity parameter is optional and defaults to the length.
// The make function creates an array and returns a slice reference to it.

func SliceUsingMake() {
	sl := make([]int, 3, 5)
	fmt.Println("length - ", len(sl), " capacity - ", cap(sl), " values - ", sl)
}
