package baseconcept

import "fmt"

// if is a statement that has a boolean condition and it executes a block of code if that condition evaluates to true.
// It executes an alternate else block if the condition evaluates to false.
// syntax -
// if condition {
// }
// Unlike in other languages like C, the braces {  } are mandatory even if there is only one line of code between the braces{  }

func IfFunc() {
	fmt.Println("I'm in if func")
	num := 10
	if num%2 == 0 {
		fmt.Println("The number", num, "is even")
		return
	}
	fmt.Println("The number", num, "is odd")
}

//  If else syntax --
// if condition {
// } else {
// }
func IfElseFunc() {
	num := 10
	if num%2 == 0 {
		fmt.Println("The number", num, "is even")
	} else {
		fmt.Println("The number", num, "is odd")
	}
}

// If.. else if... else syntax --
// if condition1 {
// ...
// } else if condition2 {
// ...
// } else {
// ...
// }
// In general, whichever if or else if's condition evaluates to true, it's corresponding code block is executed.
// If none of the conditions are true then else block is executed.
func IfElseIfElseFunc() {
	num := 10
	if num == 0 {
		fmt.Println("The number", num, "is zero")
	} else if num%2 == 0 {
		fmt.Println("The number", num, "is even")
	} else {
		fmt.Println("The number", num, "is odd")
	}
}

// If with assignment :
// There is one more variant of if which includes an optional shorthand assignment statement
// that is executed before the condition is evaluated.
// Its syntax is
// if assignment-statement; condition {
// }
func IfWithStatement() {
	// One thing to be noted is that num is available only for access from inside the if and else.
	// i.e. the scope of num is limited to the if else blocks.
	// If we try to access num from outside the if or else, the compiler will complain
	if num := 10; num%2 == 0 {
		fmt.Println("The number", num, "is even")
	} else {
		fmt.Println("The number", num, "is odd")
	}
}

// Syntax requirement
// The else statement should start in the same line after the closing curly brace } of the if statement.
// If not the compiler will complain.
// if(flag == true){
// }
// else{
// } -- this would fail since go compiler inserts ; after } in the code

// Idiomatic Go
// In Go's philosophy, it is better to avoid unnecessary branches and indentation of code.
// It is also considered better to return as early as possible.
// Example -
// In the above, IfElseFunc and IfFunc are doing the same thing, but IfElse has an extra branch which seems unneccessary,
// bcz once if is excuted, we can straight away return from the func and else becomes an extra code branch
