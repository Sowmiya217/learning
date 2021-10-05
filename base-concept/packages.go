package baseconcept

import (
	"fmt"
	simpleinterest "learnmod/simple-interest"
	"log"
)

// Using blank an identitfier to mute the error when imported package is not used
var _ = simpleinterest.Calculate // error silencer
// it is recommended to write error silencers in the package level just after the import statement.

var principal, interest, tenure = 1000.0, 5.25, 2.0

// a package can have any no of init funcs in same file or diff files in that package and they are executed in the order they are persented to the compiler
// imported packages are initialized first and it'll be initialized only once regardless of the no of packges it is imported in

func init() {
	fmt.Println("Initializing packages file")
	if principal < 0 {
		log.Fatal("Principal amount should be greater than 0")
	}
	if interest < 0 {
		log.Fatal("Interest percentage should be greater than 0")
	}
	if tenure < 0 {
		log.Fatal("Tenure period should be greater than 0")
	}
}

func Packages() {
	interestAmount := simpleinterest.Calculate(principal, tenure, interest)
	fmt.Println("interest amount - ", interestAmount)
}

// Packages
// Packages are used to organize Go source code for better reusability and readability.
// Packages are a collection of Go sources files that reside in the same directory.
// Packages provide code compartmentalization and hence it becomes easy to maintain Go projects.

// package packagename specifies that a particular source file belongs to package packagename.
// This should be the first line of every go source file.

// Go Modules
// A Go Module is nothing but a collection of Go packages.
// Now this question might come to your mind.
// Why do we need Go modules to create a custom package?
// The answer is the import path for the custom package we create is derived from the name of the go module.
// In addition to this, all the other third-party packages(such as source code from github)
// which our application uses will be present in the go.mod file along with the version.
// This go.mod file is created when we create a new module.

// Creating go module
// command - go mod init learnmod
