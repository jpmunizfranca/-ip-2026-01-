package main

import f "fmt"

func main() {
	var v [30]int
	f.Println("Digite 30 números inteiros:")
	for i := range v {
		f.Scan(&v[i])
	}

	var resultado [30]int
	for i, n := range v {
		if i%2 == 0 {
			resultado[i] = n * 2
		} else {
			resultado[i] = n * 3
		}
	}
	f.Println("Vetor resultante:", resultado)
}
