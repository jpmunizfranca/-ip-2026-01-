package main

import f "fmt"

func main() {
	var n int
	f.Print("Digite o número de termos (N): ")
	f.Scan(&n)

	if n <= 0 {
		f.Println("Número de termos inválido.")
		return
	}

	s := 0.0
	sinal := 1.0
	numerador := 1000.0

	for k := 1; k <= n; k++ {
		s += sinal * numerador / float64(k)
		sinal *= -1
		numerador -= 3
	}
	f.Printf("S = %.6f\n", s)
}
