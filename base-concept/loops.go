package baseconcept

import (
	"fmt"
)

// for loo syntax
// for initialisation; condition; post {
// }
// for is the only loop available in Go.
// Go doesn't have while or do while loops which are present in other languages like C
// All the three components namely initialisation, condition and post are optional in Go
func ForLoop() {
	for i := 0; i < 11; i++ {
		if i > 5 {
			break // loop terminates if i is greater than 5
		}
		fmt.Printf("i = %d \n", i)
	}
	fmt.Println("Loop ended")
}

// continue
// The continue statement is used to skip the current iteration of the for loop.
// All code present in a for loop after the continue statement will not be executed for the current iteration.
// The loop will move on to the next iteration.
func ForLoopWithContinue() {
	for i := 0; i < 11; i++ {
		if i < 5 {
			continue // skips next lines if i < 5
		}
		fmt.Printf("i = %d \n", i)
	}
	fmt.Println("Loop ended")
}

// Nested for loops
func NestedForLoop() {
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("#")
	}
	// Example()
	// CorrectLoopWithLabels()
	// ForOnlyWithCondition()
	// ForWithMultiVars()
	ForInfinite()
}

// Labels
// Q: What if we want to stop printing when i and j are equal,
// then we need to break out from outer loop from inner loop,
// bcz both i and j will be available only insde the inner loop

// Try this, but this break will only break out form inner loop
func Example() {
	for i := 0; i < 3; i++ {
		for j := 1; j < 5; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break
			}
		}

	}
}

// Below one helps to break out outer
func CorrectLoopWithLabels() {
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 5; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}
}

// More examples
func ForOnlyWithCondition() {
	i := 0
	for i < 10 { // initialisation and post are omitted
		fmt.Printf("%d ", i)
		i += 2
	}
}

// Multiple variables in loop
func ForWithMultiVars() {
	for num, i := 11, 1; num < 21 && i < 11; num, i = num+1, i+1 {
		fmt.Printf("%d * %d = %d\n", num, i, num*i)
	}
}

// infinite loop
// The syntax for creating an infinite loop is,
// for {
// }
func ForInfinite() {
	for {
		fmt.Print("Hi!")
	}
}
