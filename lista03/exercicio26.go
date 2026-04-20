package main

import f "fmt"

func main() {
	s := 0.0
	fatorial := 1.0

	for k := 0; k < 20; k++ {
		if k > 0 {
			fatorial *= float64(k)
		}
		s += float64(100-k) / fatorial
	}
	f.Printf("Soma dos 20 primeiros termos da série: %.6f\n", s)
}
