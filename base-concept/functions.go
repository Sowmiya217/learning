package baseconcept

import (
	"fmt"
)

// Functions - syntax
// func functionname(parametername type) returntype {
// 	//function body
//    }

func Functions() {

	billAmount := calculateBill(15, 3)
	fmt.Println(" Bill - ", billAmount)

	area, perimeter := rectProps(10.2, 8.6)
	fmt.Printf("Area - %v, Perimeter - %v \n", area, perimeter)

	// Blank Identifier
	// _ is known as the blank identifier in Go. It can be used in place of any value of any type
	area2, _ := namedRectProps(20.8, 15.7) // _ identifier is used to discard the perimeter
	fmt.Println("Area - ", area2)
}

// If consecutive parameters are of the same type,
// we can avoid writing the type each time and it is enough to be written once at the end
func calculateBill(price, nos int) int {
	return price * nos
}

// Multiple return values
type float float64

func rectProps(length, width float) (float, float) {
	area := length * width
	perimeter := 2 * (length + width)
	return area, perimeter
}

// Named return values
// It is possible to return named values from a function
// If a return value is named, it can be considered as being declared as a variable in the first line of the function

func namedRectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // no need to return explicitly since return type and values are mentioned (named return values) in func declaration
}
