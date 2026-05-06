package main

import f "fmt"

func main() {
	var a [10]int
	var b [5]int

	f.Println("Digite os 10 números do primeiro vetor:")
	for i := range a {
		f.Scan(&a[i])
	}
	f.Println("Digite os 5 números do segundo vetor:")
	for i := range b {
		f.Scan(&b[i])
	}

	somaB := 0
	for _, v := range b {
		somaB += v
	}

	var pares, impares []int
	for _, v := range a {
		if v%2 == 0 {
			pares = append(pares, v+somaB)
		} else {
			impares = append(impares, v+somaB)
		}
	}

	f.Println("Vetor resultante dos pares:", pares)
	f.Println("Vetor resultante dos ímpares:", impares)
}
