package main

import (
	"fmt"
	"math"
)

const diff = 0.001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	
	//complex numbers not supported!
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	z     := 1.0
	z0    := 0.0
	delta := 1.0
		
	for delta > diff {
		z0 = z
		z  = z - (z*z - x) / (2*z) 
		delta = math.Abs(z - z0)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

