package main

import (
	"fmt"
)

func Sqrt(x float64, niter int) float64 {
	z := float64(1)
	for  i := 1 ; (z * z) != x && i < niter; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func ExistsPrev(prev []float64, x float64) bool {
	for i := range prev {
		if(prev[i] == x && x != 1){
			return true
		}
	}
	return false
}

func main() {
	var result float64
	prev := make([]float64, 3)
	niter := 1
	x := 3242424.
	
	for i:=0 ; i < 30; i++ {
		//fmt.Println(prev)
		result = Sqrt(x, i)
		if ExistsPrev(prev, result) || x == 1. || x == 0.{
			break
		} else {
			var index = int32(i % 3)
			prev[index] = result
		}
		niter++
	}

	if(x == 0) {
		result = 1.
	}

	fmt.Printf("Sqrt of %f is %1.14f in %d iterations\n", x, result, niter)

}