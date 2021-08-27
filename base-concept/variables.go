package baseconcept

import (
	"fmt"
 	"math"
)

func Variables(){
	fmt.Println("Helooooo makkaley!")

	// Declaring single variable
	// var <name> <type>
	var num int 
	fmt.Println("Num is - ", num)
	num = 12
	fmt.Println("Num value assigned - ", num)

	// Declaring single variblae with initial value
	var num2 int = 15
	fmt.Println("Inialised value - ", num2)

	// Type inference
	//  var <name>
	var num3 = 19
	fmt.Println("Example of type inference - ", num3)

	// Declaring multiple variables
	// var <name1> <name2> <type> = <initialval1>, <initialval2>
	var length, width int = 150, 30
	fmt.Println("Length = ", length, "Width = ", width)

	var length2, width2 = 160, 40
	fmt.Println("Length2 = ", length2, "Width2 = ", width2)

	var length3, width3 int
	fmt.Println("Length and width 3 are ", length3, width3)
	length3 = 170
	width3 = 50
	fmt.Println("Assigned values of length and width 3 are ", length3, width3)

	// Declaring multiple variables of different types
	// var (
	// 		<name1> = <initialvalue1>
	// 		<name2> = <initialvalue2> 
	// )
	var (
		name = "Buck"
		age = 106
		crimeScore int
	)
	fmt.Println("Name, age and crimeScore of Mr.",name, age, crimeScore)

	// Short hand declaration
	// <name> := <initialvalue>
	// Short hand declaration requires initial values for all variables on the left-hand side of the assignment
	anotherName := "Steven"
	anotherAge, anotherCrimeScore := 105, 200
	fmt.Println("Another name, age and crimscore of Mr.",anotherName, anotherAge, anotherCrimeScore)

	// Short hand syntax can only be used when at least one of the variables on the left side of := is newly declared
	newName, anotherName, name := "Cap", "Steve", "Bucky"
	fmt.Println("Changes name with a new name", anotherName, name, newName)

	// Since Go is strongly typed, variables declared as belonging to one type cannot be assigned a value of another type
	// newName is of type string and cannot take a value of another type
	// newName = 20 will give error


	// Variables can also be assigned values which are computed during run time
	a, b := 20.98, 20.97777
	c := math.Min(a, b)
	fmt.Println("Min value - ", c)
}