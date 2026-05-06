package main

import f "fmt"

func main() {
	var v [10]int
	f.Println("Digite 10 números inteiros:")
	for i := range v {
		f.Scan(&v[i])
	}

	menor, pos := v[0], 0
	for i, n := range v {
		if n < menor {
			menor, pos = n, i
		}
	}
	f.Printf("O menor elemento do vetor é %d e sua posição dentro do vetor é: %d\n", menor, pos)
}
