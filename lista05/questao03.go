package main

import f "fmt"

func main() {
	var v [10]int
	f.Println("Digite 10 números inteiros:")
	for i := range v {
		f.Scan(&v[i])
	}

	var pares, impares []int
	somaPares := 0
	for _, n := range v {
		if n%2 == 0 {
			pares = append(pares, n)
			somaPares += n
		} else {
			impares = append(impares, n)
		}
	}

	f.Println("Números pares:", pares)
	f.Println("Soma dos pares:", somaPares)
	f.Println("Números ímpares:", impares)
	f.Println("Quantidade de ímpares:", len(impares))
}
