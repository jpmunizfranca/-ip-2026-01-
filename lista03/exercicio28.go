package main

import (
	f "fmt"
	"math"
)

func main() {
	s := 0.0
	sinal := 1.0

	for k := 0; k < 51; k++ {
		base := float64(2*k + 1)
		s += sinal / (base * base * base)
		sinal *= -1
	}

	pi := math.Sqrt(3) * s * 32
	f.Printf("Valor aproximado de π com 51 termos: %.6f\n", pi)
}
