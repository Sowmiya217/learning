package baseconcept

import (
	"fmt"
	"math/rand"
)

// A switch is a conditional statement that evaluates an expression and compares it
// against a list of possible matches and executes the corresponding block of code.
// It can be considered as an idiomatic way of replacing complex if else clauses.
func Switch() {
	numberType := "odd"
	switch numberType { // can declare the variable here itself, like switch numberType:="odd"; numberType
	case "even":
		fmt.Println("Even")
	case "odd": // can have multiple expressions here
		fmt.Println("Odd")
	default: // default case and not required
		fmt.Println("It is Zero, neither odd nor even")
	}
	// ExpressionlessSwitch()
	// SwitchFallThorugh()
	// BreakingSwitch()
	BreakingOuterForLoop()
}

// Duplicate cases with the same constant value are not allowed

// Expressionless switch
// The expression in a switch is optional and it can be omitted.
// If the expression is omitted, the switch is considered to be switch true
// and each of the case expression is evaluated for truth
// and the corresponding block of code is executed.
func ExpressionlessSwitch() {
	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less than 100", num)
	case num >= 101:
		fmt.Printf("%d is greater than 100", num)
	}

}

// Fallthrough
// In Go, the control comes out of the switch statement immediately after a case is executed.
// A fallthrough statement is used to transfer control to the first statement of the case
// that is present immediately after the case which has been executed.

func number() int {
	num := 25 * 5
	return num
}
func SwitchFallThorugh() {
	switch num := number(); { // Switch and case expressions need not be only constants. They can be evaluated at runtime too.
	case num > 200:
		fmt.Printf("%d is greater than 200\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 50:
		fmt.Printf("%d is lesser than 50", num)
	}
}

// Fallthrough will happen even when the case evaluates to false

// fallthrough should be the last statement in a case.
// If it is present somewhere in the middle, the compiler will complain that fallthrough statement out of place.

// One more thing is fallthrough cannot be used in the last case of a switch
// since there are no more cases to fallthrough

// Breaking switch
func BreakingSwitch() {
	switch num := -5; {
	case num < 50:
		if num < 0 {
			break
		}
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}

func BreakingOuterForLoop() { // using label
randloop:
	for {
		switch i := rand.Intn(100); {
		case i%2 == 0:
			fmt.Printf("Generated even number %d", i)
			break randloop
		default:
			fmt.Println("Generated number : ", i)
		}
	}
}
