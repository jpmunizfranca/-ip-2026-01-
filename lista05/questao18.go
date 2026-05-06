package main

import f "fmt"

func main() {
	var v [10]int
	f.Println("Digite 10 números em ordem crescente:")
	for i := range v {
		for {
			f.Scan(&v[i])
			if i == 0 || v[i] >= v[i-1] {
				break
			}
			f.Printf("Número deve ser >= %d. Tente novamente: ", v[i-1])
		}
	}
	f.Println("Vetor ordenado:", v)
}
