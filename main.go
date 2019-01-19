package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %d", e)
}

func Sqrt(x float64) (v float64, e error) {	
	if x < 0 {
		return 0, ErrNegativeSqrt(float64(x))
	}
	v = x
	for i := 0; i < 10; i++ {
		v = v - (v*v-x)/(2*v)
	}
	return
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
