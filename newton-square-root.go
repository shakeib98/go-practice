package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	// Xn - f(Xn)/f'(Xn) -> x^2 - 2 - > 2x -> Xn - (Xn*Xn - x)/2*Xn
	z := x / 2
	for {
		temp := z - ((z*z - x) / (2 * z))
		if z-temp < 0.0001 {
			return temp
		}
		z = temp
	}
}

func main() {
	fmt.Println(Sqrt(16))
}
