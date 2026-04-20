package main

import f "fmt"

func main() {
	s := 0.0
	sinal := 1.0
	numerador := 1.0

	for k := 0; k <= 14; k++ {
		denominador := float64((15 - k) * (15 - k))
		s += sinal * numerador / denominador
		sinal *= -1
		numerador *= 2
	}
	f.Printf("S = %.6f\n", s)
}
