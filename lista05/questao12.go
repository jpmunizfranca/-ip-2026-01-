package main

import f "fmt"

func main() {
	var notas [15]int
	freq := [11]int{}

	f.Println("Digite as 15 notas (0 a 10):")
	for i := range notas {
		f.Scan(&notas[i])
		freq[notas[i]]++
	}

	f.Println("\nNota | Freq. Absoluta | Freq. Relativa")
	for nota, fa := range freq {
		fr := float64(fa) / 15.0
		f.Printf("  %2d |      %2d        |    %.4f\n", nota, fa, fr)
	}
}
