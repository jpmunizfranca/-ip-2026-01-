package main

import f "fmt"

func main() {
	var v1, v2 [10]int
	f.Println("Digite os 10 elementos do vetor 1:")
	for i := range v1 {
		f.Scan(&v1[i])
	}
	f.Println("Digite os 10 elementos do vetor 2:")
	for i := range v2 {
		f.Scan(&v2[i])
	}

	resultado := make([]int, 0, 20)
	for i := 0; i < 10; i++ {
		resultado = append(resultado, v1[i], v2[i])
	}
	f.Println("Vetor intercalado:", resultado)
}
