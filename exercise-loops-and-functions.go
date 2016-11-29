package main

import (
	"fmt"
	"math"
)

const diff = 0.001

func Sqrt(x float64) float64 {
	
	z     := 1.0
	z0    := 0.0
	delta := 1.0
		
	for delta > diff {
		z0 = z
		z  = z - (z*z - x) / (2*z) 
		delta = math.Abs(z - z0)
	}

	return z
}

func main() {

	fmt.Println(Sqrt(16))
}


