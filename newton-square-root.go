package main

import "fmt"

func Sqrt(x float64) float64 {
	z := x / 2
	// Xn - f(Xn)/f'(Xn) -> x^2 - 2 - > 2x -> Xn - (Xn*Xn - x)/2*Xn
	for {
		temp := z - ((z*z - x) / (2 * z))
		if (z-temp > 0 && z-temp < 0.0001) || z == temp {
			return temp
		}
		z = temp

	}
}

func main() {
	fmt.Println(Sqrt(4))
}
