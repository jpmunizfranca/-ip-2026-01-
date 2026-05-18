package main

import "fmt"

func main() {
	h := 0.0
	numerador := 1.0
	for k := 1; k <= 50; k++ {
		h = h + numerador/float64(k)
		numerador = numerador + 2
	}
	fmt.Printf("H = %.6f\n", h)
}
