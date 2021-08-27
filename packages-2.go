package main

import (
	"fmt"
	// bc "learnpackage/base-concept"
	// _ "learnpackage/simple-interest"
)

// Sometimes we need to import a package just to make sure the initialization takes place
// even though we do not need to use any function or variable from the package.
// For example, we might need to ensure that the init function of the simpleinterest package is called
// even though we plan not to use that package anywhere in our code.
// The _ blank identifier can be used in this case too as shown below.

// package main

// import (
//     _ "learnpackage/simpleinterest"
// )

// func main() {

// }
// Running the above program will output Simple interest package initialized.
// We have successfully initialized the simpleinterest package even though it is not used anywhere in the code.

func init() {
	fmt.Println("Init func in packages - 2 file")
}

func Packages2() {
	// bc.Packages()
}
