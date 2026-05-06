package main

import f "fmt"

func main() {
	var jogadas [20]int
	freq := [7]int{}

	f.Println("Digite os 20 números sorteados no dado (1 a 6):")
	for i := range jogadas {
		f.Scan(&jogadas[i])
		freq[jogadas[i]]++
	}

	f.Println("Números sorteados:", jogadas)
	f.Println("\nFrequência:")
	for face := 1; face <= 6; face++ {
		f.Printf("Face %d: %d vezes\n", face, freq[face])
	}
}
