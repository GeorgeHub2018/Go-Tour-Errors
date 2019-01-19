package main

import (
	"fmt"
)

//ErrNegativeSqrt -
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %d", float64(e))
}

func sqrt(x float64) (v float64, e error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	v = x
	for i := 0; i < 10; i++ {
		v = v - (v*v-x)/(2*v)
	}
	return
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
