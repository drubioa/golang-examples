package main

import (
	"fmt"
)

func newtonmethod(x float64, niter int) float64 {
	z := float64(1)
	for  i := 1 ; (z * z) != x && i < niter; i++ {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func existsPrev(prev []float64, x float64) bool {
	for i := range prev {
		if(prev[i] == x && x != 1){
			return true
		}
	}
	return false
}

type SqrtError struct {
	val float64
}

func (e *SqrtError) Error() string {
	return fmt.Sprintf("invalid value %f", e.val)
}

func Sqrt(x float64) (float64, error) {
	var result float64
	prev := make([]float64, 3)
	niter := 1
	
	if(x < 0.) {
		return 0., &SqrtError{
			x,
		}
	}

	for i:=0 ; i < 30; i++ {
		result = newtonmethod(x, i)
		if existsPrev(prev, result) || x == 1. || x == 0.{
			break
		} else {
			index := int32(i % 3)
			prev[index] = result
		}
		niter++
	}

	if(x == 0) {
		result = 1.
	}

	//fmt.Printf("Sqrt of %f is %1.14f in %d iterations\n", x, result, niter)
	return result, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}