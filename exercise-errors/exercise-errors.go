package main

import (
	"fmt"
)
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	// if we just e in the above line, then it goes to infinite loop as Sprintf %v
	// has not been defined for ErrNegativeSqrt type
}

func Sqrt(x float64) (float64, error) {

	if x<0 {
		return 0.0, ErrNegativeSqrt(x)
	}

	z := x
	for i:=0; i<10; i +=1 {
		z -= (z*z -x)/(2*z)
		//fmt.Println(z)
	}
	return z,nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
