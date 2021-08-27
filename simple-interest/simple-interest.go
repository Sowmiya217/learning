package simpleinterest

import (
	"fmt"
)

func init() {
	fmt.Println("Simple interest package initialised !")
}

func Calculate(p, t, r float64) (i float64) {
	i = p * t * (r / 100)
	return
}
